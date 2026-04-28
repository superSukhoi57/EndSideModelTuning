package handler

import (
	"net/http"

	machine "iterative_control/internal/handler/machine"
	parameter "iterative_control/internal/handler/parameter"
	result "iterative_control/internal/handler/result"
	task "iterative_control/internal/handler/task"
	test "iterative_control/internal/handler/test"
	"iterative_control/internal/middleware"
	"iterative_control/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	authMiddleware := middleware.NewAuthMiddleware(serverCtx)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/question",
				Handler: authMiddleware.Handle(test.QuestionHandler(serverCtx)),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/machine",
				Handler: authMiddleware.Handle(machine.CreateMachineHandler(serverCtx)),
			},
			{
				Method:  http.MethodPut,
				Path:    "/machine",
				Handler: authMiddleware.Handle(machine.UpdateMachineHandler(serverCtx)),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/machine/:id",
				Handler: authMiddleware.Handle(machine.DeleteMachineHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/machine/:id",
				Handler: authMiddleware.Handle(machine.GetMachineHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/machines",
				Handler: authMiddleware.Handle(machine.ListMachineHandler(serverCtx)),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/parameter",
				Handler: authMiddleware.Handle(parameter.CreateParameterHandler(serverCtx)),
			},
			{
				Method:  http.MethodPut,
				Path:    "/parameter",
				Handler: authMiddleware.Handle(parameter.UpdateParameterHandler(serverCtx)),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/parameter/:id",
				Handler: authMiddleware.Handle(parameter.DeleteParameterHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/parameter/:id",
				Handler: authMiddleware.Handle(parameter.GetParameterHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/parameters",
				Handler: authMiddleware.Handle(parameter.ListParameterHandler(serverCtx)),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/task",
				Handler: authMiddleware.Handle(task.CreateTaskHandler(serverCtx)),
			},
			{
				Method:  http.MethodPut,
				Path:    "/task",
				Handler: authMiddleware.Handle(task.UpdateTaskHandler(serverCtx)),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/task/:id",
				Handler: authMiddleware.Handle(task.DeleteTaskHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/task/:id",
				Handler: authMiddleware.Handle(task.GetTaskHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/tasks",
				Handler: authMiddleware.Handle(task.ListTaskHandler(serverCtx)),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/result",
				Handler: authMiddleware.Handle(result.CreateResultHandler(serverCtx)),
			},
			{
				Method:  http.MethodPut,
				Path:    "/result",
				Handler: authMiddleware.Handle(result.UpdateResultHandler(serverCtx)),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/result/:id",
				Handler: authMiddleware.Handle(result.DeleteResultHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/result/:id",
				Handler: authMiddleware.Handle(result.GetResultHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/results",
				Handler: authMiddleware.Handle(result.ListResultHandler(serverCtx)),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
