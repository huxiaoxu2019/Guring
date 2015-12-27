package apimodels

type APIModel struct {
	ErrorCode int
	ErrorMsg  string
	Data      string
	LastTime  float64 // 最新内容(DATA)的时间戳(score)
}

type MsgModel struct {
	Content string
	Name    string
	Time    float64
}

const RedisRoomKey = "room_1720"
