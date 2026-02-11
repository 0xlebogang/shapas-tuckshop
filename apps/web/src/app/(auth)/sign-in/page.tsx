"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import Link from "next/link";
import { RedirectType, redirect } from "next/navigation";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import GoogleAuthButton from "@/components/google-auth-button";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { authClient } from "@/lib/auth-client";
import {
	type SignIn as SignInSchema,
	signInSchema,
} from "@/lib/validation/auth";

export default function SignIn() {
	const {
		register,
		handleSubmit,
		formState: { errors },
		reset,
	} = useForm({
		resolver: zodResolver(signInSchema),
	});

	async function onSubmit(payload: SignInSchema) {
		const { data, error } = await authClient.signIn.email({
			email: payload.email,
			password: payload.password,
		});

		if (error) {
			toast.error(`Sign in failed: ${error.message}`);
			reset({ password: "" });
			return;
		}

		toast.success(`Authenticated as ${data.user.email}`);
		reset();
		redirect("/dashboard", RedirectType.replace);
	}

	return (
		<form
			onSubmit={handleSubmit(onSubmit)}
			className="bg-muted m-auto h-fit w-full max-w-sm overflow-hidden rounded-[calc(var(--radius)+.125rem)] border shadow-md shadow-zinc-950/5 dark:[--color-muted:var(--color-zinc-900)]"
		>
			<div className="bg-card -m-px rounded-[calc(var(--radius)+.125rem)] border p-8 pb-6">
				<div className="text-center">
					<Link href="/" aria-label="go home" className="mx-auto block w-fit">
						<span className="font-bold text-3xl">Shapas</span>
					</Link>
					<h1 className="mb-1 mt-4 text-xl font-semibold">
						Sign In to Shapas Tuckshop
					</h1>
					<p className="text-sm">Welcome back! Sign in to continue</p>
				</div>

				{errors.root?.message && (
					<div className="p-4 bg-destructive my-6 rounded text-forground">
						{errors.root?.message}
					</div>
				)}

				<div className="mt-6 space-y-6">
					<div className="space-y-2">
						<Label htmlFor="email" className="block text-sm">
							Email
						</Label>
						<div className="flex flex-col gap-1">
							<Input {...register("email", { required: true })} />
							{errors.email?.message && (
								<p className="text-xs text-destructive">
									{errors.email.message}
								</p>
							)}
						</div>
					</div>

					<div className="space-y-0.5">
						<div className="flex items-center justify-between">
							<Label htmlFor="password" className="text-sm">
								Password
							</Label>
							<Button asChild variant="link" size="sm">
								<Link
									href="#"
									className="link intent-info variant-ghost text-sm"
								>
									Forgot your Password ?
								</Link>
							</Button>
						</div>
						<div className="flex flex-col gap-1">
							<Input
								{...register("password", { required: true })}
								type="password"
								className="input sz-md variant-mixed"
							/>
							{errors.email?.message && (
								<p className="text-xs text-destructive">
									{errors.password?.message}
								</p>
							)}
						</div>
					</div>

					<Button type="submit" className="w-full">
						Sign In
					</Button>
				</div>

				<div className="my-6 grid grid-cols-[1fr_auto_1fr] items-center gap-3">
					<hr className="border-dashed" />
					<span className="text-muted-foreground text-xs">
						or continue with
					</span>
					<hr className="border-dashed" />
				</div>

				<div className="grid grid-cols-2 gap-3">
					<GoogleAuthButton />
				</div>
			</div>

			<div className="p-3">
				<p className="text-accent-foreground text-center text-sm">
					Don't have an account ?
					<Button asChild variant="link" className="px-2">
						<Link href="/sign-up">Sign Up</Link>
					</Button>
				</p>
			</div>
		</form>
	);
}
