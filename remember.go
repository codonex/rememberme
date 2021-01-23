package rememberme

func RememberMe() func(string) (string, error) {
	var remember string = ""
	return func(name string) (string, error) {
		remember = remember + " " + name
		return remember, nil
	}
}
