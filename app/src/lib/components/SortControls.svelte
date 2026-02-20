<script lang="ts">
	import { tStore } from '$lib/i18n';
	import type { SortOption } from '$lib/utils/plant';

	interface Props {
		sortBy?: SortOption;
		onSortChange?: (value: SortOption) => void;
		compact?: boolean;
		iconOnly?: boolean;
	}

	let { sortBy = 'nameAsc', onSortChange, compact = false, iconOnly = false }: Props = $props();

	const sortOptions: { value: SortOption; label: string; icon: string }[] = [
		{ value: 'nameAsc', label: 'plants.sortOptions.nameAsc', icon: '🔤' },
		{ value: 'nameDesc', label: 'plants.sortOptions.nameDesc', icon: '🔤' },
		{ value: 'lastWateredAsc', label: 'plants.sortOptions.lastWateredAsc', icon: '💧' },
		{ value: 'lastWateredDesc', label: 'plants.sortOptions.lastWateredDesc', icon: '💧' },
		{ value: 'speciesAsc', label: 'plants.sortOptions.speciesAsc', icon: '🌿' },
		{ value: 'speciesDesc', label: 'plants.sortOptions.speciesDesc', icon: '🌿' }
	];

	function handleChange(e: Event) {
		const value = (e.currentTarget as HTMLSelectElement).value as SortOption;
		sortBy = value;
		onSortChange?.(value);
	}
</script>
{#if iconOnly}
	<div class="relative h-11 w-11 flex-shrink-0">
		<div
			class="pointer-events-none absolute inset-0 flex items-center justify-center rounded-full border-2 border-[var(--p-emerald)] bg-[var(--card-light)] text-[var(--text-light-main)] shadow-sm"
		>
			<svg class="h-4.5 w-4.5" viewBox="0 0 24 24" fill="none" stroke="currentColor">
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M3 5h18M6 12h12m-9 7h6"
				/>
			</svg>
		</div>
		<select
			id="sort"
			value={sortBy}
			onchange={handleChange}
			aria-label={$tStore('plants.sortBy') ?? 'Sort'}
			class="absolute inset-0 h-full w-full cursor-pointer rounded-full opacity-0"
		>
			{#each sortOptions as option (option.value)}
				<option value={option.value}>{option.icon} {$tStore(option.label)}</option>
			{/each}
		</select>
	</div>
{:else}
	<div class="flex w-full flex-col gap-2 sm:w-auto sm:flex-row sm:items-center sm:gap-3">
		<select
			id="sort"
			value={sortBy}
			onchange={handleChange}
			class={`rounded-full border-2 border-[var(--p-emerald)] bg-[var(--card-light)] text-base font-medium text-[var(--text-light-main)] shadow-sm transition hover:border-[var(--p-emerald-dark)] focus:border-[var(--p-emerald)] focus:outline-none ${
				compact
					? 'h-11 flex-shrink-0 px-3 text-xs'
					: 'w-full px-4 py-3 sm:w-auto'
			}`}
		>
			{#each sortOptions as option (option.value)}
				<option value={option.value}>{option.icon} {$tStore(option.label)}</option>
			{/each}
		</select>
	</div>
{/if}
