package apimodels

type APIModel struct {
	ErrorCode int
	ErrorMsg  string
	Data      string
}

type MsgModel struct {
	Content string
	Name    string
	Time    float64
}

const RedisRoomKey = "room_1720"
