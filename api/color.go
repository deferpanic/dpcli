//mv me

package api

func RedBold(str string) string {
	return "\033[1;31m" + str + "\033[0m"
}

func GreenBold(str string) string {
	return "\033[32m" + str + "\033[0m"
}
