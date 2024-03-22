<template>
    <div class="Select">
        <div class="spazio">
            <button class="button" @click="Select('X')" :disabled="scelta == 'O' || scelta == 'X'"> <b>Select</b>  <img src="./../assets/x.png" class="immagine"></button>
        </div>
        <div>
            <button class="button" @click="Select('O')" :disabled="scelta == 'X' || scelta == 'O'"> <b>Select</b>  <img src="./../assets/O.png" class="immagine"></button>
        </div>
        <br>
        <br>
    </div>

</template>

<script>
import {eventBus} from "../main.js";

export default {
    created () {
        eventBus.$on('inizia_nuova', (i) => {   //[ricevo dall'esterno]
            if (i == true) {    // i = true sempre (sia che inizio una nuova partita sia se scelgo di no)
                this.scelta = '';
            }
        });
    },
    data() {
        return {
        
        }
    },
    methods : {
        Select(i) {
            this.scelta = i;
            eventBus.$emit('evento_scelta', this.scelta);
        }
    },
    props: {
        scelta : {
            type : String,
            default : ''
        }
    }
    
}
</script>

<style scoped>
    .Select {
        display: flexbox;
        margin-top: 40%;
        padding-right: 20%;
    }
    .immagine {
        height: 50px;
        width: 50px;
        margin-top: 10px;
    }
    .button {
        height: 80px;
        width: 300px;
        font-size: 50px;
        background: black;
        color: white;
    }
    .spazio {
        margin-bottom: 10%;
    }

</style>