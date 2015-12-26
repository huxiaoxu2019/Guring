package apimodels

import "encoding/json"
import "log"

func GetDemoData() []byte {
	data := APIModel{
		ErrorCode: 0,
		ErrorMsg:  "successful",
		Data:      "<li class='chat-list-li'><p class='text-center chat-time'>2015-12-23 22:22:08</p><p class='text-left chat-name'><strong>GenialX</strong> <i>Said</i>: <span>Baitch....</span></p><hr class='chat-line'/></li>",
	}
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	return b
}

func GetDemoDataStr() string {

	b := GetDemoData()
	str := ""
	for i := 0; i < len(b); i++ {
		str += string(b[i])
	}
	return str

}
