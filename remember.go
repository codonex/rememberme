package rememberme

func RememberMe() func(string) string {
	var remember string = ""
	return func(name string) string {
		remember = remember + " " + name
		return remember
	}
}
