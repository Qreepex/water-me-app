<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import { fetchData } from '$lib/auth/fetch.svelte';
	import { invalidateApiCache } from '$lib/utils/cache';
	import { tStore } from '$lib/i18n';
	import type { Plant } from '$lib/types/api';
	import { SunlightRequirement, WateringMethod, WaterType, FertilizerType } from '$lib/types/api';
	import type { FormData } from '$lib/types/forms';
	import { createEmptyFormData } from '$lib/types/forms';
	import PageHeader from '$lib/components/layout/PageHeader.svelte';
	import PageContent from '$lib/components/layout/PageContent.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import LoadingSpinner from '$lib/components/ui/LoadingSpinner.svelte';
	import Alert from '$lib/components/ui/Message.svelte';
	import BasicInformationForm from '$lib/components/PlantForms/BasicInformationForm.svelte';
	import LocationForm from '$lib/components/PlantForms/LocationForm.svelte';
	import WateringForm from '$lib/components/PlantForms/WateringForm.svelte';
	import FertilizingForm from '$lib/components/PlantForms/FertilizingForm.svelte';
	import MistingForm from '$lib/components/PlantForms/MistingForm.svelte';
	import SoilForm from '$lib/components/PlantForms/SoilForm.svelte';
	import SeasonalityForm from '$lib/components/PlantForms/SeasonalityForm.svelte';
	import MetadataForm from '$lib/components/PlantForms/MetadataForm.svelte';

	type EditSection =
		| 'basic'
		| 'location'
		| 'watering'
		| 'fertilizing'
		| 'humidity'
		| 'soil'
		| 'seasonality'
		| 'metadata';

	const sectionTitles: Record<EditSection, string> = {
		basic: 'plants.basicInformation',
		location: 'plants.location',
		watering: 'plants.wateringTitle',
		fertilizing: 'plants.fertilizingTitle',
		humidity: 'plants.humidityTitle',
		soil: 'plants.soilTitle',
		seasonality: 'plants.seasonalityTitle',
		metadata: 'plants.metadata'
	};

	let plant = $state<Plant | null>(null);
	let loading = $state(true);
	let saving = $state(false);
	let error = $state<string | null>(null);
	let formData = $state<FormData>(createEmptyFormData());
	let originalFormData = $state<FormData>(createEmptyFormData());
	let newNote = $state('');
	let soilComponentInput = $state('');

	const section = $derived((page.params.section ?? 'basic') as EditSection);
	const isCreateFlow = $derived(page.url.searchParams.get('createFlow') === '1');
	const isDirty = $derived(JSON.stringify(formData) !== JSON.stringify(originalFormData));

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
			formData = initializeFormData();
			originalFormData = JSON.parse(JSON.stringify(formData));
		} catch (err) {
			error = err instanceof Error ? err.message : $tStore('plants.failedToFetchPlants');
		} finally {
			loading = false;
		}
	});

	function initializeFormData(): FormData {
		if (!plant) return createEmptyFormData();
		return {
			name: plant.name,
			species: plant.species,
			isToxic: plant.isToxic,
			sunlight: plant.sunlight as SunlightRequirement,
			preferedTemperature: plant.preferedTemperature,
			room: plant.location?.room ?? '',
			position: plant.location?.position ?? '',
			isOutdoors: plant.location?.isOutdoors ?? false,
			wateringIntervalDays: plant.watering?.intervalDays ?? 7,
			wateringMethod: plant.watering?.method ?? WateringMethod.Top,
			waterType: plant.watering?.waterType ?? WaterType.Tap,
			fertilizingType: plant.fertilizing?.type ?? FertilizerType.Liquid,
			fertilizingIntervalDays: plant.fertilizing?.intervalDays ?? 30,
			npkRatio: plant.fertilizing?.npkRatio ?? '10:10:10',
			concentrationPercent: plant.fertilizing?.concentrationPercent ?? 50,
			activeInWinter: plant.fertilizing?.activeInWinter ?? false,
			targetHumidity: plant.humidity?.targetHumidityPct ?? 50,
			requiresMisting: plant.humidity?.requiresMisting ?? false,
			mistingIntervalDays: plant.humidity?.mistingIntervalDays ?? 3,
			requiresHumidifier: plant.humidity?.requiresHumidifier ?? false,
			soilType: plant.soil?.type ?? 'Generic',
			repottingCycle: plant.soil?.repottingCycle ?? 2,
			soilComponents: plant.soil?.components ?? [],
			winterRestPeriod: plant.seasonality?.winterRestPeriod ?? false,
			winterWaterFactor: plant.seasonality?.winterWaterFactor ?? 0.5,
			minTempCelsius: plant.seasonality?.minTempCelsius ?? 15,
			flags: plant.flags ?? [],
			notes: plant.notes ?? []
		};
	}

	function buildSectionPayload(): Record<string, unknown> {
		switch (section) {
			case 'basic':
				return {
					name: formData.name,
					species: formData.species
				};
			case 'location':
				return {
					sunlight: formData.sunlight,
					location: {
						room: formData.room,
						position: formData.position,
						isOutdoors: formData.isOutdoors
					}
				};
			case 'watering':
				return {
					watering: {
						intervalDays: formData.wateringIntervalDays,
						method: formData.wateringMethod,
						waterType: formData.waterType
					}
				};
			case 'fertilizing':
				return {
					fertilizing: {
						type: formData.fertilizingType,
						intervalDays: formData.fertilizingIntervalDays,
						npkRatio: formData.npkRatio,
						concentrationPercent: formData.concentrationPercent,
						activeInWinter: formData.activeInWinter
					}
				};
			case 'humidity':
				return {
					humidity: {
						targetHumidityPct: formData.targetHumidity,
						requiresMisting: formData.requiresMisting,
						mistingIntervalDays: formData.mistingIntervalDays,
						requiresHumidifier: formData.requiresHumidifier
					}
				};
			case 'soil':
				return {
					soil: {
						type: formData.soilType,
						repottingCycle: formData.repottingCycle,
						components: formData.soilComponents
					}
				};
			case 'seasonality':
				return {
					preferedTemperature: formData.preferedTemperature,
					seasonality: {
						winterRestPeriod: formData.winterRestPeriod,
						winterWaterFactor: formData.winterWaterFactor,
						minTempCelsius: formData.minTempCelsius
					}
				};
			case 'metadata':
				return {
					isToxic: formData.isToxic,
					flags: formData.flags,
					notes: formData.notes
				};
		}
	}

	async function saveSection(): Promise<void> {
		if (!plant) return;
		if (section === 'basic' && (!formData.name.trim() || !formData.species.trim())) {
			error = $tStore('plants.requiredNameSpecies');
			return;
		}

		error = null;
		saving = true;
		try {
			const payload = buildSectionPayload();
			const res = await fetchData('/api/plants/{id}', {
				method: 'patch',
				params: { id: plant.id },
				body: payload
			});

			if (!res.ok) {
				throw new Error(res.error?.message || $tStore('plants.failedToUpdatePlant'));
			}

			await invalidateApiCache(['/api/plants', `/api/plants/${plant.id}`], {
				waitForAck: true,
				timeoutMs: 100
			});

			await goto(resolve(`/manage/${plant.id}${isCreateFlow ? '?createFlow=1' : ''}`));
		} catch (err) {
			error = err instanceof Error ? err.message : $tStore('plants.failedToUpdatePlant');
		} finally {
			saving = false;
		}
	}

	function backToHub(): void {
		if (!plant) return;
		goto(resolve(`/manage/${plant.id}${isCreateFlow ? '?createFlow=1' : ''}`));
	}
