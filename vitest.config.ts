import { defineConfig } from "vitest/config";

const vitestConfig = defineConfig({
	test: {
		name: "Shapas Tuckshop",
		coverage: {
			enabled: true,
			provider: "istanbul",
		},
		setupFiles: ["./vitest.setup.tsx"],
	},
});

export default vitestConfig;
