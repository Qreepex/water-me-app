<script>
	import { resolve } from '$app/paths';
	import FooterSection from '$lib/components/landing/FooterSection.svelte';
</script>

<svelte:head>
	<title>Terms & Conditions - Water Me</title>
	<meta name="description" content="Terms and conditions for Water Me app." />
	<meta property="og:title" content="Terms & Conditions - Water Me" />
	<meta property="og:description" content="Terms and conditions for Water Me app." />
	<link rel="canonical" href="https://water-me.app/terms" />
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
		<h1 class="text-5xl font-bold text-[#061f12]">Terms & Conditions</h1>
		<p class="mt-2 text-[#4b6658]">Last updated: February 2026</p>

		<div class="prose prose-sm mt-12 max-w-none text-[#4b6658]">
			<h2 class="text-3xl font-semibold text-[#061f12]">1. Acceptance of Terms</h2>
			<p>
				By accessing and using the Water Me application, website, and related services
				("Service"), you accept and agree to be bound by these Terms & Conditions and our Privacy
				Policy. If you do not agree to these terms, you must not use this Service. Your continued
				use of the Service constitutes your acceptance of any modifications to these terms.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">2. Service Description</h2>
			<p>
				Water Me is a plant care management application that allows users to track, manage, and
				organize their plants through plant records, watering schedules, reminders, and plant
				images. The Service is provided on an "as is" and "as available" basis for personal,
				non-commercial use only.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">3. User Accounts & Responsibilities</h2>
			<p>
				When you create an account, you agree to:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li>Provide accurate, complete, and current information</li>
				<li>
					Maintain the confidentiality of your account credentials and Firebase authentication
					tokens
				</li>
				<li>
					Notify us immediately of any unauthorized access or use of your account
				</li>
				<li>
					Accept full responsibility for all activities conducted through your account
				</li>
				<li>
					Be at least 16 years of age (or older in your jurisdiction if the legal age of digital
					consent is higher)
				</li>
			</ul>

			<h2 class="text-3xl font-semibold text-[#061f12]">4. Prohibited Uses</h2>
			<p>You agree not to use the Service to:</p>
			<ul class="list-inside list-disc space-y-2">
				<li>
					Engage in any illegal, fraudulent, or harmful activity, including hacking, malware
					distribution, or unauthorized access attempts
				</li>
				<li>
					Upload, store, or transmit any content that violates applicable laws, including but not
					limited to illegal material, stolen content, or content violating third-party rights
				</li>
				<li>
					Harass, threaten, defame, or discriminate against any individual or group
				</li>
				<li>
					Disrupt or interfere with the Service's functionality, security, or integrity
				</li>
				<li>
					Reverse engineer, decompile, modify, or attempt to derive the source code of the Service
				</li>
				<li>
					Scrape, crawl, or systematically extract data from the Service for unauthorized purposes
				</li>
				<li>
					Use the Service for commercial purposes without written permission
				</li>
				<li>
					Violate any applicable laws, regulations, or these Terms
				</li>
			</ul>

			<h2 class="text-3xl font-semibold text-[#061f12]">5. Photo Upload & Content Rights</h2>
			<p>
				When you upload plant photos to the Water Me Service, you grant Water Me a
				worldwide, non-exclusive, royalty-free license to host, store, display, and transmit
				your images for the sole purpose of operating and providing the Service to you. You
				retain full ownership and all intellectual property rights to your photos.
			</p>
			<p>
				You agree that:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li>
					You own or have permission to upload all photos, and they do not infringe any
					third-party intellectual property or privacy rights
				</li>
				<li>
					Photos do not contain identifiable images of people, faces, or individuals without
					explicit consent
				</li>
				<li>
					You will not upload content that is illegal, abusive, defamatory, obscene, or
					otherwise harmful
				</li>
				<li>
					Water Me may delete any photos that violate these terms or applicable laws
				</li>
				<li>
					You are responsible for creating backups of your photos, as Water Me is not liable
					for accidental loss or deletion
				</li>
			</ul>

			<h2 class="text-3xl font-semibold text-[#061f12]">6. Data Storage & Photo Hosting</h2>
			<p>
				Your plant data and photos are stored on our cloud providers' infrastructure. By using the
				Service, you consent to this storage arrangement and acknowledge the associated
				cloud-based data handling as detailed in our Privacy Policy. Water Me takes reasonable
				security measures to protect your data, but cannot guarantee absolute security.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">7. Authentication & Data Collection</h2>
			<p>
				The Service uses Firebase Authentication for secure login via Google Sign-In. We
				collect your Firebase unique identifier, email address, and basic profile information
				necessary for authentication and service operation. All data processing complies with our
				Privacy Policy, GDPR, and German DSGVO regulations. See our Privacy Policy for detailed
				information on how your data is processed.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">8. Push Notifications</h2>
			<p>
				Water Me uses Firebase Cloud Messaging to send plant care reminders and notifications.
				You can manage notification preferences within the application. Push notifications will
				only be sent to registered devices associated with your account. By enabling
				notifications, you consent to receiving these messages.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">9. Third-Party Services</h2>
			<p>
				The Service integrates with third-party services including Google (Firebase
				Authentication), Amazon Web Services (S3 storage), and Simple Analytics (web analytics).
				These services are subject to their own terms and privacy policies. Water Me is not
				responsible for third-party service outages, data handling, or policy changes. Please
				review their terms before using the Service.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">10. Disclaimer of Warranties</h2>
			<p>
				The Service is provided "as is" and "as available" without any warranties, express or
				implied. Water Me disclaims all warranties including merchantability, fitness for a
				particular purpose, non-infringement, and accuracy. We do not warrant that:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li>The Service will be uninterrupted, error-free, or secure</li>
				<li>Plant care information provided is always accurate or suitable for all plants</li>
				<li>Reminders and notifications will always be delivered reliably</li>
				<li>Your data will be permanently retained or recoverable</li>
			</ul>

			<h2 class="text-3xl font-semibold text-[#061f12]">11. Limitation of Liability</h2>
			<p>
				To the fullest extent permitted by law, Water Me shall not be liable for any direct,
				indirect, incidental, consequential, special, or punitive damages arising from your use
				or inability to use the Service, including:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li>Plant damage or loss resulting from missed watering reminders</li>
				<li>Data loss or corruption</li>
				<li>Service interruptions or unavailability</li>
				<li>Security breaches or unauthorized access</li>
				<li>Reliance on inaccurate plant care information</li>
			</ul>
			<p>
				This limitation applies even if Water Me has been advised of the possibility of such
				damages. Some jurisdictions do not allow limitation of liability, so this clause may
				not apply to you.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">12. Indemnification</h2>
			<p>
				You agree to indemnify and hold harmless Water Me and its officers, employees, and
				agents from any claims, damages, losses, or expenses (including legal fees) arising from:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li>Your violation of these Terms</li>
				<li>Your use of the Service in violation of applicable laws</li>
				<li>Content you upload or transmit through the Service</li>
				<li>Your infringement of third-party rights</li>
			</ul>

			<h2 class="text-3xl font-semibold text-[#061f12]">13. User-Generated Content</h2>
			<p>
				You are solely responsible for content you create, upload, or transmit through the
				Service. Water Me may monitor, review, or remove content that violates these Terms or
				applicable laws. Water Me does not endorse or guarantee the accuracy of any
				user-generated content.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">14. Intellectual Property</h2>
			<p>
				All content, features, and functionality of the Water Me Service, including but not
				limited to text, graphics, logos, code, and software, are the exclusive property of
				Water Me or its licensors and are protected by copyright and other intellectual property
				laws. You may not reproduce, modify, distribute, or create derivative works without
				explicit written permission.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">15. Account Termination</h2>
			<p>
				Water Me reserves the right to suspend or terminate your account at any time if you
				violate these Terms or applicable laws. Upon termination, your right to use the Service
				ceases immediately. You may request deletion of your account and associated data at any
				time, subject to our Data Deletion Policy outlined in the Privacy Policy.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">16. GDPR & DSGVO Compliance</h2>
			<p>
				Water Me processes personal data in accordance with the EU General Data Protection
				Regulation (GDPR) and the German Bundesdatenschutzgesetz (BDSG/DSGVO). As a data
				controller, Water Me provides:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li>
					Transparent data collection, use, and storage practices detailed in our Privacy Policy
				</li>
				<li>Your right to access, correct, or delete your personal data</li>
				<li>Your right to data portability and objection to processing</li>
				<li>Secure storage with encryption and appropriate safeguards</li>
				<li>
					Data Protection Impact Assessments and Data Processing Agreements with third-party
					processors
				</li>
			</ul>
			<p>
				For more information, see our complete Privacy Policy at{' '}
				<a href={resolve('/privacy')} class="text-[#00ee57] hover:underline">
					water-me.app/privacy
				</a>
				.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">17. Modifications to Terms</h2>
			<p>
				Water Me may modify these Terms at any time. Material changes will be notified via email
				or prominent notice on the Service. Your continued use of the Service following
				notification of changes constitutes acceptance of the modified terms. If you do not
				accept changes, you must discontinue use of the Service.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">18. Governing Law & Dispute Resolution</h2>
			<p>
				These Terms and Conditions are governed by the laws of Germany, without regard to its
				conflict of law provisions. You irrevocably submit to the exclusive jurisdiction of the
				courts located in Germany for resolution of any disputes arising from these Terms or
				your use of the Service.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">19. Entire Agreement</h2>
			<p>
				These Terms and Conditions, together with our Privacy Policy and any other policies
				referenced herein, constitute the entire agreement between you and Water Me regarding
				the Service and supersede all prior and contemporaneous agreements, understandings, and
				negotiations.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">20. Contact & Support</h2>
			<p>
				If you have questions, concerns, or wish to exercise your rights regarding these Terms
				or the Service, please contact us at:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li>Email: ben@water-me.app</li>
				<li>See our imprint for additional contact information and data protection officer details</li>
			</ul>
			<p>
				For urgent matters related to data protection, you may also file a complaint with the
				competent data protection authority in your jurisdiction.
			</p>
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
		margin-bottom: 1rem;
	}

	:global(ul) {
		margin-bottom: 1rem;
	}
</style>
