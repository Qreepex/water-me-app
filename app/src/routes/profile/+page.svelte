<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { authStore } from '$lib/stores/auth';
	import { API_BASE_URL } from '$lib/constants';

	let user = $state($authStore.user);
	let token = $state($authStore.token);

	let username = $state(user?.username || '');
	let email = $state(user?.email || '');
	let currentPassword = $state('');
	let newPassword = $state('');
	let confirmPassword = $state('');

	let message = $state('');
	let messageType = $state<'success' | 'error' | ''>('');
	let loading = $state(false);

	let initialized = $state(false);

	onMount(() => {
		const checkAuth = setInterval(() => {
			if ($authStore.initialized) {
				clearInterval(checkAuth);
				initialized = true;

				if (!$authStore.isAuthenticated) {
					goto('/');
				}

				user = $authStore.user;
				token = $authStore.token;
				if (user) {
					username = user.username || '';
					email = user.email || '';
				}
			}
		}, 50);

		return () => clearInterval(checkAuth);
	});

	async function handleUpdateProfile() {
		if (!token) {
			message = 'Not authenticated';
			messageType = 'error';
			return;
		}

		loading = true;
		message = '';

		try {
			const response = await fetch(API_BASE_URL + '/api/user', {
				method: 'PUT',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${token}`
				},
				body: JSON.stringify({
					username,
					email
				})
			});

			if (!response.ok) {
				if (response.status === 401) {
					await authStore.logout();
					goto('/');
					return;
				}
				const error = await response.json();
				throw new Error(error.error || 'Failed to update profile');
			}

			const updatedUser = await response.json();
			await authStore.setUser(updatedUser);

			user = updatedUser;
			message = 'Profile updated successfully';
			messageType = 'success';
		} catch (err) {
			message = err instanceof Error ? err.message : 'An error occurred';
			messageType = 'error';
		} finally {
			loading = false;
			setTimeout(() => {
				message = '';
			}, 3000);
		}
	}

	async function handleChangePassword() {
		if (!token) {
			message = 'Not authenticated';
			messageType = 'error';
			return;
		}

		if (newPassword !== confirmPassword) {
			message = 'Passwords do not match';
			messageType = 'error';
			return;
		}

		if (newPassword.length < 6) {
			message = 'Password must be at least 6 characters';
			messageType = 'error';
			return;
		}

		loading = true;
		message = '';

		try {
			const response = await fetch(API_BASE_URL + '/api/user/password', {
				method: 'PUT',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${token}`
				},
				body: JSON.stringify({
					currentPassword,
					newPassword
				})
			});

			if (!response.ok) {
				if (response.status === 401) {
					const error = await response.json();
					throw new Error(error.error || 'Incorrect current password');
				}
				const error = await response.json();
				throw new Error(error.error || 'Failed to change password');
			}

			currentPassword = '';
			newPassword = '';
			confirmPassword = '';

			message = 'Password changed successfully';
			messageType = 'success';
		} catch (err) {
			message = err instanceof Error ? err.message : 'An error occurred';
			messageType = 'error';
		} finally {
			loading = false;
			setTimeout(() => {
				message = '';
			}, 3000);
		}
	}
</script>

{#if initialized}
	<div class="mx-auto max-w-2xl">
		<h1 class="mb-8 text-3xl font-bold text-emerald-900">User Profile</h1>

		{#if message}
			<div
				class={`mb-6 rounded-lg p-4 ${
					messageType === 'success'
						? 'border border-green-300 bg-green-100 text-green-800'
						: 'border border-red-300 bg-red-100 text-red-800'
				}`}
			>
				{message}
			</div>
		{/if}

		<div class="mb-8 rounded-lg border border-emerald-200 bg-white p-6 shadow-md">
			<h2 class="mb-6 text-xl font-bold text-emerald-900">Profile Information</h2>

			<div class="space-y-4">
				<div>
					<label for="username" class="mb-2 block text-sm font-semibold text-emerald-800">
						Username
					</label>
					<input
						id="username"
						type="text"
						bind:value={username}
						placeholder="Your username"
						class="w-full rounded-lg border border-emerald-300 px-4 py-2 focus:ring-2 focus:ring-emerald-500 focus:outline-none"
					/>
				</div>

				<div>
					<label for="email" class="mb-2 block text-sm font-semibold text-emerald-800">
						Email
					</label>
					<input
						id="email"
						type="email"
						bind:value={email}
						placeholder="your@email.com"
						class="w-full rounded-lg border border-emerald-300 px-4 py-2 focus:ring-2 focus:ring-emerald-500 focus:outline-none"
					/>
				</div>

				<button
					onclick={handleUpdateProfile}
					disabled={loading}
					class="w-full rounded-lg bg-emerald-600 py-2 font-semibold text-white transition-colors hover:bg-emerald-700 disabled:cursor-not-allowed disabled:opacity-50"
				>
					{loading ? 'Updating...' : 'Update Profile'}
				</button>
			</div>
		</div>

		<div class="rounded-lg border border-emerald-200 bg-white p-6 shadow-md">
			<h2 class="mb-6 text-xl font-bold text-emerald-900">Change Password</h2>

			<div class="space-y-4">
				<div>
					<label for="currentPassword" class="mb-2 block text-sm font-semibold text-emerald-800">
						Current Password
					</label>
					<input
						id="currentPassword"
						type="password"
						bind:value={currentPassword}
						placeholder="Enter current password"
						class="w-full rounded-lg border border-emerald-300 px-4 py-2 focus:ring-2 focus:ring-emerald-500 focus:outline-none"
					/>
				</div>

				<div>
					<label for="newPassword" class="mb-2 block text-sm font-semibold text-emerald-800">
						New Password
					</label>
					<input
						id="newPassword"
						type="password"
						bind:value={newPassword}
						placeholder="Enter new password"
						class="w-full rounded-lg border border-emerald-300 px-4 py-2 focus:ring-2 focus:ring-emerald-500 focus:outline-none"
					/>
				</div>

				<div>
					<label for="confirmPassword" class="mb-2 block text-sm font-semibold text-emerald-800">
						Confirm New Password
					</label>
					<input
						id="confirmPassword"
						type="password"
						bind:value={confirmPassword}
						placeholder="Confirm new password"
						class="w-full rounded-lg border border-emerald-300 px-4 py-2 focus:ring-2 focus:ring-emerald-500 focus:outline-none"
					/>
				</div>

				<button
					onclick={handleChangePassword}
					disabled={loading}
					class="w-full rounded-lg bg-emerald-600 py-2 font-semibold text-white transition-colors hover:bg-emerald-700 disabled:cursor-not-allowed disabled:opacity-50"
				>
					{loading ? 'Changing...' : 'Change Password'}
				</button>
			</div>
		</div>
	</div>
{/if}
