<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		children: Snippet;
	}

	const { children }: Props = $props();
</script>

<div class="flex h-full min-h-0 flex-col overflow-hidden">
	{@render children()}
</div>
