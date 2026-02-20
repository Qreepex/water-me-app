<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import { fetchData } from '$lib/auth/fetch.svelte';
	import { getImageObjectURL, revokeObjectURL } from '$lib/utils/imageCache';
	import { invalidateApiCache } from '$lib/utils/cache';
	import { tStore } from '$lib/i18n';
	import type { Plant } from '$lib/types/api';
	import PageHeader from '$lib/components/layout/PageHeader.svelte';
	import PageContent from '$lib/components/layout/PageContent.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import LoadingSpinner from '$lib/components/ui/LoadingSpinner.svelte';
	import Alert from '$lib/components/ui/Message.svelte';

	const MAX_BYTES = 2 * 1024 * 1024;
	const allowedTypes = new Set(['image/jpeg', 'image/png', 'image/webp']);

	type PhotoItem = {
		fileName: string;
		previewUrl: string;
		status: 'pending' | 'compressing' | 'uploading' | 'uploaded' | 'error';
		error?: string;
		key?: string;
	};

	let plant = $state<Plant | null>(null);
	let loading = $state(true);
	let saving = $state(false);
	let error = $state<string | null>(null);
	let previewUrls = $state<string[]>([]);
	let photos = $state<PhotoItem[]>([]);
	let uploadedPhotoKeys = $state<string[]>([]);
	let removedPhotoIds = $state<string[]>([]);
	const existingPhotoIds = $derived(plant?.photoIds ?? []);
	const isCreateFlow = $derived(page.url.searchParams.get('createFlow') === '1');
	const hasPhotoChanges = $derived(uploadedPhotoKeys.length > 0 || removedPhotoIds.length > 0);

	onMount(async () => {
		try {
			const plantId = page.params.plant ?? '';
			const response = await fetchData('/api/plants/{id}', {
				params: { id: plantId }
			});

			if (!response.ok) {
				error = response.error?.message || $tStore('plants.failedToFetchPlants');
				return;
			}

			plant = response.data;
			await loadPhotoPreviews();
		} catch (err) {
			error = err instanceof Error ? err.message : $tStore('plants.failedToFetchPlants');
		} finally {
			loading = false;
		}
	});

	onDestroy(() => {
		previewUrls.forEach((u) => revokeObjectURL(u));
		photos.forEach((p) => p.previewUrl && URL.revokeObjectURL(p.previewUrl));
	});

	async function loadPhotoPreviews(): Promise<void> {
		if (!plant) return;
		const ids = plant.photoIds || [];
		// eslint-disable-next-line @typescript-eslint/no-explicit-any
		const urls = ((plant as any)?.photoUrls as string[] | undefined) || [];
		previewUrls = [];
		for (let i = 0; i < ids.length; i++) {
			const id = ids[i];
			const url = urls[i];
			if (!id || !url) continue;
			const objUrl = await getImageObjectURL(id, url);
			if (objUrl) previewUrls.push(objUrl);
		}
	}

	function onFilesSelected(e: Event): void {
		const input = e.target as HTMLInputElement;
		const files = input.files ? Array.from(input.files) : [];
		if (!files.length) return;
		photos = files.map((f) => ({
			fileName: f.name,
			previewUrl: URL.createObjectURL(f),
			status: 'pending'
		}));
		processUploads(files).catch((err) => {
			error = err instanceof Error ? err.message : $tStore('common.uploadError');
		});
	}

	async function processUploads(files: File[]): Promise<void> {
		for (let i = 0; i < files.length; i++) {
			const file = files[i];
			const item = photos[i];
			if (!allowedTypes.has(file.type)) {
				item.status = 'error';
				item.error = $tStore('common.unsupportedFileType');
				photos = [...photos];
				continue;
			}

			item.status = 'compressing';
			photos = [...photos];

			let blob: Blob;
			let contentType: string;
			let outName: string;
			try {
				const result = await compressToUnder2MB(file);
				blob = result.blob;
				contentType = result.contentType;
				outName = result.outName;
			} catch (err) {
				item.status = 'error';
				item.error = err instanceof Error ? err.message : $tStore('common.compressionFailed');
				photos = [...photos];
				continue;
			}

			item.status = 'uploading';
			photos = [...photos];

			// eslint-disable-next-line @typescript-eslint/no-explicit-any
			const presignRes = await (fetchData as any)('/api/uploads/presign', {
				method: 'post',
				body: { filename: outName, contentType, sizeBytes: blob.size }
			});
			if (!presignRes.ok) {
				item.status = 'error';
				item.error = presignRes.error?.message || $tStore('common.failed');
				photos = [...photos];
				continue;
			}

			const { url, headers, key } = presignRes.data as {
				url: string;
				headers: Record<string, string>;
				key: string;
			};

			const putOk = await putToS3(url, headers, blob);
			if (!putOk) {
				item.status = 'error';
				item.error = $tStore('common.failed');
				photos = [...photos];
				continue;
			}

			// eslint-disable-next-line @typescript-eslint/no-explicit-any
			const regRes = await (fetchData as any)('/api/uploads/register', {
				method: 'post',
				body: { key }
			});
			if (!regRes.ok) {
				item.status = 'error';
				item.error = regRes.error?.message || $tStore('common.failed');
				photos = [...photos];
				continue;
			}

			item.status = 'uploaded';
			item.key = key;
			uploadedPhotoKeys = [...uploadedPhotoKeys, key];
			photos = [...photos];
		}
	}

	async function putToS3(
		url: string,
		headers: Record<string, string>,
		blob: Blob
	): Promise<boolean> {
		try {
			const res = await fetch(url, {
				method: 'PUT',
				body: blob,
				headers
			});
			return res.ok;
		} catch {
			return false;
		}
	}

	async function blobFromImage(bitmap: ImageBitmap, type: string, quality: number): Promise<Blob> {
		const canvas = document.createElement('canvas');
		canvas.width = bitmap.width;
		canvas.height = bitmap.height;
		const ctx = canvas.getContext('2d');
		if (!ctx) throw new Error('Canvas unsupported');
		ctx.drawImage(bitmap, 0, 0);
		const b = await new Promise<Blob | null>((resolve) => canvas.toBlob(resolve, type, quality));
		if (!b) throw new Error('Failed to create blob');
		return b;
	}

	async function compressToUnder2MB(
		file: File
	): Promise<{ blob: Blob; contentType: string; outName: string }> {
		const targetType = 'image/webp';
		let bitmap = await createImageBitmap(file);
		const maxDim = 3000;
		if (bitmap.width > maxDim || bitmap.height > maxDim) {
			const scale = Math.min(maxDim / bitmap.width, maxDim / bitmap.height);
			bitmap = await downscaleBitmap(bitmap, scale);
		}
		let quality = 0.92;
		let blob = await blobFromImage(bitmap, targetType, quality);
		let attempts = 0;
		while (blob.size > MAX_BYTES && attempts < 6) {
			quality = Math.max(0.4, quality - 0.15);
			blob = await blobFromImage(bitmap, targetType, quality);
			attempts++;
		}
		if (blob.size > MAX_BYTES) {
			for (let i = 0; i < 3 && blob.size > MAX_BYTES; i++) {
				bitmap = await downscaleBitmap(bitmap, 0.8);
				quality = Math.max(0.5, quality - 0.1);
				blob = await blobFromImage(bitmap, targetType, quality);
			}
		}
		if (blob.size > MAX_BYTES) throw new Error($tStore('common.compressionFailed'));
		const outName = file.name.replace(/\.[^.]+$/, '') + '.webp';
		return { blob, contentType: targetType, outName };
	}

	async function downscaleBitmap(src: ImageBitmap, scale: number): Promise<ImageBitmap> {
		const canvas = document.createElement('canvas');
		canvas.width = Math.max(1, Math.floor(src.width * scale));
		canvas.height = Math.max(1, Math.floor(src.height * scale));
		const ctx = canvas.getContext('2d');
		if (!ctx) throw new Error('Canvas unsupported');
		ctx.imageSmoothingQuality = 'high';
		ctx.drawImage(src, 0, 0, canvas.width, canvas.height);
		const blob = await new Promise<Blob | null>((resolve) => canvas.toBlob(resolve));
		if (!blob) throw new Error('Downscale failed');
		return await createImageBitmap(blob);
	}

	function removeExistingPhoto(photoId: string, index: number): void {
		if (!confirm($tStore('plants.deletePhotoConfirm'))) return;
		removedPhotoIds = [...removedPhotoIds, photoId];
		const newUrls = [...previewUrls];
		const urlToRevoke = newUrls[index];
		newUrls.splice(index, 1);
		previewUrls = newUrls;
		if (urlToRevoke) revokeObjectURL(urlToRevoke);
		void deletePhotoFromS3(photoId);
	}

	async function deletePhotoFromS3(key: string): Promise<void> {
		try {
			// eslint-disable-next-line @typescript-eslint/no-explicit-any
			await (fetchData as any)(`/api/uploads/${encodeURIComponent(key)}`, {
				method: 'delete'
			});
		} catch {
			// ignore
		}
	}

	async function saveAndContinue(): Promise<void> {
		if (!plant) return;
		saving = true;
		error = null;

		try {
			const existingPhotoIds = (plant.photoIds || []).filter((id) => !removedPhotoIds.includes(id));
			const allPhotoIds = [...existingPhotoIds, ...uploadedPhotoKeys];

			if (allPhotoIds.length !== (plant.photoIds || []).length || removedPhotoIds.length > 0) {
				const res = await fetchData('/api/plants/{id}', {
					method: 'patch',
					params: { id: plant.id },
					body: { photoIds: allPhotoIds }
				});

				if (!res.ok) {
					throw new Error(res.error?.message || $tStore('plants.failedToSavePhotos'));
				}
			}

			await invalidateApiCache(['/api/plants', `/api/plants/${plant.id}`], {
				waitForAck: true,
				timeoutMs: 100
			});

			await goto(resolve(`/manage/${plant.id}?createFlow=1`));
		} catch (err) {
			error = err instanceof Error ? err.message : $tStore('plants.failedToSavePhotos');
		} finally {
			saving = false;
		}
	}

	function skipToHub(): void {
		if (!plant) return;
		goto(resolve(`/manage/${plant.id}${isCreateFlow ? '?createFlow=1' : ''}`));
	}
