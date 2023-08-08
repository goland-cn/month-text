package redisFun

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
)

var (
	ClientRedis *redis.Client
)

// NewRedisClient redis连接初始化 127.0.0.1:6379
func NewRedisClient(addr string) {
	ClientRedis = redis.NewClient(&redis.Options{
		Addr:     addr, // 更改为实际的 Redis 地址
		Password: "",   // 如有密码，添加密码
		DB:       0,    // Redis 默认数据库
	})
}

// AddHSet 添加set集合 hKey键值  field唯一字段  value内容
func AddHSet(hKey string, field string, value interface{}) error {
	val, _ := json.Marshal(value)
	_, err := ClientRedis.HSet(hKey, field, val).Result()
	if err != nil {
		return errors.New("添加集合失败")
	}
	return nil
}

// GetHSet 获取set集合中指定键的值
func GetHSet(hKey string, field string) (*User, error) {
	data := &User{} //<---------------------------------这里需要更换为你的结构体,返回值会跟着你的结构体进行变动
	val, err := ClientRedis.HGet(hKey, field).Result()
	if err != nil || len(val) == 0 {
		return nil, errors.New("获取集合失败")
	}
	_ = json.Unmarshal([]byte(val), data)
	return data, nil
}
