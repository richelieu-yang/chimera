package pulsarKit

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"time"
)

type (
	Producer struct {
		pulsar.Client
		pulsar.Producer
	}
)

func (p *Producer) SendWithTimeout(pMsg *pulsar.ProducerMessage, timeout time.Duration) (pulsar.MessageID, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	return p.Send(ctx, pMsg)
}

func (p *Producer) Close() {
	if p.Producer != nil {
		p.Producer.Close()
	}
	if p.Client != nil {
		p.Client.Close()
	}
}

// NewProducerOriginally
/*
PS: 目标Pulsar服务未启动的情况下，如果ctx不加以限制，要过约 1min 才会返回error（期间客户端日志有connection refused输出）.

@param options 	必须的属性: Topic
				建议的属性: SendTimeout
@param logPath 	客户端的日志输出（为空则输出到控制台）
*/
func NewProducerOriginally(ctx context.Context, addresses []string, options pulsar.ProducerOptions, logPath string) (*Producer, error) {
	var client pulsar.Client
	var producer pulsar.Producer

	// 写入nil: 新建Producer成功
	errCh := make(chan error, 1)

	go func() {
		var err error
		client, err = NewClient(addresses, logPath)
		if err != nil {
			err = errorKit.WithLocationInfo(err)
			errCh <- err
			return
		}
		producer, err = client.CreateProducer(options)
		if err != nil {
			err = errorKit.WithLocationInfo(err)
			errCh <- err
			return
		}

		select {
		case <-ctx.Done():
			// 新建Producer成功之前，ctx已经 超时 或 被取消 了，此时需要释放资源
			if producer == nil {
				producer.Close()
			}
			if client == nil {
				client.Close()
			}
			errCh <- errorKit.WithLocationInfo(ctx.Err())
		default:
			errCh <- nil
		}
	}()

	select {
	case <-ctx.Done():
		return nil, errorKit.WithLocationInfo(ctx.Err())
	case err := <-errCh:
		if err != nil {
			return nil, err
		}
		return &Producer{
			Client:   client,
			Producer: producer,
		}, nil
	}
}
