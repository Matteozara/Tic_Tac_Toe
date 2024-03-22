<template>
    <div class="Prova1">    
    <body>
        <div>
            <table class="tabella">
            <tr>
                <td class="colonna">
                <button class="button" @click="check(0)" :disabled="simbolo == '' || controllo != 0">
                    <img src="./../assets/x.png" class="immagine" v-if="q[0] == 'X'">
                    <img src="./../assets/O.png" class="immagine" v-if="q[0] == 'O'">
                    </button>
                </td>
                <td>
                <button class="button" @click="check(1)" :disabled="simbolo == ''|| controllo != 0">
                    <img src="./../assets/x.png" class="immagine" v-if="q[1] == 'X'">
                    <img src="./../assets/O.png" class="immagine" v-if="q[1] == 'O'">
                </button>
                </td>
                <td>
                <button class="button" @click="check(2)" :disabled="simbolo == ''|| controllo != 0">
                    <img src="./../assets/x.png" class="immagine" v-if="q[2] == 'X'">
                    <img src="./../assets/O.png" class="immagine" v-if="q[2] == 'O'">
                </button>
                </td>
            </tr>
            <tr>
                <td class="colonna">
                <button class="button" @click="check(3)" :disabled="simbolo == ''|| controllo != 0">
                    <img src="./../assets/x.png" class="immagine" v-if="q[3] == 'X'">
                    <img src="./../assets/O.png" class="immagine" v-if="q[3] == 'O'">
                    </button>
                </td>
                <td>
                <button class="button" @click="check(4)" :disabled="simbolo == ''|| controllo != 0">
                    <img src="./../assets/x.png" class="immagine" v-if="q[4] == 'X'">
                    <img src="./../assets/O.png" class="immagine" v-if="q[4] == 'O'">
                    </button>
                </td>
                <td>
                <button class="button" @click="check(5)" :disabled="simbolo == ''|| controllo != 0">
                    <img src="./../assets/x.png" class="immagine" v-if="q[5] == 'X'">
                    <img src="./../assets/O.png" class="immagine" v-if="q[5] == 'O'">
                </button>
                </td>
            </tr>
            <tr>
                <td class="colonna">
                <button class="button" @click="check(6)" :disabled="simbolo == ''|| controllo != 0">
                    <img src="./../assets/x.png" class="immagine" v-if="q[6] == 'X'">
                    <img src="./../assets/O.png" class="immagine" v-if="q[6] == 'O'">
                </button>
                </td>
                <td>
                <button class="button" @click="check(7)" :disabled="simbolo == ''|| controllo != 0">
                    <img src="./../assets/x.png" class="immagine" v-if="q[7] == 'X'">
                    <img src="./../assets/O.png" class="immagine" v-if="q[7] == 'O'">
                </button>
                </td>
                <td>
                <button class="button" @click="check(8)" :disabled="simbolo == ''|| controllo != 0">
                    <img src="./../assets/x.png" class="immagine" v-if="q[8] == 'X'">
                    <img src="./../assets/O.png" class="immagine" v-if="q[8] == 'O'">
                </button>
                </td>
            </tr>            
            </table>
          </div>
        </body>  
    </div>
    
</template>

<script>
import {eventBus} from "../main.js";    //trasposta sia la scelta del simbolo dell'utente, sia la decisione di rigiocare o meno (refresh)
import {ris} from "../main.js"          //trasporta solo il risultato della partita
//potevo utilizzare una variabile unica penso, ma così è più chiaro mi sembra

import axios from "axios";              //lo importo per le chiamate (api) backend


export default {
    name : "prova1",
    created () {
        eventBus.$on('evento_scelta', (v)=>{this.simbolo = v;});        //ricevo simbolo scelto dall'utente
        eventBus.$on('inizia_nuova', (i) => {                           //ricevo al momento in cui decide se rigiocare (si o no), e azzero tutto, pronto per una prossima partita
            if (i == true) {
                this.q = ["-","-","-","-","-","-","-","-","-"],
                this.simbolo = '',
                this.controllo = 0,

                axios.get("http://localhost:3004/api/reset",            //azzero le mosse salvate nel backend (mosse che ha fatto l'utente nella partita)
                {}
                ).then().catch(err => {
                    console.log(err)
                })
            }
        });
    },

    data() {
        return {
            q : ["-","-","-","-","-","-","-","-","-"],  //campo da gioco
            simbolo : ""                                //simbolo scelto dall'utente
        }
    },
    methods : {
        check(a) {
            if (this.q[a] == "-") {
                axios.post("http://localhost:3004/api/mossa",       //chiamata backend ogni mossa dell'utente
                {
                T: this.q,
                Index: a,
                Simbolo: this.simbolo,
                }
            ).then((response) => {
                // console.log(this.q)
                // console.log(this.simbolo)
                // console.log(response.data.Gioco)
                // console.log(response.data.Message)

                if (response.data.Message == 0) {
                    this.q = response.data.Gioco;
                    //console.log(this.q);
                }
                if (response.data.Message == 2) {
                    this.q = response.data.Gioco;
                    console.log("Pareggio")
                    this.controllo = 2;
                }
                if (response.data.Message == -1) {
                    this.q = response.data.Gioco;
                    console.log("Hai perso");
                    this.controllo = -1;
                }
                if (response.data.Message == 1) {
                    this.q = response.data.Gioco;
                    console.log("Hai vinto");
                    this.controllo = 1;

                }
                this.swap(this.controllo)
            }).catch(err => {
                console.log(err)
            })
            }
            else {
                console.log("selezionato spazio già occupato! Selezionarne un altro")
            }
        },
        swap (k) {
            if (k != 0) {
                ris.$emit('risultato', k);      //invio il risultato della partita (se l'utente ha vinto, perso o pareggiato)
            }
        },
        refresh () {
            this.viosione = true,
            this.controllo = 0,
            this.q = ["-","-","-","-","-","-","-","-","-"],
            this.simbolo = ""
        }
    },
    props: {
        controllo : {           //nei props perchè viene esportata, controllo rappresenta se il giocatore ha vinto, perso o pareggiato, o nessuna di queste
            type : Number,
            default : 0
        }
    }
} 
</script>

<style scoped>
    .Prova1 {
        display: flex;
    }

    .button {
        height: 150px;
        width: 150px;
        background: black;
        color: white;
    }
    .colonna {
        align-items: center;
        align-self: center;
        margin-left: 200px;
        margin-right: 200px;
    }
    .tabella {
        margin-top: 10%;
        margin-left: auto;
        margin-right: auto;
        height: 600px;
        width: 600px;
    }
    .immagine {
        height: 150px;
    }
</style>