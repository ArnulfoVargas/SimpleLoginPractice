import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import SignInViewVue from '@/views/SignInView.vue'
import LoginViewVue from '@/views/LoginView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/login',
      name: "login",
      component: LoginViewVue
    }, 
    {
      path: '/signin',
      name: 'sign in',
      component: SignInViewVue
    }
  ]
})

export default router
