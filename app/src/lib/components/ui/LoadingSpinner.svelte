<script lang="ts">
	import { tStore } from '$lib/i18n';

	interface Props {
		icon?: string;
		message?: string;
	}

	const { icon = '🌱', message = 'common.loading' }: Props = $props();

	// Translate the message as it might be an i18n key
	const translatedMessage = $derived(message ? $tStore(message) : 'Loading...');
</script>

<div class="flex min-h-96 items-center justify-center">
	<div class="text-center">
		<div class="mb-4 animate-spin text-5xl">{icon}</div>
		<p class="text-lg text-[var(--text-light-main)]">{translatedMessage}</p>
	</div>
</div>
