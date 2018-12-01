package model

import (
	"GO_Admin/global"
	"fmt"

	"github.com/gomodule/redigo/redis"

	Redis "github.com/gomodule/redigo/redis"
)

// redisConnect 與redis連線
func redisConnect() (redis Redis.Conn, err error) {
	// 使用redis封裝的Dial進行tcp連接
	host := global.Config.Redis.Host
	port := global.Config.Redis.Port

	// 組合連接資訊
	var connectionString = fmt.Sprintf("%s:%s", host, port)
	redis, err = Redis.Dial("tcp", connectionString)

	if err != nil {
		err = global.NewError{
			Title:   "Redis connect Fail",
			Message: fmt.Sprintf("Error message is: %s", err),
		}
		return nil, err
	}

	return redis, nil
}

// Insert 存入redis值
func Insert(key string, value interface{}, time string) error {
	r, err := redisConnect()
	if err != nil {
		return err
	}
	defer r.Close()

	if _, err = r.Do("SET", key, value, "EX", time); err != nil {
		err = global.NewError{
			Title:   "Redis insert Fail",
			Message: fmt.Sprintf("Error message is: %s", err),
		}
		return err
	}

	return nil
}

// Select 取出redis值
func Select(key string) (interface{}, error) {
	r, err := redisConnect()
	if err != nil {
		return nil, err
	}
	defer r.Close()

	value, err := redis.String(r.Do("GET", key))
	if err != nil {
		err = global.NewError{
			Title:   "Redis select Fail",
			Message: fmt.Sprintf("Error message is: %s", err),
		}
		return nil, err
	}

	return value, nil
}

// Delete 刪除redis值
func Delete(key string) error {
	r, err := redisConnect()
	if err != nil {
		return err
	}
	defer r.Close()

	if _, err := r.Do("DEL", key); err != nil {
		err = global.NewError{
			Title:   "Redis delete Fail",
			Message: fmt.Sprintf("Error message is: %s", err),
		}
		return err
	}

	return nil
}

// Append 在相同key新增多個值
func Append(key string, value interface{}) (interface{}, error) {
	r, err := redisConnect()
	if err != nil {
		return nil, err
	}
	defer r.Close()

	n, err := r.Do("APPEND", key, value)
	if err != nil {
		err = global.NewError{
			Title:   "Redis select Fail",
			Message: fmt.Sprintf("Error message is: %s", err),
		}
		return nil, err
	}

	return n, nil
}
