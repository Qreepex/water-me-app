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

<!-- Mobile-first search and controls -->
<div class="mb-3 flex flex-shrink-0 items-center gap-2 px-2">
	<div class="min-w-0 flex-1">
		<SearchInput
			bind:value={searchQuery}
			placeholder={$tStore('plants.searchPlants') ?? 'Search plants...'}
			ariaLabel={$tStore('plants.searchPlants') ?? 'Search plants'}
		/>
	</div>

	<div
		class="flex h-11 min-w-[2.5rem] flex-shrink-0 items-center justify-center rounded-full bg-[var(--p-emerald)]/10 px-2.5 text-xs font-semibold text-[var(--text-light-main)]"
	>
		{filteredPlants.length}
	</div>

	<SortControls {sortBy} {onSortChange} compact={true} iconOnly={true} />
</div>

<Scrollable multi>
	{#each filteredPlants as plant (plant.id)}
		<PlantCard {plant} />
	{/each}
</Scrollable>
