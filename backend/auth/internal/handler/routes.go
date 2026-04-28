package handler

import (
	"auth/internal/middleware"
	"auth/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/callback",
				Handler: callbackHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/state",
				Handler: stateHandler(serverCtx),
			},
		},
	)

	authMiddleware := middleware.NewAuthMiddleware(serverCtx)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/refresh",
				Handler: authMiddleware.Handle(refreshTokenHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/verify",
				Handler: verifyTokenHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/logout",
				Handler: authMiddleware.Handle(logoutHandler(serverCtx)),
			},
		},
		rest.WithPrefix("/auth/token"),
	)
}
