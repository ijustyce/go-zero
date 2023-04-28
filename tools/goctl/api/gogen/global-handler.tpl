package rest

import (
	{{.importPackages}}
)

func Handler(resp any, err error, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
	} else {
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}

func OnError(err error, w http.ResponseWriter, r *http.Request) {
	httpx.ErrorCtx(r.Context(), w, err)
}