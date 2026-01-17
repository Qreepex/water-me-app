import { Preferences } from '@capacitor/preferences';

const AUTH_KEY = 'auth_token';
const USER_KEY = 'auth_user';

export interface User {
	id: string;
	email: string;
	createdAt: string;
}

export async function saveAuth(token: string, user: User): Promise<void> {
	await Promise.all([
		Preferences.set({
			key: AUTH_KEY,
			value: token
		}),
		Preferences.set({
			key: USER_KEY,
			value: JSON.stringify(user)
		})
	]);
}

export async function getAuth(): Promise<{ token: string | null; user: User | null }> {
	try {
		const [tokenResult, userResult] = await Promise.all([
			Preferences.get({ key: AUTH_KEY }),
			Preferences.get({ key: USER_KEY })
		]);

		const token = tokenResult.value;
		let user: User | null = null;

		if (userResult.value) {
			try {
				user = JSON.parse(userResult.value);
			} catch (e) {
				console.error('Failed to parse stored user:', e);
			}
		}

		return { token, user };
	} catch (e) {
		console.error('Failed to retrieve auth from preferences:', e);
		return { token: null, user: null };
	}
}

export async function clearAuth(): Promise<void> {
	await Promise.all([Preferences.remove({ key: AUTH_KEY }), Preferences.remove({ key: USER_KEY })]);
}

export async function isLoggedIn(): Promise<boolean> {
	const { token } = await getAuth();
	return !!token;
}
