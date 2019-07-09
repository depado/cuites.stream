import Vue from 'vue'
import App from './App.vue'
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
import router from './router'
// import VueFuse from 'vue-fuse'

Vue.config.productionTip = false
Vue.use(Buefy)
// Vue.use(VueFuse)

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
