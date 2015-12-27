package viewmodels

func GetRoom(nickname string) ViewModel {
	result := ViewModel{
		Title:     "Room",
		SingnedIn: false,
		Active:    "room",
		Nickname:  nickname,
	}
	return result
}
