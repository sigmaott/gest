package kafkafx

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type HandlerMessage func(message kafka.Message) error
type HandlerError func(err error) error
type KafkaSubscriber struct {
	//Reader *kafka.Reader
}

func (r *KafkaSubscriber) Subscribe(ctx context.Context, config kafka.ReaderConfig, handler HandlerMessage, handlerErr HandlerError) {

	reader := kafka.NewReader(config)
	go func() {
		for {
			message, err := reader.ReadMessage(ctx)
			if err != nil {
				handlerErr(err)
			}
			err = handler(message)

			if err != nil {
				handlerErr(err)
			}
		}
	}()

}
