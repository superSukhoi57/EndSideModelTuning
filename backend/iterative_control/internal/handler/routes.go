package handler

import (
	"net/http"

	machine "iterative_control/internal/handler/machine"
	parameter "iterative_control/internal/handler/parameter"
	result "iterative_control/internal/handler/result"
	test "iterative_control/internal/handler/test"
	task "iterative_control/internal/handler/task"
	"iterative_control/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/question",
				Handler: test.QuestionHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/machine",
				Handler: machine.CreateMachineHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/machine",
				Handler: machine.UpdateMachineHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/machine/:id",
				Handler: machine.DeleteMachineHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/machine/:id",
				Handler: machine.GetMachineHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/machines",
				Handler: machine.ListMachineHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/parameter",
				Handler: parameter.CreateParameterHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/parameter",
				Handler: parameter.UpdateParameterHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/parameter/:id",
				Handler: parameter.DeleteParameterHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/parameter/:id",
				Handler: parameter.GetParameterHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/parameters",
				Handler: parameter.ListParameterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/task",
				Handler: task.CreateTaskHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/task",
				Handler: task.UpdateTaskHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/task/:id",
				Handler: task.DeleteTaskHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/task/:id",
				Handler: task.GetTaskHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/tasks",
				Handler: task.ListTaskHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/result",
				Handler: result.CreateResultHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/result",
				Handler: result.UpdateResultHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/result/:id",
				Handler: result.DeleteResultHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/result/:id",
				Handler: result.GetResultHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/results",
				Handler: result.ListResultHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
