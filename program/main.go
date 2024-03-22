package main

import (
	//"crypto/rand"
	"fmt"
	"math/rand"
)

func stampa(a [9]string) {
	fmt.Println("Tabellone:")
	for i := 0; i < 9; i++ {
		fmt.Print(" ", a[i], " ")
		if i == 2 || i == 5 || i == 8 {
			fmt.Print("\n")
		}
	}
}

func controlloRighe(t [9]string, a int, simbolo string) int {
	if a >= 0 && a <= 2 {
		if t[0] == simbolo && t[0] == t[2] {
			if t[1] == "-" {
				return 1
			}
		}
		if t[0] == simbolo && t[0] == t[1] {
			if t[2] == "-" {
				return 2
			}
		}
		if t[1] == simbolo && t[1] == t[2] {
			if t[0] == "-" {
				return 0
			}
		}
	}
	if a >= 3 && a <= 5 {
		if t[3] == simbolo && t[3] == t[5] {
			if t[4] == "-" {
				return 4
			}
		}
		if t[3] == simbolo && t[3] == t[4] {
			if t[5] == "-" {
				return 5
			}
		}
		if t[4] == simbolo && t[4] == t[5] {
			if t[3] == "-" {
				return 3
			}
		}
	}
	if a >= 6 && a <= 8 {
		if t[6] == simbolo && t[6] == t[8] {
			if t[7] == "-" {
				return 7
			}
		}
		if t[6] == simbolo && t[6] == t[7] {
			if t[8] == "-" {
				return 8
			}
		}
		if t[7] == simbolo && t[7] == t[8] {
			if t[6] == "-" {
				return 6
			}
		}
	}
	return -1
}

func controlloColonne(t [9]string, a int, simbolo string) int {
	if a == 0 || a == 3 || a == 6 {
		if t[0] == simbolo && t[0] == t[3] {
			if t[6] == "-" {
				return 6
			}
		}
		if t[0] == simbolo && t[0] == t[6] {
			if t[3] == "-" {
				return 3
			}
		}
		if t[3] == simbolo && t[3] == t[6] {
			if t[0] == "-" {
				return 0
			}
		}
	}
	if a == 1 || a == 4 || a == 7 {
		if t[1] == simbolo && t[1] == t[4] {
			if t[7] == "-" {
				return 7
			}
		}
		if t[1] == simbolo && t[1] == t[7] {
			if t[4] == "-" {
				return 4
			}
		}
		if t[4] == simbolo && t[4] == t[7] {
			if t[1] == "-" {
				return 1
			}
		}
	}
	if a == 2 || a == 5 || a == 8 {
		if t[2] == simbolo && t[2] == t[5] {
			if t[8] == "-" {
				return 8
			}
		}
		if t[2] == simbolo && t[2] == t[8] {
			if t[5] == "-" {
				return 5
			}
		}
		if t[5] == simbolo && t[5] == t[8] {
			if t[2] == "-" {
				return 2
			}
		}
	}
	return -1
}

func controlloDiagonali(t [9]string, a int, simbolo string) int {
	if a == 0 || a == 4 || a == 8 {
		if t[0] == simbolo && t[0] == t[4] {
			if t[8] == "-" {
				return 8
			}
		}
		if t[0] == simbolo && t[0] == t[8] {
			if t[4] == "-" {
				return 4
			}
		}
		if t[4] == simbolo && t[4] == t[8] {
			if t[0] == "-" {
				return 0
			}
		}
	}
	if a == 2 || a == 4 || a == 6 {
		if t[2] == simbolo && t[2] == t[4] {
			if t[6] == "-" {
				return 6
			}
		}
		if t[2] == simbolo && t[2] == t[6] {
			if t[4] == "-" {
				return 4
			}
		}
		if t[4] == simbolo && t[4] == t[6] {
			if t[2] == "-" {
				return 2
			}
		}
	}
	return -1
}

func sceltaCasuale(t [9]string) int {
	for z := 0; z < 9; z++ {
		if t[z] == "-" {
			return z
		}
	}
	return -1
}

