<script lang="ts">
	import { t } from '$lib/i18n.svelte';

	const faqKeys = ['q1', 'q2', 'q3', 'q4', 'q5', 'q6', 'q7'];
</script>

<section class="px-4 py-20 sm:px-6 lg:px-8">
	<div class="mx-auto max-w-4xl">
		<h2 class="text-center text-4xl font-bold text-[#061f12] sm:text-5xl">
			{$t('faq.title')}
		</h2>

		<div class="mt-12 space-y-6">
			{#each faqKeys as key (key)}
				<div class="rounded-lg border border-[#00ee57]/20 bg-white p-6">
					<h3 class="font-semibold text-[#061f12]">{$t(`faq.${key}`)}</h3>
					<p class="mt-2 text-[#4b6658]">{$t(`faq.a${key.substring(1)}`)}</p>
				</div>
			{/each}
		</div>
	</div>
</section>
