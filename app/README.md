# App (SvelteKit + Capacitor)

Mobile / Web app for the Plants project.

## Requirements

- Node.js + npm
- Android Studio (for Android builds)

## Development

```bash
npm install
npm run dev
```

## Build

```bash
npm run build
```

## Quality

```bash
npm run format
npm run lint
npm run check
```

## API Types

```bash
npx openapi-typescript ../backend/openapi.yaml -o ./src/lib/types/api.ts --root-types --root-types-no-schema-prefix --generate-path-params --enum --enum-values
```

## Config

- API base URL: src/lib/constants.ts
