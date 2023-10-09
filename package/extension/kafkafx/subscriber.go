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
type ReaderConfig struct {
	kafka.ReaderConfig
	Ack bool
}

func (r *KafkaSubscriber) Subscribe(ctx context.Context, config ReaderConfig, handler HandlerMessage, handlerErr HandlerError) {

	reader := kafka.NewReader(config.ReaderConfig)
	go func() {
		for {
			message, err := reader.ReadMessage(ctx)
			if err != nil {
				handlerErr(err)
				continue
			}
			err = handler(message)
			if err != nil && config.Ack {
				err = reader.CommitMessages(context.Background(), message)
				if err != nil {
					handlerErr(err)
					continue
				}
			}
			if err != nil {
				handlerErr(err)
				continue
			}
		}
	}()

}
