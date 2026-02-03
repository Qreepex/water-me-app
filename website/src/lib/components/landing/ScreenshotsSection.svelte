<script lang="ts">
	import { t } from '$lib/i18n.svelte';

	const cards = [
		{
			emoji: '🏠',
			titleKey: 'screenshots.cards.home.title',
			descriptionKey: 'screenshots.cards.home.description'
		},
		{
			emoji: '💧',
			titleKey: 'screenshots.cards.watering.title',
			descriptionKey: 'screenshots.cards.watering.description'
		},
		{
			emoji: '✏️',
			titleKey: 'screenshots.cards.management.title',
			descriptionKey: 'screenshots.cards.management.description'
		}
	];
</script>

<section class="px-4 py-20 sm:px-6 lg:px-8">
	<div class="mx-auto max-w-6xl">
		<h2 class="text-center text-4xl font-bold text-[#061f12] sm:text-5xl">
			{$t('screenshots.title')}
		</h2>
		<p class="mx-auto mt-4 max-w-2xl text-center text-lg text-[#4b6658]">
			{$t('screenshots.subtitle')}
		</p>

		<div class="mt-12 grid gap-8 sm:grid-cols-2 lg:grid-cols-3">
			{#each cards as card (card.titleKey)}
				<div class="overflow-hidden rounded-2xl border border-[#00ee57]/20 bg-white shadow-sm">
					<div
						class="flex h-48 items-center justify-center bg-gradient-to-br from-[#00ee57]/10 to-[#00a343]/10 text-6xl"
					>
						{card.emoji}
					</div>
					<div class="p-6">
						<h3 class="text-xl font-semibold text-[#061f12]">{$t(card.titleKey)}</h3>
						<p class="mt-2 text-[#4b6658]">{$t(card.descriptionKey)}</p>
					</div>
				</div>
			{/each}
		</div>
	</div>
</section>
