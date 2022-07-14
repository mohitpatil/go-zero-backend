package handler

import (
	"net/http"

	"EmployeeService/internal/logic"
	"EmployeeService/internal/svc"
	"EmployeeService/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func EmployeeServiceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewEmployeeServiceLogic(r.Context(), svcCtx)
		resp, err := l.EmployeeService(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
