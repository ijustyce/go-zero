package rest

import (
	{{.importPackages}}
)

func Handler(resp any, err error, writer http.ResponseWriter, req *http.Request) {
	if err != nil {
		httpx.ErrorCtx(req.Context(), writer, err)
	} else {
		httpx.OkJsonCtx(req.Context(), writer, resp)
	}
}

func OnError(err error, writer http.ResponseWriter, req *http.Request) {
	httpx.ErrorCtx(req.Context(), writer, err)
}