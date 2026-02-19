import {
	convexAuthNextjsMiddleware,
	createRouteMatcher,
	nextjsMiddlewareRedirect as redirect,
} from "@convex-dev/auth/nextjs/server";

const isSignInPage = createRouteMatcher(["/signin"]);
const isProtectedRoute = createRouteMatcher(["/", "/product(.*)"]);

export default convexAuthNextjsMiddleware(async (request, { convexAuth }) => {
	if (isSignInPage(request) && (await convexAuth.isAuthenticated)) {
		return redirect(request, "/");
	}
	if (isProtectedRoute(request) && !(await convexAuth.isAuthenticated())) {
		return redirect(request, "/auth");
	}
});

export const config = {
	matcher: ["/((?!.*\\..*|_next).*)", "/", "/(api|trpc)(.*)"],
};
