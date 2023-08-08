package consumer

//消费
import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"time"
)

func main() {
	// 创建一个新的消息消费者
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
	)
	c.Start()

	c.Subscribe("test", consumer.MessageSelector{}, func(c context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, m := range msgs {
			fmt.Printf("消息获取成功：%v \n", m.Body)
		}
		return consumer.ConsumeSuccess, nil
	})
	time.Sleep(time.Hour)
	c.Shutdown()
	//c, _ := rocketmq.NewPushConsumer(
	//	//
	//	consumer.WithGroupName("testGroup"),
	//	consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
	//)
	//// 从哪个主题里面读取信息
	//err := c.Subscribe("test", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	//	for i := range msgs {
	//		fmt.Printf("subscribe callback: %v \n", msgs[i])
	//	}
	//	return consumer.ConsumeSuccess, nil
	//})
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//// Note: start after subscribe
	//err = c.Start()
	//if err != nil {
	//	fmt.Println(err.Error())
	//	os.Exit(-1)
	//}
	//
	//time.Sleep(time.Hour)
	//c.Shutdown()

}
