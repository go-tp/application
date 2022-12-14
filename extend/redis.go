package extend

import (
	"gtp/configs"
	"strconv"
	"time"
	"github.com/garyburd/redigo/redis"
)

var Rdb redis.Conn
var Pool *redis.Pool

func RedisInit() {	
	StartRedis()
}


func RedisClose() {
	Rdb.Close()
}

// redis pool
func PoolInitRedis(server string, rdb int, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     300, //空闲数
		IdleTimeout: 600 * time.Second,
		MaxActive:   4000, //最大数
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server, redis.DialDatabase(rdb), redis.DialPassword(password))
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func StartRedis() {
	Rdb_url := configs.ReadYaml().Redis.Host + ":" + configs.ReadYaml().Redis.Port
	Rdb_db, _ := strconv.Atoi(configs.ReadYaml().Redis.Db)
	Rdb_pwd := configs.ReadYaml().Redis.Password
	Pool = PoolInitRedis(Rdb_url, Rdb_db, Rdb_pwd)
	Rdb = Pool.Get()
}

// redis 使用
// r_key := "indexapi:getindex"
// rdb := extend.Pool
// conn := rdb.Get()
// defer conn.Close()
// // redis 设置
// conn.Do("Set", r_key,"2")
// // redis 获取
// v,_ := conn.Do("GET",r_key)

// v1,_ := redis.String(v,nil)
// fmt.Printf("type=> %T\n", v)
// // unit8
// fmt.Println(v)

// // string
// fmt.Println(v1)