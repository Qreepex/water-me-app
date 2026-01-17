# Internationalization (i18n) System

## Overview

The plant management app now includes a complete i18n system supporting English (en), German (de), and Spanish (es).

## File Structure

```
app/
├── i18n/
│   ├── en/
│   │   ├── common.json      # Common strings (10 keys)
│   │   ├── menu.json        # Menu strings (9 keys)
│   │   ├── auth.json        # Authentication page strings (13 keys)
│   │   ├── plants.json      # Plant overview strings (14 keys)
│   │   └── profile.json     # User profile page strings (17 keys)
│   ├── de/                  # German translations (complete)
│   │   ├── common.json
│   │   ├── menu.json
│   │   ├── auth.json
│   │   ├── plants.json
│   │   └── profile.json
│   └── es/                  # Spanish translations (complete)
│       ├── common.json
│       ├── menu.json
│       ├── auth.json
│       ├── plants.json
│       └── profile.json
├── src/
│   └── lib/
│       ├── i18n/
│       │   ├── index.ts     # Translation function implementation
│       │   └── types.ts     # TypeScript types for translations
│       └── stores/
│           └── language.ts  # Language store for reactive switching
```

## Usage

### Basic Translation (Synchronous)

```typescript
import { tSync } from '$lib/i18n';

const label = tSync('profile.username', language);
// Returns: "Username" (en), "Benutzername" (de), "Nombre de usuario" (es)
```

### Async Translation

```typescript
import { t } from '$lib/i18n';

const label = await t('auth.title', 'en');
// Returns: "Welcome Back"
```

### Using in Svelte Components

```svelte
<script lang="ts">
  import { languageStore } from '$lib/stores/language';
  import { tSync } from '$lib/i18n';
</script>

<h1>{tSync('menu.menu', $languageStore)}</h1>
<p>{tSync('common.appDescription', $languageStore)}</p>
```

## Translation Keys

### common.json

- `app` - Application name
- `appDescription` - App description
- `language` - Language label
- `logout` - Logout button text
- `loading` - Loading message
- `error` - Error message
- `success` - Success message
- `cancel` - Cancel button text
- `save` - Save button text
- `close` - Close button text

### menu.json

- `menu` - Menu title
- `userProfile` - User profile link
- `managePlants` - Manage plants link
- `website` - Website link
- `privacyPolicy` - Privacy policy link
- `impressum` - Legal notice link
- `build` - Build info label
- `version` - Version info label
- `pushNotifications` - Push notifications debug section

### auth.json

- `title` - Login page title
- `subtitle` - Login page subtitle
- `email` - Email field label
- `password` - Password field label
- `signIn` - Sign in button
- `signingIn` - Signing in state
- `createAccountTitle` - Signup page title
- `createAccountSubtitle` - Signup page subtitle
- `signUp` - Sign up button
- `creatingAccount` - Creating account state
- `dontHaveAccount` - "Don't have account?" link text
- `alreadyHaveAccount` - "Already have account?" link text
- `authenticationFailed` - Auth failed message
- `invalidEmailOrPassword` - Invalid credentials message
- `userAlreadyExists` - User exists message
- `passwordTooShort` - Password validation message
- `demoTip` - Demo credentials tip

### plants.json

- `myPlants` - My plants title
- `sortBy` - Sort by label
- `plantName` - Plant name column
- `lastWatered` - Last watered column
- `lastFertilized` - Last fertilized column
- `wateringFrequency` - Watering frequency label
- `sprayFrequency` - Spray frequency label
- `plants` - Plants (plural)
- `plant` - Plant (singular)
- `managePlants` - Manage plants title
- `loadingPlants` - Loading message
- `errorLoadingPlants` - Error message
- `noPlants` - No plants message
- `startAddingPlants` - Call to action
- `needsWater` - Water needed status
- `waterSoon` - Water soon status
- `watered` - Watered status
- `every` - "Every" (for intervals)
- `days` - Days (unit)

### profile.json

- `userProfile` - Profile page title
- `profileInformation` - Section title
- `username` - Username field label
- `email` - Email field label
- `language` - Language field label
- `updateProfile` - Update button text
- `updating` - Updating state
- `changePassword` - Password section title
- `currentPassword` - Current password field
- `newPassword` - New password field
- `confirmPassword` - Confirm password field
- `changePasswordButton` - Change button text
- `changing` - Changing state
- `profileUpdated` - Success message
- `passwordChanged` - Password changed message
- `passwordsDoNotMatch` - Validation message
- `passwordTooShort` - Validation message
- `currentPasswordIncorrect` - Error message

## Backend Integration

The language preference is stored in the user profile:

```go
type User struct {
    ID           string
    Username     string
    Email        string
    PasswordHash string
    Language     string  // 'en', 'de', or 'es'
    CreatedAt    string
}
```

### API Endpoints

**Get User Profile**

```
GET /api/user
Authorization: Bearer {token}
```

Response includes `language` field.

**Update User Profile**

```
PUT /api/user
Authorization: Bearer {token}
Body: {
    username?: string,
    email?: string,
    language?: string  // 'en', 'de', or 'es'
}
```

## Language Switching

Users can change their language preference via:

1. **Profile Page**: Language dropdown selector
2. **Burger Menu**: EN/DE/ES buttons (language changes globally and syncs to profile)

## Initialization

On app startup (`+layout.svelte`):

1. `initializeI18n()` - Preloads all translation files
2. `initializeLanguage()` - Sets language from:
   - User profile (if logged in)
   - Capacitor Preferences (if previously selected)
   - Browser language (en/de/es) or defaults to 'en'

## Persistence

Language preference is persisted via:

1. **Capacitor Preferences** - For offline access
2. **User Profile** - Synced to backend when updated
3. **Auth Store** - Reactive updates across app

## Fallback Behavior

If a translation key is missing:

1. App checks current language's JSON file
2. If not found, falls back to English
3. If still not found, returns the key itself

## Adding New Languages

To add a new language (e.g., French):

1. Create `app/i18n/fr/` directory
2. Create JSON files: `common.json`, `menu.json`, `auth.json`, `plants.json`, `profile.json`
3. Update `app/src/lib/i18n/index.ts` - add 'fr' to `locales` array
4. Update profile page language selector
5. Update burger menu language buttons
6. Update language store type if needed
