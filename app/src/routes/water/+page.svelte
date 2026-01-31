<script lang="ts">
	import type { Plant } from '$lib/types/api';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { fetchData } from '$lib/auth/fetch.svelte';
	import WaterPlantCard from '$lib/components/WaterPlantCard.svelte';
	import { onMount } from 'svelte';

	let plants = $state<Plant[]>([]);
	let loading = $state(true);
	let error = $state<string | null>(null);
	let watering = $state(false);
	let selectedIds = $state<string[]>([]);

	async function loadPlants(): Promise<void> {
		loading = true;
		try {
			const response = await fetchData('/api/plants', {});
			if (!response.ok) {
				error = response.error?.message || 'Failed to load plants';
				loading = false;
				return;
			}
			plants = response.data || [];
		} catch (err) {
			error = err instanceof Error ? err.message : 'Failed to load plants';
		} finally {
			loading = false;
		}
	}

	function getDaysUntilWater(plant: Plant): number {
		const lastWatered = plant.watering?.lastWatered
			? new Date(plant.watering.lastWatered).getTime()
			: 0;
		const interval = plant.watering?.intervalDays ?? 0;
		const daysSinceWatered = Math.floor((Date.now() - lastWatered) / (1000 * 60 * 60 * 24));
		return interval - daysSinceWatered;
	}

	function getPlantStatus(plant: Plant): 'overdue' | 'due-soon' | 'ok' {
		const daysUntil = getDaysUntilWater(plant);
		if (daysUntil <= 0) return 'overdue';
		if (daysUntil <= 1) return 'due-soon';
		return 'ok';
	}

	function sortPlantsByWateringPriority(): Plant[] {
		return [...plants].sort((a, b) => {
			const aDays = getDaysUntilWater(a);
			const bDays = getDaysUntilWater(b);
			return aDays - bDays;
		});
	}

	function togglePlant(id: string): void {
		if (selectedIds.includes(id)) {
			selectedIds = selectedIds.filter((sid) => sid !== id);
		} else {
			selectedIds = [...selectedIds, id];
		}
	}

	function selectAll(): void {
		selectedIds = plants.map((p) => p.id);
	}

	function selectDueToday(): void {
		const due = plants.filter((p) => getDaysUntilWater(p) <= 0);
		selectedIds = due.map((p) => p.id);
	}

	function clearSelection(): void {
		selectedIds = [];
	}

	async function waterSelectedPlants(): Promise<void> {
		if (selectedIds.length === 0) return;

		watering = true;
		error = null;

		try {
			const response = await fetchData('/api/plants/water', {
				method: 'post',
				body: {
					plantIds: selectedIds
				}
			});

			if (!response.ok) {
				error = response.error?.message || 'Failed to water plants';
				return;
			}

			// Update the plants locally
			const now = new Date().toISOString();
			plants = plants.map((p) => {
				if (selectedIds.includes(p.id) && p.watering) {
					return {
						...p,
						watering: {
							...p.watering,
							lastWatered: now
						}
					};
				}
				return p;
			});

			// Clear selection
			selectedIds = [];
		} catch (err) {
			error = err instanceof Error ? err.message : 'Failed to water plants';
		} finally {
			watering = false;
		}
	}

	function getStatusIcon(status: 'overdue' | 'due-soon' | 'ok'): string {
		switch (status) {
			case 'overdue':
				return '🚨';
			case 'due-soon':
				return '⚠️';
			default:
				return '✅';
		}
	}

	function getStatusText(plant: Plant): string {
		const daysUntil = getDaysUntilWater(plant);
		if (daysUntil < 0)
			return `${Math.abs(daysUntil)} ${Math.abs(daysUntil) === 1 ? 'day' : 'days'} overdue`;
		if (daysUntil === 0) return 'Due today';
		if (daysUntil === 1) return 'Due tomorrow';
		return `Due in ${daysUntil} days`;
	}

	onMount(() => {
		loadPlants();
	});
</script>

<div class="min-h-screen bg-gradient-to-br from-emerald-50 via-green-50 to-teal-100 p-6 pb-24 md:p-10">
	<div class="mx-auto max-w-6xl">
		<!-- Header -->
		<div class="mb-8">
			<h1 class="mb-1 text-4xl font-bold text-green-900">💧 Water Plants</h1>
			<p class="text-gray-600">Quick watering view</p>
		</div>

		<!-- Error Message -->
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
				<p class="mb-6 text-gray-600">Add plants to start tracking watering schedules</p>
				<button
					onclick={() => goto(resolve('/manage/create'))}
					class="rounded-lg bg-gradient-to-r from-emerald-600 to-green-600 px-6 py-3 font-semibold text-white shadow-md transition hover:from-emerald-700 hover:to-green-700"
				>
					Add Your First Plant
				</button>
			</div>
		{:else}
			<!-- Quick Actions Bar -->
			<div
				class="mb-6 flex flex-wrap gap-2 rounded-lg border border-emerald-100 bg-white/90 p-4 shadow-md backdrop-blur"
			>
				<button
					onclick={selectDueToday}
					class="flex-1 rounded-lg bg-red-100 px-4 py-2 text-sm font-medium text-red-800 transition hover:bg-red-200"
				>
					Select Overdue
				</button>
				<button
					onclick={selectAll}
					class="flex-1 rounded-lg bg-emerald-100 px-4 py-2 text-sm font-medium text-emerald-800 transition hover:bg-emerald-200"
				>
					Select All
				</button>
				<button
					onclick={clearSelection}
					class="flex-1 rounded-lg bg-gray-100 px-4 py-2 text-sm font-medium text-gray-800 transition hover:bg-gray-200"
				>
					Clear
				</button>
			</div>

			<!-- Plant List -->
			<div class="space-y-4">
				{#each sortPlantsByWateringPriority() as plant (plant.id)}
					<WaterPlantCard
						{plant}
						isSelected={selectedIds.includes(plant.id)}
						status={getPlantStatus(plant)}
						statusText={getStatusText(plant)}
						statusIcon={getStatusIcon(getPlantStatus(plant))}
						onToggle={togglePlant}
					/>
				{/each}
			</div>
		{/if}
	</div>

	<!-- Bottom Action Bar (Fixed) -->
	{#if selectedIds.length > 0}
		<div
			class="fixed bottom-0 left-0 right-0 border-t border-emerald-200 bg-white/95 px-4 py-4 shadow-xl backdrop-blur"
		>
			<div class="mx-auto flex max-w-6xl items-center justify-between gap-4">
				<div class="text-sm font-medium text-gray-700">
					{selectedIds.length}
					{selectedIds.length === 1 ? 'plant' : 'plants'} selected
				</div>
				<button
					onclick={waterSelectedPlants}
					disabled={watering}
					class="rounded-lg bg-gradient-to-r from-emerald-600 to-green-600 px-8 py-3 font-bold text-white shadow-md transition hover:from-emerald-700 hover:to-green-700 disabled:opacity-50"
				>
					{watering ? '💧 Watering...' : '💧 Water Selected'}
				</button>
			</div>
		</div>
	{/if}
</div>

<style>
	:global(body) {
		font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
	}
</style>
