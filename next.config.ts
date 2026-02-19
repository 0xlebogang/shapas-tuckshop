import type { NextConfig } from "next";
import "@/lib/env/server";
import "@/lib/env/client";
import { env } from "@/lib/env/server";

const nextConfig: NextConfig = {
	// Set 'output' and 'transpilePackages' when 'DOCKER_IMAGE' is enabled.
	...(env.DOCKER_IMAGE
		? {
				output: "standalone",
				transpilePackages: ["@t3-oss/env-nextjs", "@t3-oss/env-core"],
			}
		: {}),
};

export default nextConfig;
