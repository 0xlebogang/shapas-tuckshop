"use client";

import { ConvexAuthNextjsProvider as ConvexProvider } from "@convex-dev/auth/nextjs";
import { ConvexReactClient } from "convex/react";
import { ThemeProvider } from "next-themes";
import type * as React from "react";
import { env } from "@/lib/env/client";

export interface ProvidersProps {
	children: Readonly<React.ReactNode>;
}

const convex = new ConvexReactClient(env.NEXT_PUBLIC_CONVEX_URL);

export default function Providers({ children }: ProvidersProps) {
	return (
		<ConvexProvider client={convex}>
			<ThemeProvider
				enableSystem
				attribute="class"
				defaultTheme="system"
				disableTransitionOnChange
			>
				{children}
			</ThemeProvider>
		</ConvexProvider>
	);
}
