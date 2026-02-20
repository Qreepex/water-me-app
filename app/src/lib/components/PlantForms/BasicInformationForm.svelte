<script lang="ts">
	import { tStore } from '$lib/i18n';
	import type { FormData } from '$lib/types/forms';

	interface Props {
		formData: FormData;
	}

	let { formData = $bindable() }: Props = $props();
</script>

<div class="space-y-4">
	<h2 class="mb-4 text-2xl font-bold text-green-800">ℹ️ {$tStore('plants.basicInformation')}</h2>

	<div class="grid grid-cols-1 gap-4 md:grid-cols-2">
		<!-- Name -->
		<div>
			<label for="name" class="mb-1 block text-base font-semibold text-gray-700">
				{$tStore('plants.plantName')} *
			</label>
			<input
				type="text"
				id="name"
				required
				bind:value={formData.name}
				class="w-full rounded-lg border-2 border-emerald-200 px-4 py-3 text-base shadow-sm focus:border-emerald-500 focus:outline-none"
			/>
		</div>

		<!-- Species -->
		<div>
			<label for="species" class="mb-1 block text-base font-semibold text-gray-700">
				{$tStore('plants.species')} *
			</label>
			<input
				type="text"
				id="species"
				required
				bind:value={formData.species}
				class="w-full rounded-lg border-2 border-emerald-200 px-4 py-3 text-base shadow-sm focus:border-emerald-500 focus:outline-none"
			/>
		</div>
	</div>
</div>
