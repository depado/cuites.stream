<template>
  <b-notification :closable="false" v-if="loading">
    <b-loading :is-full-page="true" :active.sync="loading" :can-cancel="false"></b-loading>
  </b-notification>
  <div v-else>
    <p class="control has-icons-left">
      <input class="input is-rounded" v-model="filter" type="text" placeholder="Search…" />
      <span class="icon is-small is-left">
        <b-icon icon="magnify" size="is-small"></b-icon>
      </span>
    </p>
    <div v-if="filtered" class="columns is-multiline" >
      <div class="column is-half" v-for="playlist in filtered" :key="playlist.id">
        <Playlist :playlist="playlist"></Playlist>
      </div>
    </div>
    <div v-else class="columns is-multiline">
      <div class="column is-half" v-for="playlist in playlists" :key="playlist.id">
        <Playlist :playlist="playlist"></Playlist>
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
      filter: null
    };
  },
  computed: {
     filtered: function() {
       if(this.filter) {
         return this.playlists.filter(p => p.title.toLowerCase().includes(this.filter.toLowerCase()) || p.user.username.toLowerCase().includes(this.filter.toLowerCase()));
       } else {
         return null
       }
     }
  },
  created() {
    axios
      .get(this.apiURL()+"/playlists")
      .then(response => {
        this.playlists = response.data;
      })
      .catch(() => {
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

<style>
.control {
  margin-bottom: 10px;
}
</style>
