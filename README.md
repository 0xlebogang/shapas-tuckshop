<div align='center'>

# Shapas

A simple order management system for tuckshops.

</div>

> This project is currently in active development. Changes are expected to be frequent.

## Prerequisites

- Nodejs (v20.9+)
- PNPM (v10.28.0)
- Docker/Docker Desktop (latest stable version)
- Git (Any stable version)

## Getting started

To get the project up and running for development:

> Have a look at the [contributing guidline](./CONTRIBUTING.md) for more detailed docs of the structure and procedures in the repo.

1. Clone the repo

	```bash
	git clone https://github.com/0xlebogang/shapas-tuckshopt.git shapas
	cd shapas
	```

2. Install project dependencies

	```bash
	pnpm install
	```

3. Setup environment variables

	> There is a [helper script](./scripts/setup-env-files.js) that will create copies of the `.env.example` files and name them `.env`
	>
	> Be sure to run this script from the root of the project (same directory as `pnpm-lock.yaml`)

	```bash
	node scripts/setup-env-files.js
	```

4. Run the database server

	```bash
	pnpm turbo db:start
	```

5. Push the existing databse schema

	```bash
	pnpm turbo db:push
	```

6. Run all dev servers

	```bash
	pnpm dev
	```
