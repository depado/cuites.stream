<template>
  <article class="media">
    <figure class="media-left">
      <p class="image is-128x128">
        <img :src="playlist.artwork_url" alt="Playlist Artwork" />
      </p>
    </figure>
    <div class="media-content">
      <div class="content columns">
        <div class="column">
          <p>
            <a class="is-title is-3" :href="playlist.permalink_url">{{ playlist.title }}</a>
            <br />
            <small>
              <b-icon icon="clock-outline" size="is-small"></b-icon>
              {{ duration_str }}
            </small>
          </p>
          <a :href="playlist.user.permalink_url" class="inline">
            <img class="image is-48x48 avatar" :src="playlist.user.avatar_url" />
          </a>
        </div>
        <p class="column">
          <b-taglist>
            <b-tag v-for="tag in tags" :key="tag" type="is-info">#{{ tag }}</b-tag>
          </b-taglist>
        </p>
      </div>
    </div>
  </article>
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
    tags: function() {
      if ("tag_list" in this.playlist) {
        return this.playlist.tag_list.split(" ");
      } else {
        return null;
      }
    }
  }
};
</script>

<style scoped>
.image > img {
  border-radius: 5px;
}

article {
  background-color: white;
  padding: 5px;
  border-radius: 5px;
  box-shadow: 0 3px 6px rgba(0, 0, 0, 0.16), 0 3px 6px rgba(0, 0, 0, 0.23);
}

.avatar {
  border-radius: 10%;
  display: flex;
  align-items: center;
}

.inline {
  display: inline-flex;
}
</style>
