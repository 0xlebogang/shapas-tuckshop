import Providers from "@/components/providers";
import "./globals.css";
import { Toaster } from "sonner";

export default function RootLayout({
	children,
}: {
	children: React.ReactNode;
}) {
	return (
		<html lang="en" suppressHydrationWarning>
			<body>
				<Providers>
					{children}
					<Toaster closeButton position="bottom-right" />
				</Providers>
			</body>
		</html>
	);
}
