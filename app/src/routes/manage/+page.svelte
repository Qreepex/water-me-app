<script lang="ts">
	import type { Plant } from '$lib/types/api';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { fetchData } from '$lib/auth/fetch.svelte';

	let plants: Plant[] = [];
	let loading = true;
	let error: string | null = null;
	let deleting: string | null = null;

	async function loadPlants(): Promise<void> {
		try {
			const response = await fetchData('/api/plants', {});
			if (!response.ok) {
				error = response.error?.message || 'Failed to load plants';
				return;
			}
			plants = response.data || [];
		} catch (err) {
			error = err instanceof Error ? err.message : 'Failed to load plants';
		} finally {
			loading = false;
		}
	}

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

			plants = plants.filter((p) => p.id !== id);
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

	loadPlants();
</script>

<div class="min-h-screen bg-gradient-to-br from-emerald-50 via-green-50 to-teal-100 p-6 md:p-10">
	<div class="mx-auto max-w-6xl">
		<!-- Header -->
		<div class="mb-8 flex items-center justify-between">
			<div>
				<h1 class="text-4xl font-bold text-green-900">🌿 My Plants</h1>
				<p class="mt-1 text-gray-600">Manage your plant collection</p>
			</div>
			<button
				on:click={createNewPlant}
				class="rounded-lg bg-gradient-to-r from-emerald-600 to-green-600 px-6 py-3 font-semibold text-white shadow-md transition hover:from-emerald-700 hover:to-green-700"
			>
				+ New Plant
			</button>
		</div>

		<!-- Messages -->
		{#if error}
			<div class="mb-6 rounded-lg border-2 border-red-400 bg-red-100 px-6 py-4 text-red-800">
				✕ {error}
			</div>
		{/if}

		<!-- Loading State -->
		{#if loading}
			<div class="flex min-h-96 items-center justify-center">
				<div class="text-center">
					<div class="mb-4 animate-spin text-5xl">🌱</div>
					<p class="text-lg text-gray-600">Loading your plants...</p>
				</div>
			</div>
		{:else if plants.length === 0}
			<!-- Empty State -->
			<div
				class="rounded-2xl border-2 border-dashed border-emerald-300 bg-emerald-50 p-12 text-center"
			>
				<div class="mb-4 text-5xl">🪴</div>
				<h2 class="mb-2 text-2xl font-bold text-green-900">No plants yet</h2>
				<p class="mb-6 text-gray-600">Start building your plant collection</p>
				<button
					on:click={createNewPlant}
					class="rounded-lg bg-gradient-to-r from-emerald-600 to-green-600 px-6 py-3 font-semibold text-white shadow-md transition hover:from-emerald-700 hover:to-green-700"
				>
					Add Your First Plant
				</button>
			</div>
		{:else}
			<!-- Plants Grid -->
			<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
				{#each plants as plant (plant.id)}
					<div
						class="overflow-hidden rounded-2xl border border-emerald-100 bg-white/90 shadow-md backdrop-blur transition hover:shadow-lg"
					>
						<!-- Plant Header -->
						<div class="bg-gradient-to-r from-emerald-100 to-green-100 p-4">
							<h3 class="text-xl font-bold text-green-900">{plant.name}</h3>
							<p class="text-sm text-emerald-700 italic">{plant.species}</p>
						</div>

						<!-- Plant Details -->
						<div class="space-y-2 p-4 text-sm">
							{#if plant.sunlight}
								<div class="flex items-center gap-2">
									<span>☀️</span>
									<span class="text-gray-700">{plant.sunlight}</span>
								</div>
							{/if}

							{#if plant.watering?.intervalDays}
								<div class="flex items-center gap-2">
									<span>💧</span>
									<span class="text-gray-700">Every {plant.watering.intervalDays} days</span>
								</div>
							{/if}

							{#if plant.location?.room}
								<div class="flex items-center gap-2">
									<span>📍</span>
									<span class="text-gray-700">{plant.location.room}</span>
								</div>
							{/if}

							{#if plant.isToxic}
								<div
									class="inline-block rounded bg-red-100 px-2 py-1 text-xs font-medium text-red-800"
								>
									⚠️ Toxic
								</div>
							{/if}
						</div>

						<!-- Actions -->
						<div class="flex gap-2 border-t border-emerald-100 p-4">
							<button
								on:click={() => editPlant(plant.id)}
								class="flex-1 rounded-lg bg-blue-600 px-4 py-2 font-medium text-white shadow-sm transition hover:bg-blue-700"
							>
								Edit
							</button>
							<button
								on:click={() => deletePlant(plant.id, plant.name)}
								disabled={deleting === plant.id}
								class="flex-1 rounded-lg bg-red-600 px-4 py-2 font-medium text-white shadow-sm transition hover:bg-red-700 disabled:opacity-50"
							>
								{deleting === plant.id ? 'Deleting...' : 'Delete'}
							</button>
						</div>
					</div>
				{/each}
			</div>
		{/if}
	</div>
</div>

<style>
	:global(body) {
		font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
	}
</style>
