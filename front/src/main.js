import Vue from 'vue'
import App from './App.vue'
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
import router from './router'
import store from './store'

Vue.config.productionTip = false
Vue.use(Buefy)
Vue.mixin({
  methods: {
    apiURL: function () {
      let url = process.env.VUE_APP_API_PROTOCOL + "://" + process.env.VUE_APP_API_HOST;
      if (process.env.VUE_APP_API_PORT) {
        url = url + ":" + process.env.VUE_APP_API_PORT;
      }
      return url
    }
  }
})

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')