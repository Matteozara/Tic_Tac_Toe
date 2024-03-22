package main

import (
	"log"
	"math/rand"

	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var mosse []int
var ang = [4]int{0, 2, 6, 8}

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

func controlloPosizione(t [9]string, index int) bool {
	if t[index] == "-" {
		return true
	} else {
		return false
	}
}

func sceltaComputer(t [9]string, simbolo string) ([9]string, int) {
	var b string
	pos := -1

	//giro i simbolo
	if simbolo == "O" {
		b = "X"
	} else {
		b = "O"
	}

	if len(mosse) <= 1 { //se prima mossa computer:
		if t[4] == "-" {
			t[4] = b
			return t, 0
		} else {
			angolo := rand.Intn(4)
			t[ang[angolo]] = b
			return t, 0
		}
	} else { //se non è la prima mossa
		for z := 0; z < len(mosse); z++ {
			pos = controlloRighe(t, mosse[z], simbolo)
			if pos != -1 {
				t[pos] = b
				return t, 0
			} else {
				pos = controlloColonne(t, mosse[z], simbolo)
				if pos != -1 {
					t[pos] = b
					return t, 0
				} else {
					pos = controlloDiagonali(t, mosse[z], simbolo)
					if pos != -1 {
						t[pos] = b
						return t, 0
					}
				}
			}
		}
		if pos == -1 {
			pos = sceltaCasuale(t)
			if pos != -1 {
				t[pos] = b
				return t, 0
			} else {
				return t, 2
			}
		}
	}
	return t, -1
}

func assegna(t [9]string, index int, s string) [9]string {
	t[index] = s
	return t
}

func mossa(w http.ResponseWriter, r *http.Request) {
	type Http_req struct {
		T       [9]string `json:"t"`
		Index   int       `json:"index"`
		Simbolo string    `json:"simbolo"`
	}

	var soluzione string
	var h Http_req
	_ = json.NewDecoder(r.Body).Decode(&h)

	a := controlloPosizione(h.T, h.Index)

	if a != true {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(bson.M{"Gioco": -1, "Message": "Posizione gia occupata"})
		return
	}

	h.T = assegna(h.T, h.Index, h.Simbolo)
	mosse = append(mosse, h.Index)

	//controllo
	soluzione = controllo(h.T)
	if soluzione != "continua" {
		if soluzione == h.Simbolo {
			json.NewEncoder(w).Encode(bson.M{"Gioco": h.T, "Message": 1}) //Message = 1  => Hai vinto
			return
		} else {
			json.NewEncoder(w).Encode(bson.M{"Gioco": h.T, "Message": -1}) //Message = -1  => Hai perso
			return
		}
	}

	err := 0
	h.T, err = sceltaComputer(h.T, h.Simbolo)

	if err == 2 {
		json.NewEncoder(w).Encode(bson.M{"Gioco": h.T, "Message": 2}) //caso pareggio!
		return
	}

	if err != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(bson.M{"Gioco": -1, "Message": "Si sono verificati degli errori"})
		return
	}

	//controllo
	soluzione = controllo(h.T)
	if soluzione != "continua" {
		if soluzione == h.Simbolo {
			json.NewEncoder(w).Encode(bson.M{"Gioco": h.T, "Message": 1}) //Message = 1  => Hai vinto
			return
		} else {
			json.NewEncoder(w).Encode(bson.M{"Gioco": h.T, "Message": -1}) //Message = -1  => Hai perso
			return
		}
	}

	json.NewEncoder(w).Encode(bson.M{"Gioco": h.T, "Message": 0})
}

func reset(w http.ResponseWriter, r *http.Request) {
	var m []int
	mosse = m

	json.NewEncoder(w).Encode(bson.M{"Result": "OK"})
}

func main() {

	r := mux.NewRouter()
	//r := http.NewServeMux()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"}) //cosi consento tutto

	r.HandleFunc("/api/mossa", mossa).Methods("POST")
	r.HandleFunc("/api/reset", reset).Methods("GET")

	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	log.Fatal(http.ListenAndServe(":3004", handlers.CORS(header, methods, origins)(r)))

}
