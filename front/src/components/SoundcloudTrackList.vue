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
    <div v-if="filtered" class="columns is-multiline">
      <div class="column is-half" v-for="(track, i) in filtered" :key="`${i}-${track.id}`">
        <SoundcloudTrack :track="track"></SoundcloudTrack>
      </div>
    </div>
    <div v-else class="columns is-multiline">
      <div class="column is-half" v-for="(track, i) in tracks" :key="`${i}-${track.id}`">
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
      filter: null
    };
  },
  computed: {
    filtered: function() {
      if (this.filter && this.filter.length > 2) {
        return this.tracks.filter(
          t =>
            t.title.toLowerCase().includes(this.filter.toLowerCase()) ||
            t.user.username.toLowerCase().includes(this.filter.toLowerCase())
        );
      } else {
        return null;
      }
    }
  },
  created() {
    console.log(this.apiURL());
    axios
      .get(this.apiURL()+"/tracks")
      .then(response => {
        this.tracks = response.data;
      })
      .catch(() => {
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
.control {
  margin-bottom: 10px;
}
</style>
