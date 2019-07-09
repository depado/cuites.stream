<template>
  <nav class="navbar is-info" role="navigation" aria-label="main navigation">
    <div class="navbar-menu">
      <div class="navbar-start">
          <router-link to="/" class="navbar-item"><b-icon icon="playlist-music"></b-icon>&nbsp;&nbsp;Playlists</router-link>
          <router-link to="/tracks" class="navbar-item"><b-icon icon="library-music"></b-icon>&nbsp;&nbsp;Tracks</router-link>
      </div>
    </div>
  </nav>
</template>

<script>
export default {};
</script>

<style>
nav {
  margin-bottom: 10px;
}
.router-link-exact-active {
  background-color: #0e71de;
}
</style>