</script>

<PageHeader icon="✏️" title={sectionTitles[section]} description={plant?.name || ''} />

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
		{#if error}
			<Alert type="error" title="common.error" description={error} />
		{/if}

		<div class="space-y-4 px-2 pb-[calc(env(safe-area-inset-bottom)+12rem)]">
			{#if section === 'basic'}
				<BasicInformationForm {formData} />
			{:else if section === 'location'}
				<LocationForm {formData} />
			{:else if section === 'watering'}
				<WateringForm {formData} />
			{:else if section === 'fertilizing'}
				<FertilizingForm {formData} />
			{:else if section === 'humidity'}
				<MistingForm {formData} />
			{:else if section === 'soil'}
				<SoilForm {formData} bind:soilComponentInput />
			{:else if section === 'seasonality'}
				<SeasonalityForm {formData} />
			{:else if section === 'metadata'}
				<MetadataForm {formData} bind:newNote />
			{/if}
		</div>

		<div
			class="fixed right-3 left-3 z-50 flex gap-3 rounded-2xl border border-gray-200 bg-white/95 p-3 shadow-lg backdrop-blur md:right-10 md:left-10 xl:right-32 xl:left-32"
			style="bottom: calc(env(safe-area-inset-bottom) + 5.5rem);"
		>
			<Button
				variant="secondary"
				size="lg"
				onclick={backToHub}
				text={isCreateFlow ? 'common.skip' : 'common.close'}
				class="w-full"
			/>
			<Button
				variant="primary"
				size="lg"
				disabled={saving || !isDirty}
				onclick={saveSection}
				text={saving ? 'common.loading' : 'common.save'}
				class="w-full"
			/>
		</div>
	{/if}
</PageContent>
