<script lang="ts">
	import { tStore } from '$lib/i18n';
	import type { FormData } from '$lib/types/forms';
	import { WateringMethod, WaterType } from '$lib/types/api';

	interface Props {
		formData: FormData;
	}

	let { formData = $bindable() }: Props = $props();
</script>

<div class="space-y-4">
	<h2 class="mb-4 text-xl font-bold text-green-800">💧 {$tStore('plants.wateringTitle')}</h2>

	<div class="space-y-4">
		<div>
			<label for="water-interval" class="mb-1 block text-base font-semibold text-gray-700">
				{$tStore('plants.formIntervalDays')}
			</label>
			<input
				type="number"
				id="water-interval"
				min="1"
				bind:value={formData.wateringIntervalDays}
				class="w-full rounded-lg border-2 border-emerald-200 px-4 py-3 text-base shadow-sm focus:border-emerald-500 focus:outline-none"
			/>
		</div>

		<div>
			<label for="water-method" class="mb-1 block text-base font-semibold text-gray-700">
				{$tStore('plants.wateringMethod')}
			</label>
			<select
				id="water-method"
				bind:value={formData.wateringMethod}
				class="w-full rounded-lg border-2 border-emerald-200 px-4 py-3 text-base shadow-sm focus:border-emerald-500 focus:outline-none"
			>
				{#each Object.values(WateringMethod) as method (method)}
					<option value={method}
						>{$tStore('plants.wateringMethodOptions.' + method) || method}</option
					>
				{/each}
			</select>
		</div>

		<div>
			<label for="water-type" class="mb-1 block text-base font-semibold text-gray-700">
				{$tStore('plants.waterType')}
			</label>
			<select
				id="water-type"
				bind:value={formData.waterType}
				class="w-full rounded-lg border-2 border-emerald-200 px-4 py-3 text-base shadow-sm focus:border-emerald-500 focus:outline-none"
			>
				{#each Object.values(WaterType) as type (type)}
					<option value={type}>{$tStore('plants.waterTypeOptions.' + type) || type}</option>
				{/each}
			</select>
		</div>
	</div>
</div>
