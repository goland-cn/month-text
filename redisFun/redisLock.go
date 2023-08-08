package redisFun

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type DistributedLock struct {
	client *redis.Client
}

// NewDistributedLock 实例化锁
func NewDistributedLock(client *redis.Client) *DistributedLock {
	return &DistributedLock{
		client: client,
	}
}

// AcquireLock 创建锁
func (lock *DistributedLock) AcquireLock(key string, expiration time.Duration) (bool, error) {
	result, err := lock.client.SetNX(key, "locked", expiration).Result()
	if err != nil {
		return false, err
	}

	return result, nil
}

// ReleaseLock 释放锁
func (lock *DistributedLock) ReleaseLock(key string) error {
	_, err := lock.client.Del(key).Result()
	return err
}

func Lock() {
	NewRedisClient("127.0.0.1:6379")
	lock := NewDistributedLock(ClientRedis) //使用

	key := "my_resource" //Lock text
	expiration := 5 * time.Second

	// Get Lock
	success, err := lock.AcquireLock(key, expiration)
	if err != nil {
		fmt.Println("Failed to acquire lock:", err)
		return
	}
	//Not Lock Req
	if !success {
		fmt.Println("Failed to acquire lock: resource is already locked")
		return
	}

	// 成功获得锁后执行操作
	fmt.Println("Lock acquired. Performing operation...")

	// 模拟（需要保护的）操作
	time.Sleep(10 * time.Second)

	// 释放锁
	err = lock.ReleaseLock(key)
	if err != nil {
		fmt.Println("Failed to release lock:", err)
		return
	}
}
