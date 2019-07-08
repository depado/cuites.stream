<template>
  <div>
    <b-notification :closable="false" v-if="loading">
      <b-loading :is-full-page="true" :active.sync="loading" :can-cancel="false"></b-loading>
    </b-notification>
    <div class="columns is-multiline" id="playlists" v-if="!loading">
      <div class="column is-full" v-for="playlist in playlists" :key="playlist.id">
        <Playlist :playlist=playlist></Playlist>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Playlist from "./Playlist.vue";

export default {
  components: {
    Playlist
  },
  data() {
    return {
      playlists: null,
      loading: true,
      error: null,
    };
  },
  created() {
    axios
      .get("http://localhost:8081/playlists")
      .then(response => {
        this.playlists = response.data;
        this.$toast.open({
          message: "I got the playlists",
          type: "is-success",
          position: "is-bottom"
        });
      })
      .catch(error => {
        this.error = error;
        this.$toast.open({
          duration: 5000,
          message: `Unable to retrieve playlists`,
          position: "is-bottom",
          type: "is-danger"
        });
      })
      .finally(() => (this.loading = false));
  }
};
</script>