package redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

func ZAdd(redis_key string, value string) int {
	conn := getCon()
	now := time.Now()
	millis := now.UnixNano()
	millis = millis / 1000000
	score := float64(millis)
	rs, err := redis.Int(conn.Do("ZADD", redis_key, score, value))
	if err != nil {
		log.Fatalln(err)
	}
	return rs
}

func ZRevRange(redis_key string, start int, end int) []string {
	conn := getCon()
	reply, err := conn.Do("ZREVRANGE", redis_key, start, end)
	if err != nil {
		log.Fatalln(err)
	}
	var result []string
	if rep, ok := reply.([]interface{}); ok {
		result = make([]string, len(rep))
		for key, val := range rep {
			if v, ok := val.([]byte); ok {
				str := ""
				for i := 0; i < len(v); i++ {
					str += string(v[i])
				}
				result[key] = str
			}
		}
	}
	return result
}

func ZRevRangeByScore(redis_key string, max float64, min float64) []string {
	conn := getCon()
	reply, err := conn.Do("ZREVRANGEBYSCORE", redis_key, max, min)
	if err != nil {
		log.Fatalln(err)
	}
	var result []string
	if rep, ok := reply.([]interface{}); ok {
		result = make([]string, len(rep))
		for key, val := range rep {
			if v, ok := val.([]byte); ok {
				str := ""
				for i := 0; i < len(v); i++ {
					str += string(v[i])
				}
				result[key] = str
			}
		}
	}
	return result
}

var conn redis.Conn = nil

func getCon() redis.Conn {
	if conn == nil {
		var err interface{} = nil
		conn, err = redis.DialTimeout("tcp", "rc7569.mars.grid.sina.com.cn:7569", 10*time.Second, 10*time.Second, 10*time.Second)
		if err != nil {
			panic(err)
		}
	}
	return conn
}

// hm..., some problems here to solve
func CloseCon() {
	if conn != nil {
		conn.Close()
	}
}
