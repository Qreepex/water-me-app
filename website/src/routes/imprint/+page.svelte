<script>
	import { resolve } from '$app/paths';
	import FooterSection from '$lib/components/landing/FooterSection.svelte';
</script>

<svelte:head>
	<title>Imprint - Water Me</title>
	<meta name="description" content="Imprint and legal information for Water Me app." />
	<meta property="og:title" content="Imprint - Water Me" />
	<meta property="og:description" content="Imprint and legal information for Water Me app." />
	<link rel="canonical" href="https://water-me.app/imprint" />
</svelte:head>

<div class="min-h-screen bg-gradient-to-br from-[#f2faf5] via-white to-[#e8f9f0]">
	<!-- Navigation -->
	<nav class="sticky top-0 z-50 border-b border-[#00ee57]/10 bg-white/80 backdrop-blur-md">
		<div class="mx-auto max-w-6xl px-4 py-4 sm:px-6 lg:px-8">
			<div class="flex items-center gap-3">
				<a href={resolve('/')} class="flex items-center gap-3 transition hover:opacity-80">
					<div class="text-3xl">🌿</div>
					<span class="text-2xl font-bold text-[#061f12]">Water Me</span>
				</a>
			</div>
		</div>
	</nav>

	<!-- Content -->
	<div class="mx-auto max-w-4xl px-4 py-20 sm:px-6 lg:px-8">
		<h1 class="text-5xl font-bold text-[#061f12]">Imprint</h1>
		<p class="mt-2 text-[#4b6658]">Legal information and contact details</p>

		<div class="prose prose-sm mt-12 max-w-none space-y-8 text-[#4b6658]">
			<div>
				<h2 class="text-3xl font-semibold text-[#061f12]">Provider</h2>
				<div class="mt-4 space-y-2">
					<p>
						Ben Schiemann<br />
						Reinbeker Weg 70<br />
						21465 Wentorf<br />
						Germany
					</p>
				</div>
			</div>

			<div>
				<h2 class="text-3xl font-semibold text-[#061f12]">Contact Information</h2>
				<div class="mt-4 space-y-2">
					<p>
						<strong>Phone:</strong><br />
						<a
							href="tel:+491516487726"
							class="text-[#00ee57] transition hover:text-[#00a343]"
						>
							+49 1516 4877269
						</a>
					</p>
					<p>
						<strong>Email:</strong><br />
						<a
							href="mailto:ben@water-me.app"
							class="text-[#00ee57] transition hover:text-[#00a343]"
						>
							ben@water-me.app
						</a>
					</p>
				</div>
			</div>

			<div>
				<h2 class="text-3xl font-semibold text-[#061f12]">Editorial Responsibility</h2>
				<div class="mt-4 space-y-2">
					<p>
						<strong>Responsible for editorial content:</strong><br />
						Ben Schiemann<br />
						Reinbeker Weg 70<br />
						21465 Wentorf<br />
						Germany
					</p>
				</div>
			</div>

			<div>
				<h2 class="text-3xl font-semibold text-[#061f12]">EU Dispute Resolution</h2>
				<p>
					The European Commission provides a platform for online dispute resolution (ODR):
					<a
						href="https://ec.europa.eu/consumers/odr/"
						class="text-[#00ee57] transition hover:text-[#00a343]"
						target="_blank"
						rel="noopener noreferrer"
					>
						https://ec.europa.eu/consumers/odr/
					</a>
				</p>
				<p>
					Our email address can be found above.
				</p>
			</div>

			<div>
				<h2 class="text-3xl font-semibold text-[#061f12]">Consumer Dispute Resolution</h2>
				<p>
					We are not willing or obligated to participate in dispute resolution proceedings before a consumer arbitration board.
				</p>
			</div>

			<div>
				<h2 class="text-3xl font-semibold text-[#061f12]">Image Credits</h2>
				<p>
					The images, photos, and graphics on this website are protected by copyright.
				</p>
				<p>
					Image rights are held by the following photographers and companies:
				</p>
				<ul class="mt-2 list-inside list-disc">
					<li>Discord Inc.</li>
					<li>YouTube LLC</li>
					<li>Twitch</li>
					<li>GitHub Inc.</li>
					<li>Twitter Inc.</li>
					<li>Meta Platforms, Inc.</li>
				</ul>
				<p class="mt-4">
					All text content is protected by copyright.
				</p>
			</div>

			<div>
				<h2 class="text-3xl font-semibold text-[#061f12]">Editorial Responsibility and Content Management</h2>
				<p>
					Editorial responsibility for the content of the Water Me website and application rests with:
				</p>
				<p>
					<strong>Ben Schiemann</strong><br />
					Reinbeker Weg 70<br />
					21465 Wentorf<br />
					Germany
				</p>
				<p>
					We strive to provide accurate and up-to-date information in all content. However, we do not guarantee the completeness, accuracy, timeliness, or reliability of any information presented. Any reliance you place on such information is strictly at your own risk.
				</p>
			</div>

			<div>
				<h2 class="text-3xl font-semibold text-[#061f12]">Technical Responsibility and System Operator</h2>
				<p>
					Technical responsibility for the operation and maintenance of the Water Me application and website rests with:
				</p>
				<p>
					<strong>Ben Schiemann</strong><br />
					Reinbeker Weg 70<br />
					21465 Wentorf<br />
					Germany
				</p>
				<p>
					While we maintain technical systems to the best of our ability and knowledge, we cannot guarantee uninterrupted availability of the service. We perform regular maintenance and updates that may occasionally result in service interruptions. We are not liable for any damages resulting from temporary service unavailability unless caused by our gross negligence.
				</p>
			</div>

			<div>
				<h2 class="text-3xl font-semibold text-[#061f12]">Third-Party Components and Open Source Libraries</h2>
				<p>
					The Water Me application and website incorporate components and libraries developed by third parties, many of which are open source software. We acknowledge and respect the licenses under which these components are distributed.
				</p>
				<p>
					Key open source frameworks and libraries used include:
				</p>
				<ul class="mt-2 list-inside list-disc">
					<li><strong>SvelteKit:</strong> Framework for building web applications (BSD 3-Clause License)</li>
					<li><strong>Svelte:</strong> JavaScript compiler for reactive user interfaces (MIT License)</li>
					<li><strong>Tailwind CSS:</strong> Utility-first CSS framework (MIT License)</li>
					<li><strong>Go:</strong> Open source programming language used for backend (BSD License)</li>
					<li><strong>MongoDB Driver for Go:</strong> Official MongoDB driver (Apache 2.0 License)</li>
					<li><strong>Firebase SDK:</strong> Google's application development platform (Apache 2.0 License)</li>
					<li><strong>Capacitor:</strong> Cross-platform app framework (MIT License)</li>
				</ul>
				<p class="mt-4">
					The complete source code of the Water Me application is publicly available on GitHub under an appropriate open source license. We encourage contributions, bug reports, and feature requests from the community. For access to the source code and license information, please visit our <a href="https://github.com" class="text-[#00ee57] transition hover:text-[#00a343]">GitHub repository</a>.
				</p>
				<p class="mt-4">
					All third-party libraries and their respective licenses are documented in our project's dependency management files. By using Water Me, you acknowledge your acceptance of the licenses of all incorporated open source components.
				</p>
			</div>

			<div>
				<h2 class="text-3xl font-semibold text-[#061f12]">Disclaimer of Warranties and Limitations of Liability</h2>
				<p>
					<strong>As Provided Basis:</strong> The Water Me application and website are provided on an "as is" and "as available" basis without warranties of any kind, either express or implied. We disclaim all warranties, including but not limited to, merchantability, fitness for a particular purpose, and non-infringement.
				</p>
				<p>
					<strong>No Liability:</strong> To the fullest extent permitted by law, in no event shall Water Me, its operators, contributors, or any affiliated parties be liable for any indirect, incidental, special, consequential, or punitive damages, including but not limited to loss of profits, data loss, or business interruption, arising out of or in connection with the use of or inability to use the service, even if advised of the possibility of such damages.
				</p>
				<p>
					<strong>User Responsibility:</strong> Users are solely responsible for their use of the Water Me service and for ensuring that their use complies with all applicable laws and regulations. We are not responsible for content you create, upload, or share through the service.
				</p>
			</div>

			<div>
				<h2 class="text-3xl font-semibold text-[#061f12]">Commercial Nature and Open Source Status</h2>
				<p>
					Water Me is a commercial application provided by Ben Schiemann for general use. While the application's source code is open source and freely available for inspection and contribution, the operation of the service may involve commercial aspects including potential future monetization, advertisements, or premium features.
				</p>
				<p>
					<strong>Current Status:</strong> The Water Me application is provided free of charge to users. Any future changes to pricing or premium features will be clearly communicated in advance.
				</p>
				<p>
					The open source nature of our codebase reflects our commitment to transparency and community participation. Users are encouraged to review the source code, identify security issues, and contribute improvements.
				</p>
			</div>
		</div>
	</div>

	<!-- Footer -->
	<FooterSection />
</div>

<style>
	:global(h2) {
		margin: 2rem 0 1rem;
	}

	:global(p) {
		margin-bottom: 0.5rem;
	}
</style>
