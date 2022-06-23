import { createRouter, createWebHistory } from 'vue-router'
import Main from '@/views/Main'
import LoginPage from '@/views/LoginPage'
// import news from '@/public/news'

const routes = [
  {
    path: '/',
    name: 'main',
    component: Main
  },
  {
    path: '/login',
    name: 'login',
    component: LoginPage
  },
  // {
  //   path: '/news',
  //   name: 'news',
  //   component: news
  // },
  
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
