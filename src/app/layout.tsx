import type * as React from "react";
import "./globals.css";
import Providers from "@/components/providers";

export interface RootLayoutProps {
	children: Readonly<React.ReactNode>;
}

export default function RootLayout({ children }: RootLayoutProps) {
	return (
		<html lang="en" suppressHydrationWarning>
			<body>
				<Providers>{children}</Providers>
			</body>
		</html>
	);
}
