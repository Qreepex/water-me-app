<script lang="ts">
	import { tStore } from '$lib/i18n';
	import type { Snippet } from 'svelte';

	interface Props {
		title: string;
		description?: string;
		icon?: string;
		children?: Snippet;
	}

	const { title, description, icon, children }: Props = $props();
</script>

<div class="mb-8 flex flex-wrap items-center justify-between gap-3">
	<div class="flex min-w-0 items-center gap-3">
		{#if icon}
			<span class="flex-shrink-0 text-5xl">{icon}</span>
		{/if}
		<div class="min-w-0">
			<h1 class="truncate text-4xl font-bold text-[var(--text-light-main)]">{$tStore(title)}</h1>
			{#if description}
				<p class="truncate text-[var(--text-light-main)]/60">{$tStore(description)}</p>
			{/if}
		</div>
	</div>

	{@render children?.()}
</div>
