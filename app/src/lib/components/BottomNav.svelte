<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { page } from '$app/stores';
	import BurgerMenu from './BurgerMenu.svelte';

	let showMenu = $state(false);

	function isActive(path: '/' | '/water' | '/manage'): boolean {
		return $page.url.pathname === resolve(path) || $page.url.pathname.startsWith(resolve(path) + '/');
	}

	function navigate(path: '/' | '/water' | '/manage'): void {
		goto(resolve(path));
	}

	function toggleMenu() {
		showMenu = !showMenu;
	}
</script>

<!-- Bottom Navigation Bar -->
<div class="fixed bottom-0 left-0 right-0 z-40 border-t border-emerald-200 bg-white shadow-lg">
	<div class="pb-safe flex h-20 items-center justify-around">
		<!-- Home -->
		<button
			onclick={() => navigate('/')}
			class="flex flex-col items-center justify-center gap-1 flex-1 py-2 transition-colors {isActive('/')
				? 'text-emerald-600'
				: 'text-gray-600'}"
			aria-label="Home"
		>
			<svg class="h-6 w-6" fill={isActive('/') ? 'currentColor' : 'none'} stroke="currentColor" viewBox="0 0 24 24">
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M3 12l2-3m0 0l7-4 7 4M5 9v10a1 1 0 001 1h12a1 1 0 001-1V9m-9 11l4-4"
				></path>
			</svg>
			<span class="text-xs font-medium">Home</span>
		</button>

		<!-- Water -->
		<button
			onclick={() => navigate('/water')}
			class="flex flex-col items-center justify-center gap-1 flex-1 py-2 transition-colors {isActive('/water')
				? 'text-emerald-600'
				: 'text-gray-600'}"
			aria-label="Water"
		>
			<svg class="h-6 w-6" fill={isActive('/water') ? 'currentColor' : 'none'} stroke="currentColor" viewBox="0 0 24 24">
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M12 8v4m0 4v.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
				></path>
			</svg>
			<span class="text-xs font-medium">Water</span>
		</button>

		<!-- Manage -->
		<button
			onclick={() => navigate('/manage')}
			class="flex flex-col items-center justify-center gap-1 flex-1 py-2 transition-colors {isActive('/manage')
				? 'text-emerald-600'
				: 'text-gray-600'}"
			aria-label="Manage"
		>
			<svg class="h-6 w-6" fill={isActive('/manage') ? 'currentColor' : 'none'} stroke="currentColor" viewBox="0 0 24 24">
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
				></path>
			</svg>
			<span class="text-xs font-medium">Manage</span>
		</button>

		<!-- Menu -->
		<button
			onclick={toggleMenu}
			class="flex flex-col items-center justify-center gap-1 flex-1 py-2 transition-colors {showMenu
				? 'text-emerald-600'
				: 'text-gray-600'}"
			aria-label="Menu"
		>
			<svg class="h-6 w-6" fill={showMenu ? 'currentColor' : 'none'} stroke="currentColor" viewBox="0 0 24 24">
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M4 6h16M4 12h16M4 18h16"
				></path>
			</svg>
			<span class="text-xs font-medium">Menu</span>
		</button>
	</div>
</div>

<!-- Menu Overlay -->
{#if showMenu}
	<div class="fixed inset-0 z-50 bg-white overflow-y-auto">
		<div class="pt-safe pb-safe flex items-center justify-between border-b border-emerald-200 p-6">
			<h2 class="text-2xl font-bold text-emerald-700">Settings</h2>
			<button
				onclick={() => (showMenu = false)}
				class="rounded-full p-2 transition-colors hover:bg-emerald-100"
				aria-label="Close menu"
			>
				<svg
					class="h-6 w-6 text-emerald-700"
					fill="none"
					stroke="currentColor"
					viewBox="0 0 24 24"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M6 18L18 6M6 6l12 12"
					></path>
				</svg>
			</button>
		</div>
		<BurgerMenu onClose={() => (showMenu = false)} />
	</div>
{/if}

<style>
	.pb-safe {
		padding-bottom: env(safe-area-inset-bottom);
	}

	.pt-safe {
		padding-top: env(safe-area-inset-top);
	}
</style>
