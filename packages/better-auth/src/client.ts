import { createAuthClient } from "better-auth/react";

export function createClient(baseURL: string, basePath: string) {
	return createAuthClient({
		baseURL,
		basePath,
	});
}
