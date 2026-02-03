<script lang="ts">
	import { t } from '$lib/i18n.svelte';
</script>

<section id="download" class="bg-white/50 px-4 py-20 sm:px-6 lg:px-8">
	<div class="mx-auto max-w-4xl">
		<div
			class="rounded-3xl border-2 border-[#00ee57] bg-gradient-to-br from-[#00ee57]/5 to-[#00a343]/5 p-12 text-center"
		>
			<h2 class="text-4xl font-bold text-[#061f12] sm:text-5xl">{$t('download.title')}</h2>
			<p class="mt-4 text-lg text-[#4b6658]">{$t('download.subtitle')}</p>

			<div class="mt-8 flex flex-col gap-4 sm:flex-row sm:justify-center">
				<a
					href="https://play.google.com/store/apps/details?id=app.waterme.app"
					class="inline-flex items-center justify-center gap-2 rounded-lg bg-[#061f12] px-8 py-4 font-semibold text-white transition hover:bg-[#061f12]/80"
				>
					<span>🤖</span>
					{$t('download.googleplay')}
				</a>
				<div
					class="inline-flex items-center justify-center gap-2 rounded-lg bg-[#061f12] px-8 py-4 font-semibold text-white transition hover:bg-[#061f12]/80"
				>
					<span>🍎</span>
					{$t('download.appstore')}
				</div>
			</div>
		</div>
	</div>
</section>
