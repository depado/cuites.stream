<template>
  <div id="app">
    <nav-bar></nav-bar>
    <div class="container">
      <router-view />
    </div>
    <Player></Player>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue";
import Player from "@/components/Player.vue"

export default {
  components: {
    NavBar,
    Player
  }
};
</script>


<style>
html {
  background-color: #f5f5f5 !important;
}
.container {
  margin-bottom: 70px !important;
}
.columns {
  margin-right: 0 !important;
  margin-left: 0 !important;
}
</style>
