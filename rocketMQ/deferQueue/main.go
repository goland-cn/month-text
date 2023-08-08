package deferQueue

//延迟
import (
	"context"
	"fmt"
	"os"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	//
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		//消息发送失败后重试的次数
		producer.WithRetry(2),
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}

	defer func(p rocketmq.Producer) {
		err = p.Shutdown()
		if err != nil {
		}
	}(p)
	for i := 0; i < 10; i++ {
		msg := primitive.NewMessage("test", []byte("Hello RocketMQ Go Client!"))
		//设置消息的延迟时间
		msg.WithDelayTimeLevel(3)
		//1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
		res, _ := p.SendSync(context.Background(), msg)
		fmt.Printf("send result: %v \n", res)
	}
}
