package middleware

import "github.com/danielgtaylor/huma/v2"

func AuthMiddleware(ctx huma.Context, next func(ctx huma.Context)) {
	sessionCookie := huma.ReadCookie(ctx, "session")

	// TODO: Check the session cookie to see if the user is logged in

	next(ctx)
}
