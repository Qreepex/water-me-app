<script lang="ts">
	import { tStore } from '$lib/i18n';
	import type { FormData } from '$lib/types/forms';
	import { FertilizerType } from '$lib/types/api';

	interface Props {
		formData: FormData;
	}

	let { formData = $bindable() }: Props = $props();
</script>

<div class="space-y-4">
	<h2 class="mb-4 text-xl font-bold text-green-800">{$tStore('plants.fertilizingTitle')}</h2>

	<div class="space-y-4">
		<div>
			<label for="fert-interval" class="mb-1 block text-base font-semibold text-gray-700">
				{$tStore('plants.formIntervalDays')}
			</label>
			<input
				type="number"
				id="fert-interval"
				min="1"
				bind:value={formData.fertilizingIntervalDays}
				class="w-full rounded-lg border-2 border-emerald-200 px-4 py-3 text-base shadow-sm focus:border-emerald-500 focus:outline-none"
			/>
		</div>

		<div>
			<label for="fert-type" class="mb-1 block text-base font-semibold text-gray-700"
				>{$tStore('plants.formType')}</label
			>
			<select
				id="fert-type"
				bind:value={formData.fertilizingType}
				class="w-full rounded-lg border-2 border-emerald-200 px-4 py-3 text-base shadow-sm focus:border-emerald-500 focus:outline-none"
			>
				{#each Object.values(FertilizerType) as type (type)}
					<option value={type}>{$tStore('plants.fertilizerTypeOptions.' + type) || type}</option>
				{/each}
			</select>
		</div>

		<div>
			<label for="npk-ratio" class="mb-1 block text-base font-semibold text-gray-700">
				{$tStore('plants.npkRatio')}
			</label>
			<input
				type="text"
				id="npk-ratio"
				bind:value={formData.npkRatio}
				placeholder={$tStore('plants.formNpkPlaceholder')}
				class="w-full rounded-lg border-2 border-emerald-200 px-4 py-3 text-base shadow-sm focus:border-emerald-500 focus:outline-none"
			/>
		</div>

		<div>
			<label for="concentration" class="mb-1 block text-base font-semibold text-gray-700">
				{$tStore('plants.concentration')} (%):
				<span class="font-bold text-emerald-600">{formData.concentrationPercent}</span>
			</label>
			<input
				type="range"
				id="concentration"
				min="1"
				max="100"
				bind:value={formData.concentrationPercent}
				class="w-full accent-emerald-600"
			/>
		</div>

		<label class="flex min-h-11 items-center gap-3">
			<input type="checkbox" bind:checked={formData.activeInWinter} class="h-5 w-5" />
			<span class="text-base text-gray-700">{$tStore('plants.formFertilizeInWinter')}</span>
		</label>
	</div>
</div>
