package main

func redBold(str string) string {
	return "\033[1;31m" + str + "\033[0m"
}

func greenBold(str string) string {
	return "\033[32m" + str + "\033[0m"
}
