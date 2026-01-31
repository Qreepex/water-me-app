<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { fetchData } from '$lib/auth/fetch.svelte';
	import { imageCacheStore } from '$lib/stores/imageCache.svelte';
	import PageHeader from '$lib/components/layout/PageHeader.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import Card from '$lib/components/ui/Card.svelte';
	import LoadingSpinner from '$lib/components/ui/LoadingSpinner.svelte';
	import EmptyState from '$lib/components/ui/EmptyState.svelte';
	import Alert from '$lib/components/ui/Message.svelte';
	import { getPlantsStore } from '$lib/stores/plants.svelte';

	const store = getPlantsStore();
	let error = $state<string | null>(null);
	let deleting = $state<string | null>(null);
	let previews = $state<Record<string, string>>({});

	// Load previews for all plants from cache
	$effect(() => {
		const loadPreviews = async () => {
			for (const plant of store.plants) {
				const firstId = plant.photoIds?.[0];
				// eslint-disable-next-line @typescript-eslint/no-explicit-any
				const firstUrl = (plant as any)?.photoUrls?.[0] as string | undefined;
				if (firstId && firstUrl) {
					const objUrl = await imageCacheStore.getImageURL(firstId, firstUrl);
					if (objUrl) previews[plant.id] = objUrl;
				}
			}
		};

		loadPreviews();
	});

	async function deletePlant(id: string, name: string): Promise<void> {
		if (!confirm(`Delete "${name}"? This cannot be undone.`)) return;

		deleting = id;
		try {
			const response = await fetchData('/api/plants/{id}', {
				method: 'delete',
				params: { id }
			});

			if (!response.ok) {
				error = response.error?.message || 'Failed to delete plant';
				return;
			}

			// Update store
			store.setPlants(store.plants.filter((p) => p.id !== id));
		} catch (err) {
			error = err instanceof Error ? err.message : 'Failed to delete plant';
		} finally {
			deleting = null;
		}
	}

	function editPlant(id: string): void {
		goto(resolve(`/manage/${id}`));
	}

	function createNewPlant(): void {
		goto(resolve('/manage/create'));
	}
</script>

	<!-- Header -->
	<PageHeader icon="🌿" title="plants.myPlants" description="plants.manageDescription">
		<Button variant="primary" onclick={createNewPlant} text="plants.newPlant" />
	</PageHeader>

	<!-- Messages -->
	{#if error}
		<Alert type="error" title="Error" description={error} />
	{/if}

	<!-- Loading State -->
	{#if store.loading}
		<LoadingSpinner message="Loading your plants..." icon="🌱" />
	{:else if store.plants.length === 0}
		<!-- Empty State -->
		<EmptyState icon="🪴" title="No plants yet" description="Start building your plant collection">
			<Button variant="primary" onclick={createNewPlant} text="addPlant" />
		</EmptyState>
	{:else}
		<!-- Plants Grid -->
		<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
			{#each store.plants as plant (plant.id)}
				<Card rounded="2xl">
					{#if previews[plant.id]}
						<img
							src={previews[plant.id]}
							alt={plant.name}
							class="h-40 w-full rounded-t-2xl object-cover"
						/>
					{/if}
					<!-- Plant Header -->
					<div class="bg-gradient-to-r from-[var(--bg-light)] to-[var(--p-emerald)]/20 p-4">
						<h3 class="text-xl font-bold text-[var(--text-light-main)]">{plant.name}</h3>
						<p class="text-sm text-[var(--status-success)] italic">{plant.species}</p>
					</div>

					<!-- Plant Details -->
					<div class="space-y-2 p-4 text-sm">
						{#if plant.sunlight}
							<div class="flex items-center gap-2">
								<span>☀️</span>
								<span class="text-[var(--text-light-main)]">{plant.sunlight}</span>
							</div>
						{/if}

						{#if plant.watering?.intervalDays}
							<div class="flex items-center gap-2">
								<span>💧</span>
								<span class="text-[var(--text-light-main)]"
									>Every {plant.watering.intervalDays} days</span
								>
							</div>
						{/if}

						{#if plant.location?.room}
							<div class="flex items-center gap-2">
								<span>📍</span>
								<span class="text-[var(--text-light-main)]">{plant.location.room}</span>
							</div>
						{/if}

						{#if plant.isToxic}
							<div
								class="inline-block rounded bg-[var(--status-error)]/20 px-2 py-1 text-xs font-medium text-[var(--status-error)]"
							>
								⚠️ Toxic
							</div>
						{/if}
					</div>

					<!-- Actions -->
					<div class="flex gap-2 border-t border-[var(--p-emerald)]/30 p-4">
						<Button variant="primary" size="md" onclick={() => editPlant(plant.id)} text="edit" />

						<Button
							variant="danger"
							size="md"
							disabled={deleting === plant.id}
							onclick={() => deletePlant(plant.id, plant.name)}
							text={deleting === plant.id ? 'deleting' : 'delete'}
						/>
					</div>
				</Card>
			{/each}
		</div>
	{/if}
