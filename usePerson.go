package mymod

import (
	"github.com/ao-leaf/go-mod-sample/person"
)

// personパッケージで定義した構造体Personとそのメソッドを使う関数
func UsePerson() {
	taro := person.NewPerson("Taro", 0)
	Jiro := person.NewPerson("Jiro", 50)

	taro.SayName()
	Jiro.SayName()

	taro.SayScore()
	Jiro.SayScore()

	taro.AddScore(100)
	Jiro.AddScore(50)

	taro.SayScore()
	Jiro.SayScore()
}
