package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Journal struct {
	entries []string
	index   int
}

func (j *Journal) AddEntry(note string) {
	j.index++
	j.entries = append(j.entries, fmt.Sprintf("%d: %s", j.index, note))
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

// Bad Example
//func (j *Journal) Save(filename string) {
//	_ = ioutil.WriteFile(filename,
//		[]byte(strings.Join(j.entries, "\n")), 0644)
//}

// ---
// 新增一個 persistence 物件來幫忙 journal 這個物件作 save to file 的動作
// 而不是把所有工作都攬在 journal 這個物件身上

type Persistence struct{}

func (p *Persistence) SaveToFile(journal Journal, filename string) error {
	return ioutil.WriteFile(filename, []byte(strings.Join(journal.entries, "\n")), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("Learn algorithm")
	j.AddEntry("Learn OS")
	p := Persistence{}
	if err := p.SaveToFile(j, "journal.txt"); err != nil {
		fmt.Println(err)
	}
}
