<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import type { Plant } from '$lib/types/api';
	import { SunlightRequirement, WateringMethod, WaterType, FertilizerType } from '$lib/types/api';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { page } from '$app/stores';
	import { fetchData } from '$lib/auth/fetch.svelte';
	import { getImageObjectURL, revokeObjectURL } from '$lib/utils/imageCache';
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

	let plant: Plant | null = null;
	let loading = true;
	let error: string | null = null;
	let success: string | null = null;
	let submitting = false;
	let newNote = '';
	let soilComponentInput = '';
	let previewUrls: string[] = [];

	let formData: FormData = createEmptyFormData();

	onMount(async () => {
		try {
			const plantId = $page.params.plant ?? '';
			const response = await fetchData('/api/plants/{id}', {
				params: { id: plantId }
			});

			if (!response.ok) {
				error = response.error?.message || 'Failed to load plant';
				return;
			}

			plant = response.data;
			formData = initializeFormData();
			await loadPhotoPreviews();
		} catch (err) {
			error = err instanceof Error ? err.message : 'Failed to load plant';
		} finally {
			loading = false;
		}
	});

	async function loadPhotoPreviews(): Promise<void> {
		if (!plant) return;
		const ids = plant.photoIds || [];
		// eslint-disable-next-line @typescript-eslint/no-explicit-any
		const urls = ((plant as any)?.photoUrls as string[] | undefined) || [];
		previewUrls = [];
		for (let i = 0; i < ids.length; i++) {
			const id = ids[i];
			const url = urls[i];
			if (!id || !url) continue;
			const objUrl = await getImageObjectURL(id, url);
			if (objUrl) previewUrls.push(objUrl);
		}
	}

	onDestroy(() => {
		previewUrls.forEach((u) => revokeObjectURL(u));
	});

	function initializeFormData(): FormData {
		if (!plant) return createEmptyFormData();
		return {
			// Basic info
			name: plant.name,
			species: plant.species,
			isToxic: plant.isToxic,
			sunlight: plant.sunlight as SunlightRequirement,
			preferedTemperature: plant.preferedTemperature,

			// Location
			room: plant.location?.room ?? 'Unknown',
			position: plant.location?.position ?? 'Unknown',
			isOutdoors: plant.location?.isOutdoors ?? false,

			// Watering
			wateringIntervalDays: plant.watering?.intervalDays ?? 7,
			wateringMethod: plant.watering?.method ?? WateringMethod.Top,
			waterType: plant.watering?.waterType ?? WaterType.Tap,

			// Fertilizing
			fertilizingType: plant.fertilizing?.type ?? FertilizerType.Liquid,
			fertilizingIntervalDays: plant.fertilizing?.intervalDays ?? 30,
			npkRatio: plant.fertilizing?.npkRatio ?? '10:10:10',
			concentrationPercent: plant.fertilizing?.concentrationPercent ?? 50,
			activeInWinter: plant.fertilizing?.activeInWinter ?? false,

			// Humidity
			targetHumidity: plant.humidity?.targetHumidityPct ?? 50,
			requiresMisting: plant.humidity?.requiresMisting ?? false,
			mistingIntervalDays: plant.humidity?.mistingIntervalDays ?? 3,
			requiresHumidifier: plant.humidity?.requiresHumidifier ?? false,

			// Soil
			soilType: plant.soil?.type ?? 'Generic',
			repottingCycle: plant.soil?.repottingCycle ?? 2,
			soilComponents: plant.soil?.components ?? [],

			// Seasonality
			winterRestPeriod: plant.seasonality?.winterRestPeriod ?? false,
			winterWaterFactor: plant.seasonality?.winterWaterFactor ?? 0.5,
			minTempCelsius: plant.seasonality?.minTempCelsius ?? 15,

			// Metadata
			flags: plant.flags ?? [],
			notes: plant.notes ?? []
		};
	}

	async function submitForm(): Promise<void> {
		if (!formData.species.trim() || !formData.name.trim()) {
			error = 'Species and name are required';
			return;
		}

		submitting = true;
		error = null;
		success = null;

		try {
			const updatePayload = {
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
				flags: formData.flags,
				notes: formData.notes
			};

			const res = await fetchData('/api/plants/{id}', {
				method: 'patch',
				params: { id: plant?.id ?? '' },
				body: updatePayload
			});

			if (!res.ok) {
				throw new Error(res.error?.message || 'Failed to update plant');
			}

			success = 'Plant updated successfully!';
			if (res.data) {
				plant = res.data;
			}
			setTimeout(() => goto(resolve('/manage')), 1500);
		} catch (err) {
			error = err instanceof Error ? err.message : 'Unknown error';
		} finally {
			submitting = false;
		}
	}

	function resetForm(): void {
		formData = initializeFormData();
		error = null;
	}
</script>

<div class="min-h-screen bg-gradient-to-br from-emerald-50 via-green-50 to-teal-100 p-6 md:p-10">
	<div class="mx-auto max-w-4xl">
		{#if loading}
			<div class="flex min-h-screen items-center justify-center">
				<div class="text-center">
					<div class="mb-4 animate-spin text-4xl">🌱</div>
					<p class="text-lg text-gray-600">Loading plant details...</p>
				</div>
			</div>
		{:else if !plant}
			<div class="flex min-h-screen items-center justify-center">
				<div class="text-center">
					<p class="mb-4 text-lg text-red-600">{error || 'Plant not found'}</p>
					<a
						href={resolve('/manage')}
						class="rounded-xl bg-gray-600 px-4 py-2 font-medium text-white shadow-sm transition hover:bg-gray-700"
					>
						← Back to Plants
					</a>
				</div>
			</div>
		{:else}
			<div class="mb-8">
				<div class="mb-4 flex items-center justify-between">
					<div>
						<h1 class="flex items-center gap-3 text-4xl font-bold text-green-900">
							🌿 {plant?.name || 'Plant'}
						</h1>
						<p class="mt-1 text-sm text-emerald-700 italic">{plant?.species || ''}</p>
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
				<div
					class="mb-6 rounded-lg border-2 border-green-400 bg-green-100 px-6 py-4 text-green-800"
				>
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
					{#if previewUrls.length}
						<div class="grid grid-cols-2 gap-3 md:grid-cols-3">
							{#each previewUrls as u (u)}
								<img src={u} alt="" class="h-32 w-full rounded object-cover" />
							{/each}
						</div>
					{:else}
						<div
							class="flex h-48 items-center justify-center rounded-lg border-2 border-dashed border-emerald-300 bg-emerald-50"
						>
							<div class="text-center">
								<div class="mb-2 text-4xl">🖼️</div>
								<p class="text-sm text-emerald-700">No photos yet</p>
							</div>
						</div>
					{/if}
				</div>

				<!-- Form Sections -->
				<BasicInformationForm {formData} />
				<LocationForm {formData} />

				<!-- Watering & Fertilizing -->
				<div class="grid grid-cols-1 gap-6 md:grid-cols-2">
					<WateringForm {formData} />
					<FertilizingForm {formData} />
				</div>

				<MistingForm {formData} />

				<!-- Advanced Settings -->
				<SoilForm {formData} bind:soilComponentInput />
				<SeasonalityForm {formData} />
				<MetadataForm {formData} bind:newNote />

				<!-- Action Buttons -->
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
						{submitting ? 'Saving...' : 'Save Changes'}
					</button>
				</div>
			</div>
		{/if}
	</div>
</div>

<style>
	:global(body) {
		font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
	}
</style>
