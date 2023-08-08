package nacoss

import (
	"fmt"
)

func main() {
	config := ReadConfig("goods_srv.json", "dev")
	fmt.Println(config)
}
