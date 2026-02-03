<script lang="ts">
	import LandingNav from './LandingNav.svelte';
	import HeroSection from './HeroSection.svelte';
	import FeaturesSection from './FeaturesSection.svelte';
	import StatsSection from './StatsSection.svelte';
	import ScreenshotsSection from './ScreenshotsSection.svelte';
	import DownloadSection from './DownloadSection.svelte';
	import FaqSection from './FaqSection.svelte';
	import FooterSection from './FooterSection.svelte';
</script>

<main class="min-h-screen bg-gradient-to-br from-[#f2faf5] via-white to-[#e8f9f0]">
	<LandingNav />
	<HeroSection />
	<FeaturesSection />
	<StatsSection />
	<ScreenshotsSection />
	<DownloadSection />
	<FaqSection />
	<FooterSection />
</main>
