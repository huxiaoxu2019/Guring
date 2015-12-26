package viewmodels

func GetRoom() ViewModel {
	result := ViewModel{
		Title:     "Room",
		SingnedIn: false,
		Active:    "room",
	}
	return result
}
