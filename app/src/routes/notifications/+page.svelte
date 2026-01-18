<script lang="ts">
	import { onMount } from 'svelte';
	import { requestNotificationPermissions, getNotificationState } from '$lib/notifications';
	import { tStore } from '$lib/i18n';

	let state = getNotificationState();
	let requesting = false;
	let message: string | null = null;

	function refreshState() {
		state = getNotificationState();
	}

	async function enableNotifications() {
		requesting = true;
		message = null;
		try {
			await requestNotificationPermissions();
			refreshState();
			if (state.isRegistered && state.token) {
				message = 'Notifications enabled';
			} else {
				message = 'Permission denied or registration failed';
			}
		} catch (e) {
			console.error(e);
			message = 'Error while enabling notifications';
		} finally {
			requesting = false;
		}
	}

	onMount(() => {
		refreshState();
	});
</script>

<div class="min-h-screen bg-gradient-to-br from-green-50 via-emerald-50 to-teal-50 p-8">
	<div class="mx-auto max-w-3xl">
		<div class="mb-8">
			<h1 class="text-4xl font-bold text-green-800">
				{$tStore('menu.notifications') || 'Notifications'}
			</h1>
			<p class="text-green-700">Configure your notification preferences</p>
		</div>

		<div class="space-y-6">
			<!-- Permission CTA -->
			<div class="rounded-2xl bg-white p-6 shadow">
				<h2 class="mb-2 text-xl font-semibold text-green-800">Enable Push Notifications</h2>
				<p class="mb-4 text-green-700">Ask for permission and register your device.</p>
				<button
					disabled={requesting}
					on:click={enableNotifications}
					class="rounded-lg bg-green-600 px-5 py-2 text-white transition hover:bg-green-700 disabled:opacity-70"
				>
					{requesting ? 'Requesting…' : 'Enable Notifications'}
				</button>
				{#if message}
					<p class="mt-3 text-sm text-green-800">{message}</p>
				{/if}
				{#if state.token}
					<p class="mt-3 text-xs text-gray-600">Token: {state.token}</p>
				{/if}
			</div>

			<!-- Placeholder settings -->
			<div class="rounded-2xl bg-white p-6 shadow">
				<h2 class="mb-2 text-xl font-semibold text-green-800">Preferences (Placeholder)</h2>
				<div class="space-y-3 text-green-800">
					<label class="flex items-center gap-3">
						<input type="checkbox" />
						<span>Watering reminders</span>
					</label>
					<label class="flex items-center gap-3">
						<input type="checkbox" />
						<span>Fertilizing reminders</span>
					</label>
					<label class="flex items-center gap-3">
						<input type="checkbox" />
						<span>Spray reminders</span>
					</label>
				</div>
			</div>
		</div>
	</div>
</div>
