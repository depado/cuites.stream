<template>
  <div>
    <div class="flex-container">
      <div class="flex-item cover">
        <div v-if="playing || currentTrack">
          <img :src="currentArtwork" />
          <p class="metadata">
            <a :href="currentTrack.permalink_url">{{ currentTrack.title }}</a>
            <br />
            <small>
              <b-icon icon="clock-outline" size="is-small"></b-icon>
              {{ duration_str }}
              <b-icon icon="account" size="is-small"></b-icon>
              {{ currentTrack.user.username }}
            </small>
          </p>
        </div>
      </div>
      <div class="flex-item controls">
        <a v-if="hasPrev" @click="previous">
          <b-icon icon="skip-previous"></b-icon>
        </a>
        <span v-else class="disabled">
          <b-icon icon="skip-previous"></b-icon>
        </span>

        <a v-if="playing" @click="togglePlay">
          <b-icon icon="pause"></b-icon>
        </a>
        <b-loading
          v-else-if="loading"
          :is-full-page="false"
          :active.sync="loading"
          :can-cancel="false"
        ></b-loading>
        <a v-else @click="togglePlay">
          <b-icon icon="play"></b-icon>
        </a>
        <a v-if="hasNext" @click="next">
          <b-icon icon="skip-next"></b-icon>
        </a>
        <span v-else class="disabled">
          <b-icon icon="skip-next"></b-icon>
        </span>
      </div>
      <a v-if="trackList.length > 0" class="flex-item playlist" @click="showTracks=!showTracks">
        <b-icon icon="playlist-music-outline"></b-icon>
      </a>
    </div>
    <div v-if="showTracks" class="playlist-tracklist">
      <PlayerTrackList
        v-for="(track, i) in trackList"
        :key="`${i}-${track.id}`"
        :track="track"
        :active="i===current"
        :position="i"
      ></PlayerTrackList>
    </div>
  </div>
</template>

<script>
import PlayerTrackList from "./PlayerTrackList";
import { mapState, mapGetters } from "vuex";
import { Howl } from "howler";
import axios from "axios";

export default {
  name: "Player",
  components: {
    PlayerTrackList
  },
  methods: {
    next: function() {
      this.$store.commit("next");
    },
    previous: function() {
      this.$store.commit("previous");
    },
    togglePlay: function() {
      this.$store.commit("setPlaying", !this.playing);
    }
  },
  data() {
    return {
      showTracks: false,
      player: null,
      loading: false
    };
  },
  watch: {
    currentTrack: function() {
      this.loading = true;
      if (this.player) {
        this.player.pause();
        this.player = null;
      }
      var _this = this;
      axios
        .get(this.apiURL() + "/track/" + this.currentTrack.id + "/stream")
        .then(response => {
          this.player = new Howl({
            src: response.data.http_mp3_128_url,
            html5: true,
            volume: 1.0,
            onend: function() {
              _this.$store.commit("next");
            }
          });
          this.player.play();
        })
        .catch(() => {
          this.$toast.open({
            duration: 5000,
            message: `Unable to retrieve stream`,
            position: "is-bottom",
            type: "is-danger"
          });
        })
        .finally(() => (this.loading = false));
    },
    playing: function() {
      if (this.player) {
        if (this.playing) {
          this.player.play();
        } else {
          this.player.pause();
        }
      }
    }
  },
  computed: {
    ...mapState(["trackList", "playing", "current", "currentTrack"]),
    ...mapGetters(["hasNext", "hasPrev"]),
    currentArtwork: function() {
      if (this.currentTrack.artwork_url) {
        return this.currentTrack.artwork_url;
      } else {
        return this.currentTrack.user.avatar_url;
      }
    },
    duration_str: function() {
      var seconds = Math.floor((this.currentTrack.duration / 1000) % 60),
        minutes = Math.floor((this.currentTrack.duration / (1000 * 60)) % 60),
        hours = Math.floor(
          (this.currentTrack.duration / (1000 * 60 * 60)) % 24
        );

      hours = hours < 10 ? "0" + hours : hours;
      minutes = minutes < 10 ? "0" + minutes : minutes;
      seconds = seconds < 10 ? "0" + seconds : seconds;

      return hours + ":" + minutes + ":" + seconds;
    }
  }
};
</script>

<style scoped>
a {
  color: white;
}

.disabled {
  color: lightgrey;
}

.flex-item {
  display: flex;
}

.flex-container {
  display: -ms-flexbox;
  display: -webkit-flex;
  display: flex;
  -webkit-flex-direction: row;
  -ms-flex-direction: row;
  flex-direction: row;
  -webkit-flex-wrap: nowrap;
  -ms-flex-wrap: nowrap;
  flex-wrap: nowrap;
  -webkit-justify-content: flex-start;
  -ms-flex-pack: start;
  justify-content: flex-start;
  -webkit-align-content: stretch;
  -ms-flex-line-pack: stretch;
  align-content: stretch;
  -webkit-align-items: center;
  -ms-flex-align: center;
  align-items: center;
  color: white;
  position: fixed;
  bottom: 0;
  height: 58px;
  width: 100%;
  padding: 0px 20px;
  background-color: rgb(22, 125, 240, 0.9);
}

.cover {
  -webkit-order: 0;
  -ms-flex-order: 0;
  order: 0;
  -webkit-flex: 0 1 auto;
  -ms-flex: 0 1 auto;
  flex: 0 1 auto;
  -webkit-align-self: auto;
  -ms-flex-item-align: auto;
  align-self: auto;
  max-height: 100%;
}

.cover img {
  border-radius: 4px;
  max-height: 100%;
  float: left;
}

.cover .metadata {
  float: right;
}

.controls {
  -webkit-order: 0;
  -ms-flex-order: 0;
  order: 0;
  -webkit-flex: 1 1 auto;
  -ms-flex: 1 1 auto;
  flex: 1 1 auto;
  -webkit-align-self: auto;
  -ms-flex-item-align: auto;
  align-self: auto;
  align-items: center;
  justify-content: right;
}

.playlist {
  -webkit-order: 0;
  -ms-flex-order: 0;
  order: 0;
  -webkit-flex: 0 1 auto;
  -ms-flex: 0 1 auto;
  flex: 0 1 auto;
  -webkit-align-self: auto;
  -ms-flex-item-align: auto;
  align-self: auto;
}

@media screen and (max-width: 768px) {
  .playlist-tracklist {
    width: 100% !important;
  }
}

.playlist-tracklist {
  color: white;
  position: fixed;
  float: right;
  bottom: 58px;
  right: 0px;
  height: 50%;
  width: 40%;
  background-color: rgb(22, 125, 240, 0.9);
  border-top-left-radius: 6px;
  border-top-right-radius: 6px;
  overflow-y: auto;
}
</style>
