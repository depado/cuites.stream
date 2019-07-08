<template>
  <div>
    <b-notification :closable="false" v-if="loading">
      <b-loading :is-full-page="true" :active.sync="loading" :can-cancel="false"></b-loading>
    </b-notification>
    <div class="columns is-multiline" id="tracks" v-if="!loading">
      <div class="column is-half" v-for="track in tracks" :key="track.id">
        <SoundcloudTrack :track="track"></SoundcloudTrack>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import SoundcloudTrack from "./SoundcloudTrack.vue";

export default {
  components: {
    SoundcloudTrack
  },
  data() {
    return {
      tracks: null,
      loading: true,
      error: null
    };
  },
  created() {
    axios
      .get("http://localhost:8081/tracks")
      .then(response => {
        this.tracks = response.data;
        this.$toast.open({
          message: "I got the tracks",
          type: "is-success",
          position: "is-bottom"
        });
      })
      .catch(error => {
        this.error = error;
        this.$toast.open({
          duration: 5000,
          message: `Unable to retrieve tracks`,
          position: "is-bottom",
          type: "is-danger"
        });
      })
      .finally(() => (this.loading = false));
  }
};
</script>

<style>
</style>
