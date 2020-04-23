import Vue from 'vue'
import App from './App.vue'
import '@/assets/main.scss'
import router from './router'
import store from './store'
import VModal from 'vue-js-modal'
import Toasted from 'vue-toasted';
// @ts-ignore
import { VueSpinners } from '@saeris/vue-spinners'
// @ts-ignore
import VueLocalforage from 'v-localforage'
// @ts-ignore
import VTooltip from 'v-tooltip'
// @ts-ignore
import VueFuse from 'vue-fuse'
// @ts-ignore
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
// you can also pass options, check options reference below
Vue.use(VueLocalforage, {
    instances: [
        {
            storeName: 'projects'
        },
        {
            storeName: 'config'
        }
    ]
})



new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')








