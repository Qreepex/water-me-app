<script lang="ts">
	import { resolve } from '$app/paths';
	import { GITHUB_URL, INSTAGRAM_URL, TIKTOK_URL } from '$lib';
	import { t } from '$lib/i18n.svelte';
</script>

<footer class="border-t border-[#00ee57]/10 bg-[#061f12] px-4 py-12 sm:px-6 lg:px-8">
	<div class="mx-auto max-w-6xl">
		<div class="grid gap-8 sm:grid-cols-4">
			<div>
				<div class="flex items-center gap-2">
					<div class="text-2xl">🌿</div>
					<span class="text-xl font-bold text-white">Water Me</span>
				</div>
				<p class="mt-2 text-sm text-white/70">{$t('footer.tagline')}</p>
			</div>

			<div>
				<h4 class="font-semibold text-white">{$t('footer.product')}</h4>
				<ul class="mt-2 space-y-1">
					<li>
						<a
							href={resolve("/#features")}
							class="text-sm text-white/70 outline-none transition hover:text-white focus-visible:ring-0 focus-visible:text-white"
						>
							{$t('nav.features')}
						</a>
					</li>
					<li>
						<a
							href={resolve("/#download")}  
							class="text-sm text-white/70 outline-none transition hover:text-white focus-visible:ring-0 focus-visible:text-white"
						>
							{$t('nav.download')}
						</a>
					</li>
					<li>
						<a
							href="mailto:support@water-me.app"
							class="text-sm text-white/70 outline-none transition hover:text-white focus-visible:ring-0 focus-visible:text-white"
						>
							{$t('footer.support')}
						</a>
					</li>
				</ul>
			</div>

			<div>
				<h4 class="font-semibold text-white">{$t('footer.legal')}</h4>
				<ul class="mt-2 space-y-1">
					<li>
						<a
							href={resolve('/privacy')}
							data-sveltekit-preload-data="off"
							class="text-sm text-white/70 outline-none transition hover:text-white focus-visible:ring-0 focus-visible:text-white"
						>
							{$t('footer.privacy')}
						</a>
					</li>
					<li>
						<a
							href={resolve('/terms')}
							data-sveltekit-preload-data="off"
							class="text-sm text-white/70 outline-none transition hover:text-white focus-visible:ring-0 focus-visible:text-white"
						>
							{$t('footer.terms')}
						</a>
					</li>
					<li>
						<a
							href={resolve('/imprint')}
							data-sveltekit-preload-data="off"
							class="text-sm text-white/70 outline-none transition hover:text-white focus-visible:ring-0 focus-visible:text-white"
						>
							{$t('footer.imprint')}
						</a>
					</li>
				</ul>
			</div>

			<div>
				<h4 class="font-semibold text-white">{$t('footer.connect')}</h4>
				<ul class="mt-2 space-y-1">
					<li>
						<a
							rel="external"
							target="_blank"
							href={TIKTOK_URL}
							class="text-sm text-white/70 outline-none transition hover:text-white focus-visible:ring-0 focus-visible:text-white"
						>
							{$t('footer.social.tiktok')}
						</a>
					</li>
					<li>
						<a
							rel="external"
							target="_blank"
							href={INSTAGRAM_URL}
							class="text-sm text-white/70 outline-none transition hover:text-white focus-visible:ring-0 focus-visible:text-white"
						>
							{$t('footer.social.instagram')}
						</a>
					</li>
					<li>
						<a
							rel="external"
							target="_blank"
							href={GITHUB_URL}
							class="text-sm text-white/70 outline-none transition hover:text-white focus-visible:ring-0 focus-visible:text-white"
						>
							{$t('footer.social.github')}
						</a>
					</li>
				</ul>
			</div>
		</div>

		<div class="mt-8 border-t border-white/10 pt-8">
			<p class="text-center text-sm text-white/70">{$t('footer.copyright')}</p>
		</div>
	</div>
</footer>
