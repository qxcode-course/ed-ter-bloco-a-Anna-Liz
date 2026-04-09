package main

import (
	"fmt"
)

func circulos(pen *Pen, raio float64) {
	if raio < 60 {
		return
	}
	for range 6 {
		pen.DrawCircle(raio)

		pen.Left(60)
		pen.Walk(raio)
		circulos(pen, raio-120)
	}
}

func main() {
	pen := NewPen(600, 500)
	pen.SetHeading(90)
	pen.SetPosition(300, 250)
	circulos(pen, 160)
	pen.SavePNG("circulo.png")
	fmt.Println("PNG file created successfully.")
}
