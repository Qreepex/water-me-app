<script lang="ts">
	import type { Plant } from '$lib/types/api';
	import { sortPlants, daysAgo, getWateringStatus, type SortOption } from '$lib/utils/plant';
	import PlantCard from './PlantCard.svelte';
	import { tStore } from '$lib/i18n';

	interface Props {
		plants: Plant[];
		sortBy: SortOption;
		onSortChange: (sort: SortOption) => void;
	}

	const { plants, sortBy, onSortChange }: Props = $props();

	let showSortMenu = $state(false);
	let searchQuery = $state('');

	const sortOptions: { value: SortOption; label: string; icon: string }[] = [
		{ value: 'name', label: 'plants.sortByName', icon: '🔤' },
		{ value: 'wateringDate', label: 'plants.sortByWatering', icon: '💧' },
		{ value: 'species', label: 'plants.sortBySpecies', icon: '🌿' }
	];

	const filteredPlants = $derived.by(() => {
		let filtered = sortPlants(plants, sortBy);

		if (searchQuery.trim()) {
			const query = searchQuery.toLowerCase().trim();
			filtered = filtered.filter((plant) => {
				const name = plant.name?.toLowerCase() ?? '';
				const species = plant.species?.toLowerCase() ?? '';
				const room = plant.location?.room?.toLowerCase() ?? '';
				const position = plant.location?.position?.toLowerCase() ?? '';

				return (
					name.includes(query) ||
					species.includes(query) ||
					room.includes(query) ||
					position.includes(query)
				);
			});
		}

		return filtered;
	});

	function selectSort(option: SortOption) {
		onSortChange(option);
		showSortMenu = false;
	}

	function getCurrentSortLabel(): string {
		const option = sortOptions.find((o) => o.value === sortBy);
		return option ? option.icon : '🔤';
	}
</script>

<div class="flex h-full min-h-0 flex-col">
	<!-- Header Bar with Search and Sort Button -->
	<div class="mb-4 flex flex-shrink-0 flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
		<!-- Search Input -->
		<div class="relative flex-1">
			<input
				type="text"
				placeholder={$tStore('plants.searchPlants') ?? 'Search plants...'}
				bind:value={searchQuery}
				class="w-full rounded-lg border border-[var(--p-emerald)]/30 bg-white py-2 pl-10 pr-4 text-sm placeholder-[var(--text-light-main)]/40 transition-all focus:border-[var(--p-emerald)]/60 focus:outline-none focus:ring-2 focus:ring-[var(--p-emerald)]/20"
			/>
			<svg
				class="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-[var(--p-emerald)]/50"
				fill="none"
				stroke="currentColor"
				viewBox="0 0 24 24"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
				/>
			</svg>
		</div>

		<!-- Count and Sort Button -->
		<div class="flex items-center gap-3">
			<div class="text-sm font-medium text-[var(--text-light-main)]">
				{filteredPlants.length}{filteredPlants.length === 1
					? ' ' + $tStore('common.plant')
					: ' ' + $tStore('common.plants')}
			</div>

			<!-- Sort Button -->
			<div class="relative">
				<button
					onclick={() => (showSortMenu = !showSortMenu)}
					class="flex items-center gap-2 rounded-lg bg-[var(--p-emerald)]/10 px-4 py-2 text-sm font-medium text-[var(--p-emerald-dark)] transition-colors hover:bg-[var(--p-emerald)]/20"
				>
					<span class="text-lg">{getCurrentSortLabel()}</span>
					<span class="hidden sm:inline">{$tStore('plants.sort')}</span>
					<svg
						class="h-4 w-4 transition-transform {showSortMenu ? 'rotate-180' : ''}"
						fill="none"
						stroke="currentColor"
						viewBox="0 0 24 24"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M19 9l-7 7-7-7"
						/>
					</svg>
				</button>

				<!-- Sort Menu -->
				{#if showSortMenu}
					<div
						class="absolute top-full right-0 z-50 mt-2 w-56 rounded-lg border border-[var(--p-emerald)]/30 bg-white shadow-lg"
					>
						{#each sortOptions as option}
							<button
								onclick={() => selectSort(option.value)}
								class="flex w-full items-center gap-3 px-4 py-3 text-left text-sm transition-colors hover:bg-[var(--p-emerald)]/10 {sortBy ===
								option.value
									? 'bg-[var(--p-emerald)]/20 font-semibold text-[var(--p-emerald-dark)]'
									: 'text-[var(--text-light-main)]'}"
							>
								<span class="text-lg">{option.icon}</span>
								<span>{$tStore(option.label)}</span>
								{#if sortBy === option.value}
									<svg class="ml-auto h-5 w-5" fill="currentColor" viewBox="0 0 20 20">
										<path
											fill-rule="evenodd"
											d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
											clip-rule="evenodd"
										/>
									</svg>
								{/if}
							</button>
						{/each}
					</div>
				{/if}
			</div>
		</div>
	</div>

	<!-- Scrollable Plant Grid -->
	<div class="min-h-0 flex-1 overflow-y-auto pb-24">
		<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
			{#each filteredPlants as plant (plant.id)}
				<PlantCard {plant} {daysAgo} {getWateringStatus} />
			{/each}
		</div>
	</div>
</div>

<!-- Click outside to close sort menu -->
{#if showSortMenu}
	<button class="fixed inset-0 z-40" onclick={() => (showSortMenu = false)} aria-label="Close menu"
	></button>
{/if}
