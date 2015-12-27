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
	_, err := pushMsgToRedis(msg)

	var errorCode int
	var errorMsg string
	if err != nil {
		errorCode = 1
		errorMsg = "fail"
	} else {
		errorCode = 0
		errorMsg = "successful"
	}

	result := APIModel{
		ErrorCode: errorCode,
		ErrorMsg:  errorMsg,
		Data:      "",
		LastTime:  0,
	}
	b, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}
	return b
}

func pushMsgToRedis(Msg MsgModel) (bool, error) {
	b, err := json.Marshal(Msg)
	if err != nil {
		log.Fatalln(err)
	}
	value := ""
	for i := 0; i < len(b); i++ {
		value += string(b[i])
	}
	_, _ = redis.ZAdd(RedisRoomKey, Msg.Time, value)

	return true, nil

}
