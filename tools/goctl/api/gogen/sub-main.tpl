package {{.packageName}}

import (
	{{.importPackages}}
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	handler.RegisterHandlers(server, serverCtx)
}
