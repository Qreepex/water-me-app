<script lang="ts">
	type SortOption =
		| 'name'
		| 'lastWatered'
		| 'lastFertilized'
		| 'sprayInterval'
		| 'wateringInterval';
	export let sortBy: SortOption = 'name';
</script>

<div class="flex items-center gap-3">
	<label for="sort" class="font-semibold text-green-800">Sort by:</label>
	<select
		id="sort"
		bind:value={sortBy}
		class="rounded-lg border-2 border-green-300 bg-white px-4 py-2 font-medium text-green-800 transition hover:border-green-400 focus:border-green-500 focus:outline-none"
	>
		<option value="name">Plant Name</option>
		<option value="lastWatered">Last Watered</option>
		<option value="lastFertilized">Last Fertilized</option>
		<option value="wateringInterval">Watering Frequency</option>
		<option value="sprayInterval">Spray Frequency</option>
	</select>
</div>
