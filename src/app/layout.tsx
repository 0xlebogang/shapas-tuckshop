import type * as React from "react";
import "./globals.css";
import { ConvexAuthNextjsServerProvider as ConvexAuthProvider } from "@convex-dev/auth/nextjs/server";
import Providers from "@/components/providers";

export interface RootLayoutProps {
	children: Readonly<React.ReactNode>;
}

export default function RootLayout({ children }: RootLayoutProps) {
	return (
		<ConvexAuthProvider>
			<html lang="en" suppressHydrationWarning>
				<body>
					<Providers>{children}</Providers>
				</body>
			</html>
		</ConvexAuthProvider>
	);
}
