<template>
    <div class="Scritte">
        <div v-if="r == 0">
        <p v-if="s == ''">
            <b>
                <h2 class="start">INIZIAMO</h2>
                <br>
                <div class="scritta">
                Scegli il tuo simbolo
                <br>
                <small>(sulla destra)</small>
                </div>
            </b>
        </p>
        <p v-if="s == 'O'">
            <b>
                <h2 class="blue">TOCCA  <br>A TE!</h2>
                
                <br>
                <div class="scritta">
                Scegli dove posizionare il simbolo
                <div>
                <img src="./../assets/O.png" class="immagine">
                </div>
                </div>
            </b>
        </p>
        <p v-if="s == 'X'">
            <b>
                <h2 class="red">TOCCA  <br>A TE!</h2>
                
                <br>
                <div class="scritta">
                Scegli dove posizionare il simbolo
                <div>
                <img src="./../assets/x.png" class="immagine">
                </div>
                </div>
            </b>
        </p>
        </div>

        <div v-if="r==-1">
            <div class="tit">Hai Perso!</div>
            <br>
            <div class="testo">
                Vuoi giocare un altra partita?
            </div>
            <div style="padding-top:10%">
                <button class="si" @click="refresh(true)"><b>SI</b></button>
                <button class="si" @click="refresh(false)"><b>NO</b></button>
            </div>
        </div>

        <div v-if="r==1">
            <div class="tit">Hai Vinto!</div>
            <br>
            <div class="testo">
                Vuoi giocare un altra partita?
            </div>
            <div style="padding-top:10%">
                <button class="si" @click="refresh(true)"><b>SI</b></button>
                <button class="si" @click="refresh(false)"><b>NO</b></button>
            </div>
        </div>

        <div v-if="r==2">
            <div class="tit">Pareggio</div>
            <br>
            <div class="testo">
                Vuoi giocare un altra partita?
            </div>
            <div style="padding-top:10%">
                <button class="si" @click="refresh(true)"><b>SI</b></button>
                <button class="si" @click="refresh(false)"><b>NO</b></button>
            </div>
        </div>
        <div v-if="s=='a'">
            <div class="tit">A presto!</div>
            <br>
            <div class="piccolo">
                (Se cambi idea, scegli un simbolo sulla destra)
            </div>
        </div>
    </div>
</template>

<script>
import {eventBus} from "../main.js" //trasposta sia la scelta del simbolo dell'utente, sia la decisione di rigiocare o meno (refresh)
import {ris} from "../main.js"      //trasporta solo il risultato della partita
//potevo utilizzare una variabile unica penso, ma così è più chiaro mi sembra

export default {
    created () {
        eventBus.$on('evento_scelta', (v)=>{this.s = v;});  //ricevo la scelta del simbolo 
        ris.$on('risultato', (v)=>{this.r = v;});   //ricevo il risultato della partita (1 = vittoria, -1 = sconfitta, 2 = pareggio)
    },
    
    data() {
        return {
            s : "", //simbolo scelto dall'utente (s == 'a', simbolo scelto come decisione di non volere rigiocare)
            r : 0   //risultato della partita
        }
    },
    props : {
        inizia : {                  //perchè la invio fuori, quetsa variabile
            type : Boolean
        }
    },
    methods : {
        refresh(i) {
            this.inizia = true;                             //invio se l'utente vuole fare un altra partita o meno
            eventBus.$emit('inizia_nuova', this.inizia);    //in ogni caso invio la segnalazione di azzerare tutti campi (pronto per iniziare una nuova partita o meno)
            this.r = 0;
            if (i == false) {
                this.s = 'a'
            } else {
                this.s = "";

            }
        }
    }
    
}
</script>

<style scoped>
    .scritta {
        font-size: 50px;
        text-align: center;
        color: white;
    }
    .start {
        font-size: 60px;
        text-align: center;
    }
    .blue {
        font-size: 60px;
        text-align: center;
        color: rgb(57, 168, 202);
    }
    .red {
        font-size: 60px;
        text-align: center;
        color: red;
    }
    .fondo {
        align-items: center;
        padding-top: 10%;
    }
    .immagine {
        height: 50px;
        width: 50px;
    }
    .tit {
        padding-top: 20%;
        padding-left: 20px;
        font-size: 100px;
        color: gold;
        text-align: center;
    }
    .testo {
        font-size: 50px;
        color: gold;
        text-align: center;
    }
    .si {
        color: black;
        background-color: gold;
        width: 100px;
        height: 50px;
    }
    .piccolo {
        font-size: 30px;
        color: gold;
        text-align: center;
        padding-top: 20%;
    }
</style>