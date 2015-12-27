package redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

func ZAdd(redis_key string, score float64, value string) (int, error) {
	conn := getCon()
	_, err := conn.Do("ZADD", redis_key, score, value)
	/**
	尽管抛异常，但是ZADD是成功的。
	redigo: illegal bytes in length (possible server error or unsupported concurrent read by application)
	**/
	if err != nil {
		log.Println("ZADD ERR:", err)
		setCon()
		return 0, err
	}
	return 1, nil
}

func ZRevRange(redis_key string, start int, end int) ([]string, error) {
	conn := getCon()
	reply, err := conn.Do("ZREVRANGE", redis_key, start, end)
	if err != nil {
		log.Println("ZREVRANGE ERR:", err)
		setCon()
		return nil, err
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
	return result, nil
}

func ZRevRangeByScore(redis_key string, max float64, min float64) ([]string, error) {
	conn := getCon()
	reply, err := conn.Do("ZREVRANGEBYSCORE", redis_key, max, min)
	if err != nil {
		log.Println("ZREVRANGEBYSCORE ERR:", err)
		setCon()
		return nil, err
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
	return result, nil
}

var conn redis.Conn = nil

func setCon() redis.Conn {
	var err interface{} = nil
	conn, err = redis.Dial("tcp", "rc7569.mars.grid.sina.com.cn:7569")
	if err != nil {
		log.Println("setCon ERR:", err)
	}

	return conn
}

func getCon() redis.Conn {
	if conn == nil {
		var err interface{} = nil
		conn, err = redis.Dial("tcp", "rc7569.mars.grid.sina.com.cn:7569")
		if err != nil {
			log.Println("getCon ERR:", err)
		}
		// defer conn.Close()
	}
	return conn
}

// hm..., some problems here to solve
func CloseCon() {
	if conn != nil {
		conn.Close()
	}
}
