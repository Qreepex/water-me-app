<script lang="ts">
	import { tStore } from '$lib/i18n';
	import type { FormData } from '$lib/types/forms';

	interface Props {
		formData: FormData;
	}

	let { formData = $bindable() }: Props = $props();
</script>

<div class="space-y-4">
	<h2 class="mb-4 text-xl font-bold text-green-800">💦 {$tStore('plants.humidityTitle')}</h2>

	<div class="grid grid-cols-1 gap-4 md:grid-cols-2">
		<div>
			<label for="humidity" class="mb-1 block text-base font-semibold text-gray-700">
				{$tStore('plants.targetHumidity')} (%):
				<span class="font-bold text-emerald-600">{formData.targetHumidity}</span>
			</label>
			<input
				type="range"
				id="humidity"
				min="0"
				max="100"
				bind:value={formData.targetHumidity}
				class="w-full accent-emerald-600"
			/>
		</div>

		<label class="flex min-h-11 items-center gap-3">
			<input type="checkbox" bind:checked={formData.requiresMisting} class="h-5 w-5" />
			<span class="text-base font-semibold text-gray-700"
				>{$tStore('plants.formRequiresMisting')}</span
			>
		</label>

		<label class="flex min-h-11 items-center gap-3">
			<input type="checkbox" bind:checked={formData.requiresHumidifier} class="h-5 w-5" />
			<span class="text-base font-semibold text-gray-700"
				>{$tStore('plants.formNeedsHumidifier')}</span
			>
		</label>

		{#if formData.requiresMisting}
			<div>
				<label for="mist-interval" class="mb-1 block text-base font-semibold text-gray-700">
					{$tStore('plants.mistingInterval')} ({$tStore('plants.days')})
				</label>
				<input
					type="number"
					id="mist-interval"
					min="1"
					bind:value={formData.mistingIntervalDays}
					class="w-full rounded-lg border-2 border-emerald-200 px-4 py-3 text-base shadow-sm focus:border-emerald-500 focus:outline-none"
				/>
			</div>
		{/if}
	</div>
</div>
