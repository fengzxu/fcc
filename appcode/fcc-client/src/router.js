import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Netcon',
      component: () => import('@/views/Netcon.vue')
    },
    {
      path: '/estatebook',
      name: 'Estatebook',
      component: () => import('@/views/EstateBook.vue')
    },
    {
      path: '/bigdata',
      name: 'Bigdata',
      component: () => import('@/views/Bigdata.vue')
    },
    {
      path: '/estatetax',
      name: 'EstateTax',
      component: () => import('@/views/EstateTax.vue')
    },
  ]
})