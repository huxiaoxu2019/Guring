package viewmodels

func GetAbout() ViewModel {
	result := ViewModel{
		Title:     "About",
		SingnedIn: false,
		Active:    "about",
	}
	return result
}
