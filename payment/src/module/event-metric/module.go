package event_metric

import (
	//"github.com/gestgo/gest/package/core/repository"
	"go.uber.org/fx"
	"payment/src/module/event-metric/controller"
	"payment/src/module/event-metric/repository"
	"payment/src/module/event-metric/service"
	"reflect"
)

func Module() fx.Option {
	return fx.Module("ssai-event-metric",
		fx.Provide(
			controller.NewController,
			service.NewEventService,
			repository.NewEventRepository,
		),
	)
}
func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	// Get the reflect.Value of the struct object
	value := reflect.ValueOf(obj)

	// Ensure that the provided object is a struct
	if value.Kind() != reflect.Struct {
		panic("Object is not a struct")
	}

	// Iterate over struct fields
	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		fieldValue := value.Field(i).Interface()
		result[field.Name] = fieldValue
	}

	return result
}

//fx.ResultTags(`group:"controllers"`)
