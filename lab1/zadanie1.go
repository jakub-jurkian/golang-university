package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var proby int = 10000

	// 1 czesc
	var n int = 3
	var liczbaTrafionychLosowan int = 0
	for i := 0; i < proby; i++ {

		nagroda := rand.Intn(n) + 1
		wyborGracza := rand.Intn(n) + 1

		wyborProwadzacego := rand.Intn(n) + 1
		for {
			if wyborProwadzacego != wyborGracza && wyborProwadzacego != nagroda {
				break
			}
			wyborProwadzacego = rand.Intn(n) + 1
		}

		if nagroda == wyborGracza {
			liczbaTrafionychLosowan++
		}
	}
	fmt.Println("Bez zmiany (N=3):", liczbaTrafionychLosowan)

	// 2 czesc
	n = 3
	k := 1
	liczbaTrafionychLosowanZeZmiana := 0

	for i := 0; i < proby; i++ {
		nagroda := rand.Intn(n) + 1
		wyborGracza := rand.Intn(n) + 1

		var otwartePudla []int

		// proces otwierania pudel przez prowadzacego
		for len(otwartePudla) < k {
			p := rand.Intn(n) + 1

			if p != nagroda && p != wyborGracza {
				juzOtwarte := false
				for j := 0; j < len(otwartePudla); j++ {
					if otwartePudla[j] == p {
						juzOtwarte = true
						break
					}
				}

				if !juzOtwarte {
					otwartePudla = append(otwartePudla, p)
				}
			}
		}

		// proces wyboru nowego pudla przez gracza
		staryWyborGracza := wyborGracza
		for {
			wyborGracza = rand.Intn(n) + 1

			if wyborGracza != staryWyborGracza {
				pudloJestOtwarte := false
				for j := 0; j < len(otwartePudla); j++ {
					if otwartePudla[j] == wyborGracza {
						pudloJestOtwarte = true
						break
					}
				}

				if !pudloJestOtwarte {
					break
				}
			}
		}

		if nagroda == wyborGracza {
			liczbaTrafionychLosowanZeZmiana++
		}
	}

	fmt.Println("Ze zmiana (N=3, K=1):", liczbaTrafionychLosowanZeZmiana)
}
