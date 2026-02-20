<script lang="ts">
	import type { FormData } from '$lib/types/forms';

	interface Props {
		formData: FormData;
	}

	let { formData = $bindable() }: Props = $props();
</script>

<div class="space-y-4">
	<h2 class="mb-4 text-2xl font-bold text-green-800">📍 Location</h2>

	<div class="grid grid-cols-1 gap-4 md:grid-cols-3">
		<div>
			<label for="room" class="mb-1 block text-base font-semibold text-gray-700">Room</label>
			<input
				type="text"
				id="room"
				bind:value={formData.room}
				placeholder="e.g., Living Room"
				class="w-full rounded-lg border-2 border-emerald-200 px-4 py-3 text-base shadow-sm focus:border-emerald-500 focus:outline-none"
			/>
		</div>

		<div>
			<label for="position" class="mb-1 block text-base font-semibold text-gray-700">
				Position
			</label>
			<input
				type="text"
				id="position"
				bind:value={formData.position}
				placeholder="e.g., Windowsill"
				class="w-full rounded-lg border-2 border-emerald-200 px-4 py-3 text-base shadow-sm focus:border-emerald-500 focus:outline-none"
			/>
		</div>

		<div>
			<fieldset>
				<legend class="mb-1 block text-base font-semibold text-gray-700">Environment</legend>
				<label class="flex min-h-11 items-center gap-3">
					<input type="checkbox" bind:checked={formData.isOutdoors} class="h-5 w-5" />
					<span class="text-base text-gray-700">Outdoors</span>
				</label>
			</fieldset>
		</div>
	</div>
</div>
