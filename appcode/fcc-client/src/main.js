import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import router from './router'
import axios from './axios'
import querystring from 'querystring'
import 'roboto-fontface/css/roboto/roboto-fontface.css'
import '@fortawesome/fontawesome-free/css/all.css'

Vue.config.productionTip = false
//set axios form-data
Vue.prototype.$qs = querystring;
Vue.prototype.$axios = axios

new Vue({
  vuetify,
  router,
  axios,
  render: h => h(App)
}).$mount('#app')
