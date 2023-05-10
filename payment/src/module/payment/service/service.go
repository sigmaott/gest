package service

import (
	"fmt"
	"github.com/gestgo/gest/package/extension/i18nfx"
	"go.uber.org/fx"
)

type IUserService interface {
	Log()
}
type userService struct {
}

func (u *userService) Log() {
	//TODO implement me
	panic("implement me")
}

type UserServiceParams struct {
	fx.In
	I18nService i18nfx.I18nService
}

func NewUserService(params UserServiceParams) IUserService {
	a, _ := params.I18nService.T("en", "ok")
	fmt.Println(a)

	return &userService{}

}
