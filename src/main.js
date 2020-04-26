import Vue from 'vue'
import App from './App.vue'
import '@/assets/main.scss'
import router from './router'
import store from './store'
import VModal from 'vue-js-modal'
import Toasted from 'vue-toasted';

import { VueSpinners } from '@saeris/vue-spinners'

import VTooltip from 'v-tooltip'

import VueFuse from 'vue-fuse'

import VueMousetrap from 'vue-mousetrap'


Vue.config.productionTip = false
Vue.use(VModal)
Vue.use(VTooltip)
Vue.use(VueFuse)
Vue.use(Toasted, {
  position: "top-center",
  duration: 3000,
  fullWidth: true,
})
Vue.use(VueMousetrap)

Vue.use(VueSpinners)


new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')








