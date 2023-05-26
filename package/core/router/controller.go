package router

import (
	"reflect"
)

type IRouter interface {
	InitRouter()
}

type Router struct {
	controller any
}

func (b *Router) InitRouter() {
	//t := reflect.TypeOf((*I)(nil)).Elem()
	//for i := 0; i < t.NumMethod(); i++ {
	//	reflect.ValueOf(b.controller).MethodByName(t.Method(i).Name).Call([]reflect.Value{})
	//}

	value := reflect.ValueOf(b.controller)

	// Iterate over the struct's methods
	for i := 0; i < value.NumMethod(); i++ {
		// Get the method value
		methodValue := value.Method(i)

		// Call the method dynamically
		methodValue.Call(nil)
	}
}
func NewBaseRouter[I any](controller I) *Router {
	return &Router{
		controller: controller,
	}
}

func InitRouter(controllers []any) {

	controllerWapper := make([]Router, len(controllers))
	for i, router := range controllerWapper {
		router.controller = controllers[i]
		router.InitRouter()

	}
	//for _, controller := range controllers {
	//	controller.InitRouter()
	//}
}
