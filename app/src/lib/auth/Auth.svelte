<script lang="ts">
	import Spinner from '$lib/components/Spinner.svelte';
	import { FirebaseAuthentication } from '@capacitor-firebase/authentication';
	import { Capacitor } from '@capacitor/core';
	import { initializeApp } from 'firebase/app';
	import {
		getAuth,
		setPersistence,
		browserLocalPersistence,
		onAuthStateChanged,
		type User
	} from 'firebase/auth';
	import { onMount } from 'svelte';
	import { FIREBASE_CONFIG } from './firebase';
	import { tStore } from '$lib/i18n';
	import Button from '$lib/components/ui/Button.svelte';
	import Message from '$lib/components/ui/Message.svelte';
	import { fetchData } from '$lib/auth/fetch.svelte';
	import { getPlantsStore } from '$lib/stores/plants.svelte';
	import { imageCacheStore } from '$lib/stores/imageCache.svelte';
	import type { Plant } from '$lib/types/api';
	import type { Snippet } from 'svelte';

	interface Props {
		children: Snippet;
	}

	const { children }: Props = $props();

	const app = initializeApp(FIREBASE_CONFIG);
	const auth = getAuth(app);

	let user = $state<User | null>(null);
	let loading = $state(false);
	let initializing = $state(true);
	let error = $state<string | null>(null);
	let loadingImages = $state(false);

	const platform = Capacitor.getPlatform();
	const store = getPlantsStore();
	let hasLoadedPlants = $state(false);

	// Load plants data when user is authenticated
	$effect(() => {
		if (user && !hasLoadedPlants) {
			hasLoadedPlants = true;
			loadPlants();
		}
	});

	async function loadPlants(): Promise<void> {
		store.setLoading(true);
		loadingImages = true;
		store.setError(null);

		try {
			const result = await fetchData('/api/plants', {});

			if (!result.ok) {
				const errorMsg = result.error?.message || $tStore('plants.failedToFetchPlants');
				store.setError(errorMsg);
				return;
			}

			const plants = result.data || [];
			store.setPlants(plants);

			// Prefetch all plant images into cache
			await prefetchAllImages(plants);
		} catch (err) {
			const errorMsg = err instanceof Error ? err.message : 'Unknown error';
			store.setError(errorMsg);
		} finally {
			store.setLoading(false);
			loadingImages = false;
		}
	}

	async function prefetchAllImages(plants: Plant[]): Promise<void> {
		const prefetchPromises: Promise<void>[] = [];

		for (const plant of plants) {
			const photoIds = plant.photoIds || [];
			// eslint-disable-next-line @typescript-eslint/no-explicit-any
			const photoUrls = ((plant as any)?.photoUrls as string[] | undefined) || [];

			for (let i = 0; i < photoIds.length; i++) {
				const photoId = photoIds[i];
				const photoUrl = photoUrls[i];

				if (photoId && photoUrl) {
					// Fire and forget - load all images into persistent cache
					prefetchPromises.push(
						imageCacheStore
							.getImageURL(photoId, photoUrl)
							.then(() => {
								// Image now in cache
							})
							.catch(() => {
								// Ignore errors during prefetch
							})
					);
				}
			}
		}

		// Wait for all images to be prefetched (or fail gracefully)
		await Promise.allSettled(prefetchPromises);
	}

	onMount(async () => {
		if (platform === 'web') {
			setPersistence(auth, browserLocalPersistence).catch(console.error);

			const unsubscribe = onAuthStateChanged(auth, (firebaseUser) => {
				user = firebaseUser ?? null;
				initializing = false;
			});

			return unsubscribe;
		} else {
			const result = await FirebaseAuthentication.getCurrentUser();
			user = result.user ?? null;
			initializing = false;

			const listener = await FirebaseAuthentication.addListener('authStateChange', (res) => {
				user = res.user;
			});

			return async () => listener.remove();
		}
	});

	async function loginWithGoogle() {
		try {
			loading = true;
			error = null;

			const result = await FirebaseAuthentication.signInWithGoogle();

			// Check if user cancelled (result exists but no user)
			if (!result || !result.user) {
				error = 'auth.signInCancelled';
				return;
			}

			user = result.user;

			const idToken = await FirebaseAuthentication.getIdToken();
			// eslint-disable-next-line @typescript-eslint/no-explicit-any
		} catch (err: any) {
			console.error('Login fehlgeschlagen', err);
			// Check if user cancelled the popup
			if (
				err?.message?.includes('popup_closed_by_user') ||
				err?.code === 'popup-closed-by-user' ||
				err?.message?.includes('cancelled')
			) {
				error = 'auth.signInCancelled';
			} else {
				error = 'auth.signInError';
			}
		} finally {
			loading = false;
		}
	}
</script>

{#if initializing || loadingImages}
	<Spinner />
{:else if user}
	{@render children()}
{:else}
	<div
		class="flex h-full items-center justify-center bg-gradient-to-br from-green-50 via-emerald-50 to-teal-50 p-4"
	>
		<div class="w-full max-w-md">
			<!-- Logo/Title -->
			<div class="mb-8 text-center">
				<h1 class="mb-2 text-5xl font-bold text-green-800">{$tStore('common.app')}</h1>
				<p class="text-green-700">{$tStore('common.appDescription')}</p>
			</div>

			<!-- Card -->
			<div class="rounded-2xl bg-white p-8 shadow-lg">
				<!-- Mode Indicator -->
				<div class="mb-8">
					<h2 class="mb-4 text-2xl font-bold text-green-800">{$tStore('auth.signIn')}</h2>
					<p class="text-gray-600">{$tStore('auth.signInToContinue')}</p>
				</div>

				<!-- Error Message -->
				{#if error}
					<Message message={error} type="error" />
				{/if}

				{#if loading}
					<Message title="auth.signingIn" />
				{/if}

				<!-- Submit Button -->
				<Button
					disabled={loading}
					{loading}
					onclick={loginWithGoogle}
					text="auth.signInWithGoogle"
					loadingText="auth.signingIn"
				/>
			</div>
		</div>
	</div>
{/if}
