<template>
  <div class="box">
    <div class="columns is-mobile">
      <div class="column is-one-fifth">
        <img :src="artwork" class="image" alt="Track Artwork" />
      </div>
      <div class="column">
        <p class="metadata">
          <a :href="track.permalink_url">{{ track.title }}</a>
          <br />
          <small>
            <b-icon icon="clock-outline" size="is-small"></b-icon>
            {{ duration_str }}
            <br />
            <b-icon icon="account" size="is-small"></b-icon>
            {{ track.user.username }}
          </small>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    track: Object
  },
  computed: {
    artwork: function() {
      if (this.track.artwork_url) {
        return this.track.artwork_url;
      } else {
        return this.track.user.avatar_url;
      }
    },
    duration_str: function() {
      var seconds = Math.floor((this.track.duration / 1000) % 60),
        minutes = Math.floor((this.track.duration / (1000 * 60)) % 60),
        hours = Math.floor((this.track.duration / (1000 * 60 * 60)) % 24);

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

