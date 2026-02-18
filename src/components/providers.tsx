import { ThemeProvider } from "next-themes";
import type * as React from "react";

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
