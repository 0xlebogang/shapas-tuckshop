import type * as React from "react";

export default function AuthLayout({
	children,
}: {
	children: React.ReactNode;
}) {
	return (
		<section className="flex min-h-screen px-4 py-16 md:py-32">
			{children}
		</section>
	);
}
