<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		children: Snippet;
	}

	const { children }: Props = $props();
</script>

<div class="min-h-0 flex-1 overflow-hidden">
	{@render children?.()}
</div>
