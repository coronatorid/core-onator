package adapter

import (
	"time"

	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

// KafkaPublisher adapt segmentio kafka publisher
type KafkaPublisher struct {
	writer *kafka.Writer
}

// AdaptKafkaPublisher create new struct of kafkapublisher
func AdaptKafkaPublisher(writer *kafka.Writer) *KafkaPublisher {
	return &KafkaPublisher{
		writer: writer,
	}
}

// WriteMessages write message to kafka
func (k *KafkaPublisher) WriteMessages(ctx provider.Context, msgs ...kafka.Message) error {
	startTime := time.Now()
	if err := k.writer.WriteMessages(ctx.Ctx(), msgs...); err != nil {
		log.Error().
			Err(err).
			Stack().
			Str("request_id", util.GetRequestID(ctx)).
			Dur("duration", time.Since(startTime)).
			Msg("success publishing message to kafka")
		return err
	}

	log.Info().
		Str("request_id", util.GetRequestID(ctx)).
		Dur("duration", time.Since(startTime)).
		Msg("success publishing message to kafka")
	return nil
}
