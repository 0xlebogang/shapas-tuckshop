import type * as React from "react";
import { ThemeProvider } from "next-themes";

export interface ProvidersProps {
	children: Readonly<React.ReactNode>;
}

export default function Providers({ children }: ProvidersProps) {
	return (
		<ThemeProvider
			enableSystem
			attribute="class"
			defaultTheme="system"
			disableTransitionOnChange
		>
			{children}
		</ThemeProvider>
	);
}
