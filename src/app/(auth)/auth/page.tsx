import Link from "next/link";
import AppleAuthButton from "@/components/apple-auth-button";
import GoogleAuthButton from "@/components/google-auth-button";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

export default function Auth() {
	return (
		<form
			action=""
			className="bg-muted m-auto h-fit w-full max-w-sm overflow-hidden rounded-[calc(var(--radius)+.125rem)] border shadow-md shadow-zinc-950/5 dark:[--color-muted:var(--color-zinc-900)]"
		>
			<div className="bg-card -m-px rounded-[calc(var(--radius)+.125rem)] border p-8 pb-6">
				<div className="text-center">
					<Link href="/" aria-label="go home" className="mx-auto block w-fit">
						<span className="font-bold text-3xl md:text-5xl">Shapas</span>
					</Link>
					<h1 className="mb-1 mt-4 text-xl font-semibold">
						Sign In to Shapas Tuckshop
					</h1>
					<p className="text-sm">Welcome back! Sign in to continue</p>
				</div>

				<div className="mt-6 space-y-6">
					<div className="space-y-2">
						<Label htmlFor="email" className="block text-sm">
							Email
						</Label>
						<Input type="email" required name="email" id="email" />
					</div>

					<Button className="w-full">Sign In</Button>
				</div>

				<div className="my-6 grid grid-cols-[1fr_auto_1fr] items-center gap-3">
					<hr className="border-dashed" />
					<span className="text-muted-foreground text-xs">
						Or continue With
					</span>
					<hr className="border-dashed" />
				</div>

				<div className="grid grid-cols-2 gap-3">
					<GoogleAuthButton />
					<AppleAuthButton />
				</div>
			</div>
		</form>
	);
}
