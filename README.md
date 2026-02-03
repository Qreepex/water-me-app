# Plants

Plant management app with a Go backend and SvelteKit app.

## Structure

- backend/ - Go API server
- app/ - Mobile + Web app (SvelteKit + Capacitor)
- website/ - Marketing site (SvelteKit)

## Quick Start

```bash
# Backend
cd backend
go run .

# App
cd ../app
npm install
npm run dev

# Website
cd ../website
npm install
npm run dev
```

## OpenAPI Types

```bash
npx openapi-typescript ./backend/openapi.yaml -o ./app/src/lib/types/api.ts --root-types --root-types-no-schema-prefix --generate-path-params --enum --enum-values
```
