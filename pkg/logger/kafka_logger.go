package logger

import (
	"github.com/IBM/sarama"
	"go.uber.org/zap/zapcore"
	"log"
)

type KafkaCore struct {
	zapcore.LevelEnabler
	producer sarama.SyncProducer
	topic    string
	encoder  zapcore.Encoder
}

func NewKafkaCore(producer sarama.SyncProducer, topic string, level zapcore.LevelEnabler, encoder zapcore.Encoder) *KafkaCore {
	return &KafkaCore{
		LevelEnabler: level,
		producer:     producer,
		topic:        topic,
		encoder:      encoder,
	}
}

func (c *KafkaCore) With(fields []zapcore.Field) zapcore.Core {
	clone := *c
	clone.encoder = c.encoder.Clone()
	for _, field := range fields {
		field.AddTo(clone.encoder)
	}
	return &clone
}

func (c *KafkaCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *KafkaCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	buffer, err := c.encoder.EncodeEntry(ent, fields)
	if err != nil {
		return err
	}
	defer buffer.Free()

	_, _, err = c.producer.SendMessage(&sarama.ProducerMessage{
		Topic: c.topic,
		Value: sarama.ByteEncoder(buffer.Bytes()),
	})
	if err != nil {
		log.Printf("kafka logger error: %v", err)
	}
	return err
}

func (c *KafkaCore) Sync() error {
	return nil
}
