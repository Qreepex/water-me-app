<script lang="ts">
	import { onMount } from 'svelte';
	import { getNotificationState, checkNotificationPermissions } from '$lib/notifications';
	import { Preferences } from '@capacitor/preferences';
	import type { NotificationState } from '$lib/notifications';

	let notificationState = $state<NotificationState>({
		token: null,
		isRegistered: false,
		isSupported: false
	});
	let permissionGranted = $state(false);
	let showToken = $state(false);

	onMount(async () => {
		// Get current notification state
		notificationState = getNotificationState();

		// Check if token is in Capacitor Preferences (in case it was registered before)
		if (!notificationState.token) {
			try {
				const { value } = await Preferences.get({ key: 'fcm_token' });
				if (value) {
					notificationState.token = value;
				}
			} catch (err) {
				console.error('Failed to get stored token:', err);
			}
		}

		// Check permissions
		permissionGranted = await checkNotificationPermissions();
	});

	function copyToken() {
		if (notificationState.token) {
			navigator.clipboard.writeText(notificationState.token);
			alert('Token copied to clipboard!');
		}
	}
</script>

<div class="rounded-lg border border-emerald-200 bg-white p-6 shadow-md">
	<h3 class="mb-4 flex items-center gap-2 text-lg font-bold text-emerald-900">
		<svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
			<path
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
				d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
			></path>
		</svg>
		Push Notifications
	</h3>

	<div class="space-y-3">
		<!-- Status -->
		<div class="grid grid-cols-2 gap-3 text-sm">
			<div>
				<p class="font-semibold text-gray-700">Supported:</p>
				<p class={notificationState.isSupported ? 'text-green-600' : 'text-gray-500'}>
					{notificationState.isSupported ? '✓ Yes' : '✗ No (Web only)'}
				</p>
			</div>
			<div>
				<p class="font-semibold text-gray-700">Permission:</p>
				<p class={permissionGranted ? 'text-green-600' : 'text-red-600'}>
					{permissionGranted ? '✓ Granted' : '✗ Denied'}
				</p>
			</div>
			<div>
				<p class="font-semibold text-gray-700">Registered:</p>
				<p class={notificationState.isRegistered ? 'text-green-600' : 'text-gray-500'}>
					{notificationState.isRegistered ? '✓ Yes' : '✗ No'}
				</p>
			</div>
			<div>
				<p class="font-semibold text-gray-700">Token:</p>
				<p class={notificationState.token ? 'text-green-600' : 'text-gray-500'}>
					{notificationState.token ? '✓ Available' : '✗ None'}
				</p>
			</div>
		</div>

		<!-- FCM Token -->
		{#if notificationState.token}
			<div class="border-t border-emerald-200 pt-3">
				<div class="mb-2 flex items-center justify-between">
					<p class="text-sm font-semibold text-gray-700">FCM Token:</p>
					<button
						onclick={() => (showToken = !showToken)}
						class="text-xs font-medium text-emerald-600 hover:text-emerald-700"
					>
						{showToken ? 'Hide' : 'Show'}
					</button>
				</div>

				{#if showToken}
					<div class="rounded border border-gray-200 bg-gray-50 p-3">
						<p class="font-mono text-xs break-all text-gray-700">
							{notificationState.token}
						</p>
						<button
							onclick={copyToken}
							class="mt-2 w-full rounded-lg bg-emerald-600 px-3 py-2 text-sm text-white transition-colors hover:bg-emerald-700"
						>
							📋 Copy Token
						</button>
					</div>
				{/if}
			</div>

			<!-- Instructions -->
			<div class="border-t border-emerald-200 pt-3">
				<p class="mb-2 text-xs text-gray-600">
					<strong>To send test notification:</strong>
				</p>
				<ol class="list-inside list-decimal space-y-1 text-xs text-gray-600">
					<li>Copy the token above</li>
					<li>Go to Firebase Console → Cloud Messaging</li>
					<li>Click "Send test message"</li>
					<li>Paste the token and send</li>
				</ol>
			</div>
		{:else}
			<div class="border-t border-emerald-200 pt-3">
				<p class="text-xs text-gray-600">
					{#if !notificationState.isSupported}
						Push notifications are only available on native platforms (iOS/Android).
					{:else if !permissionGranted}
						Please grant notification permissions to receive push notifications.
					{:else}
						Waiting for FCM token registration...
					{/if}
				</p>
			</div>
		{/if}
	</div>
</div>
