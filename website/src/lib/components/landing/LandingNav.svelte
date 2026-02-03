<script lang="ts">
	import { resolve } from '$app/paths';
	import { WEB_APP_URL } from '$lib';
	import { t, languageStore } from '$lib/i18n.svelte';
</script>

<nav class="sticky top-0 z-50 border-b border-[#00ee57]/10 bg-white/80 backdrop-blur-md">
	<div class="mx-auto max-w-6xl px-4 py-4 sm:px-6 lg:px-8">
		<div class="flex items-center justify-between">
			<div class="flex items-center gap-3">
				<div class="text-3xl">🌿</div>
				<span class="text-2xl font-bold text-[#061f12]">Water Me</span>
			</div>
			<div class="flex gap-4">
				<div class="hidden gap-8 sm:flex">
					<a href="#features" class="text-[#4b6658] transition hover:text-[#00ee57]">
						{$t('nav.features')}
					</a>
					<a href="#stats" class="text-[#4b6658] transition hover:text-[#00ee57]">
						{$t('nav.stats')}
					</a>
					<a href="#download" class="text-[#4b6658] transition hover:text-[#00ee57]">
						{$t('nav.download')}
					</a>
					<a
						href={WEB_APP_URL}
						target="_blank"
						rel="noopener noreferrer external"
						class="inline-flex items-center gap-2 rounded-full bg-[#00ee57] px-6 py-2 font-semibold text-[#061f12] transition hover:bg-[#00a343] hover:text-white"
					>
						{$t('nav.open_app')}
					</a>
				</div>
				<div>
					<label for="language-select" class="sr-only">{$t('nav.select_language')}</label>
					<select
						id="language-select"
						bind:value={$languageStore}
						onchange={() => {
							const lang = $languageStore;
							window.location.href = resolve(`/${lang}`);
						}}
						class="rounded border border-[#00ee57]/20 bg-white px-2 py-1 text-[#061f12]"
					>
						<option value="en">{$t('nav.languages.en')}</option>
						<option value="de">{$t('nav.languages.de')}</option>
						<option value="es">{$t('nav.languages.es')}</option>
					</select>
				</div>
			</div>
		</div>
	</div>
</nav>
