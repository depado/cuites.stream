<template>
  <div id="app">
    <div class="container">
      <nav-bar></nav-bar>
      <CuiteList></CuiteList>
    </div>
  </div>
</template>

<script>
import NavBar from "./components/NavBar.vue";
import CuiteList from "./components/CuiteList.vue";

export default {
  name: "app",
  components: {
    NavBar,
    CuiteList
  }
};
</script>

<style>
</style>
