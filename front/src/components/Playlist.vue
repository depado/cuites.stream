<template>
  <div class="box">
    <div class="columns is-mobile">
      <div class="column is-one-quarter rm-leftpad">
        <div class="overlay-image">
          <img class="image" :src="playlist.artwork_url" alt="Alt text" />
          <div class="hover">
            <div class="text">
              <a href="#" @click="addPlaylist">
                <b-icon icon="play-circle-outline" size="is-large"></b-icon>
              </a>
            </div>
          </div>
        </div>
      </div>
      <div class="column rm-rightpad">
        <p>
          <a :href="playlist.user.permalink_url" class="is-pulled-right">
            <img class="image is-48x48" :src="playlist.user.avatar_url" />
          </a>
          <strong>
            <a :href="playlist.permalink_url">{{ playlist.title }}</a>
          </strong>
          <br />
          <small>
            <b-icon icon="clock-outline" size="is-small"></b-icon>
            {{ duration_str }}
            <br />
            <b-icon icon="account" size="is-small"></b-icon>
            {{ playlist.user.username }}
            <br />
            <b-icon icon="calendar" size="is-small"></b-icon>
            {{ just_date }}
          </small>
          <br />
          <a @click="fetch_tracks">
            <b-icon icon="library-music" size="is-small"></b-icon>Tracklist
          </a>
        </p>
        <b-notification :closable="false" v-if="loading">
          <b-loading :is-full-page="true" :active.sync="loading" :can-cancel="false"></b-loading>
        </b-notification>
        <div v-if="modal_active" class="modal" :class="{'is-active': modal_active}">
          <div class="modal-background"></div>
          <div class="modal-content">
            <div class="container breathe">
              <div class="columns is-multiline">
                <div class="column is-half" v-for="(track, i) in tracks" :key="`${i}-${track.id}`">
                  <SoundcloudTrack :track="track" :small="true"></SoundcloudTrack>
                </div>
              </div>
            </div>
          </div>
          <button class="modal-close is-large" aria-label="close" @click="close_modal"></button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import SoundcloudTrack from "./SoundcloudTrack";

export default {
  props: {
    playlist: Object
  },
  components: {
    SoundcloudTrack
  },
  data() {
    return {
      loading: false,
      tracks: null,
      modal_active: false
    };
  },
  methods: {
    addPlaylist: function() {
      if (!this.tracks) {
        axios
          .get(this.apiURL() + "/playlist/" + this.playlist.id + "/tracks")
          .then(response => {
            this.tracks = response.data;
            this.$store.commit("cleanTracks");
            this.tracks.forEach(t => {
              this.$store.commit("pushTrack", t);
            });
            this.$store.commit("setPlaying", true);
          })
          .catch(() => {
            this.$toast.open({
              duration: 5000,
              message: `Unable to retrieve tracks`,
              position: "is-bottom",
              type: "is-danger"
            });
          });
      } else {
        this.$store.commit("cleanTracks");
        this.tracks.forEach(t => {
          this.$store.commit("pushTrack", t);
        });
        this.$store.commit("setPlaying", true);
      }
    },
    close_modal: function() {
      this.modal_active = false;
      document.querySelector("html").classList.remove("is-clipped");
    },
    fetch_tracks: function() {
      if (!this.tracks) {
        this.loading = true;
        document.querySelector("html").classList.add("is-clipped");
        axios
          .get(this.apiURL() + "/playlist/" + this.playlist.id + "/tracks")
          .then(response => {
            this.modal_active = true;
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
          .finally(() => {
            this.loading = false;
          });
      } else {
        document.querySelector("html").classList.add("is-clipped");
        this.modal_active = true;
      }
    }
  },
  computed: {
    duration_str: function() {
      var milliseconds = parseInt((this.playlist.duration % 1000) / 100),
        seconds = Math.floor((this.playlist.duration / 1000) % 60),
        minutes = Math.floor((this.playlist.duration / (1000 * 60)) % 60),
        hours = Math.floor((this.playlist.duration / (1000 * 60 * 60)) % 24);

      hours = hours < 10 ? "0" + hours : hours;
      minutes = minutes < 10 ? "0" + minutes : minutes;
      seconds = seconds < 10 ? "0" + seconds : seconds;

      return hours + ":" + minutes + ":" + seconds + "." + milliseconds;
    },
    just_date: function() {
      return this.playlist.created_at.split(" ", 1)[0];
    }
  }
};
</script>

<style scoped>
a {
  color: #167df0;
}
.hover > .text > a {
  color: white;
}
img {
  border-radius: 4px;
}
small {
  color: grey;
}
.rm-leftpad {
  padding-left: 0;
}
.rm-rightpad {
  padding-right: 0;
}
.box {
  padding: 0.3rem;
}
@media screen and (min-width: 768px) and (max-width: 1024px) {
  .breathe {
    margin-right: 50px;
  }
}
.overlay-image {
  position: relative;
  width: 100%;
}
.overlay-image .image {
  display: block;
  width: 100%;
  height: auto;
}
.overlay-image .hover {
  border-radius: 4px;
  width: 100%;
  background-color: rgba(0, 0, 0, 0.3);
  position: absolute;
  text-align: center;
  height: 100%;
  opacity: 0;
  color: white;
  transition: 0.5s ease;
  top: 0;
}
.overlay-image .text {
  color: #fff;
  font-size: 30px;
  line-height: 1.5em;
  text-shadow: 2px 2px 2px #000;
  text-align: center;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 100%;
}
.overlay-image:hover .hover {
  opacity: 1;
}
</style>
