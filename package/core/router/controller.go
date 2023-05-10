package router

import (
	"reflect"
)

type IRouter interface {
	InitRouter()
}

type Router[I any] struct {
	controller I
}

func (b *Router[I]) InitRouter() {
	t := reflect.TypeOf((*I)(nil)).Elem()
	for i := 0; i < t.NumMethod(); i++ {
		reflect.ValueOf(b.controller).MethodByName(t.Method(i).Name).Call([]reflect.Value{})
	}
}
func NewBaseRouter[I any](controller I) *Router[I] {
	return &Router[I]{
		controller: controller,
	}
}

func InitRouter(controllers []IRouter) {
	for _, controller := range controllers {

		controller.InitRouter()
	}
}
