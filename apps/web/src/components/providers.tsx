import { ThemeProvider } from "next-themes";
import type * as React from "react";

export default function Providers({ children }: { children: React.ReactNode }) {
	return (
		<ThemeProvider enableSystem defaultTheme="system" attribute="class">
			{children}
		</ThemeProvider>
	);
}
