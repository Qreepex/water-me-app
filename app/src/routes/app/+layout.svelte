<script lang="ts">
	import BurgerMenu from '$lib/components/BurgerMenu.svelte';

	let { children } = $props();
</script>

<!-- Floating Burger Menu -->
<div class="pt-safe pr-safe fixed top-0 right-0 z-50">
	<div class="p-4">
		<BurgerMenu />
	</div>
</div>

{@render children()}

<style>
	/* Safe area insets for mobile notches and status bars */
	.pt-safe {
		padding-top: env(safe-area-inset-top);
	}
	.pr-safe {
		padding-right: env(safe-area-inset-right);
	}
</style>
