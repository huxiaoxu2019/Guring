package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {
	conn, err := redis.DialTimeout("tcp", "rc7569.mars.grid.sina.com.cn", 0, 1*time.Second, 1*time.Second)
	if err != nil {
		panic(err)
	}
	size, err := conn.Do("DBSIZE")
	fmt.Printf("size is %d \n", size)
	conn.Close()

}
