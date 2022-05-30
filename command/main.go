package main

func main() {
	tv := &tv{}

	onCmd := onCommand{device: tv}
	offCmd := offCommand{device: tv}

	onBtn := button{cmd: onCmd}
	offBtn := button{cmd: offCmd}

	onBtn.press()
	offBtn.press()
}
