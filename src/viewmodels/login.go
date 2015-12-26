package viewmodels

func GetLogin() ViewModel {
	result := ViewModel{
		Title:     "Login",
		SingnedIn: false,
		Active:    "login",
	}
	return result
}
