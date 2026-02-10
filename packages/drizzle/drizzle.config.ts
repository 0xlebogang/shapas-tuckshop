import "dotenv/config";
import { defineConfig } from "drizzle-kit";

const drizzleConfig = defineConfig({
	dialect: "postgresql",
	out: "./drizzle",
	schema: "./src/schema",
	dbCredentials: {
		url: process.env.DATABASE_URL as string,
		ssl: process.env.NODE_ENV === "production",
	},
});

export default drizzleConfig;
