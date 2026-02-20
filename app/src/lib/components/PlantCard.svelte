<script lang="ts">
	import { tStore } from '$lib/i18n';
	import type { Plant } from '$lib/types/api';
	import { imageCacheStore } from '$lib/stores/imageCache.svelte';
	import { onDestroy } from 'svelte';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { formatPastTimestamp } from '$lib/utils/timestamp.svelte';
	import { getWateringStatus } from '$lib/utils/plant';

	interface Props {
		plant: Plant;
	}

	const { plant }: Props = $props();

	const firstId = $derived(plant.photoIds?.[0]);
	// Get the URL from cache once (already preloaded in Auth)
	const previewUrl = $derived(firstId ? imageCacheStore.getImageURLSync(firstId) : null);

	const wateringStatus = $derived(getWateringStatus(plant));

	function openPlant() {
		goto(resolve(`/plant/${plant.id}`));
	}

	function onKeydown(event: KeyboardEvent) {
		if (event.key === 'Enter' || event.key === ' ') {
			event.preventDefault();
			openPlant();
		}
	}

	onDestroy(() => {
		if (firstId) {
			imageCacheStore.releaseImage(firstId);
		}
	});
</script>

<div
	role="button"
	tabindex="0"
	onclick={openPlant}
	onkeydown={onKeydown}
	class="group cursor-pointer touch-manipulation overflow-hidden rounded-2xl bg-white shadow transition-all duration-200 active:scale-[0.98] active:shadow-sm"
>
	<!-- Image -->
	<div
		class="relative flex h-40 items-center justify-center overflow-hidden bg-gradient-to-br from-[var(--p-emerald)] to-[var(--p-emerald-dark)]"
	>
		{#if previewUrl}
			<img
				src={previewUrl}
				alt={plant.name}
				class="h-full w-full object-cover transition-transform duration-300 group-active:scale-105"
			/>
		{:else}
			<div class="text-6xl">🌱</div>
		{/if}
	</div>

	<!-- Content -->
	<div class="p-3.5">
		<!-- Name and Species -->
		<h3 class="line-clamp-1 text-base font-bold text-[var(--text-light-main)]">{plant.name}</h3>
		<p class="mb-2 line-clamp-1 text-xs text-[var(--status-success)]">{plant.species}</p>

		<!-- Watering Status -->
		{#if wateringStatus}
			<div class="mb-3">
				<div
					class={`mb-1.5 flex items-center gap-1.5 text-xs font-semibold ${wateringStatus.color}`}
				>
					<span>{wateringStatus.emoji}</span>
					<span>{$tStore(wateringStatus.text, wateringStatus.args)}</span>
				</div>

				{#if plant.watering?.lastWatered}
					<p class="text-xs text-[var(--text-light-main)]/60">
						{$tStore('plants.lastWatered')}: {formatPastTimestamp(
							new Date(plant.watering?.lastWatered ?? '')
						)}
					</p>
				{/if}
			</div>
		{/if}

		<!-- Metadata Grid -->
		<div class="grid grid-cols-2 gap-2 text-xs">
			{#if plant.watering?.intervalDays}
				<div class="rounded-lg bg-[var(--p-emerald)]/20 p-1.5">
					<div class="font-semibold text-[var(--p-emerald-dark)]">💧</div>
					<p class="mt-0.5 text-xs text-[var(--text-light-main)]/80">
						{$tStore('plants.every')}
						{plant.watering?.intervalDays}
						{$tStore('plants.days')}
					</p>
				</div>
			{/if}
			{#if plant.fertilizing?.intervalDays}
				<div class="rounded-lg bg-[var(--status-warn)]/20 p-1.5">
					<div class="font-semibold text-[var(--status-warn)]">🥗</div>
					<p class="mt-0.5 text-xs text-[var(--text-light-main)]/80">
						{$tStore('plants.every')}
						{plant.fertilizing?.intervalDays}
						{$tStore('plants.days')}
					</p>
				</div>
			{/if}
			{#if plant.sunlight}
				<div class="rounded-lg bg-[var(--status-info)]/20 p-1.5">
					<div class="font-semibold text-[var(--status-info)]">☀️</div>
					<p class="mt-0.5 text-xs text-[var(--text-light-main)]/80">
						{$tStore('plants.sunlight.' + plant.sunlight)}
					</p>
				</div>
			{/if}
			{#if plant.humidity?.targetHumidityPct}
				<div class="rounded-lg bg-[var(--p-emerald)]/10 p-1.5">
					<p class="text-xs font-semibold text-[var(--text-light-main)]/80">
						{plant.humidity?.targetHumidityPct}%
					</p>
				</div>
			{/if}
		</div>

		<!-- Spray Info -->
		{#if plant.humidity?.requiresMisting && plant.humidity?.mistingIntervalDays}
			<div class="mt-2 rounded-lg bg-[var(--status-info)]/20 p-1.5">
				<p class="text-xs text-[var(--text-light-main)]/70">
					💦 {$tStore('plants.sprayEvery')}
					<span class="font-semibold text-[var(--status-info)]"
						>{plant.humidity?.mistingIntervalDays}</span
					>
					{$tStore('plants.days')}
				</p>
				<p class="mt-0.5 text-xs text-[var(--text-light-main)]/70">
					{$tStore('plants.lastSprayedStatus')}:
					<span class="font-semibold"
						>{formatPastTimestamp(new Date(plant.humidity?.lastMisted ?? ''))}</span
					>
				</p>
			</div>
		{/if}

		<!-- Flags -->
		{#if plant.flags && plant.flags.length > 0}
			<div class="mt-2 flex flex-wrap gap-1.5">
				{#each plant.flags as flag (flag)}
					<span
						class="rounded-full bg-[var(--status-warn)]/30 px-2 py-0.5 text-xs font-medium text-[var(--status-warn)]"
						>⚡{flag}</span
					>
				{/each}
			</div>
		{/if}

		<!-- Notes Preview -->
		{#if plant.notes && plant.notes.length > 0}
			<div class="mt-2 border-t border-[var(--p-emerald)]/20 pt-2">
				<p class="line-clamp-2 text-xs text-[var(--text-light-main)]/60">📝 {plant.notes[0]}</p>
			</div>
		{/if}
	</div>
</div>
