import { DefaultSystemBrowserOptions, InAppBrowser } from '@capacitor/inappbrowser';

export async function openExternalLink(url: string): Promise<void> {
	await InAppBrowser.openInSystemBrowser({ url, options: DefaultSystemBrowserOptions });
}
