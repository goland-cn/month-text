package sendMsg

import (
	"fmt"
	"net/http"
)

// SendMsg 定义发送短信的，上下文信息和策略
type SendMsg struct {
	context  *MessageContext //消息
	strategy MessageStrategy //策略
}

// MessageContext 定义消息上下文
type MessageContext struct {
	Username string //username
	Password string //password
	Mobile   string //mobile
	Msg      string //message
}

// MessageStrategy 定义短信策略接口
type MessageStrategy interface {
	Msg(*MessageContext)
}

// MsgInstance 1.消息实例化
func MsgInstance(userName, passWord, mobile, message string, strategy MessageStrategy) *SendMsg {
	return &SendMsg{
		context: &MessageContext{
			Username: userName,
			Password: passWord,
			Mobile:   mobile,
			Msg:      message,
		},
		strategy: strategy,
	}
}

// Msg 2.继承发送消息，执行短信操作，根据策略调用不同的短信方式
func (p *SendMsg) Msg() {
	p.strategy.Msg(p.context)
}

// Aly 阿里云
type Aly struct{}

// Msg 3.实现阿里云发送短信的具体操作
func (*Aly) Msg(ctx *MessageContext) {
	fmt.Println("阿里云"+ctx.Username, ctx.Password)
}

// Dxb 短信宝
type Dxb struct{}

// Msg 3.实现短信宝发送短信的具体操作
func (*Dxb) Msg(ctx *MessageContext) {
	url := "https://api.smsbao.com/sms?u=" + ctx.Username + "&p=" + ctx.Password + "&m=" + ctx.Mobile + "&c=" + ctx.Msg
	req, err := http.Post(url, "", nil)
	if err != nil {
		fmt.Println("发送短信失败")
	}
	fmt.Println(req)
}

func SendMessageAly() {
	payment := MsgInstance(
		"1838393649",
		"d09c29eea08845a8a2d5bb4d710b3902",
		"17519385442",
		"[小老八有限公司],验证码30分钟内有效:845268",
		&Aly{}, //*根据传入结构体判断
	)
	payment.Msg()
}

func SendMessageDxb() {
	payment := MsgInstance(
		"1838393649",
		"d09c29eea08845a8a2d5bb4d710b3902",
		"17519385442",
		"[小老八有限公司],验证码30分钟内有效:845268",
		&Dxb{}, //*根据传入结构体判断
	)
	payment.Msg()
}
