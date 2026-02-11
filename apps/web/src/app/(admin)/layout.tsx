import type * as React from "react";

export interface AdminRootLayoutProps {
	children: React.ReactNode;
}

export default function AdminRootLayout({ children }: AdminRootLayoutProps) {
	return <main>{children}</main>;
}
