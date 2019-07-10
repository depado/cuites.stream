<template>
  <nav class="navbar is-info" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
      <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" :class="{ 'is-active': active }" @click="active=!active">
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
      </a>
    </div>
    <div class="navbar-menu" :class="{'is-active': active}">
      <div class="navbar-start">
        <router-link to="/" class="navbar-item">
          <b-icon icon="playlist-music"></b-icon>&nbsp;&nbsp;Playlists
        </router-link>
        <router-link to="/tracks" class="navbar-item">
          <b-icon icon="library-music"></b-icon>&nbsp;&nbsp;Tracks
        </router-link>
      </div>
      <div class="navbar-end">
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  data() {
    return {
      active: false
    }
  },
};
</script>

<style>
nav {
  margin-bottom: 10px;
}
.router-link-exact-active {
  color: #0e71de;
}
</style>
