// Persistent in-memory cache for object URLs to prevent reloading on navigation
// This keeps the object URLs alive across page transitions

import { getImageObjectURL as fetchImageObjectURL } from '$lib/utils/imageCache';

type ImageCacheEntry = {
	url: string;
	timestamp: number;
	refCount: number;
};

class ImageCacheStore {
	private cache = $state<Map<string, ImageCacheEntry>>(new Map());
	private loadingPromises = new Map<string, Promise<string | null>>();

	// Synchronous URL lookup - returns immediately if available
	getImageURLSync(photoId: string): string | null {
		const existing = this.cache.get(photoId);
		if (existing) {
			existing.refCount++;
			existing.timestamp = Date.now();
			return existing.url;
		}
		return null;
	}

	async getImageURL(photoId: string, remoteUrl?: string): Promise<string | null> {
		// Check if we already have it cached
		const existing = this.cache.get(photoId);
		if (existing) {
			existing.refCount++;
			existing.timestamp = Date.now();
			return existing.url;
		}

		// Check if we're already loading it
		const loadingPromise = this.loadingPromises.get(photoId);
		if (loadingPromise) {
			return loadingPromise;
		}

		// Load the image
		const promise = fetchImageObjectURL(photoId, remoteUrl).then((url) => {
			this.loadingPromises.delete(photoId);
			if (url) {
				this.cache.set(photoId, {
					url,
					timestamp: Date.now(),
					refCount: 1
				});
			}
			return url;
		});

		this.loadingPromises.set(photoId, promise);
		return promise;
	}

	releaseImage(photoId: string): void {
		const entry = this.cache.get(photoId);
		if (entry) {
			entry.refCount--;
			// Don't actually revoke or remove - keep it cached for reuse
			// We could implement LRU cleanup here if needed
		}
	}

	// Optional: cleanup old entries that haven't been used in a while
	cleanup(maxAge: number = 30 * 60 * 1000): void {
		const now = Date.now();
		for (const [photoId, entry] of this.cache.entries()) {
			if (entry.refCount === 0 && now - entry.timestamp > maxAge) {
				// Only revoke if it's a blob: URL (browser)
				if (entry.url.startsWith('blob:')) {
					URL.revokeObjectURL(entry.url);
				}
				this.cache.delete(photoId);
			}
		}
	}
}

export const imageCacheStore = new ImageCacheStore();
