import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/prefecture',
      name: 'prefecture',
      component: () => import('../views/PrefectureView.vue')
    },
    {
      path: '/signup',
      name: 'signup',
      component: () => import('../views/SignUpView.vue')
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue')
    },
    {
      path: '/logout',
      name: 'logout',
      component: () => import('../views/LogoutView.vue')
    },
    {
      path: '/trip-plan',
      name: 'trip-plan',
      component: () => import('../views/TripPlanView.vue')
    },
    {
      path: '/trip-plan/new',
      name: 'register-trip-plan',
      component: () => import('../views/RegisterNew.vue')
    },
  ]
})

export default router
