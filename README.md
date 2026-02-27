<div align='center'>

# Shapas

A simple order management system built specically for tuckshops. You could probably use this if you are selling any goods.

</div>

## Development setup

### Prerequisites

- Node.js (v20.9+)
- PNPM (v10.28.0)
- Go (v1.24.11)
- air (latest stable version)
- GolangCI-Lint (latest stable version)
- Docker & Docker compose (latest stable versions)

### Running the system

1. Clone the repo

	```bash
	git clone https://github.com/0xlebogang/shapas-tuckshop.git shapas

	cd shapas
	```

2. Install project dependencies

	```bash
	pnpm install
	```

	> This will also automatically install the dependencies for the backend package (Go)

3. Create copies of all `.env.example` files and rename them to .env then populate the values as necessary

	> You can use the [helper script](./scripts/setup-env-files.js)

	```bash
	node ./scripts/setup-env-files.js
	```

4. Run all servers

	```bash
	pnpm run dev
	```

	refer to the [contributing guidline](./CONTRIBUTING.md) to get into more detail about development and practices.

---

<div align='center'>
	<small>Built by [Lebogang Phoshoko](https://lebophoshoko.dedyn.io)
</div>
