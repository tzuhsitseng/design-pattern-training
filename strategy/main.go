package main

import "fmt"

type OutputFormat int

const (
	Markdown OutputFormat = iota + 1
	HTML
)

type ListStrategy interface {
	List(data []string)
}

type TextProcessor struct {
	data   []string
	lister ListStrategy
}

func (p *TextProcessor) Print() {
	p.lister.List(p.data)
}

type MarkdownLister struct {
}

func (l *MarkdownLister) List(data []string) {
	for _, d := range data {
		fmt.Printf("- %s\n", d)
	}
}

type HTMLLister struct {
}

func (l *HTMLLister) List(data []string) {
	fmt.Println("<ul>")

	for _, d := range data {
		fmt.Printf("\t<li>%s</li>\n", d)
	}

	fmt.Println("</ul>")
}

func main() {
	tp := TextProcessor{
		data:   []string{"A", "B", "C"},
		lister: &MarkdownLister{},
	}
	tp.Print()

	tp.lister = &HTMLLister{}
	tp.Print()
}
