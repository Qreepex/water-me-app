// Unified image cache: prefers Capacitor Filesystem on native (Android),
// falls back to IndexedDB in browsers. Provides object/file URLs for rendering.

import { Capacitor } from '@capacitor/core';
import { Filesystem, Directory } from '@capacitor/filesystem';
import { Preferences } from '@capacitor/preferences';

const DB_NAME = 'plants-image-cache';
const STORE_NAME = 'images';
const DB_VERSION = 1;
const MAX_AGE_MS = 7 * 24 * 60 * 60 * 1000; // 7 days
const NATIVE_DIR = Directory.Data;
const INDEX_KEY = 'image-cache-index';

type CacheIndex = Record<string, number>; // photoId -> storedAt

type CacheEntry = {
	id: string; // photoId
	blob: Blob;
	storedAt: number; // epoch ms
};

function isBrowser(): boolean {
	return typeof window !== 'undefined';
}

function isNative(): boolean {
	return Capacitor.isNativePlatform();
}

async function getIndex(): Promise<CacheIndex> {
	try {
		const { value } = await Preferences.get({ key: INDEX_KEY });
		return value ? (JSON.parse(value) as CacheIndex) : {};
	} catch {
		return {};
	}
}

async function setIndex(index: CacheIndex): Promise<void> {
	try {
		await Preferences.set({ key: INDEX_KEY, value: JSON.stringify(index) });
	} catch {
		// ignore
	}
}

async function blobToBase64(blob: Blob): Promise<string> {
	const buf = await blob.arrayBuffer();
	const bytes = new Uint8Array(buf);
	let binary = '';
	for (let i = 0; i < bytes.length; i++) binary += String.fromCharCode(bytes[i]);
	return btoa(binary);
}

function nativePathFor(photoId: string): string {
	return `images/${photoId}.webp`;
}

async function getNativeUri(photoId: string): Promise<string | null> {
	try {
		const path = nativePathFor(photoId);
		await Filesystem.stat({ path, directory: NATIVE_DIR });
		const { uri } = await Filesystem.getUri({ path, directory: NATIVE_DIR });
		return Capacitor.convertFileSrc(uri);
	} catch {
		return null;
	}
}

async function putNativeBlob(photoId: string, blob: Blob): Promise<string | null> {
	try {
		const b64 = await blobToBase64(blob);
		const path = nativePathFor(photoId);

		// Ensure all parent directories exist
		try {
			await Filesystem.mkdir({ path: 'images', directory: NATIVE_DIR, recursive: true });
		} catch {
			// Directory might already exist, ignore
		}

		// Create any nested directories if needed
		const parts = path.split('/');
		for (let i = 1; i < parts.length - 1; i++) {
			const dir = parts.slice(0, i + 1).join('/');
			try {
				await Filesystem.mkdir({ path: dir, directory: NATIVE_DIR, recursive: true });
			} catch {
				// Directory might already exist, ignore
			}
		}

		await Filesystem.writeFile({ path, directory: NATIVE_DIR, data: b64 });
		const index = await getIndex();
		index[photoId] = Date.now();
		await setIndex(index);
		const { uri } = await Filesystem.getUri({ path, directory: NATIVE_DIR });
		return Capacitor.convertFileSrc(uri);
	} catch (err) {
		console.error('Failed to cache image natively:', err);
		return null;
	}
}

function openDB(): Promise<IDBDatabase> {
	return new Promise((resolve, reject) => {
		if (!isBrowser() || typeof indexedDB === 'undefined') {
			reject(new Error('IndexedDB unavailable'));
			return;
		}
		const req = indexedDB.open(DB_NAME, DB_VERSION);
		req.onupgradeneeded = () => {
			const db = req.result;
			if (!db.objectStoreNames.contains(STORE_NAME)) {
				db.createObjectStore(STORE_NAME, { keyPath: 'id' });
			}
		};
		req.onsuccess = () => resolve(req.result);
		req.onerror = () => reject(req.error);
	});
}

async function getEntry(db: IDBDatabase, id: string): Promise<CacheEntry | null> {
	return new Promise((resolve, reject) => {
		const tx = db.transaction(STORE_NAME, 'readonly');
		const store = tx.objectStore(STORE_NAME);
		const req = store.get(id);
		req.onsuccess = () => resolve((req.result as CacheEntry) || null);
		req.onerror = () => reject(req.error);
	});
}

async function putEntry(db: IDBDatabase, entry: CacheEntry): Promise<void> {
	return new Promise((resolve, reject) => {
		const tx = db.transaction(STORE_NAME, 'readwrite');
		const store = tx.objectStore(STORE_NAME);
		const req = store.put(entry);
		req.onsuccess = () => resolve();
		req.onerror = () => reject(req.error);
	});
}

export async function getImageObjectURL(photoId: string, url?: string): Promise<string | null> {
	try {
		const index = await getIndex();
		const storedAt = index[photoId] ?? 0;

		if (isNative()) {
			const nativeUri = await getNativeUri(photoId);
			if (nativeUri && Date.now() - storedAt < MAX_AGE_MS) return nativeUri;
			if (!url) return nativeUri;
			const res = await fetch(url);
			if (!res.ok) return nativeUri;
			const blob = await res.blob();
			const newUri = await putNativeBlob(photoId, blob);
			return newUri || nativeUri || URL.createObjectURL(blob);
		}

		// Browser fallback: IndexedDB
		if (!isBrowser()) return url ?? null;
		const db = await openDB();
		const existing = await getEntry(db, photoId);
		if (existing && Date.now() - existing.storedAt < MAX_AGE_MS) {
			return URL.createObjectURL(existing.blob);
		}
		if (!url) return null;
		const res = await fetch(url);
		if (!res.ok) return null;
		const blob = await res.blob();
		await putEntry(db, { id: photoId, blob, storedAt: Date.now() });
		return URL.createObjectURL(blob);
	} catch {
		return url ?? null;
	}
}

export function revokeObjectURL(url?: string): void {
	try {
		if (url && isBrowser()) URL.revokeObjectURL(url);
	} catch {
		// ignore
	}
}
