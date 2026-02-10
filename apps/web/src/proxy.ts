import { type NextRequest, NextResponse } from "next/server";

export function proxy(request: NextRequest) {
	if (request.nextUrl.pathname === "/") {
		return NextResponse.redirect(`${request.nextUrl.origin}/sign-in`);
	}
	return NextResponse.next();
}

export const config = {
	matcher: ["/:path*"],
};
