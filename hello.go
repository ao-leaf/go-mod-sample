package mymod

import "fmt"

func Hello() {
	fmt.Printf("Hello.\n")
}

//パッケージ外に公開されない関数
func helloJP() {
	fmt.Printf("こんにちは。\n")
}

//enかjpなら挨拶する関数
func Greet(s string) {
	switch s {
	case "en", "EN":
		Hello()
	case "jp", "JP":
		helloJP()
	default:
		fmt.Printf("(´・ω・`)\n")
	}
}
