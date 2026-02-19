import type * as React from "react";

export interface AuthLayoutProps {
	children: Readonly<React.ReactNode>;
}

export default function AuthLayout({ children }: AuthLayoutProps) {
	return (
		<section className="flex min-h-screen bg-background px-4 py-16 md:py-32 dark:bg-transparent">
			{children}
		</section>
	);
}
