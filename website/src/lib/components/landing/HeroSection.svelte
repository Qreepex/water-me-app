<script lang="ts">
	import { WEB_APP_URL } from '$lib';
	import { t } from '$lib/i18n.svelte';
</script>

<section class="relative overflow-hidden px-4 py-20 sm:px-6 lg:px-8">
	<div class="mx-auto max-w-6xl">
		<div class="grid items-center gap-12 lg:grid-cols-2">
			<div>
				<h1 class="text-5xl leading-tight font-bold text-[#061f12] sm:text-6xl">
					{$t('hero.title')}
				</h1>
				<p class="mt-6 text-xl text-[#4b6658]">
					{$t('hero.subtitle')}
				</p>

				<div class="mt-8 flex gap-4">
					<a
						href="#download"
						class="inline-flex items-center gap-2 rounded-full bg-[#00ee57] px-8 py-4 font-semibold text-[#061f12] shadow-lg shadow-[#00ee57]/40 transition hover:shadow-xl hover:shadow-[#00ee57]/50"
					>
						📥 {$t('hero.cta_primary')}
					</a>
					<a
						href={WEB_APP_URL}
						target="_blank"
						rel="noopener noreferrer external"
						class="inline-flex items-center gap-2 rounded-full border-2 border-[#00ee57] px-8 py-4 font-semibold text-[#00ee57] transition hover:bg-[#00ee57]/10"
					>
						🌐 {$t('hero.cta_secondary')}
					</a>
				</div>

				<div class="mt-12 flex items-center gap-4">
					<div class="text-4xl">⭐⭐⭐⭐⭐</div>
					<div>
						<p class="font-semibold text-[#061f12]">{$t('hero.trust_badge')}</p>
					</div>
				</div>
			</div>

			<div class="relative h-96 lg:h-full">
				<div
					class="absolute inset-0 rounded-3xl bg-gradient-to-br from-[#00ee57] to-[#00a343] shadow-2xl"
				>
					<div class="flex h-full items-center justify-center text-8xl">🌱</div>
				</div>
			</div>
		</div>
	</div>
</section>
