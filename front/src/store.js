import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    trackList: [],
    playing: false,
    current: 0,
  },
  getters: {
    currentTrack: state => {
      return state.trackList[state.current];
    },
    hasNext: state => {
      return state.current < state.trackList.length-1
    },
    hasPrev: state => {
      return state.current > 0 && state.current > state.trackList.length;
    }
  },
  mutations: {
    next(state) {
      state.current++;
    },
    previous(state) {
      state.current--;
    },
    pushTrack(state, track) {
      state.trackList.push(track);
    },
    cleanTracks(state) {
      state.trackList.length = 0;
      state.current = 0;
    },
    setPlaying(state, playing) {
      state.playing = playing;
    }
  },
  actions: {

  }
})