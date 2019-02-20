package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// insertando la semilla para mis números aleatorios
	rand.Seed(time.Now().UnixNano())
	scoreTotal := 0

	for ronda := 0; ronda < 9; ronda++ {

		tiro, scoreRound, pinosTotal := 0, 0, 10
		//strike, spare := false, false
		strike := false

		for turno := 0; turno < 2; turno++ {

			tiro = rand.Intn(pinosTotal + 1)

			if tiro == 10 {
				//strike = true
				//scoreTotal += 10
				strikeNo()
				break
			}

			pinosTotal -= tiro
			scoreRound += tiro

			//if scoreRound == 10 {
			//
			//	break
			//}

			fmt.Println(tiro)

		}

		scoreTotal += scoreRound
		if strike == false {
			fmt.Println("Score: ", scoreTotal)
		} else {
			fmt.Println("Hiciste un STRIKE", scoreTotal)
		}
	}
}


func spare() {

}


func strikeNo() {
	fmt.Println("hola")
}


func generadorNumeros(rango int) {


}