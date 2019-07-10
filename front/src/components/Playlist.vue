<template>
  <div class="box">
    <div class="columns is-mobile">
      <div class="column is-one-quarter">
        <img :src="playlist.artwork_url" class="image" alt="Playlist Artwork" />
      </div>
      <div class="column">
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
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    playlist: Object
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
img {
  border-radius: 4px;
}
small {
  color: grey;
}
.box {
  padding: 0.5rem;
}
</style>
