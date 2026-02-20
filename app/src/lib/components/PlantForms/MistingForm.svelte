<script lang="ts">
	import type { FormData } from '$lib/types/forms';

	interface Props {
		formData: FormData;
	}

	let { formData = $bindable() }: Props = $props();
</script>

<div class="space-y-4">
	<h2 class="mb-4 text-xl font-bold text-green-800">💦 Misting & Humidity</h2>

	<div class="grid grid-cols-1 gap-4 md:grid-cols-2">
		<label class="flex min-h-11 items-center gap-3">
			<input type="checkbox" bind:checked={formData.requiresMisting} class="h-5 w-5" />
			<span class="text-base font-semibold text-gray-700">Requires misting</span>
		</label>

		<label class="flex min-h-11 items-center gap-3">
			<input type="checkbox" bind:checked={formData.requiresHumidifier} class="h-5 w-5" />
			<span class="text-base font-semibold text-gray-700">Needs humidifier</span>
		</label>

		{#if formData.requiresMisting}
			<div>
				<label for="mist-interval" class="mb-1 block text-base font-semibold text-gray-700">
					Misting Interval (days)
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
