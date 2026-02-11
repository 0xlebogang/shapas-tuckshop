/**@type {import('next').NextConfig} */
const nextConfig = {
	reactCompiler: true,
	transpilePackages: ["@repo/better-auth"],
	rewrites: async () => {
		return {
			beforeFiles: [
				{
					source: "/api/:path*",
					destination: `${process.env.NEXT_PUBLIC_API_URL}/:path*`,
				},
			],
		};
	},
};

export default nextConfig;
