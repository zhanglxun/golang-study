package main

import (
	"golangStudy/2.web/21.work04/internal/biz"
	"golangStudy/2.web/21.work04/internal/data"

	"github.com/google/wire"
)

func InitUserUsecase() *biz.UserUsecase {
	wire.Build(biz.NewUserUsecase, data.NewUserRepo)
	return &biz.UserUsecase{}
}
