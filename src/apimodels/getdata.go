package apimodels

import (
	"encoding/json"
	"fmt"
	"io"
	"library/redis"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetData(req *http.Request) []byte {
	req.ParseForm()
	startMillis, err := strconv.ParseFloat(req.FormValue("msgstarttimestamp"), 64)
	if err != nil {
		log.Fatalln(err)
	}

	msgList := getMsgFromRedisByTime(float64(1651138431853), float64(startMillis))

	htmlItemFormat := "<li class='chat-list-li'><p class='text-center chat-time'>%s</p><p class='text-left chat-name'><strong>%s</strong> <i>Said</i>: <span>%s</span></p><class='chat-line'/></li>"
	htmlStr := ""
	for i := len(msgList) - 1; i >= 0; i-- {
		tm := time.Unix(int64(msgList[i].Time/1000), 0)
		htmlStr += fmt.Sprintf(htmlItemFormat, tm.Format("2006-01-02 15:04:05"), msgList[i].Name, msgList[i].Content)
	}

	data := APIModel{
		ErrorCode: 0,
		ErrorMsg:  "successful",
		Data:      htmlStr,
	}
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	return b
}

func GetDataStr(req *http.Request) string {

	b := GetData(req)
	str := ""
	for i := 0; i < len(b); i++ {
		str += string(b[i])
	}
	return str

}

func getMsgFromRedisByTime(timestampMax float64, timestampMin float64) []MsgModel {
	list := redis.ZRevRangeByScore(RedisRoomKey, timestampMax, timestampMin)
	msg := make([]MsgModel, len(list))
	for i := 0; i < len(list); i++ {
		decoder := json.NewDecoder(strings.NewReader(list[i]))
		for {
			var itemMsg MsgModel
			if err := decoder.Decode(&itemMsg); err == io.EOF {
				break
			} else if err != nil {
				log.Fatalln(err)
			}
			msg[i] = itemMsg
		}
	}
	return msg
}

func getMsgFromRedisByIndex(start int, end int) []MsgModel {
	list := redis.ZRevRange(RedisRoomKey, start, end)
	msg := make([]MsgModel, len(list))
	for i := 0; i < len(list); i++ {
		decoder := json.NewDecoder(strings.NewReader(list[i]))
		for {
			var itemMsg MsgModel
			if err := decoder.Decode(&itemMsg); err == io.EOF {
				break
			} else if err != nil {
				log.Fatalln(err)
			}
			msg[i] = itemMsg
		}
	}
	return msg
}
