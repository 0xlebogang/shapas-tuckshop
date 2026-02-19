<div align='center'>

# Shapas Tuckshop

A simple order management platform for small tuckshops to better align with their customers.

</div>

> This project is still under active development and my go through frequent breaking changes

## Prerequisites

- Nodejs (v20.9+)
- PNPM (v10.28.0)
- A [Convex.dev](https://convex.dev) account
- Git (latest stable version)

## Running for development

1. Clone the repo

	```bash
	git clone https://github.com/0xlebogang/shapas-tuckshop.git shapas

	cd shapas
	```

2. Install dependencies

	```bash
	pnpm install
	```

3. Run all dev servers.

	> On the first run, you will be prompted by the Convex CLI to either create a new
	> account or to attach the current project to an existing one.

	```bash
	pnpm run dev
	```

	> Uses concurrently cli tool to run `convex dev` & `next dev` concurrently.
