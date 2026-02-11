import { z } from "zod";

export const signInSchema = z.object({
	email: z.email("Provided email is invalid"),
	password: z.string("Password is required"),
});
export type SignIn = z.infer<typeof signInSchema>;

export const signUpSchema = z
	.object({
		name: z.string().min(2, "Name must have at least 2 characters").optional(),
		email: z.email("Provided email is invalid"),
		password: z.string().min(8),
		confirmPassword: z.string(),
	})
	.refine((data) => data.password === data.confirmPassword, {
		message: "Passwords do not match",
		path: ["confirmPassword"],
	});
export type SignUp = z.infer<typeof signUpSchema>;
