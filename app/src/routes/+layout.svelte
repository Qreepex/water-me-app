<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import { initializeI18n } from '$lib/i18n';
	import { initializeLanguage } from '$lib/stores/language';
	import Auth from '$lib/auth/Auth.svelte';
	import { Capacitor } from '@capacitor/core';
	import { SplashScreen } from '@capacitor/splash-screen';
	import BottomNav from '$lib/components/BottomNav.svelte';

	let { children } = $props();

	onMount(async () => {
		// hide splash screen once the app is ready
		if (Capacitor.isNativePlatform()) {
			await SplashScreen.hide();
		}

		// Initialize language from user profile or preferences
		await initializeLanguage();

		// Initialize i18n translations for the selected language
		await initializeI18n();

		// Register service worker for image caching (web only)
		if (typeof navigator !== 'undefined' && 'serviceWorker' in navigator) {
			try {
				await navigator.serviceWorker.register('/sw.js');
			} catch {
				// ignore
			}
		}

		// Do not auto-request notification permissions on startup.
		// Use $lib/notifications.requestNotificationPermissions() when user opts in.
	});

	onDestroy(() => {
		// no-op for now
	});
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<div class="fixed inset-0 bg-gradient-to-br from-emerald-50 to-green-50">
	<main class="pt-safe h-full overflow-y-auto pb-32" style="overscroll-behavior-y: none;">
		<Auth>
			<!-- Bottom Navigation: only visible when authenticated (inside Auth slot) -->
			<BottomNav />
			{@render children()}
		</Auth>
	</main>
</div>

<style>
	/* Safe area insets for mobile notches and status bars */
	.pt-safe {
		padding-top: env(safe-area-inset-top);
	}
</style>
