package main

import (
	"fmt"
)

type Lamp struct {
	color string
	state string
}

func (l *Lamp) NewLamp(color, state string) {
	l = new(Lamp)
	l.SetColor(color)
	l.SetState(state)
	fmt.Println(l.state)
}

func (l Lamp) GetColor() string {
	return l.color
}

func (l *Lamp) SetColor(color string) {
	switch color {
	case "red", "green", "orange", "yellow", "blue", "purple":
		l.color = color
	default:
		l.color = "black"
	}
}

func (l Lamp) GetState() string {
	fmt.Println(l.state)
	return l.state
}

func (l *Lamp) SetState(state string) {
	switch state {
	case "on", "off":
		l.state = state
	default:
		l.state = "off"
	}
}

func (l Lamp) SummonGenie() bool {
	return l.color == "red" && l.state == "on"
}

func genie() {
	var lamp Lamp
	lamp.NewLamp("black", "off")

	for i := 0; i < 6; i++ {
		var input string
		fmt.Scan(&input)
		switch input {
		case "on", "off":
			lamp.SetState(input)
		default:
			lamp.SetColor(input)
		}
		fmt.Println(lamp.GetColor(), lamp.GetState())
	}

	if lamp.SummonGenie() {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