func controllo(t [9]string) string {
	//righe
	for i := 0; i < 9; {
		if t[i] != "-" {
			if t[i] == t[i+1] && t[i] == t[i+2] {
				return t[i]
			}
		}
		i = i + 3
	}

	//colonne
	for i := 0; i < 3; i++ {
		if t[i] != "-" {
			if t[i] == t[i+3] && t[i] == t[i+6] {
				return t[i]
			}
		}

	}
	//diagonali
	if t[0] != "-" && t[0] == t[4] && t[0] == t[8] {
		return t[0]
	}
	if t[2] != "-" && t[2] == t[4] && t[2] == t[6] {
		return t[2]
	}
	return "continua"
}

func main() {
	t := [9]string{"-", "-", "-", "-", "-", "-", "-", "-", "-"} //tabellone di gioco
	ang := [4]int{0, 2, 6, 8}                                   //posizioni negli angoli
	k := 0                                                      //conta mosse
	var index int                                               //indice passato dall'utente
	pos := -1                                                   //posizione scelta dal computer
	var mosse []int                                             //mosse fatte dal mio avversario
	var a string                                                //simbolo scelto dal mio avversario
	sol := "continua"                                           //soluzione del gioco

	b := "b" //simbolo del computer

	stampa(t)

	//SCELTA SIMBOLO
	for b == "b" {
		fmt.Println("inizia a giocare! \nScegli il simbolo: O o X?")
		fmt.Scan(&a)

		switch a {
		case "O":
			b = "X"
		case "X":
			b = "O"
		default:
			fmt.Println("Errore, il sombolo digitato non è valido!")
		}
	}

	//GIOCO
	for sol == "continua" && k < 10 {
		val := true
		var x int
		var y int

		if k%2 == 0 {

			//controllo che non vengano sovrascritte posizioni
			for val == true {
				//RIGA
				for val == true {
					fmt.Println("Segli la riga (tra 0 e 2):")
					fmt.Scan(&x)
					if x >= 0 && x <= 2 {
						val = false
					} else {
						fmt.Println("Errore!")
					}
				}

				//COLONNA
				val = true
				for val == true {
					fmt.Println("Segli la colonna (tra 0 e 2):")
					fmt.Scan(&y)
					if y >= 0 && y <= 2 {
						val = false
					} else {
						fmt.Println("Errore!")
					}
				}

				index = (x * 3) + y
				if t[index] == "-" {
					val = false
				} else {
					fmt.Println("Errore, le coordinate digitate sono gia occupate da un simbolo!")
					val = true
				}
			}

			//stampe da eliminare
			fmt.Println("")
			fmt.Println("riga ", x)
			fmt.Println("colonna ", y)
			fmt.Println("")

			t[index] = a

			mosse = append(mosse, index)
			fmt.Println(mosse) //stampa da eliminare
		}

		if k%2 != 0 {

			//SCELTA DEL COMPUTER
			if k == 1 {
				if t[4] == "-" {
					t[4] = b
				} else {
					angolo := rand.Intn(4)
					fmt.Println(ang[angolo]) //stampa da eliminare
					t[ang[angolo]] = b
				}
			} else {
				for z := 0; z < len(mosse); z++ {
					pos = controlloRighe(t, mosse[z], a)
					if pos != -1 {
						t[pos] = b
						z = 100
					} else {
						pos = controlloColonne(t, mosse[z], a)
						if pos != -1 {
							t[pos] = b
							z = 100
						} else {
							pos = controlloDiagonali(t, mosse[z], a)
							if pos != -1 {
								t[pos] = b
								z = 100
							}
						}
					}
				}
				if pos == -1 {
					pos = sceltaCasuale(t)
					if pos != -1 {
						t[pos] = b
					} else {
						sol = "break"
					}
				}
			}
		}

		fmt.Println("")
		stampa(t)
		fmt.Println("")

		k++

		//CONTROLLO CHE NESSUNO VINCA
		if sol != "break" {
			sol = controllo(t)
		}
	}

	if sol == a {
		fmt.Println("Complimenti, hai vinto!!")
	} else {
		if sol == b {
			fmt.Println(("Game over, hai perso!"))
		} else {
			fmt.Println("Parità")
		}
	}

}
