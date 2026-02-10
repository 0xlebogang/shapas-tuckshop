import { db } from "@repo/drizzle";
import { betterAuth } from "better-auth";
import { drizzleAdapter } from "better-auth/adapters/drizzle";

export const auth = betterAuth({
	appName: "Shapas",
	baseURL: process.env.BASE_URL || "http://localhost:5000",
	basePath: "/auth",
	secret: process.env.BETTER_AUTH_SECRET as string,
	trustedOrigins: process.env.CORS_ALLOWED_ORIGINS
		? process.env.CORS_ALLOWED_ORIGINS.split(",")
		: ["http://localhost:3000"],

	database: drizzleAdapter(db, {
		provider: "pg",
	}),
});
