import { writable, derived, get } from 'svelte/store';
import { Preferences } from '@capacitor/preferences';
import { authStore } from './auth';

type Language = 'en' | 'de' | 'es';

export const languageStore = writable<Language>('en');

// Derived store for the current t() function
export const currentLanguage = derived(languageStore, ($language) => $language);

// Initialize language from user profile or preferences
export async function initializeLanguage() {
	try {
		// First, check user profile if logged in
		const auth = get(authStore);
		if (auth.user?.language) {
			languageStore.set(auth.user.language as Language);
			return;
		}

		// Otherwise, check preferences
		const stored = await Preferences.get({ key: 'language' });
		if (stored.value) {
			languageStore.set(stored.value as Language);
		} else {
			// Default to English when no user is logged in
			languageStore.set('en');
		}
	} catch (error) {
		console.error('Failed to initialize language:', error);
		languageStore.set('en');
	}
}

// Update language and persist
export async function setLanguage(language: Language) {
	languageStore.set(language);
	try {
		await Preferences.set({ key: 'language', value: language });
	} catch (error) {
		console.error('Failed to save language preference:', error);
	}
}

// Subscribe to auth changes and sync language
authStore.subscribe((auth) => {
	if (auth.user?.language) {
		languageStore.set(auth.user.language as Language);
	}
});
