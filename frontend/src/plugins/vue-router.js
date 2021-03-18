import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '',
    name: 'ray-tracing',
    props: false,
    component: () => import('@/views/RayTracingView'),
    meta: {
      title: 'Distributed Ray Tracing'
    }
  }
]

const router = new VueRouter({
  mode: 'history',
  routes: routes
})

export default router
