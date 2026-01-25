import { PushNotifications } from '@capacitor/push-notifications';
import type { Token, PushNotificationSchema, ActionPerformed } from '@capacitor/push-notifications';
import { Capacitor } from '@capacitor/core';
import { Preferences } from '@capacitor/preferences';

export interface NotificationState {
	token: string | null;
	isRegistered: boolean;
	isSupported: boolean;
}

const notificationState: NotificationState = {
	token: null,
	isRegistered: false,
	isSupported: false
};

let listenersAdded = false;

/**
 * Initialize push notifications
 * Request permissions, register for notifications, and set up listeners
 */
async function addPushListeners() {
	if (listenersAdded) return;

	// Listen for registration success
	await PushNotifications.addListener('registration', async (token: Token) => {
		console.log('Push registration success, token:', token.value);
		notificationState.token = token.value;
		notificationState.isRegistered = true;

		// Store token in Capacitor Preferences for debugging
		try {
			await Preferences.set({ key: 'fcm_token', value: token.value });
		} catch (err) {
			console.error('Failed to store token:', err);
		}
	});

	// Listen for registration errors
	await PushNotifications.addListener('registrationError', (error: unknown) => {
		console.error('Push registration error:', error);

		// Check for common Firebase initialization error
		if (error && typeof error === 'object' && 'error' in error) {
			const errorMsg = (error as Record<string, unknown>).error;
			if (typeof errorMsg === 'string' && errorMsg.includes('FirebaseApp is not initialized')) {
				console.error('❌ Firebase Configuration Missing!');
			}
		}
	});

	// Listen for push notifications received while app is in foreground
	await PushNotifications.addListener(
		'pushNotificationReceived',
		(notification: PushNotificationSchema) => {
			console.log('Push notification received (foreground):', notification);
			console.log(`Title: ${notification.title}, Body: ${notification.body}`);
		}
	);

	// Listen for push notifications when user taps on them
	await PushNotifications.addListener(
		'pushNotificationActionPerformed',
		(action: ActionPerformed) => {
			console.log('Push notification action performed:', action);
			const data = action.notification.data;
			console.log('Notification data:', data);
		}
	);

	listenersAdded = true;
	console.log('Push notification listeners added');
}

export async function initializePushNotifications(): Promise<NotificationState> {
	// Check if push notifications are supported on this platform
	if (!Capacitor.isNativePlatform()) {
		return notificationState;
	}

	notificationState.isSupported = true;

	try {
		await addPushListeners();

		// Request permission to use push notifications
		const permStatus = await PushNotifications.requestPermissions();

		if (permStatus.receive === 'granted') {
			console.log('Push notification permission granted');

			// Register with Apple / Google to receive push notifications
			try {
				await PushNotifications.register();
			} catch (registerError: unknown) {
				console.error('Failed to register for push notifications:', registerError);

				// Provide helpful error message for Firebase setup
				if (registerError && typeof registerError === 'object' && 'message' in registerError) {
					const msgVal = (registerError as Record<string, unknown>).message;
					if (typeof msgVal === 'string' && msgVal.includes('FirebaseApp is not initialized')) {
						console.error('❌ Firebase Configuration Missing!');
					}
				}

				return notificationState;
			}
		} else {
			console.log('Push notification permission denied');
			return notificationState;
		}

		console.log('Push notifications initialized successfully');
		return notificationState;
	} catch (error) {
		console.error('Error initializing push notifications:', error);
		return notificationState;
	}
}

/**
 * Request notification permissions and register without auto-triggering on app start.
 * Call this when the user opts in (e.g., opens notifications page or creates a plant).
 */
export async function requestNotificationPermissions(): Promise<NotificationState> {
	if (!Capacitor.isNativePlatform()) {
		return notificationState;
	}

	notificationState.isSupported = true;

	try {
		await addPushListeners();

		const permStatus = await PushNotifications.requestPermissions();
		if (permStatus.receive === 'granted') {
			try {
				await PushNotifications.register();
				notificationState.isRegistered = true;
			} catch (registerError: unknown) {
				console.error('Failed to register for push notifications:', registerError);
				return notificationState;
			}
		} else {
			console.log('User denied push notification permissions');
		}

		return notificationState;
	} catch (error) {
		console.error('Error requesting notification permissions:', error);
		return notificationState;
	}
}

/**
 * Get the current FCM token
 */
export function getNotificationToken(): string | null {
	return notificationState.token;
}

/**
 * Get notification state
 */
export function getNotificationState(): NotificationState {
	return { ...notificationState };
}

/**
 * Check if notifications are enabled
 */
export async function checkNotificationPermissions(): Promise<boolean> {
	if (!Capacitor.isNativePlatform()) {
		return false;
	}

	try {
		const permStatus = await PushNotifications.checkPermissions();
		return permStatus.receive === 'granted';
	} catch (error) {
		console.error('Error checking notification permissions:', error);
		return false;
	}
}

/**
 * Clean up notification listeners (call on app unmount)
 */
export async function cleanupPushNotifications(): Promise<void> {
	if (!Capacitor.isNativePlatform()) {
		return;
	}

	try {
		await PushNotifications.removeAllListeners();
		listenersAdded = false;
		console.log('Push notification listeners removed');
	} catch (error) {
		console.error('Error removing notification listeners:', error);
	}
}
