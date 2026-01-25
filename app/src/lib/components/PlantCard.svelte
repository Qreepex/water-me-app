<script lang="ts">
	import type { Plant } from '$lib/types/api';

	export let plant: Plant;
	export let daysAgo: (dateString: string) => string;
	export let getWateringStatus: (plant: Plant) => { text: string; color: string };
</script>

<div
	class="group cursor-pointer overflow-hidden rounded-2xl bg-white shadow-md transition-all duration-300 hover:scale-105 hover:shadow-xl"
>
	<!-- Image -->
	<div
		class="relative flex h-48 items-center justify-center overflow-hidden bg-gradient-to-br from-green-200 to-emerald-300"
	>
		{#if plant.photoIds && plant.photoIds.length > 0}
			<img
				src={plant.photoIds[0]}
				alt={plant.name}
				class="h-full w-full object-cover transition-transform duration-300 group-hover:scale-110"
			/>
		{:else}
			<div class="text-7xl transition-transform duration-300 group-hover:scale-110">🌱</div>
		{/if}
	</div>

	<!-- Content -->
	<div class="p-5">
		<!-- Name and Species -->
		<h3 class="mb-1 line-clamp-2 text-xl font-bold text-green-800">{plant.name}</h3>
		<p class="mb-4 line-clamp-1 text-sm text-green-600">{plant.species}</p>

		<!-- Watering Status -->
		<div class="mb-4">
			<div class={`mb-2 text-sm font-semibold ${getWateringStatus(plant).color}`}>
				{getWateringStatus(plant).text}
			</div>
			<p class="text-xs text-gray-600">Watered {daysAgo(plant.watering?.lastWatered ?? '')}</p>
		</div>

		<!-- Metadata Grid -->
		<div class="mb-4 grid grid-cols-2 gap-3 text-xs">
			<div class="rounded-lg bg-blue-50 p-2">
				<div class="font-semibold text-blue-600">💧</div>
				<p class="mt-1 text-xs text-gray-700">Every {plant.watering?.intervalDays}d</p>
			</div>
			<div class="rounded-lg bg-yellow-50 p-2">
				<div class="font-semibold text-yellow-600">🥗</div>
				<p class="mt-1 text-xs text-gray-700">Every {plant.fertilizing?.intervalDays}d</p>
			</div>
			<div class="rounded-lg bg-purple-50 p-2">
				<div class="font-semibold text-purple-600">☀️</div>
				<p class="mt-1 text-xs text-gray-700">{plant.sunlight.split(' ').slice(0, 1).join('')}</p>
			</div>
			<div class="rounded-lg bg-teal-50 p-2">
				<div class="font-semibold text-teal-600">💨</div>
				<p class="mt-1 text-xs text-gray-700">{plant.humidity?.targetHumidityPct}%</p>
			</div>
		</div>

		<!-- Spray Info -->
		{#if plant.humidity?.requiresMisting}
			<div class="mb-3 rounded-lg bg-cyan-50 p-2">
				<p class="text-xs text-gray-600">
					💦 Spray every <span class="font-semibold text-cyan-700"
						>{plant.humidity?.mistingIntervalDays}</span
					> days
				</p>
				<p class="mt-1 text-xs text-gray-600">
					Last: <span class="font-semibold">{daysAgo(plant.fertilizing?.lastFertilized ?? '')}</span
					>
				</p>
			</div>
		{/if}

		<!-- Flags -->
		{#if plant.flags && plant.flags.length > 0}
			<div class="mb-3 flex flex-wrap gap-2">
				{#each plant.flags as flag (flag)}
					<span class="rounded-full bg-orange-100 px-2 py-1 text-xs font-medium text-orange-800"
						>⚡ {flag}</span
					>
				{/each}
			</div>
		{/if}

		<!-- Notes Preview -->
		{#if plant.notes && plant.notes.length > 0}
			<div class="border-t border-gray-200 pt-3">
				<p class="line-clamp-2 text-xs text-gray-600">📝 {plant.notes[0]}</p>
			</div>
		{/if}
	</div>
</div>
