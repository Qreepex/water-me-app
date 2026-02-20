<script lang="ts">
	import type { Snippet } from 'svelte';
	import List from '../List.svelte';

	interface Props {
		children: Snippet;
		multi?: boolean;
	}

	const { children, multi = false }: Props = $props();

	let isScrolled = $state(false);
	let scrollContainer: HTMLDivElement | undefined = $state();

	function handleScroll() {
		isScrolled = (scrollContainer?.scrollTop ?? 0) > 0;
	}
</script>

<!-- Outer wrapper with flex column layout -->
<div class="relative flex min-h-0 flex-col flex-1">
	<!-- Fade out overlay at top (absolutely positioned, doesn't affect layout) -->
	<div
		class="pointer-events-none absolute top-0 left-0 right-0 h-4 bg-gradient-to-b from-white/90 via-white/40 to-transparent z-10 transition-opacity duration-100 {isScrolled
			? 'opacity-100'
			: 'opacity-0'}"
	></div>

	<!-- Scrollable content area -->
	<div
		bind:this={scrollContainer}
		onscroll={handleScroll}
		class="relative min-h-0 flex-1 overflow-y-auto overscroll-y-contain"
	>
		<List {multi}>
			{@render children?.()}
		</List>

		<!-- Fade out overlay at bottom to indicate scrollable content -->
		<div
			class="pointer-events-none sticky bottom-0 left-0 right-0 h-4 bg-gradient-to-b from-transparent via-white/40 to-white/90 z-10"
		></div>
	</div>
</div>
