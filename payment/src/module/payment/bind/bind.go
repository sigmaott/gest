package bind

import (
	"encoding/json"
	"github.com/gestgo/gest/package/common/mapstructure"
	"github.com/labstack/echo/v4"
	"log"
)

type (
	// Binder is the interface that wraps the Bind method.
	Binder interface {
		Bind(i interface{}, c echo.Context) error
	}

	// DefaultBinder is the default implementation of the Binder interface.
	DefaultBinder struct{}

	// BindUnmarshaler is the interface used to wrap the UnmarshalParam method.
	// Types that don't implement this, but do implement encoding.TextUnmarshaler
	// will use that interface instead.
	BindUnmarshaler interface {
		// UnmarshalParam decodes and assigns a value from an form or query param.
		UnmarshalParam(param string) error
	}
)

// BindPathParams binds path params to bindable object
//func (b *DefaultBinder) BindPathParams(c echo.Context, i interface{}) error {
//	names := c.ParamNames()
//	values := c.ParamValues()
//	params := map[string][]string{}
//	for i, name := range names {
//		params[name] = []string{values[i]}
//	}
//	if err := b.bindData(i, params, "param"); err != nil {
//		return NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
//	}
//	return nil
//}

// BindQueryParams binds query params to bindable object
func (b *DefaultBinder) BindBody(c echo.Context, i interface{}) error {
	req := c.Request()
	var body map[string]any
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		return err
	}
	err = b.bindData(i, body, "json")
	log.Print(body)
	return err
}

func (b *DefaultBinder) bindData(object any, i map[string]any, tagName string) error {
	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   object,
		TagName:  tagName,
	}
	decoder, err := mapstructure.NewDecoder(cfg)
	if err != nil {
		return err
	}

	return decoder.Decode(i)
}
func NewBind() *DefaultBinder {
	return &DefaultBinder{}
}
