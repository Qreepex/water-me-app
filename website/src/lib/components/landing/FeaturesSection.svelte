<script lang="ts">
	import { t } from '$lib/i18n.svelte';

	const features = [
		{
			icon: '💧',
			border: 'border-[#00ee57]/20',
			titleKey: 'features.watering.title',
			descriptionKey: 'features.watering.description'
		},
		{
			icon: '🎯',
			border: 'border-[#ffb703]/20',
			titleKey: 'features.multicare.title',
			descriptionKey: 'features.multicare.description'
		},
		{
			icon: '📸',
			border: 'border-[#219ebc]/20',
			titleKey: 'features.growth.title',
			descriptionKey: 'features.growth.description'
		},
		{
			icon: '📊',
			border: 'border-[#e63946]/20',
			titleKey: 'features.history.title',
			descriptionKey: 'features.history.description'
		},
		{
			icon: '🔔',
			border: 'border-[#008f41]/20',
			titleKey: 'features.notifications.title',
			descriptionKey: 'features.notifications.description'
		},
		{
			icon: '🌍',
			border: 'border-[#00ee57]/20',
			titleKey: 'features.multilang.title',
			descriptionKey: 'features.multilang.description'
		}
	];
</script>

<section id="features" class="bg-white/50 px-4 py-20 sm:px-6 lg:px-8">
	<div class="mx-auto max-w-6xl">
		<div class="mb-16 text-center">
			<h2 class="text-4xl font-bold text-[#061f12] sm:text-5xl">
				{$t('features.section_title')}
			</h2>
			<p class="mt-4 text-xl text-[#4b6658]">{$t('features.section_subtitle')}</p>
		</div>

		<div class="grid gap-8 sm:grid-cols-2 lg:grid-cols-3">
			{#each features as feature (feature.titleKey)}
				<div
					class="rounded-2xl border {feature.border} bg-white p-8 shadow-sm transition hover:shadow-lg"
				>
					<div class="mb-4 text-5xl">{feature.icon}</div>
					<h3 class="mb-2 text-2xl font-semibold text-[#061f12]">
						{$t(feature.titleKey)}
					</h3>
					<p class="text-[#4b6658]">{$t(feature.descriptionKey)}</p>
				</div>
			{/each}
		</div>
	</div>
</section>
