<template>
  <div id="gestione">
    <OX/> 
    <Grid/>

  </div>
</template>

<script>
import OX from './components/OX.vue'
import Grid from './components/Grid.vue'


export default {
  name: 'Gestione',
  components: {
    OX,
    Grid
  }
}
</script>

<style>
#gestione {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