</script>

<PageHeader icon="📸" title="plants.photos" description="plants.managePhotosDescription" />

<PageContent>
	{#if loading}
		<LoadingSpinner message="common.loadingPlantDetails" icon="🌱" />
	{:else if !plant}
		<Alert
			type="error"
			title="common.error"
			description={error || $tStore('common.plantNotFound')}
		/>
	{:else}
		{#if error}
			<Alert type="error" title="common.error" description={error} />
		{/if}

		<div class="space-y-4 pb-[calc(env(safe-area-inset-bottom)+12rem)]">
			<label class="block rounded-xl border border-emerald-200 bg-white p-4">
				<span class="text-base font-semibold text-emerald-900">{$tStore('plants.addImages')}</span>
				<input
					type="file"
					accept="image/jpeg,image/png,image/webp"
					multiple
					onchange={onFilesSelected}
					aria-label={$tStore('plants.addImages')}
					class="mt-2 w-full cursor-pointer rounded-lg border border-emerald-300 bg-white p-3 text-base"
				/>
			</label>

			{#if photos.length}
				<div class="rounded-xl border border-emerald-200 bg-white p-3">
					<p class="mb-2 text-sm font-semibold text-emerald-700">{$tStore('plants.newUploads')}</p>
					<div class="grid grid-cols-2 gap-2">
						{#each photos as p (p.previewUrl)}
							<div class="overflow-hidden rounded-lg border border-emerald-200 bg-emerald-50">
								<img src={p.previewUrl} alt={p.fileName} class="h-20 w-full object-cover" />
								<div class="p-2 text-sm">
									<div class="mb-1 truncate font-medium text-emerald-900">{p.fileName}</div>
									{#if p.status === 'uploaded'}
										<span class="font-semibold text-green-700"
											>✓ {$tStore('plants.uploadUploaded')}</span
										>
									{:else if p.status === 'error'}
										<span class="text-red-600">✕ {p.error || $tStore('common.error')}</span>
									{:else}
										<span class="text-gray-600"
											>{p.status === 'pending'
												? $tStore('plants.uploadPending')
												: p.status === 'compressing'
													? $tStore('plants.uploadCompressing')
													: $tStore('plants.uploadUploading')}</span
										>
									{/if}
								</div>
							</div>
						{/each}
					</div>
				</div>
			{/if}

			{#if previewUrls.length}
				<div class="rounded-xl border border-emerald-200 bg-white p-3">
					<p class="mb-2 text-sm font-semibold text-emerald-700">
						{$tStore('plants.existingPhotos')}
					</p>
					<div class="grid grid-cols-2 gap-2">
						{#each previewUrls as u, i (u)}
							<div class="group relative overflow-hidden rounded-lg border border-emerald-200">
								<img
									src={u}
									alt={plant.name || $tStore('common.plant')}
									class="h-20 w-full object-cover"
								/>
								<button
									onclick={() => removeExistingPhoto(existingPhotoIds[i] ?? '', i)}
									aria-label={$tStore('plants.deletePhoto')}
									class="absolute inset-0 flex items-center justify-center bg-red-600/80 text-sm font-bold text-white opacity-0 transition-all group-hover:opacity-100"
								>
									{$tStore('common.delete')}
								</button>
							</div>
						{/each}
					</div>
				</div>
			{/if}
		</div>

		<div
			class="fixed right-3 left-3 z-50 flex gap-3 rounded-2xl border border-gray-200 bg-white/95 p-3 shadow-lg backdrop-blur md:right-10 md:left-10 xl:right-32 xl:left-32"
			style="bottom: calc(env(safe-area-inset-bottom) + 5.5rem);"
		>
			<Button
				variant="secondary"
				size="lg"
				onclick={skipToHub}
				text={isCreateFlow ? 'common.skip' : 'common.close'}
				class="w-full"
			/>
			<Button
				variant="primary"
				size="lg"
				disabled={saving || !hasPhotoChanges}
				onclick={saveAndContinue}
				text={saving ? 'common.loading' : 'common.save'}
				class="w-full"
			/>
		</div>
	{/if}
</PageContent>
