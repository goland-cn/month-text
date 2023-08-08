package main

import "fmt"

// 定义一个接口
type SendMessage interface {
	Send(phoneNumber string, message string) error
}

// 阿里云的短信策略
type AliyunMessageSender struct {
	AccessKey    string
	SecretKey    string
	Region       string
	SignName     string
	TemplateCode string
}

func NewAliyunMessageSender(accessKey string, secretKey string, region string, signName string, templateCode string) *AliyunMessageSender {
	return &AliyunMessageSender{
		AccessKey:    accessKey,
		SecretKey:    secretKey,
		Region:       region,
		SignName:     signName,
		TemplateCode: templateCode,
	}
}

func (sender *AliyunMessageSender) Send(phoneNumber string, message string) error {
	// 在这里实现调用阿里云短信接口发送短信的逻辑
	return nil
}

// 短信宝的策略
type DuanxinbaoMessageSender struct {
	Username string
	Password string
	Gateway  string
}

func NewDuanxinbaoMessageSender(username string, password string, gateway string) *DuanxinbaoMessageSender {
	return &DuanxinbaoMessageSender{
		Username: username,
		Password: password,
		Gateway:  gateway,
	}
}

func (sender *DuanxinbaoMessageSender) Send(phoneNumber string, message string) error {
	// 在这里实现调用短信宝接口发送短信的逻辑
	return nil
}

// 创建一个MessageSender结构体，该结构体包含了不同策略的发送者
type MessageSender struct {
	AliyunSendMessage     SendMessage
	DuanxinbaoSendMessage SendMessage
}

// 定使用多态的方式添加另外一种策略
func (sender *MessageSender) Send(phoneNumber string, message string, platform string) error {
	switch platform {
	case "aliyun":
		return sender.AliyunSendMessage.Send(phoneNumber, message)
	case "duanxinbao":
		return sender.DuanxinbaoSendMessage.Send(phoneNumber, message)
	default:
		return fmt.Errorf("Invalid platform")
	}
}

func main() {
	aliyunSender := NewAliyunMessageSender("yourAccessKey", "yourSecretKey", "yourRegion", "yourSignName", "yourTemplateCode")
	duanxinbaoSender := NewDuanxinbaoMessageSender("yourUsername", "yourPassword", "yourGateway")

	messageSender := &MessageSender{
		AliyunSendMessage:     aliyunSender,
		DuanxinbaoSendMessage: duanxinbaoSender,
	}

	phoneNumber := "1234567890"
	message := "Hello, this is a test message."

	err := messageSender.Send(phoneNumber, message, "alibaba") // 指定平台发送短信
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}
