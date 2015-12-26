package main

import (
	"apimodels"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"io"
	"log"
	"strings"
	"time"
)

const redis_room_key = "room_1720"

func main() {
	GetMsg()
	//	PushMsg()
	closeCon()
}

func PushMsg(Msg apimodels.MsgModel) bool {
	//func PushMsg() bool {
	/**
		now := time.Now()
		millis := now.UnixNano()
		millis = millis / 1000000
		score := float64(millis)

		Msg := apimodels.MsgModel{
			Name:    "genialx",
			Content: "hm......,she is a girl, so i love it",
			Time:    score,
		}
	**/
	b, err := json.Marshal(Msg)
	if err != nil {
		log.Fatalln(err)
	}

	value := ""
	for i := 0; i < len(b); i++ {
		value += string(b[i])
	}

	count := zAdd(value)
	if count < 1 {
		return false
	}

	return true

}

func GetMsg(start int, end int) []apimodels.MsgModel {
	list := zRevRange(start, end)
	msg := make([]apimodels.MsgModel, len(list))
	for i := 0; i < len(list); i++ {
		decoder := json.NewDecoder(strings.NewReader(list[i]))
		for {
			var itemMsg apimodels.MsgModel
			if err := decoder.Decode(&itemMsg); err == io.EOF {
				break
			} else if err != nil {
				log.Fatalln(err)
			}
			msg[i] = itemMsg
		}
	}
	log.Println(msg)
	return msg
}

func zAdd(value string) int {
	conn := getCon()
	now := time.Now()
	millis := now.UnixNano()
	millis = millis / 1000000
	score := float64(millis)
	rs, err := redis.Int(conn.Do("ZADD", redis_room_key, score, value))
	if err != nil {
		log.Fatalln(err)
	}
	return rs
}

func zRevRange(start int, end int) []string {
	conn := getCon()
	reply, err := conn.Do("ZREVRANGE", redis_room_key, start, end)
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

func zRevRangeByScore() []string {
	conn := getCon()
	reply, err := conn.Do("ZREVRANGEBYSCORE", redis_room_key, "+inf", "-inf")
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
		conn, err = redis.DialTimeout("tcp", "rc7569.mars.grid.sina.com.cn:7569", 0, 1*time.Second, 1*time.Second)
		if err != nil {
			panic(err)
		}
	}
	return conn
}

func closeCon() {
	if conn != nil {
		conn.Close()
	}
}
