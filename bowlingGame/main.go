package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	scoreTotal :=  0
	bonusStrike := 0
	strike, spare := false, false
	bandera := false


	for ronda := 0; ronda < 9; ronda++ {

		scoreRound, pinosTotal := 0, 10

		guardarScore := ""
		var tiros[2]int

		for turno := 0; turno < 2; turno++ {

			tiro := generadorNumeros(pinosTotal)
			tiros[turno] = tiro

			if tiro == 10 {
				strike = true

				if bonusStrike == 2 {
					bonusStrike = 1
					bandera = true
				} else {
					bonusStrike = 2
				}

				scoreTotal += 10

				if spare != true && bandera != true {
					fmt.Println(tiro)
				}
				break
			}

			if bonusStrike != 0 {
				guardarScore += fmt.Sprintf(" %d \n", tiro)
				scoreTotal += tiro
				bonusStrike--

				if bonusStrike == 0 {
					fmt.Println("STRIKE, score:", scoreTotal)

					if bandera {
						fmt.Println(10)
						scoreTotal += 10
						bandera = false
						scoreTotal += scoreRound
						bonusStrike = 1
					} else {
						fmt.Println(guardarScore)
					}

					scoreRound += tiro
					if scoreRound == 10 {
						spare = true
					}
					scoreTotal += scoreRound
					strike = false
					break
				}
			}

			if spare == true {
				scoreTotal += tiro
				fmt.Println("SPARE, score:",scoreTotal)
				spare = false
			}

			if strike == false && spare == false {
				fmt.Println(tiro)
				scoreTotal += tiro

			}
			scoreRound += tiro

			if scoreRound == 10 {
				spare = true
			}

			pinosTotal -= tiro
		}



		if spare == false && strike == false {
			fmt.Println("Score: ", scoreTotal)
		} else if spare == true {
			//fmt.Println("SPARE")
		} else if strike == true {
			//fmt.Println("STRIKE")
		} else if strike == true && spare == true {
			fmt.Println("SPARE, score:",scoreTotal)
			spare = false
		}
	}
}



func generadorNumeros(rango int) int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(rango + 1)
	return num
}
