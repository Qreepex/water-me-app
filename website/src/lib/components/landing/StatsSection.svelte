<script lang="ts">
	import { API_BASE_URL } from '$lib';
	import { t } from '$lib/i18n.svelte';
	import { onMount } from 'svelte';

	let displayStats = $state({
		users: 0,
		plants: 0,
		reminders: 0
	});
	let targetStats = $state({
		users: 0,
		plants: 0,
		reminders: 0
	});
	let hasAnimated = $state(false);
	let statsReady = $state(false);

	onMount(async () => {
		try {
			const response = await fetch(`${API_BASE_URL}/stats`);
			if (!response.ok) return;
			const data = await response.json();
			if (typeof data?.users !== 'number' || typeof data?.plants !== 'number') return;
			targetStats = {
				users: data.users,
				plants: data.plants,
				reminders: typeof data?.reminders === 'number' ? data.reminders : 0
			};
			statsReady = true;
		} catch {
			return;
		}
	});

	$effect(() => {
		if (!statsReady || hasAnimated) return;
		hasAnimated = true;

		const duration = 2000;
		const startTime = Date.now();

		const animate = () => {
			const elapsed = Date.now() - startTime;
			const progress = Math.min(elapsed / duration, 1);

			displayStats.users = Math.floor(targetStats.users * progress);
			displayStats.plants = Math.floor(targetStats.plants * progress);
			displayStats.reminders = Math.floor(targetStats.reminders * progress);

			if (progress < 1) {
				requestAnimationFrame(animate);
			}
		};

		animate();
	});
</script>

<section
	id="stats"
	class="bg-gradient-to-r from-[#00ee57] to-[#00a343] px-4 py-20 sm:px-6 lg:px-8"
>
	<div class="mx-auto max-w-6xl">
		<div class="mb-12 text-center">
			<h2 class="text-4xl font-bold text-white sm:text-5xl">{$t('stats.title')}</h2>
		</div>

		<div class="grid gap-8 sm:grid-cols-3">
			<div class="rounded-2xl bg-white/20 p-8 backdrop-blur-sm">
				<div class="text-5xl font-bold text-white">{displayStats.users.toLocaleString()}</div>
				<p class="mt-2 text-lg text-white/90">{$t('stats.users')}</p>
			</div>

			<div class="rounded-2xl bg-white/20 p-8 backdrop-blur-sm">
				<div class="text-5xl font-bold text-white">{displayStats.plants.toLocaleString()}</div>
				<p class="mt-2 text-lg text-white/90">{$t('stats.plants')}</p>
			</div>

			<div class="rounded-2xl bg-white/20 p-8 backdrop-blur-sm">
				<div class="text-5xl font-bold text-white">
					{displayStats.reminders.toLocaleString()}+
				</div>
				<p class="mt-2 text-lg text-white/90">{$t('stats.reminders')}</p>
			</div>
		</div>
	</div>
</section>
