import { createEnv } from "@t3-oss/env-nextjs";
import * as z from "zod";

export const env = createEnv({
	emptyStringAsUndefined: true,
	server: {
		DOCKER_IMAGE: z.string().optional(),
		CONVEX_DEPLOYMENT: z
			.string()
			.regex(/[A-Za-z]+:[A-Za-z0-9]+/, "Invalid Convex deployment string"),
	},
	experimental__runtimeEnv: process.env,
});
