package affair

//事务队列
import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type DemoListener struct{}

func (*DemoListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	fmt.Println("开始执行本地事务")
	time.Sleep(time.Second * 3)

	fmt.Println("本地事务执行失败")
	return primitive.UnknowState
}

func (*DemoListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Println("rocketmq的消息回查")
	time.Sleep(time.Second * 15)
	return primitive.CommitMessageState
}

func main() {
	// 开启半消息生产者
	p, _ := rocketmq.NewTransactionProducer(
		&DemoListener{},
		// 设置rocketmq的地址
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		// 设置重试机制
		producer.WithRetry(2),
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("启动生产者错误: %s\n", err.Error())
		os.Exit(1)
	}
	//在事务中发送消息
	for i := 0; i < 10; i++ {
		// 发送事务消息
		res, _ := p.SendMessageInTransaction(context.Background(),
			primitive.NewMessage("TopicTest5", []byte("开始 RocketMQ 事务回查机制 "+strconv.Itoa(i))))
		fmt.Println(res.String())

	}

	time.Sleep(5 * time.Minute)
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
