<script lang="ts">
	import type { Plant } from '$lib/types/api';
	import { sortPlants, type SortOption } from '$lib/utils/plant';
	import PlantCard from './PlantCard.svelte';
	import SortControls from './SortControls.svelte';
	import SearchInput from './SearchInput.svelte';
	import { tStore } from '$lib/i18n';
	import Scrollable from './layout/Scrollable.svelte';

	interface Props {
		plants: Plant[];
		sortBy: SortOption;
		onSortChange: (sort: SortOption) => void;
	}

	const { plants, sortBy, onSortChange }: Props = $props();

	let searchQuery = $state('');

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
</script>

<!-- Header Bar with Search and Sort Button -->
<div
	class="mb-4 flex flex-shrink-0 flex-col gap-3 rounded-2xl border border-[var(--p-emerald)]/20 bg-white/80 p-3 shadow-sm backdrop-blur sm:flex-row sm:items-center sm:justify-between"
>
	<!-- Search Input -->
	<SearchInput
		bind:value={searchQuery}
		placeholder={$tStore('plants.searchPlants') ?? 'Search plants...'}
		ariaLabel={$tStore('plants.searchPlants') ?? 'Search plants'}
	/>

	<!-- Count and Sort Button -->
	<div class="flex items-center justify-between gap-3 sm:justify-start">
		<div
			class="rounded-full bg-[var(--p-emerald)]/10 px-2.5 py-1 text-xs font-semibold text-[var(--text-light-main)]"
		>
			{filteredPlants.length}{filteredPlants.length === 1
				? ' ' + $tStore('common.plant')
				: ' ' + $tStore('common.plants')}
		</div>

		<SortControls {sortBy} {onSortChange} />
	</div>
</div>

<Scrollable multi>
	{#each filteredPlants as plant (plant.id)}
		<PlantCard {plant} />
	{/each}
</Scrollable>
