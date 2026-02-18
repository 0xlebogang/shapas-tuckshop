import { defineConfig } from "vitest/config";

const vitestConfig = defineConfig({
	test: {
		name: "Shapas Tuckshop",
		coverage: {
			enabled: true,
			provider: "istanbul",
			reporter: ["text", "json", "json-summary", "lcov"],
		},
		setupFiles: ["./vitest.setup.tsx"],
	},
});

export default vitestConfig;
