<template>

  <div id="app">
    
    <Gestione/>
  </div>
</template>

<script>
import Gestione from './Gestione_prova.vue'

export default {
  name: 'App',
  components: {
    Gestione
  }
}
</script>

<style>
#app {
  text-align: center;
  background-image: linear-gradient(to top right, #f8f8f8 0%, #414040 50%, #000000 100%);
  
  /* background tutto schermo */
  position: fixed; 
  top: 0; 
  left: 0; 
  min-width: 100%;
  min-height: 100%;
}
</style>
