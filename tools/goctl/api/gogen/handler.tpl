package {{.PkgName}}

import (
	"net/http"

	{{.ImportPackages}}

	{{if .HasRequest}}"github.com/zeromicro/go-zero/rest/httpx"{{end}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		{{if .HasRequest}}var params types.{{.RequestType}}
		if err := httpx.Parse(req, &params); err != nil {
			rest.OnError(err, writer, req)
			return
		}

		{{end}}logicObj := {{.LogicName}}.New{{.LogicType}}(req, svcCtx)
		{{if .HasResp}}resp, {{end}}err := logicObj.{{.Call}}({{if .HasRequest}}&params{{end}})
		rest.Handler({{if .HasResp}}resp, {{else}}nil, {{end}}err, writer, req)
	}
}
