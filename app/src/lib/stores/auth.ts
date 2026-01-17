import { writable } from 'svelte/store';
import { browser } from '$app/environment';
import { getAuth, saveAuth, clearAuth } from '$lib/auth/auth';

export interface User {
	id: string;
	email: string;
	createdAt: string;
}

export interface AuthState {
	user: User | null;
	token: string | null;
	isAuthenticated: boolean;
	initialized: boolean;
}

function createAuthStore() {
	// Initialize with uninitialized state
	let initial: AuthState = {
		user: null,
		token: null,
		isAuthenticated: false,
		initialized: false
	};

	const { subscribe, set, update } = writable<AuthState>(initial);

	// Load auth from Capacitor preferences on browser init
	if (browser) {
		getAuth().then(({ token, user }) => {
			set({
				user,
				token,
				isAuthenticated: !!token,
				initialized: true
			});
		});
	}

	return {
		subscribe,
		login: async (user: User, token: string) => {
			await saveAuth(token, user);
			set({
				user,
				token,
				isAuthenticated: true,
				initialized: true
			});
		},
		logout: async () => {
			await clearAuth();
			set({
				user: null,
				token: null,
				isAuthenticated: false,
				initialized: true
			});
		},
		setUser: (user: User) => {
			update((state) => ({ ...state, user }));
		}
	};
}

export const authStore = createAuthStore();
