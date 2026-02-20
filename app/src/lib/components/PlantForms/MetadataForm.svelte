<script lang="ts">
	import type { FormData } from '$lib/types/forms';
	import { PlantFlag } from '$lib/types/api';
	import { tStore } from '$lib/i18n';
	import Button from '../ui/Button.svelte';

	interface Props {
		formData: FormData;
		newNote: string;
	}

	let { formData = $bindable(), newNote = $bindable() }: Props = $props();

	function toggleFlag(flag: PlantFlag): void {
		if (formData.flags.includes(flag)) {
			formData.flags = formData.flags.filter((f) => f !== flag);
		} else {
			formData.flags = [...formData.flags, flag];
		}
	}

	function addNote(): void {
		if (newNote.trim()) {
			formData.notes = [...formData.notes, newNote.trim()];
			newNote = '';
		}
	}

	function removeNote(index: number): void {
		formData.notes = formData.notes.filter((_, i) => i !== index);
	}
</script>

<div class="space-y-4">
	<h2 class="mb-4 text-xl font-bold text-green-800">{$tStore('plants.metadataTitle')}</h2>

	<div class="space-y-4">
		<div>
			<span class="mb-2 block text-base font-semibold text-gray-700">Flags</span>
			<div class="space-y-2">
				{#each Object.values(PlantFlag) as flag (flag)}
					<label class="flex min-h-11 items-center gap-3">
						<input
							type="checkbox"
							checked={formData.flags.includes(flag)}
							onchange={() => toggleFlag(flag)}
							class="h-5 w-5"
						/>
						<span class="text-base text-gray-700">{flag}</span>
					</label>
				{/each}
			</div>
		</div>

		<fieldset>
			<legend class="mb-2 block text-base font-semibold text-gray-700">Notes</legend>
			<div class="mb-2 flex flex-col gap-2 sm:flex-row">
				<input
					type="text"
					bind:value={newNote}
					onkeydown={(e) => e.key === 'Enter' && addNote()}
					placeholder="Add a note..."
					class="flex-1 rounded-lg border-2 border-emerald-200 px-4 py-3 text-base shadow-sm focus:border-emerald-500 focus:outline-none"
				/>
				<Button onclick={addNote} text="Add" variant="primary" class="w-full sm:w-auto" />
			</div>

			{#if formData.notes.length > 0}
				<div class="space-y-2">
					{#each formData.notes as note, i (i)}
						<div class="flex items-start justify-between rounded-lg bg-blue-50 p-3">
							<p class="flex-1 text-base text-gray-800">{note}</p>
							<Button text="Remove" variant="danger" size="sm" onclick={() => removeNote(i)} />
						</div>
					{/each}
				</div>
			{:else}
				<p class="text-base text-gray-500 italic">No notes yet</p>
			{/if}
		</fieldset>
	</div>
</div>
