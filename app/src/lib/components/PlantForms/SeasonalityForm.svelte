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
		<label class="flex min-h-11 items-center gap-3">
			<input type="checkbox" bind:checked={formData.winterRestPeriod} class="h-5 w-5" />
			<span class="text-base font-semibold text-gray-700">Has winter rest period</span>
		</label>

		<div class="grid grid-cols-1 gap-4 md:grid-cols-2">
			<div>
				<label for="winter-factor" class="mb-1 block text-base font-semibold text-gray-700">
					Winter Water Factor: <span class="font-bold text-emerald-600"
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
				<p class="mt-1 text-sm text-gray-500">Multiply normal watering by this factor</p>
			</div>

			<div>
				<label for="min-temp" class="mb-1 block text-base font-semibold text-gray-700">
					Minimum Temperature (°C)
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
