<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import BurgerMenu from '$lib/components/BurgerMenu.svelte';
	import { initializePushNotifications, cleanupPushNotifications, getNotificationToken } from '$lib/notifications';
	import { browser } from '$app/environment';

	let { children } = $props();
	let fcmToken = $state<string | null>(null);

	onMount(async () => {
		if (browser) {
			// Initialize push notifications
			const result = await initializePushNotifications();
			fcmToken = result.token;
			
			if (fcmToken) {
				console.log('✅ FCM Token registered:', fcmToken);
			}
		}
	});

	onDestroy(() => {
		if (browser) {
			cleanupPushNotifications();
		}
	});
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<div class="min-h-screen bg-gradient-to-br from-emerald-50 to-green-50 relative">
	<!-- Floating Burger Menu -->
	<div class="fixed top-0 right-0 z-50 pt-safe pr-safe">
		<div class="p-4">
			<BurgerMenu />
		</div>
	</div>

	<main class="px-4 pt-safe pb-safe">
		{@render children()}
	</main>
</div>

<style>
	/* Safe area insets for mobile notches and status bars */
	.pt-safe {
		padding-top: env(safe-area-inset-top);
	}
	.pr-safe {
		padding-right: env(safe-area-inset-right);
	}
	.pb-safe {
		padding-bottom: env(safe-area-inset-bottom);
	}
	.pl-safe {
		padding-left: env(safe-area-inset-left);
	}
</style>
