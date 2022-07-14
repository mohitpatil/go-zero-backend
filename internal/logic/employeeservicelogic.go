package logic

import (
	"context"

	"EmployeeService/internal/svc"
	"EmployeeService/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmployeeServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmployeeServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmployeeServiceLogic {
	return &EmployeeServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmployeeServiceLogic) EmployeeService(req *types.Request) (resp *types.Response, err error) {

	employee := new(types.Request)
	employee.Name = req.Name
	employee.DOB = req.DOB
	employee.Email = req.Email
	employee.Mobile = req.Mobile

	resp = new(types.Response)
	resp.Message = "Saved data successfully."

	return
}
