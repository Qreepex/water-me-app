<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { fetchData } from '$lib/auth/fetch.svelte';
	import { invalidateApiCache } from '$lib/utils/cache';
	import { getPlantsStore } from '$lib/stores/plants.svelte';
	import { tStore } from '$lib/i18n';
	import type { FormData } from '$lib/types/forms';
	import { createEmptyFormData } from '$lib/types/forms';
	import PageHeader from '$lib/components/layout/PageHeader.svelte';
	import PageContent from '$lib/components/layout/PageContent.svelte';
	import BasicInformationForm from '$lib/components/PlantForms/BasicInformationForm.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import Alert from '$lib/components/ui/Message.svelte';

	const plantsStore = getPlantsStore();
	const initialFormData = createEmptyFormData();
	let formData = $state<FormData>(createEmptyFormData());
	let submitting = $state(false);
	let error = $state<string | null>(null);
	const isDirty = $derived(JSON.stringify(formData) !== JSON.stringify(initialFormData));

	async function saveBasicInformation(): Promise<void> {
		if (!formData.name.trim() || !formData.species.trim()) {
			error = $tStore('plants.requiredNameSpecies');
			return;
		}

		submitting = true;
		error = null;

		try {
			const createPayload = {
				name: formData.name,
				species: formData.species,
				isToxic: formData.isToxic,
				sunlight: formData.sunlight,
				preferedTemperature: formData.preferedTemperature,
				location: {
					room: formData.room,
					position: formData.position,
					isOutdoors: formData.isOutdoors
				},
				watering: {
					intervalDays: formData.wateringIntervalDays,
					method: formData.wateringMethod,
					waterType: formData.waterType
				},
				fertilizing: {
					type: formData.fertilizingType,
					intervalDays: formData.fertilizingIntervalDays,
					npkRatio: formData.npkRatio,
					concentrationPercent: formData.concentrationPercent,
					activeInWinter: formData.activeInWinter
				},
				humidity: {
					targetHumidityPct: formData.targetHumidity,
					requiresMisting: formData.requiresMisting,
					mistingIntervalDays: formData.mistingIntervalDays,
					requiresHumidifier: formData.requiresHumidifier
				},
				soil: {
					type: formData.soilType,
					repottingCycle: formData.repottingCycle,
					components: formData.soilComponents
				},
				seasonality: {
					winterRestPeriod: formData.winterRestPeriod,
					winterWaterFactor: formData.winterWaterFactor,
					minTempCelsius: formData.minTempCelsius
				},
				photoIds: [],
				flags: formData.flags,
				notes: formData.notes,
				pestHistory: [],
				growthHistory: []
			};

			const res = await fetchData('/api/plants', {
				method: 'post' as const,
				body: createPayload
			});

			if (!res.ok) {
				throw new Error(res.error?.message || $tStore('plants.failedToCreatePlant'));
			}

			if (!res.data?.id) {
				throw new Error($tStore('plants.invalidResponse'));
			}

			plantsStore.setPlants([...plantsStore.plants, res.data]);
			await invalidateApiCache(['/api/plants', `/api/plants/${res.data.id}`], {
				waitForAck: true,
				timeoutMs: 100
			});

			await goto(resolve(`/manage/${res.data.id}/photos?createFlow=1`));
		} catch (err) {
			error = err instanceof Error ? err.message : $tStore('plants.failedToCreatePlant');
		} finally {
			submitting = false;
		}
	}

	function cancelCreate(): void {
		goto(resolve('/'));
	}
</script>

<PageHeader
	icon="🌱"
	title="plants.basicInformation"
	description="plants.requiredInfoDescription"
/>

<PageContent>
	{#if error}
		<Alert type="error" title="common.error" description={error} />
	{/if}

	<div class="space-y-4 pb-[calc(env(safe-area-inset-bottom)+12rem)]">
		<BasicInformationForm {formData} />
		<div
			class="fixed right-3 left-3 z-50 flex gap-3 rounded-2xl border border-gray-200 bg-white/95 p-3 shadow-lg backdrop-blur md:right-10 md:left-10 xl:right-32 xl:left-32"
			style="bottom: calc(env(safe-area-inset-bottom) + 5.5rem);"
		>
			<Button
				variant="secondary"
				size="lg"
				onclick={cancelCreate}
				text="common.cancel"
				class="w-full"
			/>
			<Button
				variant="primary"
				size="lg"
				disabled={submitting || !isDirty}
				onclick={saveBasicInformation}
				text={submitting ? 'common.loading' : 'common.save'}
				class="w-full"
			/>
		</div>
	</div>
</PageContent>
