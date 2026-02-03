<script>
	import { resolve } from '$app/paths';
	import FooterSection from '$lib/components/landing/FooterSection.svelte';
</script>

<svelte:head>
	<title>Privacy Policy - Water Me</title>
	<meta
		name="description"
		content="Privacy policy for Water Me app. GDPR and DSGVO compliant. Learn how we protect your data."
	/>
	<meta property="og:title" content="Privacy Policy - Water Me" />
	<meta
		property="og:description"
		content="Privacy policy for Water Me app. GDPR and DSGVO compliant. Learn how we protect your data."
	/>
	<link rel="canonical" href="https://water-me.app/privacy" />
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
		<h1 class="text-5xl font-bold text-[#061f12]">Privacy Policy</h1>
		<p class="mt-2 text-[#4b6658]">Last updated: February 3, 2026</p>

		<div class="prose prose-sm mt-12 max-w-none text-[#4b6658]">
			<h2 class="text-3xl font-semibold text-[#061f12]">1. Overview of Data Protection</h2>
			
			<h3 class="text-2xl font-semibold text-[#061f12]">General Information</h3>
			<p>
				The following information will provide you with an overview of what happens with your personal data when you use our Water Me application and website. The term "personal data" comprises all data that can be used to personally identify you.
			</p>
			<p>
				This privacy policy is in full compliance with the European Union's General Data Protection Regulation (GDPR) and the German Data Protection Act (DSGVO).
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">2. Responsible Party (Controller)</h2>
			<p>
				The data processing controller responsible for this service is:
			</p>
			<p>
				<strong>Ben Schiemann</strong><br />
				Reinbeker Weg 70<br />
				21465 Wentorf<br />
				Germany<br /><br />
				Phone: +49 1516 4877269<br />
				Email: ben@water-me.app
			</p>
			<p>
				The controller is the natural person or legal entity that, individually or jointly with others, determines the purposes of and means for processing personal data.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">3. Data Collection and Recording</h2>

			<h3 class="text-2xl font-semibold text-[#061f12]">How We Collect Your Data</h3>
			<p>
				We collect data through two primary methods:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li><strong>Data you provide directly:</strong> Information you voluntarily enter into the application (email address, plant information, care preferences)</li>
				<li><strong>Automatically collected data:</strong> Technical information recorded by our systems when you use the service (device information, usage patterns, timestamps)</li>
			</ul>

			<h3 class="text-2xl font-semibold text-[#061f12]">Types of Data Collected</h3>
			<ul class="list-inside list-disc space-y-2">
				<li>
					<strong>Authentication Data:</strong> Firebase user ID and username (managed through Firebase Authentication). We do not store your email address.
				</li>
				<li>
					<strong>Plant Data:</strong> Plant names, species, care schedules, watering frequency, care notes, and plant photos that you explicitly upload
				</li>
				<li>
					<strong>Notification Settings:</strong> Your preferences for care reminders and push notification configurations
				</li>
				<li>
					<strong>Device Information:</strong> Device platform (Android/iOS), operating system version, and notification permissions status. This information is collected only to deliver push notifications reliably.
				</li>
				<li>
					<strong>Server Log Data:</strong> IP address, browser type, operating system, referrer URL, and server access times
				</li>
			</ul>

			<h3 class="text-2xl font-semibold text-[#061f12]">What We Do NOT Collect</h3>
			<ul class="list-inside list-disc space-y-2">
				<li>Email addresses (we use Firebase user IDs for identification)</li>
				<li>Precise location data (GPS coordinates)</li>
				<li>Health or personal medical information</li>
				<li>Biometric data</li>
				<li>Payment information (we do not process payments)</li>
				<li>Image analysis or metadata from uploaded plant photos</li>
			</ul>
			
			<p>
				Uploaded plant photos are stored as-is for your personal plant collection management. We do not analyze, process, or derive any insights from these images.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">4. Purposes of Data Processing</h2>
			<p>We process your data exclusively for the following purposes:</p>
			<ul class="list-inside list-disc space-y-2">
				<li>To provide and maintain the Water Me service</li>
				<li>To store and manage your personal plant collection</li>
				<li>To deliver personalized push notifications for plant care reminders based on your settings</li>
				<li>To ensure technical functionality and error-free operation of the service</li>
				<li>To diagnose issues and provide technical support</li>
				<li>To gather analytics to improve and optimize the service</li>
				<li>To comply with legal obligations under GDPR and DSGVO</li>
			</ul>

			<h2 class="text-3xl font-semibold text-[#061f12]">5. Legal Basis for Data Processing</h2>
			<p>
				The processing of your personal data is based on the following legal grounds under GDPR:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li><strong>Article 6(1)(a) GDPR:</strong> Processing based on your explicit consent (e.g., account creation, notification preferences)</li>
				<li><strong>Article 6(1)(b) GDPR:</strong> Processing necessary for the performance of a contract (providing the service)</li>
				<li><strong>Article 6(1)(f) GDPR:</strong> Processing based on our legitimate interests (service improvement, security, analytics)</li>
				<li><strong>Article 6(1)(c) GDPR:</strong> Processing required to comply with legal obligations</li>
			</ul>

			<h2 class="text-3xl font-semibold text-[#061f12]">6. Data Storage and Hosting</h2>
			<p>
				Your data is stored securely by our hosting providers. We use multiple providers to ensure reliability and data redundancy:
			</p>

			<h3 class="text-2xl font-semibold text-[#061f12]">Hetzner</h3>
			<p>
				<strong>Provider:</strong> Hetzner Online GmbH, Industriestr. 25, 91710 Gunzenhausen, Germany<br />
				<strong>Purpose:</strong> Cloud infrastructure and database hosting
			</p>
			<p>
				We use Hetzner on the basis of Art. 6(1)(f) GDPR. We have a legitimate interest in the most reliable depiction of our service possible. If appropriate consent has been obtained, the processing is carried out exclusively on the basis of Art. 6(1)(a) GDPR and § 25 (1) TDDDG, insofar the consent includes the storage of cookies or the access to information in the user's end device (e.g., device fingerprinting) within the meaning of the TDDDG. This consent can be revoked at any time.
			</p>
			<p>
				<strong>Data Protection:</strong> We have concluded a data processing agreement (DPA) for the use of Hetzner. This is a contract mandated by data privacy laws that guarantees that they process personal data of our service's users only based on our instructions and in compliance with the GDPR. For details, please view the data privacy policy of Hetzner: <a href="https://www.hetzner.com/de/rechtliches/datenschutz" class="text-[#00ee57] hover:underline">Hetzner's Privacy Policy</a>.
			</p>

			<h3 class="text-2xl font-semibold text-[#061f12]">IONOS</h3>
			<p>
				<strong>Provider:</strong> IONOS SE, Elgendorfer Str. 57, 56410 Montabaur, Germany<br />
				<strong>Purpose:</strong> Website and application server hosting
			</p>
			<p>
				Whenever you visit our service, IONOS records various logfiles along with your IP addresses. We use IONOS on the basis of Art. 6(1)(f) GDPR. Our company has a legitimate interest in presenting a service that is as dependable as possible. If appropriate consent has been obtained, the processing is carried out exclusively on the basis of Art. 6(1)(a) GDPR and § 25 (1) TDDDG, insofar the consent includes the storage of cookies or the access to information in the user's end device (e.g., device fingerprinting) within the meaning of the TDDDG. This consent can be revoked at any time.
			</p>
			<p>
				<strong>Data Protection:</strong> We have concluded a data processing agreement (DPA) for the use of IONOS. This is a contract mandated by data privacy laws that guarantees that they process personal data of our service's users only based on our instructions and in compliance with the GDPR. For details, please consult the data privacy policy of IONOS: <a href="https://www.ionos.de/terms-gtc/terms-privacy" class="text-[#00ee57] hover:underline">IONOS's Privacy Policy</a>.
			</p>

			<h3 class="text-2xl font-semibold text-[#061f12]">OVH</h3>
			<p>
				<strong>Provider:</strong> OVH GmbH, Kelsterbacher Str. 8, 65451 Kelsterbach, Germany<br />
				<strong>Purpose:</strong> Backup and disaster recovery infrastructure
			</p>
			<p>
				We use OVH on the basis of Art. 6(1)(f) GDPR. We have a legitimate interest in the most reliable depiction of our service possible. If appropriate consent has been obtained, the processing is carried out exclusively on the basis of Art. 6(1)(a) GDPR and § 25 (1) TDDDG, insofar the consent includes the storage of cookies or the access to information in the user's end device (e.g., device fingerprinting) within the meaning of the TDDDG. This consent can be revoked at any time.
			</p>
			<p>
				<strong>Data Protection:</strong> We have concluded a data processing agreement (DPA) for the use of OVH. This is a contract mandated by data privacy laws that guarantees that they process personal data of our service's users only based on our instructions and in compliance with the GDPR. For details, please view the data privacy policy of OVH: <a href="https://www.ovhcloud.com/de/personal-data-protection/" class="text-[#00ee57] hover:underline">OVH's Privacy Policy</a>.
			</p>

			<h3 class="text-2xl font-semibold text-[#061f12]">Cloudflare</h3>
			<p>
				<strong>Provider:</strong> Cloudflare Inc., 101 Townsend St., San Francisco, CA 94107, USA
			</p>
			<p>
				We use the "Cloudflare" service provided by Cloudflare Inc. Cloudflare offers a content delivery network with DNS that is available worldwide. As a result, the information transfer that occurs between your browser and our service is technically routed via Cloudflare's network. This enables Cloudflare to analyze data transactions between your browser and our service and to work as a filter between our servers and potentially malicious data traffic from the Internet. In this context, Cloudflare may also use cookies or other technologies deployed to recognize Internet users, which shall, however, only be used for the herein described purpose.
			</p>
			<p>
				The use of Cloudflare is based on our legitimate interest in a provision of our service offerings that is as error free and secure as possible (Art. 6(1)(f) GDPR).
			</p>
			<p>
				Data transmission to the US is based on the Standard Contractual Clauses (SCC) of the European Commission. For more information on Cloudflare's security precautions and data privacy policies, please follow this link: <a href="https://www.cloudflare.com/privacypolicy/" class="text-[#00ee57] hover:underline">Cloudflare's Privacy Policy</a>.
			</p>
			<p>
				The company is certified in accordance with the "EU-US Data Privacy Framework" (DPF). The DPF is an agreement between the European Union and the US, which is intended to ensure compliance with European data protection standards for data processing in the US. Every company certified under the DPF is obliged to comply with these data protection standards. For more information, please contact the provider under the following link: <a href="https://www.dataprivacyframework.gov/s/participant-search/participant-detail?contact=true&id=a2zt0000000GnZKAA0&status=Active" class="text-[#00ee57] hover:underline">Data Privacy Framework Participant Search</a>.
			</p>
			<p>
				<strong>Data Protection:</strong> We have concluded a data processing agreement (DPA) for the use of Cloudflare. This is a contract mandated by data privacy laws that guarantees that they process personal data of our service's users only based on our instructions and in compliance with the GDPR.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">7. Authentication and Notifications</h2>

			<h3 class="text-2xl font-semibold text-[#061f12]">Firebase Authentication</h3>
			<p>
				We use Firebase Authentication (provided by Google LLC, 1600 Amphitheatre Parkway, Mountain View, CA 94043, USA) to securely manage user authentication and account management. Firebase handles your login credentials, authentication tokens, and user identity verification. Your authentication data is processed according to Google's Firebase data processing terms and the Standard Contractual Clauses (SCC) for international data transfers to the United States.
			</p>
			<p>
				The processing of authentication data through Firebase is based on Art. 6(1)(b) GDPR, as it is necessary for the performance of our service contract with you. Firebase acts as a data processor on our behalf and has committed to processing personal data only according to our instructions and in compliance with the GDPR.
			</p>
			<p>
				For detailed information about how Google processes data through Firebase, please see <a href="https://firebase.google.com/support/privacy" class="text-[#00ee57] hover:underline">Firebase's Privacy Policy</a> and <a href="https://cloud.google.com/terms/data-processing-terms" class="text-[#00ee57] hover:underline">Google's Data Processing Terms</a>.
			</p>
			<p>
				Data transmission to the United States is based on the Standard Contractual Clauses (SCC) of the European Commission and is necessary for the provision of the authentication service. Google LLC is certified under the EU-US Data Privacy Framework (DPF), which ensures an adequate level of data protection.
			</p>

			<h3 class="text-2xl font-semibold text-[#061f12]">Firebase Cloud Messaging (Push Notifications)</h3>
			<p>
				Push notifications for plant care reminders are delivered through Firebase Cloud Messaging (FCM), provided by Google LLC. FCM is a messaging service that allows us to send timely, personalized notifications to your device based on your plant care schedules and notification preferences that you have configured within the application.
			</p>
			<p>
				To enable push notifications, the application obtains a device token from FCM. This token is stored in our database and associated with your user account. The device token is a unique identifier used only for the purpose of delivering notifications to your specific device. Device tokens may be refreshed periodically by FCM, and we manage these tokens according to your notification preferences.
			</p>
			<p>
				The processing of notification data through Firebase Cloud Messaging is based on Art. 6(1)(a) GDPR (your explicit consent to receive notifications) and Art. 6(1)(b) GDPR (necessary for performance of the service). You can withdraw your consent to receive notifications at any time by disabling notifications in your app settings or through your device notification settings. Disabling notifications will not affect your ability to use the rest of the Water Me service.
			</p>
			<p>
				When you uninstall the application or delete your account, your device tokens are automatically deleted from our systems and no further notifications will be sent to that device.
			</p>
			<p>
				For more information about Firebase Cloud Messaging and how Google processes notification data, please see <a href="https://firebase.google.com/support/privacy" class="text-[#00ee57] hover:underline">Firebase's Privacy Policy</a>.
			</p>

			<h3 class="text-2xl font-semibold text-[#061f12]">Google Services - Data Processing and Legal Basis</h3>
			<p>
				All data processing through Google services (Firebase Authentication and Firebase Cloud Messaging) is subject to a data processing agreement (DPA) between us and Google LLC. This agreement ensures that Google processes personal data only on our instructions and in compliance with the GDPR.
			</p>
			<p>
				Data transfers to the United States and other countries where Google operates servers are based on the Standard Contractual Clauses (SCC) of the European Commission. These clauses provide appropriate safeguards for the transfer of personal data outside the European Economic Area. Google LLC is certified under the EU-US Data Privacy Framework (DPF), which provides an additional layer of protection for data transferred to the United States. For more information, visit the <a href="https://www.dataprivacyframework.gov/s/participant-search/participant-detail?contact=true&id=a2zt0000000L0ZJAA4&status=Active" class="text-[#00ee57] hover:underline">Data Privacy Framework Participant Search</a>.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">8. Data Retention and Deletion</h2>
			<p>
				Your personal data is retained only as long as necessary to provide the Water Me service. Upon deletion of your account, all associated personal data is permanently deleted from our active systems within 30 days.
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li><strong>Account Data:</strong> Deleted immediately upon account deletion</li>
				<li><strong>Plant Data:</strong> Deleted immediately upon account deletion</li>
				<li><strong>Notification Settings:</strong> Deleted immediately upon account deletion</li>
				<li><strong>Server Logs:</strong> Typically retained for 30 days for security and diagnostics purposes, then deleted</li>
				<li><strong>Backup Data:</strong> Retained for up to 90 days for disaster recovery purposes only, then permanently deleted</li>
			</ul>
			<p>
				If you request deletion of your account, we will erase all associated personal data within 30 days. Backup copies may persist for up to 90 days to support disaster recovery but will not be accessible or used beyond this period.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">9. Data Security</h2>
			<p>
				The security of your data is extremely important to us. We implement industry-standard security measures:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li>HTTPS/TLS encryption for all data transmitted over the internet</li>
				<li>Encrypted storage of sensitive data in our databases</li>
				<li>Regular security assessments and updates</li>
				<li>Access controls limiting employee access to personal data</li>
				<li>Data Processing Agreements with all service providers</li>
			</ul>
			<p>
				However, no method of transmission over the internet or electronic storage is 100% secure. While we use commercially acceptable means to protect your data, we cannot guarantee absolute security.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">10. Your Rights Under GDPR and DSGVO</h2>
			<p>
				Under the GDPR and DSGVO, you have the following rights regarding your personal data:
			</p>

			<h3 class="text-2xl font-semibold text-[#061f12]">Right to Access (Article 15 GDPR)</h3>
			<p>
				You have the right to receive confirmation of whether we process your personal data and to obtain a copy of the data we hold about you.
			</p>

			<h3 class="text-2xl font-semibold text-[#061f12]">Right to Rectification (Article 16 GDPR)</h3>
			<p>
				You have the right to request that we correct any inaccurate or incomplete personal data we hold about you.
			</p>

			<h3 class="text-2xl font-semibold text-[#061f12]">Right to Erasure (Article 17 GDPR)</h3>
			<p>
				You have the right to request the deletion of your personal data in certain circumstances, including when the data is no longer necessary for the purposes for which it was collected. This is known as the "right to be forgotten."
			</p>

			<h3 class="text-2xl font-semibold text-[#061f12]">Right to Restrict Processing (Article 18 GDPR)</h3>
			<p>
				You have the right to request that we restrict the processing of your personal data while we verify its accuracy or in other circumstances.
			</p>

			<h3 class="text-2xl font-semibold text-[#061f12]">Right to Data Portability (Article 20 GDPR)</h3>
			<p>
				You have the right to receive your personal data in a structured, commonly-used, and machine-readable format, and to transmit that data to another controller.
			</p>

			<h3 class="text-2xl font-semibold text-[#061f12]">Right to Object (Article 21 GDPR)</h3>
			<p>
				You have the right to object to processing of your personal data based on our legitimate interests.
			</p>

			<h3 class="text-2xl font-semibold text-[#061f12]">Right to Lodge a Complaint</h3>
			<p>
				If you believe we have violated your data protection rights, you have the right to lodge a complaint with your national data protection authority. In Germany, this is the state data protection officer of your state.
			</p>

			<h3 class="text-2xl font-semibold text-[#061f12]">Right to Withdraw Consent</h3>
			<p>
				Where we rely on your consent for processing data, you have the right to withdraw that consent at any time. This will not affect the lawfulness of processing carried out before your withdrawal.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">11. How to Exercise Your Rights</h2>
			<p>
				To exercise any of your data protection rights, please contact us at:
			</p>
			<p>
				<strong>Email:</strong> ben@water-me.app<br />
				<strong>Phone:</strong> +49 1516 4877269<br />
				<strong>Mail:</strong> Ben Schiemann, Reinbeker Weg 70, 21465 Wentorf, Germany
			</p>
			<p>
				We will respond to your request within 30 days as required by law. We may need to verify your identity before processing your request.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">12. Third-Party Links</h2>
			<p>
				Our website and application may contain links to external websites and services that are not operated by us. This privacy policy does not apply to these third-party services, and we have no control over their privacy practices. We encourage you to review the privacy policies of any external services before providing them with personal information.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">13. Web Analytics</h2>
			<p>
				We use the web analytics tool Simple Analytics for our website to understand how users interact with our services and to improve the functionality and user experience of our website. The service provider is the Dutch company Simple Analytics, Hooftlaan 4, 1401 ED, Bussum, Noord-Holland, The Netherlands.
			</p>
			<p>
				Through Simple Analytics, we collect data about your browsing patterns and interactions with our website. There is a possibility that your browsing patterns will be statistically analyzed when you visit our website. Such analysis is performed to understand usage patterns and to optimize the website for a better user experience.
			</p>
			<p>
				The processing of data through Simple Analytics is based on Art. 6(1)(f) GDPR. We have a legitimate interest in analyzing website usage to improve our services. Simple Analytics is designed with privacy in mind and collects minimal data while providing valuable insights. You can learn more about the data processed through the use of Simple Analytics and their commitment to privacy in their <a href="https://simpleanalytics.com/privacy-policy" class="text-[#00ee57] hover:underline">Privacy Policy</a>.
			</p>
			<p>
				Simple Analytics processes data according to their privacy policy and applicable data protection laws. We have ensured that appropriate data protection measures are in place for this third-party service.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">14. Server Log Files</h2>
			<p>
				The provider of our website and application automatically collects and stores information in so-called server log files, which your browser communicates to us automatically. The information recorded in server log files comprises:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li>The type and version of browser used</li>
				<li>The used operating system</li>
				<li>Referrer URL</li>
				<li>The hostname of the accessing computer</li>
				<li>The time of the server inquiry</li>
				<li>The IP address</li>
			</ul>
			<p>
				This data is not merged with other data sources. Server log data is recorded on the basis of Art. 6(1)(f) GDPR. The operator of the website and application has a legitimate interest in the technically error free depiction and the optimization of our services. In order to achieve this, server log files must be recorded. Server log files are typically retained for 30 days for security and diagnostic purposes, then deleted.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">14. Server Log Files</h2>
			<p>
				The provider of our website and application automatically collects and stores information in so-called server log files, which your browser communicates to us automatically. The information recorded in server log files comprises:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li>The type and version of browser used</li>
				<li>The used operating system</li>
				<li>Referrer URL</li>
				<li>The hostname of the accessing computer</li>
				<li>The time of the server inquiry</li>
				<li>The IP address</li>
			</ul>
			<p>
				This data is not merged with other data sources. Server log data is recorded on the basis of Art. 6(1)(f) GDPR. The operator of the website and application has a legitimate interest in the technically error free depiction and the optimization of our services. In order to achieve this, server log files must be recorded. Server log files are typically retained for 30 days for security and diagnostic purposes, then deleted.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">15. SSL and TLS Encryption</h2>
			<p>
				For security reasons and to protect the transmission of confidential content, such as your plant data, authentication information, and personal preferences that you submit to us, our website and application use SSL (Secure Sockets Layer) or TLS (Transport Layer Security) encryption protocols.
			</p>
			<p>
				You can recognize an encrypted connection by checking whether the address line of the browser switches from "http://" to "https://" and also by the appearance of a lock icon in the browser line. If SSL or TLS encryption is activated, data you transmit to us cannot be read by third parties during transmission.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">15. SSL and TLS Encryption</h2>
			<p>
				For security reasons and to protect the transmission of confidential content, such as your plant data, authentication information, and personal preferences that you submit to us, our website and application use SSL (Secure Sockets Layer) or TLS (Transport Layer Security) encryption protocols.
			</p>
			<p>
				You can recognize an encrypted connection by checking whether the address line of the browser switches from "http://" to "https://" and also by the appearance of a lock icon in the browser line. If SSL or TLS encryption is activated, data you transmit to us cannot be read by third parties during transmission.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">16. Cookies and Tracking Technologies</h2>
			<p>
				The Water Me website and mobile application do not use cookies. The mobile app uses browser localStorage only to store your language preference and Firebase authentication tokens locally on your device. This data remains entirely on your device and is not transmitted to our servers.
			</p>
			<p>
				Firebase authentication tokens are managed by Firebase and are subject to Firebase's privacy policies and security standards.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">16. Cookies and Tracking Technologies</h2>
			<p>
				The Water Me website and mobile application do not use cookies. The mobile app uses browser localStorage only to store your language preference and Firebase authentication tokens locally on your device. This data remains entirely on your device and is not transmitted to our servers.
			</p>
			<p>
				Firebase authentication tokens are managed by Firebase and are subject to Firebase's privacy policies and security standards.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">17. Children's Privacy</h2>
			<p>
				The Water Me service is not intended for children under the age of 16. We do not knowingly collect personal information from children under 16. If we become aware that we have collected personal data from a child under 16 without verifiable parental consent, we will promptly delete such information.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">17. Children's Privacy</h2>
			<p>
				The Water Me service is not intended for children under the age of 16. We do not knowingly collect personal information from children under 16. If we become aware that we have collected personal data from a child under 16 without verifiable parental consent, we will promptly delete such information.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">18. Marketing and Advertising</h2>
			<p>
				We use social media platforms to promote and advertise the Water Me application. These platforms operate under their own privacy policies which are separate from this Privacy Policy. When you interact with our content on these platforms, please be aware that their terms and data practices apply.
			</p>
			<p>
				Our presence on these platforms includes:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li>TikTok</li>
				<li>Instagram</li>
				<li>YouTube</li>
				<li>GitHub</li>
				<li>Bluesky</li>
			</ul>
			<p>
				We do not collect additional personal data through these promotional activities beyond what the respective social media platforms collect through their own systems. The social media platforms themselves determine how they process data collected through interactions with our accounts and content. For information about how these platforms handle your data, please consult their respective privacy policies:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li><a href="https://www.tiktok.com/legal/privacy-policy" class="text-[#00ee57] hover:underline">TikTok Privacy Policy</a></li>
				<li><a href="https://www.instagram.com/legal/privacy/" class="text-[#00ee57] hover:underline">Instagram Privacy Policy</a></li>
				<li><a href="https://www.youtube.com/howyoutubeworks/our-commitments/protecting-user-data/" class="text-[#00ee57] hover:underline">YouTube Privacy Policy</a></li>
				<li><a href="https://docs.github.com/en/site-policy/privacy-policies/github-privacy-statement" class="text-[#00ee57] hover:underline">GitHub Privacy Statement</a></li>
				<li><a href="https://bsky.social/about/support/privacy-policy" class="text-[#00ee57] hover:underline">Bluesky Privacy Policy</a></li>
			</ul>
			<p>
				When you visit our profiles or interact with our content on social media platforms, those platforms may collect data about your visit, your interactions, and your device information according to their own data collection practices. We have no control over these activities.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">18. Marketing and Advertising</h2>
			<p>
				We use social media platforms to promote and advertise the Water Me application. These platforms operate under their own privacy policies which are separate from this Privacy Policy. When you interact with our content on these platforms, please be aware that their terms and data practices apply.
			</p>
			<p>
				Our presence on these platforms includes:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li>TikTok</li>
				<li>Instagram</li>
				<li>YouTube</li>
				<li>GitHub</li>
				<li>Bluesky</li>
			</ul>
			<p>
				We do not collect additional personal data through these promotional activities beyond what the respective social media platforms collect through their own systems. The social media platforms themselves determine how they process data collected through interactions with our accounts and content. For information about how these platforms handle your data, please consult their respective privacy policies:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li><a href="https://www.tiktok.com/legal/privacy-policy" class="text-[#00ee57] hover:underline">TikTok Privacy Policy</a></li>
				<li><a href="https://www.instagram.com/legal/privacy/" class="text-[#00ee57] hover:underline">Instagram Privacy Policy</a></li>
				<li><a href="https://www.youtube.com/howyoutubeworks/our-commitments/protecting-user-data/" class="text-[#00ee57] hover:underline">YouTube Privacy Policy</a></li>
				<li><a href="https://docs.github.com/en/site-policy/privacy-policies/github-privacy-statement" class="text-[#00ee57] hover:underline">GitHub Privacy Statement</a></li>
				<li><a href="https://bsky.social/about/support/privacy-policy" class="text-[#00ee57] hover:underline">Bluesky Privacy Policy</a></li>
			</ul>
			<p>
				When you visit our profiles or interact with our content on social media platforms, those platforms may collect data about your visit, your interactions, and your device information according to their own data collection practices. We have no control over these activities.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">19. Information on Data Transfer to Third Countries</h2>
			<p>
				We use services from companies located in the United States and other countries outside the European Union and European Economic Area. These countries may not have the same level of data protection as the EU.
			</p>
			<p>
				We would like to point out that the United States, as a country with which an adequacy decision exists through the EU-US Data Privacy Framework (DPF), generally has a level of data protection comparable to that of the EU. Data transfer to the US is permitted if the recipient is certified under the "EU-US Data Privacy Framework" (DPF) or has appropriate additional safeguards such as Standard Contractual Clauses (SCC). All transfers to third countries in this Privacy Policy are protected by appropriate legal mechanisms including Standard Contractual Clauses and/or adequacy decisions.
			</p>
			<p>
				Information on transfers to third countries, including the data recipients and legal safeguards, can be found in the relevant sections of this Privacy Policy where each service is described.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">19. Information on Data Transfer to Third Countries</h2>
			<p>
				We use services from companies located in the United States and other countries outside the European Union and European Economic Area. These countries may not have the same level of data protection as the EU.
			</p>
			<p>
				We would like to point out that the United States, as a country with which an adequacy decision exists through the EU-US Data Privacy Framework (DPF), generally has a level of data protection comparable to that of the EU. Data transfer to the US is permitted if the recipient is certified under the "EU-US Data Privacy Framework" (DPF) or has appropriate additional safeguards such as Standard Contractual Clauses (SCC). All transfers to third countries in this Privacy Policy are protected by appropriate legal mechanisms including Standard Contractual Clauses and/or adequacy decisions.
			</p>
			<p>
				Information on transfers to third countries, including the data recipients and legal safeguards, can be found in the relevant sections of this Privacy Policy where each service is described.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">20. Revocation of Your Consent to the Processing of Data</h2>
			<p>
				A wide range of data processing transactions are possible only subject to your express consent. You can revoke at any time any consent you have already given us. This shall be without prejudice to the lawfulness of any data collection that occurred prior to your revocation.
			</p>
			<p>
				To revoke your consent, please contact us using the contact information provided in section 29 (Contact Information) of this Privacy Policy. For consent related to push notifications, you can also withdraw consent through your app settings or device notification settings.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">20. Right to Object to the Collection of Data</h2>
			<p>
				IN THE EVENT THAT DATA ARE PROCESSED ON THE BASIS OF ART. 6(1)(E) OR (F) GDPR, YOU HAVE THE RIGHT TO AT ANY TIME OBJECT TO THE PROCESSING OF YOUR PERSONAL DATA BASED ON GROUNDS ARISING FROM YOUR UNIQUE SITUATION. THIS ALSO APPLIES TO ANY PROFILING BASED ON THESE PROVISIONS. TO DETERMINE THE LEGAL BASIS ON WHICH ANY PROCESSING OF DATA IS BASED, PLEASE CONSULT THIS DATA PROTECTION DECLARATION. IF YOU LOG AN OBJECTION, WE WILL NO LONGER PROCESS YOUR AFFECTED PERSONAL DATA, UNLESS WE ARE IN A POSITION TO PRESENT COMPELLING PROTECTION WORTHY GROUNDS FOR THE PROCESSING OF YOUR DATA THAT OUTWEIGH YOUR INTERESTS, RIGHTS AND FREEDOMS OR IF THE PURPOSE OF THE PROCESSING IS THE CLAIMING, EXERCISING OR DEFENCE OF LEGAL ENTITLEMENTS (OBJECTION PURSUANT TO ART. 21(1) GDPR).
			</p>
			<p>
				IF YOUR PERSONAL DATA IS BEING PROCESSED IN ORDER TO ENGAGE IN DIRECT ADVERTISING, YOU HAVE THE RIGHT TO OBJECT TO THE PROCESSING OF YOUR AFFECTED PERSONAL DATA FOR THE PURPOSES OF SUCH ADVERTISING AT ANY TIME. THIS ALSO APPLIES TO PROFILING TO THE EXTENT THAT IT IS AFFILIATED WITH SUCH DIRECT ADVERTISING. IF YOU OBJECT, YOUR PERSONAL DATA WILL SUBSEQUENTLY NO LONGER BE USED FOR DIRECT ADVERTISING PURPOSES (OBJECTION PURSUANT TO ART. 21(2) GDPR).
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">21. Right to Demand Processing Restrictions</h2>
			<p>
				You have the right to demand the imposition of restrictions as far as the processing of your personal data is concerned. To do so, you may contact us at any time using the contact information provided in section 18 (Contact Information).
			</p>
			<p>
				The right to demand restriction of processing applies in the following cases:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li>In the event that you should dispute the correctness of your data archived by us, we will usually need some time to verify this claim. During the time that this investigation is ongoing, you have the right to demand that we restrict the processing of your personal data.</li>
				<li>If the processing of your personal data was/is conducted in an unlawful manner, you have the option to demand the restriction of the processing of your data instead of demanding the eradication of this data.</li>
				<li>If we do not need your personal data any longer and you need it to exercise, defend or claim legal entitlements, you have the right to demand the restriction of the processing of your personal data instead of its eradication.</li>
				<li>If you have raised an objection pursuant to Art. 21(1) GDPR, your rights and our rights will have to be weighed against each other. As long as it has not been determined whose interests prevail, you have the right to demand a restriction of the processing of your personal data.</li>
				<li>If you have restricted the processing of your personal data, these data – with the exception of their archiving – may be processed only subject to your consent or to claim, exercise or defend legal entitlements or to protect the rights of other natural persons or legal entities or for important public interest reasons cited by the European Union or a member state of the EU.</li>
			</ul>

			<h2 class="text-3xl font-semibold text-[#061f12]">22. Editorial Responsibility and Content Management</h2>
			<p>
				Editorial responsibility for the content of the Water Me website and application rests with:
			</p>
			<p>
				<strong>Ben Schiemann</strong><br />
				Reinbeker Weg 70<br />
				21465 Wentorf<br />
				Germany<br />
			</p>
			<p>
				We strive to provide accurate and up-to-date information in all content. However, we do not guarantee the completeness, accuracy, timeliness, or reliability of any information presented. Any reliance you place on such information is strictly at your own risk.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">23. Technical Responsibility and System Operator</h2>
			<p>
				Technical responsibility for the operation and maintenance of the Water Me application and website rests with:
			</p>
			<p>
				<strong>Ben Schiemann</strong><br />
				Reinbeker Weg 70<br />
				21465 Wentorf<br />
				Germany<br />
			</p>
			<p>
				While we maintain technical systems to the best of our ability and knowledge, we cannot guarantee uninterrupted availability of the service. We perform regular maintenance and updates that may occasionally result in service interruptions. We are not liable for any damages resulting from temporary service unavailability unless caused by our gross negligence.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">24. Third-Party Components and Open Source Libraries</h2>
			<p>
				The Water Me application and website incorporate components and libraries developed by third parties, many of which are open source software. We acknowledge and respect the licenses under which these components are distributed.
			</p>
			<p>
				Key open source frameworks and libraries used include but are not limited to:
			</p>
			<ul class="list-inside list-disc space-y-2">
				<li><strong>SvelteKit:</strong> Framework for building web applications (BSD 3-Clause License)</li>
				<li><strong>Svelte:</strong> JavaScript compiler for reactive user interfaces (MIT License)</li>
				<li><strong>Tailwind CSS:</strong> Utility-first CSS framework (MIT License)</li>
				<li><strong>Go:</strong> Open source programming language used for backend (BSD License)</li>
				<li><strong>MongoDB Driver for Go:</strong> Official MongoDB driver (Apache 2.0 License)</li>
				<li><strong>Firebase SDK:</strong> Google's application development platform (Apache 2.0 License)</li>
				<li><strong>Capacitor:</strong> Cross-platform app framework (MIT License)</li>
			</ul>
			<p>
				The complete source code of the Water Me application is publicly available on GitHub under an appropriate open source license. We encourage contributions, bug reports, and feature requests from the community. For access to the source code and license information, please visit our <a href="https://github.com" class="text-[#00ee57] hover:underline">GitHub repository</a>.
			</p>
			<p>
				All third-party libraries and their respective licenses are documented in our project's dependency management files. By using Water Me, you acknowledge your acceptance of the licenses of all incorporated open source components.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">25. Disclaimer of Warranties and Limitations of Liability</h2>
			<p>
				<strong>As Provided Basis:</strong> The Water Me application and website are provided on an "as is" and "as available" basis without warranties of any kind, either express or implied. We disclaim all warranties, including but not limited to, merchantability, fitness for a particular purpose, and non-infringement.
			</p>
			<p>
				<strong>No Liability:</strong> To the fullest extent permitted by law, in no event shall Water Me, its operators, contributors, or any affiliated parties be liable for any indirect, incidental, special, consequential, or punitive damages, including but not limited to loss of profits, data loss, or business interruption, arising out of or in connection with the use of or inability to use the service, even if advised of the possibility of such damages.
			</p>
			<p>
				<strong>User Responsibility:</strong> Users are solely responsible for their use of the Water Me service and for ensuring that their use complies with all applicable laws and regulations. We are not responsible for content you create, upload, or share through the service.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">26. Commercial Nature and Open Source Status</h2>
			<p>
				Water Me is a commercial application provided by Ben Schiemann for general use. While the application's source code is open source and freely available for inspection and contribution, the operation of the service may involve commercial aspects including potential future monetization, advertisements, or premium features.
			</p>
			<p>
				<strong>Current Status:</strong> The Water Me application is provided free of charge to users. Any future changes to pricing or premium features will be clearly communicated in advance.
			</p>
			<p>
				The open source nature of our codebase reflects our commitment to transparency and community participation. Users are encouraged to review the source code, identify security issues, and contribute improvements.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">27. Changes to This Privacy Policy</h2>
			<p>
				We may update this privacy policy from time to time to reflect changes in our practices, technology, legal requirements, or other factors. We will notify you of any material changes by updating the "Last updated" date at the top of this policy. Your continued use of the service after such modifications constitutes your acceptance of the updated privacy policy.
			</p>
			<p>
				When required by law, we will obtain your consent or provide you with the opportunity to consent to or decline any new use of your personal information.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">28. Data Protection Officer and Supervisory Authority</h2>
			<p>
				If you have questions about data protection matters or wish to file a complaint, you can contact the data protection authority in your country. For residents of Germany, you can contact your state's data protection officer (Landesdatenschutzbeauftragte). In the event of violations of the GDPR, data subjects are entitled to log a complaint with a supervisory agency, in particular in the member state where they usually maintain their domicile, place of work or at the place where the alleged violation occurred. The right to log a complaint is in effect regardless of any other administrative or court proceedings available as legal recourses.
			</p>

			<h2 class="text-3xl font-semibold text-[#061f12]">29. Contact Information</h2>
			<p>
				If you have any questions about this privacy policy or our data protection practices, please contact us:
			</p>
			<p>
				<strong>Ben Schiemann</strong><br />
				Reinbeker Weg 70<br />
				21465 Wentorf<br />
				Germany<br /><br />
				Email: ben@water-me.app<br />
				Phone: +49 1516 4877269
			</p>
			<p>
				We will respond to your requests, questions, and complaints regarding data protection matters within 30 days as required by law. We may need to verify your identity before processing your request to exercise your data subject rights.
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

	:global(h3) {
		margin: 1.5rem 0 0.75rem;
	}

	:global(p) {
		margin-bottom: 1rem;
	}

	:global(ul) {
		margin-bottom: 1rem;
	}

	a {
		color: #00ee57;
	}

	a:hover {
		text-decoration: underline;
	}
</style>
