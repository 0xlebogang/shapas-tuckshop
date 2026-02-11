import { createClient } from "@repo/better-auth/client";

export const authClient = createClient(
	process.env.NEXT_PUBLIC_API_URL as string,
	"/auth",
);
