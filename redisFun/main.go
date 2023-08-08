package redisFun

import "fmt"

type User struct {
	Name string `json:"name"`
	Pawd string `json:"pawd"`
}

func main() {
	NewRedisClient("127.0.0.1:6379") //初始化redis
	req, err := GetHSet("text", "1")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(req)
}
