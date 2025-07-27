package logger

import (
	"fmt"
	"github.com/IBM/sarama"
)

type KafkaConfig struct {
	Host string `env:"KAFKA_LOCAL_HOST" default:"localhost"`
	Port uint16 `env:"KAFKA_LOCAL_PORT" default:"9092"`
}

func NewSyncProducer(cfg *KafkaConfig) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	fmt.Printf("\n\n-----------------------------------%s:%d-------------------------\n\n", cfg.Host, cfg.Port)
	return sarama.NewSyncProducer([]string{fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)}, config)
}
