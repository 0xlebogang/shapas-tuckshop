import "dotenv/config";
import { auth } from "@repo/better-auth";
import { Hono } from "hono";
import { cors } from "hono/cors";
import { logger } from "hono/logger";

const hono = new Hono();

hono.use("*", logger());
hono.use(
	"*",
	cors({
		credentials: true,
		allowMethods: ["GET", "POST", "PATCH", "DELETE", "OPTIONS"],
		origin: process.env.CORS_ALLOWED_ORIGINS?.split(",") || [
			"http://localhost:3000",
		],
	}),
);

hono.get("/health", (c) => {
	return c.json({
		status: "healthy",
		timestamp: new Date().toISOString(),
	});
});

hono.on(["POST", "GET"], "/auth/*", (c) => auth.handler(c.req.raw));

export default hono;
export type AppType = typeof hono;
