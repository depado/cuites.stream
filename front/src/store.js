import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    trackList: [],
    currentTrack: null,
    playing: false,
    current: 0,
  },
  getters: {
    hasNext: state => {
      return state.current < state.trackList.length-1
    },
    hasPrev: state => {
      return state.current > 0 && state.current < state.trackList.length;
    }
  },
  mutations: {
    next(state) {
      if(state.current < state.trackList.length) {
        state.current++;
        state.currentTrack =state.trackList[state.current];
      }
    },
    previous(state) {
      if (state.current > 0) {
        state.current--;
        state.currentTrack =state.trackList[state.current];
      } 
    },
    pushTrack(state, track) {
      state.trackList.push(track);
    },
    cleanTracks(state) {
      state.trackList.length = 0;
      state.current = 0;
    },
    reset(state) {
      state.current = 0;
      state.currentTrack = state.trackList[state.current];
    },
    setPlaying(state, playing) {
      if(!state.currentTrack && playing) {
        state.currentTrack =state.trackList[state.current];
      }
      state.playing = playing;
    },
    changeTrack(state, position) {
      state.current = position;
      state.currentTrack = state.trackList[state.current];
    }
  },
  actions: {

  }
})