<template>
  <article class="media">
    <figure class="media-left">
      <p class="image is-64x64">
        <a :href="track.permalink_url"><img :src="artwork" alt="Track Artwork" /></a>
      </p>
    </figure>
    <div class="media-content">
      <div class="content">
        <p>
          <a class="is-title is-2" :href="track.permalink_url">{{ track.title }}</a>
          <a class="is-pulled-right" :href="track.user.permalink_url">
            <small><b-icon icon="account" size="is-small"></b-icon>
            {{ track.user.username }}</small>
          </a>
          <br />
          <small>
            <b-icon icon="clock-outline" size="is-small"></b-icon>
            {{ duration_str }}
          </small>
        </p>
      </div>
    </div>
  </article>
</template>

<script>
export default {
  props: {
    track: Object
  },
  computed: {
    artwork: function() {
      if(this.track.artwork_url){
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
.image > img {
  border-radius: 5px;
}

article {
  padding: 5px;
  border-radius: 5px;
  box-shadow: 0 3px 6px rgba(0, 0, 0, 0.16), 0 3px 6px rgba(0, 0, 0, 0.23);
}

.avatar {
  border-radius: 10%;
  display: flex;
  align-items: center;
}
</style>

