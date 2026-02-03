<script lang="ts">
	import LandingPage from '$lib/components/landing/LandingPage.svelte';
</script>

<svelte:head>
	<title>Water Me - Never Miss a Plant Care Reminder</title>
	<meta
		name="description"
		content="Smart plant care companion app. Get reminders for watering, spraying, and fertilizing. Track your plant collection with ease."
	/>
	<meta name="keywords" content="plant care, app, reminders, watering, gardening, plants" />
	<meta property="og:title" content="Water Me - Smart Plant Care Companion" />
	<meta
		property="og:description"
		content="Never forget to water your plants again. Track watering schedules, get smart reminders, and keep your plants thriving."
	/>
	<meta property="og:type" content="website" />
	<meta property="og:image" content="/og-image.jpg" />
	<link rel="canonical" href="https://water-me.app/es" />
	<link rel="alternate" hreflang="en" href="https://water-me.app/en" />
	<link rel="alternate" hreflang="de" href="https://water-me.app/de" />
	<link rel="alternate" hreflang="es" href="https://water-me.app/es" />
	<link rel="alternate" hreflang="x-default" href="https://water-me.app/en" />
</svelte:head>

<LandingPage />
