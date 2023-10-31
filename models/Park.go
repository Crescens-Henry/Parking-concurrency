package models

import (
	"sync"

	"fyne.io/fyne/v2/canvas"
)

type Park struct {
	Space          chan bool
	DrawCar        chan *canvas.Image
	CarExit        chan bool // Canal para notificar la salida de carros
	mutex          sync.Mutex
	OccupiedSpaces int // Agrega este campo para contar los espacios ocupados
}

func NewPark(nS int) *Park {
	return &Park{
		Space:          make(chan bool, nS+1),
		DrawCar:        make(chan *canvas.Image, 1),
		CarExit:        make(chan bool, nS), // Tamaño igual al número máximo de carros
		OccupiedSpaces: 0,                   // Inicializa en 0
	}
}
