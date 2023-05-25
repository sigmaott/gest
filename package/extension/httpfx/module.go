package httpfx

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("httpfx", fx.Provide(resty.New))
}
func ParserData[T any](res *resty.Response) (*T, error) {
	data := new(T)
	err := json.NewDecoder(res.RawBody()).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, err

}
