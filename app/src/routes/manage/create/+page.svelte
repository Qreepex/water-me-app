<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { fetchData } from '$lib/auth/fetch.svelte';
	import type { FormData } from '$lib/types/forms';
	import { createEmptyFormData } from '$lib/types/forms';
	import BasicInformationForm from '$lib/components/PlantForms/BasicInformationForm.svelte';
	import LocationForm from '$lib/components/PlantForms/LocationForm.svelte';
	import WateringForm from '$lib/components/PlantForms/WateringForm.svelte';
	import FertilizingForm from '$lib/components/PlantForms/FertilizingForm.svelte';
	import MistingForm from '$lib/components/PlantForms/MistingForm.svelte';
	import SoilForm from '$lib/components/PlantForms/SoilForm.svelte';
	import SeasonalityForm from '$lib/components/PlantForms/SeasonalityForm.svelte';
	import MetadataForm from '$lib/components/PlantForms/MetadataForm.svelte';

	let formData: FormData = createEmptyFormData();
	let error: string | null = null;
	let success: string | null = null;
	let submitting = false;

	let newNote = '';
	let soilComponentInput = '';

	// Upload state
	const MAX_BYTES = 2 * 1024 * 1024; // 2MB
	const allowedTypes = new Set(['image/jpeg', 'image/png', 'image/webp']);
	type PhotoItem = {
		fileName: string;
		previewUrl: string;
		status: 'pending' | 'compressing' | 'uploading' | 'uploaded' | 'error';
		error?: string;
		key?: string;
	};
	let photos: PhotoItem[] = [];
	let uploadedPhotoKeys: string[] = [];

	function revokePreviews(): void {
		photos.forEach((p) => p.previewUrl && URL.revokeObjectURL(p.previewUrl));
	}

	function onFilesSelected(e: Event): void {
		const input = e.target as HTMLInputElement;
		const files = input.files ? Array.from(input.files) : [];
		if (!files.length) return;
		revokePreviews();
		photos = files.map((f) => ({
			fileName: f.name,
			previewUrl: URL.createObjectURL(f),
			status: 'pending'
		}));
		// Process sequentially to avoid spikes
		processUploads(files).catch((err) => {
			error = err instanceof Error ? err.message : 'Upload error';
		});
	}

	async function processUploads(files: File[]): Promise<void> {
		for (let i = 0; i < files.length; i++) {
			const file = files[i];
			const item = photos[i];
			if (!allowedTypes.has(file.type)) {
				item.status = 'error';
				item.error = 'Unsupported file type';
				continue;
			}
			item.status = 'compressing';
			const { blob, contentType, outName } = await compressToUnder2MB(file);
			item.status = 'uploading';
			// eslint-disable-next-line @typescript-eslint/no-explicit-any
			const presignRes = await (fetchData as any)('/api/uploads/presign', {
				method: 'post',
				body: { filename: outName, contentType, sizeBytes: blob.size }
			});
			if (!presignRes.ok) {
				item.status = 'error';
				item.error = presignRes.error?.message || 'Failed to presign';
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
				item.error = 'Upload failed';
				continue;
			}
			// eslint-disable-next-line @typescript-eslint/no-explicit-any
			const regRes = await (fetchData as any)('/api/uploads/register', {
				method: 'post',
				body: { key }
			});
			if (!regRes.ok) {
				item.status = 'error';
				item.error = regRes.error?.message || 'Register failed';
				continue;
			}
			item.status = 'uploaded';
			item.key = key;
			uploadedPhotoKeys.push(key);
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
		// Convert to WebP for better compression
		const targetType = 'image/webp';
		let bitmap = await createImageBitmap(file);
		// Downscale if extremely large
		const maxDim = 3000;
		if (bitmap.width > maxDim || bitmap.height > maxDim) {
			const scale = Math.min(maxDim / bitmap.width, maxDim / bitmap.height);
			bitmap = await downscaleBitmap(bitmap, scale);
		}
		let quality = 0.92;
		let blob = await blobFromImage(bitmap, targetType, quality);
		// Iteratively reduce quality, then dimensions if needed
		let attempts = 0;
		while (blob.size > MAX_BYTES && attempts < 6) {
			quality = Math.max(0.4, quality - 0.15);
			blob = await blobFromImage(bitmap, targetType, quality);
			attempts++;
		}
		if (blob.size > MAX_BYTES) {
			// Reduce dimensions by 80% and retry up to 3 times
			for (let i = 0; i < 3 && blob.size > MAX_BYTES; i++) {
				bitmap = await downscaleBitmap(bitmap, 0.8);
				quality = Math.max(0.5, quality - 0.1);
				blob = await blobFromImage(bitmap, targetType, quality);
			}
		}
		if (blob.size > MAX_BYTES) {
			throw new Error('Unable to compress under 2MB');
		}
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

	async function submitForm(): Promise<void> {
		if (!formData.name.trim()) {
			error = 'Plant name is required';
			return;
		}

		submitting = true;
		error = null;
		success = null;

		try {
			// eslint-disable-next-line @typescript-eslint/no-explicit-any
			const createPayload: any = {
				name: formData.name,
				isToxic: formData.isToxic,
				preferedTemperature: formData.preferedTemperature,
				watering: {
					intervalDays: formData.wateringIntervalDays,
					method: formData.wateringMethod,
					waterType: formData.waterType,
					lastWatered: null
				},
				pestHistory: [],
				flags: formData.flags,
				notes: formData.notes,
				photoIds: uploadedPhotoKeys,
				growthHistory: []
			};

			// Only include optional fields if they're set
			if (formData.species) {
				createPayload.species = formData.species;
			}

			if (formData.sunlight) {
				createPayload.sunlight = formData.sunlight;
			}

			if (formData.room || formData.position) {
				createPayload.location = {
					room: formData.room,
					position: formData.position,
					isOutdoors: formData.isOutdoors
				};
			}

			// Check if fertilizing has non-default values
			if (
				formData.fertilizingIntervalDays !== 30 ||
				formData.npkRatio !== '10:10:10' ||
				formData.concentrationPercent !== 50 ||
				formData.activeInWinter
			) {
				createPayload.fertilizing = {
					type: formData.fertilizingType,
					intervalDays: formData.fertilizingIntervalDays,
					npkRatio: formData.npkRatio,
					concentrationPercent: formData.concentrationPercent,
					lastFertilized: null,
					activeInWinter: formData.activeInWinter
				};
			}

			// Check if humidity has non-default values
			if (
				formData.requiresMisting ||
				formData.requiresHumidifier ||
				formData.targetHumidity !== 50
			) {
				createPayload.humidity = {
					requiresMisting: formData.requiresMisting,
					mistingIntervalDays: formData.mistingIntervalDays,
					requiresHumidifier: formData.requiresHumidifier,
					targetHumidityPct: formData.targetHumidity
				};
			}

			// Check if soil has non-default values
			if (
				formData.soilType !== 'Generic' ||
				formData.repottingCycle !== 2 ||
				formData.soilComponents.length > 0
			) {
				createPayload.soil = {
					type: formData.soilType,
					components: formData.soilComponents,
					repottingCycle: formData.repottingCycle
				};
			}

			// Check if seasonality has non-default values
			if (
				formData.winterRestPeriod ||
				formData.winterWaterFactor !== 0.5 ||
				formData.minTempCelsius !== 15
			) {
				createPayload.seasonality = {
					winterRestPeriod: formData.winterRestPeriod,
					winterWaterFactor: formData.winterWaterFactor,
					minTempCelsius: formData.minTempCelsius
				};
			}

			const createRes = await fetchData('/api/plants', {
				method: 'post',
				body: createPayload
			});
			if (!createRes.ok) throw new Error(createRes.error?.message || 'Failed to create plant');
			success = 'Plant created successfully!';
			setTimeout(() => goto(resolve('/manage')), 1500);
		} catch (err) {
			error = err instanceof Error ? err.message : 'Unknown error';
		} finally {
			submitting = false;
		}
	}

	function resetForm(): void {
		formData = createEmptyFormData();
		newNote = '';
		soilComponentInput = '';
		error = null;
	}
</script>

<div class="min-h-screen bg-gradient-to-br from-emerald-50 via-green-50 to-teal-100 p-6 md:p-10">
	<div class="mx-auto max-w-4xl">
		<!-- Header -->
		<div class="mb-8">
			<div class="mb-4 flex items-center justify-between">
				<div>
					<h1 class="flex items-center gap-3 text-4xl font-bold text-green-900">
						🌱 Create New Plant
					</h1>
					<p class="mt-1 text-sm text-emerald-700 italic">Add a new plant to your collection</p>
				</div>
				<a
					href={resolve('/manage')}
					class="rounded-xl bg-gray-600 px-4 py-2 font-medium text-white shadow-sm transition hover:bg-gray-700"
				>
					← Back
				</a>
			</div>
		</div>

		<!-- Messages -->
		{#if success}
			<div class="mb-6 rounded-lg border-2 border-green-400 bg-green-100 px-6 py-4 text-green-800">
				✓ {success}
			</div>
		{/if}

		{#if error}
			<div class="mb-6 rounded-lg border-2 border-red-400 bg-red-100 px-6 py-4 text-red-800">
				✕ {error}
			</div>
		{/if}

		<div class="space-y-6">
			<!-- Images Section -->
			<div class="rounded-2xl border border-emerald-100 bg-white/90 p-6 shadow-md backdrop-blur">
				<h2 class="mb-4 text-2xl font-bold text-green-800">📸 Photos</h2>
				<div class="space-y-4">
					<label class="block">
						<span class="text-sm font-medium text-green-800"
							>Add images (JPEG/PNG/WebP, auto-compressed ≤ 2MB)</span
						>
						<input
							type="file"
							accept="image/jpeg,image/png,image/webp"
							multiple
							on:change={onFilesSelected}
							class="mt-2 w-full rounded-lg border border-emerald-200 bg-white p-2 text-sm"
						/>
					</label>

					{#if photos.length}
						<div class="grid grid-cols-2 gap-3 md:grid-cols-4">
							{#each photos as p (p.previewUrl)}
								<div class="rounded-md border border-emerald-200 bg-emerald-50 p-2">
									<img
										src={p.previewUrl}
										alt={p.fileName}
										class="h-24 w-full rounded object-cover"
									/>
									<div class="mt-1 text-xs text-emerald-800">
										{p.fileName}
									</div>
									<div class="text-xs">
										{#if p.status === 'pending'}
											<span class="text-gray-600">Pending</span>
										{:else if p.status === 'compressing'}
											<span class="text-blue-600">Compressing...</span>
										{:else if p.status === 'uploading'}
											<span class="text-emerald-600">Uploading...</span>
										{:else if p.status === 'uploaded'}
											<span class="text-green-700">Uploaded</span>
										{:else}
											<span class="text-red-600">{p.error || 'Error'}</span>
										{/if}
									</div>
								</div>
							{/each}
						</div>
					{/if}
				</div>
			</div>

			<!-- Basic Information -->
			<BasicInformationForm {formData} />

			<!-- Location -->
			<LocationForm {formData} />

			<!-- Watering & Fertilizing -->
			<div class="grid grid-cols-1 gap-6 md:grid-cols-2">
				<WateringForm {formData} />
				<FertilizingForm {formData} />
			</div>

			<!-- Misting -->
			<MistingForm {formData} />

			<!-- Advanced Section -->
			<div
				class="space-y-6 rounded-2xl border border-emerald-100 bg-white/90 p-6 shadow-md backdrop-blur"
			>
				<h2 class="text-xl font-bold text-green-800">⚙️ Advanced Settings</h2>

				<div class="space-y-6">
					<SoilForm {formData} bind:soilComponentInput />
					<SeasonalityForm {formData} />
					<MetadataForm {formData} bind:newNote />
				</div>
			</div>

			<div class="flex justify-between gap-3">
				<button
					on:click={resetForm}
					class="rounded-lg bg-gray-200 px-6 py-3 font-semibold text-gray-800 transition hover:bg-gray-300"
				>
					Reset
				</button>
				<button
					on:click={submitForm}
					disabled={submitting}
					class="rounded-lg bg-gradient-to-r from-emerald-600 to-green-600 px-8 py-3 font-semibold text-white shadow-md transition hover:from-emerald-700 hover:to-green-700 disabled:opacity-50"
				>
					{submitting ? 'Creating...' : 'Create Plant'}
				</button>
			</div>
		</div>
	</div>
</div>

<style>
	:global(body) {
		font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
	}
</style>
