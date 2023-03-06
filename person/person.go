package person

import "fmt"

//構造体Person
type Person struct {
	Name  string
	Score int
}

//Person型の構造体を作成する関数
func NewPerson(name string, score int) *Person {
	return &Person{Name: name, Score: score}
}

//名前を言うメソッド
func (p *Person) SayName() {
	fmt.Printf("My name is %s.\n", p.Name)
}

//名前とスコアを言うメソッド
func (p *Person) SayScore() {
	fmt.Printf("I'm %s. My score is %d.\n", p.Name, p.Score)
}

//スコアに加算するメソッド
func (p *Person) AddScore(n int) {
	p.Score += n
}
