<script lang="ts">
	import type { Snippet } from 'svelte';
	import List from '../List.svelte';

	interface Props {
		children: Snippet;
		multi?: boolean;
	}

	const { children, multi = false }: Props = $props();
</script>

<div class="min-h-0 flex-1 overflow-y-auto">
	<List {multi}>
		{@render children?.()}
	</List>
</div>
