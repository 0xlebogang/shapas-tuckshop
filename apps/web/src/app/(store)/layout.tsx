import type * as React from "react";

export interface StoreRootLayoutProps {
	children: React.ReactNode;
}

export default function StoreRootLayout({ children }: StoreRootLayoutProps) {
	return <main>{children}</main>;
}
