<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { fetchData } from '$lib/auth/fetch.svelte';
	import type { FormData } from '$lib/types/forms';
	import { createEmptyFormData } from '$lib/types/forms';
	import BasicInformationForm from '$lib/components/PlantForms/BasicInformationForm.svelte';
	import LocationForm from '$lib/components/PlantForms/LocationForm.svelte';
	import WateringForm from '$lib/components/PlantForms/WateringForm.svelte';
	import FertilizingForm from '$lib/components/PlantForms/FertilizingForm.svelte';
	import MistingForm from '$lib/components/PlantForms/MistingForm.svelte';
	import SoilForm from '$lib/components/PlantForms/SoilForm.svelte';
	import SeasonalityForm from '$lib/components/PlantForms/SeasonalityForm.svelte';
	import MetadataForm from '$lib/components/PlantForms/MetadataForm.svelte';

	let formData: FormData = createEmptyFormData();
	let error: string | null = null;
	let success: string | null = null;
	let submitting = false;

	let newNote = '';
	let soilComponentInput = '';

	async function submitForm(): Promise<void> {
		if (!formData.species.trim() || !formData.name.trim()) {
			error = 'Species and name are required';
			return;
		}

		submitting = true;
		error = null;
		success = null;

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
					waterType: formData.waterType,
					lastWatered: null
				},
				fertilizing: {
					type: formData.fertilizingType,
					intervalDays: formData.fertilizingIntervalDays,
					npkRatio: formData.npkRatio,
					concentrationPercent: formData.concentrationPercent,
					lastFertilized: null,
					activeInWinter: formData.activeInWinter
				},
				humidity: {
					requiresMisting: formData.requiresMisting,
					mistingIntervalDays: formData.mistingIntervalDays,
					requiresHumidifier: formData.requiresHumidifier,
					targetHumidityPct: formData.targetHumidity
				},
				soil: {
					type: formData.soilType,
					components: formData.soilComponents,
					repottingCycle: formData.repottingCycle
				},
				seasonality: {
					winterRestPeriod: formData.winterRestPeriod,
					winterWaterFactor: formData.winterWaterFactor,
					minTempCelsius: formData.minTempCelsius
				},
				pestHistory: [],
				flags: formData.flags,
				notes: formData.notes,
				photoIds: [],
				growthHistory: []
			};

			const createRes = await fetchData('/api/plants', {
				method: 'post',
				body: createPayload
			});
			if (!createRes.ok) throw new Error(createRes.error?.message || 'Failed to create plant');
			success = 'Plant created successfully!';
			setTimeout(() => goto(resolve('/manage')), 1500);
		} catch (err) {
			error = err instanceof Error ? err.message : 'Unknown error';
		} finally {
			submitting = false;
		}
	}

	function resetForm(): void {
		formData = createEmptyFormData();
		newNote = '';
		soilComponentInput = '';
		error = null;
	}
</script>

<div class="min-h-screen bg-gradient-to-br from-emerald-50 via-green-50 to-teal-100 p-6 md:p-10">
	<div class="mx-auto max-w-4xl">
		<!-- Header -->
		<div class="mb-8">
			<div class="mb-4 flex items-center justify-between">
				<div>
					<h1 class="flex items-center gap-3 text-4xl font-bold text-green-900">
						🌱 Create New Plant
					</h1>
					<p class="mt-1 text-sm text-emerald-700 italic">Add a new plant to your collection</p>
				</div>
				<a
					href={resolve('/manage')}
					class="rounded-xl bg-gray-600 px-4 py-2 font-medium text-white shadow-sm transition hover:bg-gray-700"
				>
					← Back
				</a>
			</div>
		</div>

		<!-- Messages -->
		{#if success}
			<div class="mb-6 rounded-lg border-2 border-green-400 bg-green-100 px-6 py-4 text-green-800">
				✓ {success}
			</div>
		{/if}

		{#if error}
			<div class="mb-6 rounded-lg border-2 border-red-400 bg-red-100 px-6 py-4 text-red-800">
				✕ {error}
			</div>
		{/if}

		<div class="space-y-6">
			<!-- Images Section -->
			<div class="rounded-2xl border border-emerald-100 bg-white/90 p-6 shadow-md backdrop-blur">
				<h2 class="mb-4 text-2xl font-bold text-green-800">📸 Photos</h2>
				<div
					class="flex h-48 items-center justify-center rounded-lg border-2 border-dashed border-emerald-300 bg-emerald-50"
				>
					<div class="text-center">
						<div class="mb-2 text-4xl">🖼️</div>
						<p class="text-sm text-emerald-700">Photo management coming soon</p>
					</div>
				</div>
			</div>

			<!-- Basic Information -->
			<BasicInformationForm {formData} />

			<!-- Location -->
			<LocationForm {formData} />

			<!-- Watering & Fertilizing -->
			<div class="grid grid-cols-1 gap-6 md:grid-cols-2">
				<WateringForm {formData} />
				<FertilizingForm {formData} />
			</div>

			<!-- Misting -->
			<MistingForm {formData} />

			<!-- Advanced Section -->
			<div
				class="space-y-6 rounded-2xl border border-emerald-100 bg-white/90 p-6 shadow-md backdrop-blur"
			>
				<h2 class="text-xl font-bold text-green-800">⚙️ Advanced Settings</h2>

				<div class="space-y-6">
					<SoilForm {formData} bind:soilComponentInput />
					<SeasonalityForm {formData} />
					<MetadataForm {formData} bind:newNote />
				</div>
			</div>

			<div class="flex justify-between gap-3">
				<button
					on:click={resetForm}
					class="rounded-lg bg-gray-200 px-6 py-3 font-semibold text-gray-800 transition hover:bg-gray-300"
				>
					Reset
				</button>
				<button
					on:click={submitForm}
					disabled={submitting}
					class="rounded-lg bg-gradient-to-r from-emerald-600 to-green-600 px-8 py-3 font-semibold text-white shadow-md transition hover:from-emerald-700 hover:to-green-700 disabled:opacity-50"
				>
					{submitting ? 'Creating...' : 'Create Plant'}
				</button>
			</div>
		</div>
	</div>
</div>

<style>
	:global(body) {
		font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
	}
</style>
