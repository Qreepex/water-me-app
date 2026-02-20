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

<div class="mb-4 flex flex-wrap items-center justify-between gap-3 md:mb-8 px-1">
	<div class="flex min-w-0 items-center gap-3">
		{#if icon}
			<span class="flex-shrink-0 text-4xl md:text-5xl">{icon}</span>
		{/if}
		<div class="min-w-0">
			<h1 class="truncate text-3xl font-bold text-[var(--text-light-main)] md:text-4xl">
				{$tStore(title)}
			</h1>
			{#if description}
				<p class="truncate text-sm text-[var(--text-light-main)]/65 md:text-base">
					{$tStore(description)}
				</p>
			{/if}
		</div>
	</div>
<div class="px-2">
		{@render children?.()}
</div>

</div>
