<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { tStore } from '$lib/i18n';
	import { fetchData } from '$lib/auth/fetch.svelte';
	import SortControls from '$lib/components/SortControls.svelte';
	import PlantCard from '$lib/components/PlantCard.svelte';
	import type { Plant } from '$lib/types/api';

	let plants: Plant[] = [];
	let loading = true;
	let error: string | null = null;
	type SortOption =
		| 'name'
		| 'lastWatered'
		| 'lastFertilized'
		| 'wateringIntervalDays'
		| 'mistingIntervalDays';
	let sortBy: SortOption = 'name';

	async function loadPlants() {
		try {
			const result = await fetchData('/api/plants', {});
			if (!result.ok) {
				error = result.error?.message || 'Failed to fetch plants';
				return;
			}
			plants = result.data;
		} catch (err) {
			error = err instanceof Error ? err.message : 'Unknown error';
		} finally {
			loading = false;
		}
	}

	function getSortedPlants(): Plant[] {
		const sorted = [...plants];
		switch (sortBy) {
			case 'name':
				return sorted.sort((a, b) => a.name.localeCompare(b.name));
			case 'lastWatered':
				return sorted.sort((a, b) => {
					const aw = a.watering?.lastWatered ? new Date(a.watering.lastWatered).getTime() : 0;
					const bw = b.watering?.lastWatered ? new Date(b.watering.lastWatered).getTime() : 0;
					return bw - aw;
				});
			case 'lastFertilized':
				return sorted.sort((a, b) => {
					const af = a.fertilizing?.lastFertilized
						? new Date(a.fertilizing.lastFertilized).getTime()
						: 0;
					const bf = b.fertilizing?.lastFertilized
						? new Date(b.fertilizing.lastFertilized).getTime()
						: 0;
					return bf - af;
				});
			case 'wateringIntervalDays':
				return sorted.sort(
					(a, b) => (a.watering?.intervalDays ?? 999) - (b.watering?.intervalDays ?? 999)
				);
			case 'mistingIntervalDays':
				return sorted.sort(
					(a, b) =>
						(a.humidity?.mistingIntervalDays ?? 999) - (b.humidity?.mistingIntervalDays ?? 999)
				);
			default:
				return sorted;
		}
	}

	function daysAgo(dateString: string): string {
		const days = Math.floor((Date.now() - new Date(dateString).getTime()) / (1000 * 60 * 60 * 24));
		if (days === 0) return 'Today';
		if (days === 1) return 'Yesterday';
		return `${days} days ago`;
	}

	function getWateringStatus(plant: Plant): { text: string; color: string } {
		const last = plant.watering?.lastWatered
			? new Date(plant.watering.lastWatered).getTime()
			: Date.now();
		const interval = plant.watering?.intervalDays ?? 0;
		const days = Math.floor((Date.now() - last) / (1000 * 60 * 60 * 24));
		const daysUntilWater = interval - days;
		if (daysUntilWater <= 0) return { text: '🌵 Needs water!', color: 'text-red-600' };
		if (daysUntilWater <= 1) return { text: '⚠️ Water soon', color: 'text-yellow-600' };
		return { text: `✓ In ${daysUntilWater} days`, color: 'text-green-600' };
	}

	onMount(() => {
		loadPlants();
	});
</script>

<div class="min-h-screen bg-gradient-to-br from-green-50 via-emerald-50 to-teal-50 p-8">
	<div class="mx-auto max-w-7xl">
		<!-- Header -->
		<div class="mb-12">
			<div class="mb-4 flex items-center justify-between">
				<div>
					<h1 class="mb-2 flex items-center gap-3 text-5xl font-bold text-green-800">
						{$tStore('common.app')}
					</h1>
					<p class="text-lg text-green-700">{$tStore('common.appDescription')}</p>
				</div>
			</div>
			<div class="flex gap-3">
				<button
					on:click={() => goto(resolve('/manage'))}
					class="rounded-lg bg-green-600 px-5 py-2 text-white transition hover:bg-green-700 focus:ring-2 focus:ring-green-500 focus:outline-none"
				>
					{$tStore('menu.managePlants')}
				</button>
			</div>
		</div>

		<!-- Controls -->
		<div class="mb-8 flex items-center justify-between">
			<SortControls bind:sortBy />
			<div class="font-medium text-green-800">
				{plants.length}{plants.length === 1 ? ' plant' : ' plants'}
			</div>
		</div>

		<!-- Loading & Error States -->
		{#if loading}
			<div class="flex min-h-96 items-center justify-center">
				<div class="text-center">
					<div class="mb-4 animate-bounce text-6xl">🌿</div>
					<p class="text-lg font-medium text-green-700">Loading your plants...</p>
				</div>
			</div>
		{:else if error}
			<div class="rounded-lg border-2 border-red-400 bg-red-100 px-6 py-4 text-red-800">
				<p class="font-bold">Error loading plants</p>
				<p>{error}</p>
			</div>
		{:else if plants.length === 0}
			<div class="py-16 text-center">
				<div class="mb-4 text-8xl">🪴</div>
				<p class="text-xl font-medium text-green-800">{$tStore('plants.noPlants')}</p>
				<p class="mt-2 text-green-700">{$tStore('plants.startAddingPlants')}</p>
			</div>
		{:else}
			<!-- Plant Grid -->
			<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
				{#each getSortedPlants() as plant (plant.id)}
					<PlantCard {plant} {daysAgo} {getWateringStatus} />
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
