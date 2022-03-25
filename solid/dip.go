package main

import "fmt"

// the key point is that HLM and LLM should depend on abstractions or interfaces

type Database interface {
	Load(key string)
	Save(data []byte)
}

type MySQL struct{}

func (m MySQL) Load(key string) {
	fmt.Println("load from MySQL")
}

func (m MySQL) Save(data []byte) {
	fmt.Println("save to MySQL")
}

type Postgres struct{}

func (p Postgres) Load(key string) {
	fmt.Println("load from Postgres")
}

func (p Postgres) Save(data []byte) {
	fmt.Println("save to Postgres")
}

type BadRepo struct {
	mysql MySQL
}

type BetterRepo struct {
	database Database
}

func main() {
	m := MySQL{}
	p := Postgres{}
	badRepo := BadRepo{mysql: m}
	fmt.Println(badRepo)

	//betterRepo := BetterRepo{database: m}
	betterRepo := BetterRepo{database: p}
	fmt.Println(betterRepo)
}
