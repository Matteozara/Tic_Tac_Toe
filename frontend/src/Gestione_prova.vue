<template>
  <div id="gestione">
    <Title/>
    <div class="struttura">
      <el-card class="box-card0">
        <div>
          <Scritte/>
        </div>
      </el-card>
      <el-card class="box-card1">
        <div>
          <Prova1/>
        </div>
      </el-card>
      <el-card class="box-card2">
        <div>
          <Select/>
          <div class="fondo">
            <img alt="Vue logo" src="@/assets/logo.png" class="immagine">
          </div>
        </div>
      </el-card>
    </div>

  </div>
</template>

<script>
import Prova1 from './components/Prova1.vue'
import Select from './components/Select.vue'
import Title from './components/Title.vue'
import  Scritte from './components/Scritte.vue'



export default {
  name: 'Gestione',
  components: {
    Prova1,
    Select,
    Title,
    Scritte
  }
}
</script>

<style>

.struttura {
  display: flex;
}

.immagine {
  height: 100px;
  width: 100px;
}

.fondo {
  padding-top: 60%;
  padding-left: 80%;
}

.box-card0 {
  width: 20%;
}

.box-card1 {
  width: 40%;
  padding-left: 12%;
}

.box-card2 {
  width: 20%;  
}
</style>
