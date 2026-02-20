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
			const hasChanged = <K extends keyof FormData>(key: K): boolean => {
				if (Array.isArray(formData[key]) && Array.isArray(initialFormData[key])) {
					return JSON.stringify(formData[key]) !== JSON.stringify(initialFormData[key]);
				}

				return formData[key] !== initialFormData[key];
			};

			const now = new Date().toISOString();
			const createPayload: Record<string, unknown> = {
				name: formData.name.trim(),
				species: formData.species.trim()
			};

			if (hasChanged('isToxic')) createPayload.isToxic = formData.isToxic;
			if (hasChanged('sunlight')) createPayload.sunlight = formData.sunlight;
			if (hasChanged('preferedTemperature')) {
				createPayload.preferedTemperature = formData.preferedTemperature;
			}

			if (hasChanged('room') || hasChanged('position') || hasChanged('isOutdoors')) {
				createPayload.location = {
					room: formData.room.trim(),
					position: formData.position.trim(),
					isOutdoors: formData.isOutdoors
				};
			}

			if (
				hasChanged('wateringIntervalDays') ||
				hasChanged('wateringMethod') ||
				hasChanged('waterType')
			) {
				createPayload.watering = {
					intervalDays: formData.wateringIntervalDays,
					method: formData.wateringMethod,
					waterType: formData.waterType,
					lastWatered: now
				};
			}

			if (
				hasChanged('fertilizingType') ||
				hasChanged('fertilizingIntervalDays') ||
				hasChanged('npkRatio') ||
				hasChanged('concentrationPercent') ||
				hasChanged('activeInWinter')
			) {
				createPayload.fertilizing = {
					type: formData.fertilizingType,
					intervalDays: formData.fertilizingIntervalDays,
					npkRatio: formData.npkRatio.trim(),
					concentrationPercent: formData.concentrationPercent,
					activeInWinter: formData.activeInWinter,
					lastFertilized: now
				};
			}

			if (
				hasChanged('targetHumidity') ||
				hasChanged('requiresMisting') ||
				hasChanged('mistingIntervalDays') ||
				hasChanged('requiresHumidifier')
			) {
				createPayload.humidity = {
					targetHumidityPct: formData.targetHumidity,
					requiresMisting: formData.requiresMisting,
					mistingIntervalDays: formData.mistingIntervalDays,
					requiresHumidifier: formData.requiresHumidifier
				};
			}

			if (hasChanged('soilType') || hasChanged('repottingCycle') || hasChanged('soilComponents')) {
				createPayload.soil = {
					type: formData.soilType.trim(),
					repottingCycle: formData.repottingCycle,
					components: formData.soilComponents
				};
			}

			if (
				hasChanged('winterRestPeriod') ||
				hasChanged('winterWaterFactor') ||
				hasChanged('minTempCelsius')
			) {
				createPayload.seasonality = {
					winterRestPeriod: formData.winterRestPeriod,
					winterWaterFactor: formData.winterWaterFactor,
					minTempCelsius: formData.minTempCelsius
				};
			}

			if (hasChanged('flags')) createPayload.flags = formData.flags;
			if (hasChanged('notes')) createPayload.notes = formData.notes;

			const res = await fetchData('/api/plants', {
				method: 'post' as const,
				body: createPayload as never
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

	<div class="space-y-4 px-2 pb-[calc(env(safe-area-inset-bottom)+12rem)]">
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
