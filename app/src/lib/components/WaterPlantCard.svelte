<script lang="ts">
	import type { Plant } from '$lib/types/api';
	import { getImageObjectURL, revokeObjectURL } from '$lib/utils/imageCache';
	import { onMount, onDestroy } from 'svelte';

	interface Props {
		plant: Plant;
		isSelected: boolean;
		status: 'overdue' | 'due-soon' | 'ok';
		statusText: string;
		statusIcon: string;
		onToggle: (id: string) => void;
	}

	let { plant, isSelected, status, statusText, statusIcon, onToggle }: Props = $props();

	let previewUrl = $state<string | null>(null);

	function getStatusColor(status: 'overdue' | 'due-soon' | 'ok'): string {
		switch (status) {
			case 'overdue':
				return 'bg-red-100 border-red-300';
			case 'due-soon':
				return 'bg-yellow-100 border-yellow-300';
			default:
				return 'bg-green-100 border-green-300';
		}
	}

	onMount(async () => {
		const firstId = plant.photoIds?.[0];
		const firstUrl = (plant as any)?.photoUrls?.[0] as string | undefined;
		if (firstId && firstUrl) {
			previewUrl = await getImageObjectURL(firstId, firstUrl);
		}
	});

	onDestroy(() => {
		if (previewUrl) revokeObjectURL(previewUrl);
	});
</script>

<button
	onclick={() => onToggle(plant.id)}
	class="w-full rounded-2xl border-2 bg-white/90 p-4 shadow-md backdrop-blur transition {isSelected
		? 'border-emerald-500 bg-emerald-50'
		: `border-emerald-100 ${getStatusColor(status)}`}"
>
	<div class="flex items-center gap-4">
		<!-- Checkbox -->
		<div
			class="flex h-10 w-10 flex-shrink-0 items-center justify-center rounded-lg border-2 transition {isSelected
				? 'border-emerald-600 bg-emerald-600 text-white'
				: 'border-gray-300 bg-white text-gray-400'}"
		>
			{#if isSelected}
				<span class="text-xl">✓</span>
			{:else}
				<span class="text-xl">○</span>
			{/if}
		</div>

		<!-- Photo -->
		{#if previewUrl}
			<img
				src={previewUrl}
				alt={plant.name}
				class="h-16 w-16 flex-shrink-0 rounded-lg object-cover"
			/>
		{:else}
			<div
				class="flex h-16 w-16 flex-shrink-0 items-center justify-center rounded-lg bg-green-200 text-2xl"
			>
				🌿
			</div>
		{/if}

		<!-- Plant Info -->
		<div class="flex-1 text-left">
			<div class="mb-1 flex items-center gap-2">
				<span class="text-xl">{statusIcon}</span>
				<h3 class="text-lg font-bold text-gray-900">{plant.name}</h3>
			</div>
			<p class="mb-1 text-sm italic text-gray-600">{plant.species}</p>
			<p class="text-sm font-medium text-gray-700">{statusText}</p>
		</div>
	</div>
</button>
