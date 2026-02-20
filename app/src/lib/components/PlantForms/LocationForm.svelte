<script lang="ts">
	import { tStore } from '$lib/i18n';
	import type { FormData } from '$lib/types/forms';
	import { SunlightRequirement } from '$lib/types/api';

	interface Props {
		formData: FormData;
	}

	let { formData = $bindable() }: Props = $props();
</script>

<div class="space-y-4">
	<h2 class="mb-4 text-2xl font-bold text-green-800">📍 {$tStore('plants.location')}</h2>

	<div class="grid grid-cols-1 gap-4 md:grid-cols-3">
		<div>
			<label for="sunlight" class="mb-1 block text-base font-semibold text-gray-700">
				{$tStore('plants.formSunlightRequirements')}
			</label>
			<select
				id="sunlight"
				bind:value={formData.sunlight}
				class="w-full rounded-lg border-2 border-emerald-200 px-4 py-3 text-base shadow-sm focus:border-emerald-500 focus:outline-none"
			>
				{#each Object.values(SunlightRequirement) as req (req)}
					<option value={req}>{$tStore('plants.sunlight.' + req) || req}</option>
				{/each}
			</select>
		</div>

		<div>
			<label for="room" class="mb-1 block text-base font-semibold text-gray-700">
				{$tStore('plants.formRoom')}
			</label>
			<input
				type="text"
				id="room"
				bind:value={formData.room}
				placeholder={$tStore('plants.formRoomPlaceholder')}
				class="w-full rounded-lg border-2 border-emerald-200 px-4 py-3 text-base shadow-sm focus:border-emerald-500 focus:outline-none"
			/>
		</div>

		<div>
			<label for="position" class="mb-1 block text-base font-semibold text-gray-700">
				{$tStore('plants.formPosition')}
			</label>
			<input
				type="text"
				id="position"
				bind:value={formData.position}
				placeholder={$tStore('plants.formPositionPlaceholder')}
				class="w-full rounded-lg border-2 border-emerald-200 px-4 py-3 text-base shadow-sm focus:border-emerald-500 focus:outline-none"
			/>
		</div>

		<div>
			<fieldset>
				<legend class="mb-1 block text-base font-semibold text-gray-700">
					{$tStore('plants.formEnvironment')}
				</legend>
				<label class="flex min-h-11 items-center gap-3">
					<input type="checkbox" bind:checked={formData.isOutdoors} class="h-5 w-5" />
					<span class="text-base text-gray-700">{$tStore('plants.outdoors')}</span>
				</label>
			</fieldset>
		</div>
	</div>
</div>
