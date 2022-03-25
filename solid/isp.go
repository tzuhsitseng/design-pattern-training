package main

import "fmt"

type Doc string

type AllFunctionMachine interface {
	Print(d Doc)
	Fax(d Doc)
	Scan(d Doc)
}

type OldSchoolPrinter struct {
	// ...
}

func (p OldSchoolPrinter) Print(d Doc) {
	fmt.Println("old school printer:", d)
}

func (p OldSchoolPrinter) Fax(d Doc) {
	panic("fax not supported")
}

func (p OldSchoolPrinter) Scan(d Doc) {
	panic("scan not supported")
}

type Printer interface {
	Print(d Doc)
}

type Faxer interface {
	Fax(d Doc)
}

type Scanner interface {
	Scan(d Doc)
}

type MyPrinter struct {
	// ...
}

func (p MyPrinter) Print(d Doc) {
	fmt.Println("my printer:", d)
}

func main() {
	d := Doc("my doc")
	op := OldSchoolPrinter{}
	op.Print(d)
	//op.Scan(d) // someone thought it could scan but it would panic

	mp := MyPrinter{}
	mp.Print(d) // there is only print func can use
}
