package controller

import (
	"context"
	"github.com/gestgo/gest/package/extension/kafkafx"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"log"
	config2 "payment/config"
	"payment/src/module/event-metric/service"
	"time"
)

type IUserController interface {
	FindAll()
}
type Params struct {
	fx.In
	//Router  *echo.Group
	Logger          *zap.SugaredLogger
	Service         service.ISSAIEventService
	KafkaSubscriber *kafkafx.KafkaSubscriber `name:"platformKafka"`
}
type Controller struct {
	//router  *echo.Group
	kafkaSubscriber *kafkafx.KafkaSubscriber
	logger          *zap.SugaredLogger
	service         service.ISSAIEventService
}

type Result struct {
	fx.Out
	Controller any `group:"kafkaTopics"`
}

func NewController(params Params) Result {
	return Result{
		Controller: &Controller{
			//router:  params.Router,
			logger:          params.Logger,
			service:         params.Service,
			kafkaSubscriber: params.KafkaSubscriber,
		},
	}

}

func (b *Controller) FindAll() {
	dialer := &kafka.Dialer{
		ClientID:        uuid.New().String(),
		Timeout:         10 * time.Second,
		Deadline:        time.Time{},
		LocalAddr:       nil,
		DualStack:       false,
		FallbackDelay:   0,
		KeepAlive:       0,
		Resolver:        nil,
		TLS:             nil,
		SASLMechanism:   nil,
		TransactionalID: "",
	}

	config := kafka.ReaderConfig{
		Brokers:                config2.GetConfiguration().Kafka.Urls,
		GroupID:                config2.GetConfiguration().Kafka.GroupId,
		GroupTopics:            nil,
		Topic:                  "monitor_event",
		Partition:              0,
		Dialer:                 dialer,
		QueueCapacity:          0,
		MinBytes:               0,
		MaxBytes:               0,
		MaxWait:                0,
		ReadBatchTimeout:       0,
		ReadLagInterval:        0,
		GroupBalancers:         nil,
		HeartbeatInterval:      0,
		CommitInterval:         0,
		PartitionWatchInterval: 0,
		WatchPartitionChanges:  false,
		SessionTimeout:         0,
		RebalanceTimeout:       0,
		JoinGroupBackoff:       0,
		RetentionTime:          0,
		StartOffset:            0,
		ReadBackoffMin:         0,
		ReadBackoffMax:         0,
		Logger:                 nil,
		ErrorLogger:            nil,
		IsolationLevel:         0,
		MaxAttempts:            0,
		OffsetOutOfRangeError:  false,
	}

	b.kafkaSubscriber.Subscribe(context.TODO(), config, func(message kafka.Message) error {
		log.Printf("%+v", message)
		return nil
	}, func(err error) error {
		log.Print(err)
		return err
	})
	log.Print("login controller")
	//b.router.POST("/users", func(c echo.Context) error {
	//	//c.Request().Header.Get("")
	//	//appId := common.GetAppId(c)
	//	//lang := common.GetAcceptLanguage(c)
	//	//log.Print(appId, lang)
	//	//query := new(dto.GetListUserQuery)
	//	//err := c.Bind(query)
	//	//if err != nil {
	//	//	log.Print(err)
	//	//	return err
	//	//}
	//	//err = c.Validate(query)
	//	//if err != nil {
	//	//	log.Print(err)
	//	//	return err
	//	//}
	//	//log.Print(c.Get("body"))
	//	//message, err := b.i18nService.T("en", locales.CARDINAL_TEST)
	//	//result, sort, err := queryBuilder.MongoParserQuery[model.Payment](c.Request().URL.Query())
	//	//log.Print(result, sort, err)
	//	//b.logger.Info()
	//	//b.repository.FindAll()
	//	//b.repository.FindAll()
	//	//return errors.New("error")
	//	b.service.Create(dto.CreateEvent{
	//		Total: 4,
	//		AppId: "default-app",
	//	})
	//	return c.JSON(http.StatusOK, "ok")
	//})

}
