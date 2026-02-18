import { createEnv } from "@t3-oss/env-nextjs";
import * as z from "zod";

export const env = createEnv({
	emptyStringAsUndefined: true,
	client: {
		NEXT_PUBLIC_CONVEX_URL: z.url("Invalid Convex URL"),
		NEXT_PUBLIC_CONVEX_SITE_URL: z.url("Invalid Convex site URL"),
	},
	runtimeEnv: {
		NEXT_PUBLIC_CONVEX_URL: process.env.NEXT_PUBLIC_CONVEX_URL,
		NEXT_PUBLIC_CONVEX_SITE_URL: process.env.NEXT_PUBLIC_CONVEX_SITE_URL,
	},
});
