<template>
  <div class="card">
    <div class="card-image">
      <figure class="image is-4by3">
        <img :src="playlist.artwork_url" alt="Playlist Artwork" />
      </figure>
    </div>
    <div class="card-content">
      <p class="title is-3">{{ playlist.title }}</p>
      <div class="media">
        <div class="media-left">
          <figure class="image is-48x48">
            <img :src="playlist.user.avatar_url" alt="Placeholder image" />
          </figure>
        </div>
        <div class="media-content">
          <a class="title is-4" :href="playlist.user.permalink_url">{{ playlist.user.username }}</a>
        </div>
      </div>

      <div class="content">
        <p>{{ playlist.description }}</p>
        <a href="#" v-for="tag in tags" :key="tag"> #{{ tag }} </a>
        <br />
        <time datetime="2016-1-1">11:09 PM - 1 Jan 2016</time>
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
    tags: function() {
        if('tag_list' in this.playlist) {
            return this.playlist.tag_list.split(" ");
        } else {
            return null
        }
    }
  }
};
</script>

<style>
</style>
