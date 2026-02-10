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
		origin:
			process.env.NODE_ENV !== "production" ? "*" : "http://localhost:3000",
	}),
);

hono.get("/health", (c) => {
	return c.json({
		status: "healthy",
		timestamp: new Date().toISOString(),
	});
});

export default hono;
export type AppType = typeof hono;
