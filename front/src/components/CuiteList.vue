<template>
  <div>
    <b-notification :closable="false" v-if="loading">
      <b-loading :is-full-page="true" :active.sync="loading" :can-cancel="false"></b-loading>
    </b-notification>
    <div class="columns is-multiline" id="playlists" v-if="!loading">
      <div class="column is-one-quarter" v-for="playlist in playlists" :key="playlist.id">
        <Cuite :playlist=playlist></Cuite>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Cuite from "./Cuite.vue";

export default {
  components: {
    Cuite
  },
  data() {
    return {
      playlists: null,
      loading: true
    };
  },
  created() {
    axios
      .get("http://localhost:8081/cuites")
      .then(response => {
        this.playlists = response.data;
        this.$toast.open({
          message: "I got the playlists",
          type: "is-success",
          position: "is-bottom"
        });
      })
      .catch(error => {
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