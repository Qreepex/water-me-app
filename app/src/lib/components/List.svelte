<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		children: Snippet;
		multi?: boolean;
		noPadding?: boolean;
	}

	const { children, multi = false, noPadding = false }: Props = $props();
</script>

<div
	class="grid grid-cols-1 gap-6 {multi
		? 'md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5'
		: ''} {noPadding ? "" : 'px-2'}"
>
	{@render children?.()}
</div>
