<script lang="ts">
	import { tStore } from '$lib/i18n';
	import type { FormData } from '$lib/types/forms';

	interface Props {
		formData: FormData;
	}

	let { formData = $bindable() }: Props = $props();
</script>

<div class="space-y-4">
	<h2 class="mb-4 text-xl font-bold text-green-800">{$tStore('plants.seasonalityTitle')}</h2>

	<div class="space-y-4">
		<div>
			<label for="temp" class="mb-1 block text-base font-semibold text-gray-700">
				{$tStore('plants.formPreferredTemperature')} (°C):
				<span class="font-bold text-emerald-600">{formData.preferedTemperature}</span>
			</label>
			<input
				type="range"
				id="temp"
				min="-50"
				max="100"
				bind:value={formData.preferedTemperature}
				class="w-full accent-emerald-600"
			/>
		</div>

		<label class="flex min-h-11 items-center gap-3">
			<input type="checkbox" bind:checked={formData.winterRestPeriod} class="h-5 w-5" />
			<span class="text-base font-semibold text-gray-700"
				>{$tStore('plants.formHasWinterRest')}</span
			>
		</label>

		<div class="grid grid-cols-1 gap-4 md:grid-cols-2">
			<div>
				<label for="winter-factor" class="mb-1 block text-base font-semibold text-gray-700">
					{$tStore('plants.winterWaterFactor')}:
					<span class="font-bold text-emerald-600"
						>{(formData.winterWaterFactor * 100).toFixed(0)}%</span
					>
				</label>
				<input
					type="range"
					id="winter-factor"
					min="0"
					max="1"
					step="0.1"
					bind:value={formData.winterWaterFactor}
					class="w-full accent-blue-600"
				/>
				<p class="mt-1 text-sm text-gray-500">{$tStore('plants.formWinterWaterFactorHint')}</p>
			</div>

			<div>
				<label for="min-temp" class="mb-1 block text-base font-semibold text-gray-700">
					{$tStore('plants.minTemp')} (°C)
				</label>
				<input
					type="number"
					id="min-temp"
					bind:value={formData.minTempCelsius}
					class="w-full rounded-lg border-2 border-emerald-200 px-4 py-3 text-base shadow-sm focus:border-emerald-500 focus:outline-none"
				/>
			</div>
		</div>
	</div>
</div>
