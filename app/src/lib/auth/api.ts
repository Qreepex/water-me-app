import { getToken } from './auth';

export async function fetchWithAuth(url: string, options: RequestInit = {}) {
	const token = await getToken();
	const headers = {
		...options.headers,
		Authorization: token ? `Bearer ${token}` : ''
	};

	const response = await fetch(url, {
		...options,
		headers
	});

	return response;
}
