import { z } from "zod";

export const signInSchema = z.object({
	email: z.email("The provided email is invalid"),
	password: z.string("Password is required"),
});
export type SignIn = z.infer<typeof signInSchema>;
