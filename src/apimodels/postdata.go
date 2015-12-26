package apimodels

import "net/http"
import "encoding/json"
import "log"
import "library/redis"
import "time"

func PostData(req *http.Request) []byte {
	req.ParseForm()
	content := req.PostFormValue("content")
	name := req.PostFormValue("name")
	now := time.Now()
	nanos := now.UnixNano()
	millis := float64(nanos / 1000000)

	msg := MsgModel{
		Name:    name,
		Content: content,
		Time:    millis,
	}
	pushRs := pushMsgToRedis(msg)

	if pushRs == false {
		log.Println("Push msg failed")
	}

	result := APIModel{
		ErrorCode: 0,
		ErrorMsg:  "successful",
		Data:      "",
	}
	b, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}
	return b
}

func pushMsgToRedis(Msg MsgModel) bool {
	b, err := json.Marshal(Msg)
	if err != nil {
		log.Fatalln(err)
	}

	value := ""
	for i := 0; i < len(b); i++ {
		value += string(b[i])
	}

	count := redis.ZAdd(RedisRoomKey, value)
	if count < 1 {
		return false
	}

	return true

}
