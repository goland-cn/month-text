### Install and use

```go
go get github.com/goland-cn/month-text
```

### Nacos config msg

------

- nacos text

```go
func main() {//nacos text
	config := ReadConfig("goods_srv.json", "dev")
	fmt.Println(config)
}
```

### Redis use

------

- init Redis

```go
NewRedisClient("127.0.0.1:6379") //init redis
```

- Add Set

```go
func main() {
	NewRedisClient("127.0.0.1:6379") //init redis
	user := &User{
		Name: "张三",
		Pawd: "123456",
	}
	err := AddHSet("text", "1", user) //add set
	if err != nil {
		fmt.Println(err.Error())
	}
}
```

- Get Set

```
type User struct {
	Name string `json:"name"`
	Pawd string `json:"pawd"`
}

func main() {
	NewRedisClient("127.0.0.1:6379") //init redis
	req, err := GetHSet("text", "1")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(req)
}
```

### Send Message dxb&aly

------

- Send Message Dxb

```go
func SendMessageAly() {//aly
	payment := MsgInstance(
		"1838393649",
		"d09c29eea08845a8a2d5bb4d710b3902",
		"17519385442",
		"[小老八有限公司],验证码30分钟内有效:845268",
		&Aly{}, //*根据传入结构体判断
	)
	payment.Msg()
}

func SendMessageDxb() {//dxb
	payment := MsgInstance(
		"1838393649",
		"d09c29eea08845a8a2d5bb4d710b3902",
		"17519385442",
		"[小老八有限公司],验证码30分钟内有效:845268",
		&Dxb{}, //*根据传入结构体判断
	)
	payment.Msg()
}

func main() {
	SendMessageAly() //Aly
	SendMessageDxb() //Dxb
}
```

## Fun Use

```go
// Nacos 读取nacos配置
func Nacos() {
	config := nacoss.ReadConfig("goods_srv.json", "dev")
	fmt.Println(config)
}

type User struct {
	Name string
	Pawd string
}

// AddRedis 添加到redis
func AddRedis() {
	redisFun.NewRedisClient("127.0.0.1:6379") //init redis
	user := &User{
		Name: "李四",
		Pawd: "123456",
	}
	err := redisFun.AddHSet("text", "2", user) //add set
	if err != nil {
		fmt.Println(err.Error())
	}
}

// GetSet 获取redis集合
func GetSet() {
	redisFun.NewRedisClient("127.0.0.1:6379") //init redis
	req, err := redisFun.GetHSet("text", "2")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(req)
}

// MsgDxb 短信宝
func MsgDxb() {
	sendMsg.MsgInstance("1838393649",
		"d09c29eea08845a8a2d5bb4d710b3902",
		"17519385442",
		"[小老八有限公司],验证码30分钟内有效:845268",
		&sendMsg.Dxb{},
	)
}
```

### Redis Sentry

------

- check sentry

```go
 cat sentinel.conf
```

- check client-server

```go
redis-cli -h 127.0.0.1 -p 6380
info replication
```
