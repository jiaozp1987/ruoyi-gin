package redisUtils

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	iniUtils "ruoyi-gin/src/ruoyi-common/core/ini"
)

//var RedisClient *redis.Client

// 定义一个全局的pool
var RedisClient *redis.Pool

// 当启动程序时，就初始化连接池
func init() {

	RedisClient = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   // 表示和数据库的最大链接数， 0 表示没有限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码， 链接哪个ip的redis
			return redis.Dial("tcp", iniUtils.REDIS_ADDR,
				redis.DialDatabase(iniUtils.REDIS_DB),
				redis.DialPassword(iniUtils.REDIS_PASSWORD),
			)
		},
	}

}

func Set(key string, value string, expiration int) error {
	//ctx := context.Background()
	//return RedisClient.Set(ctx, key, value, expiration)

	conn := RedisClient.Get()
	defer conn.Close()

	_, err := conn.Do("set", key, value, "EX", expiration) //redis set命令
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func Get(key string) (string, error) {
	conn := RedisClient.Get()
	defer conn.Close()
	str, err := redis.String(conn.Do("Get", key))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return "", err
	}
	return str, nil
}
