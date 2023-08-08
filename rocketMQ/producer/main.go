package producer

//生产
import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

// Package main implements a simple producer to send message.
func main() {
	//创建一个普通的消息生产者
	// 解析NS地址
	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"10.3.90.2:9876"})))
	if err != nil {
		fmt.Printf("链接失败%v\n", err.Error())
	}
	//
	p.Start()
	for i := 0; i < 10; i++ {
		sync, _ := p.SendSync(context.TODO(), &primitive.Message{
			Topic: "test",
			Body:  []byte(fmt.Sprintf("生产的消息%d", i)),
		})
		fmt.Println(sync.Status)
	}

	//p, _ := rocketmq.NewProducer(
	//	//WithNsResolver用于设置RocketMQ实例的命名服务解析器（NameServer Resolver），
	//	//这里使用了primitive.NewPassthroughResolver来创建一个简单的解析器，
	//	//它将NameServer设置为"127.0.0.1:9876"，表示本地的RocketMQ服务器地址。
	//	producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
	//	//WithRetry用于设置消息发送失败时的重试次数，这里设置为2次。
	//	producer.WithRetry(2),
	//)
	//
	////接下来，调用p.Start()方法来启动生产者。如果启动过程中发生了错误，将会打印错误信息并退出程序。
	//err := p.Start()
	//
	////在发送完所有消息后，调用p.Shutdown()方法关闭生产者。如果关闭过程中出现错误，将会打印错误信息。
	//deferQueue func(p rocketmq.Producer) {
	//	err = p.Shutdown()
	//	if err != nil {
	//
	//	}
	//}(p)
	//if err != nil {
	//	fmt.Printf("起订生产失败: %s", err.Error())
	//	os.Exit(1)
	//}
	////定义一个变量topic，它表示要发送消息的主题。
	//topic := "test"
	////，创建一个primitive.Message实例msg，设置其主题（Topic）和消息内容（Body）。
	//
	//// 调用p.SendSync方法同步发送消息，并返回结果。
	//res, err := p.SendSync(context.Background(), &primitive.Message{
	//	Topic: topic,
	//	Body:  []byte("这是一条测试消息"),
	//})
	//fmt.Println(res)

}
