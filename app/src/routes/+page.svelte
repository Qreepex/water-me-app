<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { authStore } from '$lib/stores/auth';
	import type { User } from '$lib/auth/auth';
	import {
		API_BASE_URL,
		IMPRINT_URL,
		PRIVACY_POLICY_URL,
		SOURCE_CODE_URL,
		WEBSITE_URL
	} from '$lib/constants';
	import { tStore } from '$lib/i18n';
	import { resolve } from '$app/paths';
	import { openExternalLink } from '$lib/os/browser';
	import { SplashScreen } from '@capacitor/splash-screen';

	let mode: 'login' | 'signup' = 'login';
	let email = '';
	let password = '';
	let loading = false;
	let error: string | null = null;

	async function hideSplashScreen() {
		await SplashScreen.hide();
	}

	onMount(() => {
		hideSplashScreen();

		// Wait for auth store to initialize from Capacitor preferences
		const unsubscribe = authStore.subscribe((state) => {
			if (state.isAuthenticated) {
				goto(resolve('/app'));
			}
		});

		return unsubscribe;
	});

	async function handleSubmit(e: Event) {
		e.preventDefault();
		loading = true;
		error = null;

		try {
			const endpoint = mode === 'login' ? '/api/login' : '/api/signup';
			const response = await fetch(API_BASE_URL + `${endpoint}`, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ email, password })
			});

			const data = await response.json();

			if (!response.ok) {
				error = data.error || 'Authentication failed';
				return;
			}

			// Store auth data and redirect
			const user: User = data.user;
			const token: string = data.token;

			authStore.login(user, token);
			goto(resolve('/app'));
		} catch (err) {
			error = err instanceof Error ? err.message : 'An error occurred';
		} finally {
			loading = false;
		}
	}

	function toggleMode() {
		mode = mode === 'login' ? 'signup' : 'login';
		error = null;
	}
</script>

<div
	class="flex min-h-screen items-center justify-center bg-gradient-to-br from-green-50 via-emerald-50 to-teal-50 p-4"
>
	<div class="w-full max-w-md">
		<!-- Logo/Title -->
		<div class="mb-8 text-center">
			<h1 class="mb-2 text-5xl font-bold text-green-800">{$tStore('app')}</h1>
			<p class="text-green-700">{$tStore('appDescription')}</p>
		</div>

		<!-- Card -->
		<div class="rounded-2xl bg-white p-8 shadow-lg">
			<!-- Mode Indicator -->
			<div class="mb-8">
				<h2 class="mb-4 text-2xl font-bold text-green-800">
					{mode === 'login' ? $tStore('auth.signIn') : $tStore('auth.signUp')}
				</h2>
				<p class="text-gray-600">
					{mode === 'login' ? $tStore('auth.subtitle') : $tStore('auth.createAccountSubtitle')}
				</p>
			</div>

			<!-- Form -->
			<form onsubmit={handleSubmit} class="space-y-5">
				<!-- Email Input -->
				<div>
					<label for="email" class="mb-2 block text-sm font-semibold text-green-800"
						>{$tStore('auth.email')}</label
					>
					<input
						type="email"
						id="email"
						bind:value={email}
						placeholder="you@example.com"
						autocomplete="email"
						required
						disabled={loading}
						class="w-full rounded-lg border-2 border-green-300 px-4 py-3 transition hover:border-green-400 focus:border-green-500 focus:outline-none disabled:bg-gray-100"
					/>
				</div>

				<!-- Password Input -->
				<div>
					<label for="password" class="mb-2 block text-sm font-semibold text-green-800">
						{$tStore('auth.password')}
					</label>
					<input
						type="password"
						autocomplete={mode === 'signup' ? 'new-password' : 'current-password'}
						id="password"
						bind:value={password}
						placeholder={mode === 'signup' ? 'At least 6 characters' : '••••••'}
						required
						disabled={loading}
						class="w-full rounded-lg border-2 border-green-300 px-4 py-3 transition hover:border-green-400 focus:border-green-500 focus:outline-none disabled:bg-gray-100"
					/>
				</div>

				<!-- Error Message -->
				{#if error}
					<div class="rounded-lg border-2 border-red-400 bg-red-100 px-4 py-3 text-red-800">
						<p class="text-sm font-semibold">{error}</p>
					</div>
				{/if}

				<!-- Submit Button -->
				<button
					type="submit"
					disabled={loading}
					class="w-full rounded-lg bg-green-600 px-4 py-3 font-semibold text-white transition hover:bg-green-700 focus:ring-2 focus:ring-green-500 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50"
				>
					{#if loading}
						<span class="inline-flex items-center gap-2">
							<svg
								class="h-4 w-4 animate-spin"
								xmlns="http://www.w3.org/2000/svg"
								fill="none"
								viewBox="0 0 24 24"
							>
								<circle
									class="opacity-25"
									cx="12"
									cy="12"
									r="10"
									stroke="currentColor"
									stroke-width="4"
								></circle>
								<path
									class="opacity-75"
									fill="currentColor"
									d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
								></path>
							</svg>
							{mode === 'login' ? $tStore('auth.signingIn') : $tStore('auth.creatingAccount')}
						</span>
					{:else}
						{mode === 'login' ? $tStore('auth.signIn') : $tStore('auth.signUp')}
					{/if}
				</button>
			</form>

			<!-- Toggle Mode -->
			<div class="mt-6 text-center">
				<p class="text-gray-600">
					{mode === 'login' ? $tStore('auth.dontHaveAccount') : $tStore('auth.alreadyHaveAccount')}
					<button
						type="button"
						onclick={toggleMode}
						class="font-semibold text-green-600 transition hover:text-green-700"
					>
						{mode === 'login' ? $tStore('auth.signUp') : $tStore('auth.signIn')}
					</button>
				</p>
			</div>
		</div>

		<!-- Footer -->
		<div class="mt-8 text-center text-sm text-gray-600">
			<p>{$tStore('common.madeWith')}</p>
			<p>
				<button class="underline" onclick={() => openExternalLink(WEBSITE_URL)}
					>{$tStore('menu.website')}</button
				>
				|
				<button class="underline" onclick={() => openExternalLink(SOURCE_CODE_URL)}
					>{$tStore('menu.sourceCode')}</button
				>
				|
				<button class="underline" onclick={() => openExternalLink(PRIVACY_POLICY_URL)}
					>{$tStore('menu.privacyPolicy')}</button
				>
				|
				<button class="underline" onclick={() => openExternalLink(IMPRINT_URL)}
					>{$tStore('menu.imprint')}</button
				>
			</p>
		</div>
	</div>
</div>

<style>
	:global(body) {
		font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
	}
</style>
