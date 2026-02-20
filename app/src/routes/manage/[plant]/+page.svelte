<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import { fetchData } from '$lib/auth/fetch.svelte';
	import { tStore } from '$lib/i18n';
	import PageHeader from '$lib/components/layout/PageHeader.svelte';
	import PageContent from '$lib/components/layout/PageContent.svelte';
	import Card from '$lib/components/ui/Card.svelte';
	import LoadingSpinner from '$lib/components/ui/LoadingSpinner.svelte';
	import Alert from '$lib/components/ui/Message.svelte';
	import type { Plant } from '$lib/types/api';

	let plant = $state<Plant | null>(null);
	let loading = $state(true);
	let error = $state<string | null>(null);

	const sectionItems = [
		{ key: 'basic', emoji: '📋', label: 'plants.basicInformation' },
		{ key: 'photos', emoji: '📸', label: 'plants.photos' },
		{ key: 'location', emoji: '📍', label: 'plants.location' },
		{ key: 'watering', emoji: '💧', label: 'plants.wateringTitle' },
		{ key: 'fertilizing', emoji: '🍯', label: 'plants.fertilizingTitle' },
		{ key: 'humidity', emoji: '💨', label: 'plants.humidityTitle' },
		{ key: 'soil', emoji: '🌍', label: 'plants.soilTitle' },
		{ key: 'seasonality', emoji: '❄️', label: 'plants.seasonalityTitle' },
		{ key: 'metadata', emoji: '🏷️', label: 'plants.metadata' }
	] as const;

	const showCreateNextStep = $derived(page.url.searchParams.get('createFlow') === '1');

	onMount(async () => {
		try {
			const plantId = page.params.plant ?? '';
			const res = await fetchData('/api/plants/{id}', {
				params: { id: plantId }
			});

			if (!res.ok) {
				error = res.error?.message || $tStore('plants.failedToFetchPlants');
				return;
			}

			plant = res.data;
		} catch (err) {
			error = err instanceof Error ? err.message : $tStore('plants.failedToFetchPlants');
		} finally {
			loading = false;
		}
	});

	function goToSection(section: string): void {
		if (!plant) return;
		const flowQuery = showCreateNextStep ? '?createFlow=1' : '';
		if (section === 'photos') {
			goto(resolve(`/manage/${plant.id}/photos${flowQuery}`));
			return;
		}
		goto(resolve(`/manage/${plant.id}/edit/${section}${flowQuery}`));
	}
</script>

<PageHeader
	icon="🧭"
	title={plant?.name || $tStore('plants.editPlant')}
	description={plant?.species || $tStore('plants.manageHubDescription')}
/>

<PageContent>
	{#if loading}
		<LoadingSpinner message="common.loadingPlantDetails" icon="🌱" />
	{:else if !plant}
		<Alert
			type="error"
			title="common.error"
			description={error || $tStore('common.plantNotFound')}
		/>
	{:else}
		<div class="space-y-4">
			<Card rounded="2xl">
				<div class="space-y-3 p-4">
					<p class="text-base font-semibold text-[var(--text-light-main)]">
						{$tStore('plants.manageOptionalSections')}
					</p>
					{#each sectionItems as item (item.key)}
						<button
							onclick={() => goToSection(item.key)}
							class="flex min-h-12 w-full cursor-pointer items-center justify-between rounded-xl border border-[var(--p-emerald)]/25 bg-white px-4 py-3 text-left text-base font-medium text-[var(--text-light-main)] hover:bg-[var(--bg-light)]"
						>
							<span>{item.emoji} {$tStore(item.label)}</span>
							<span aria-hidden="true">›</span>
						</button>
					{/each}
				</div>
			</Card>
		</div>
	{/if}
</PageContent>
