import { createRouter, createWebHistory } from 'vue-router'
import Main from '@/views/Main'
import LoginPage from '@/views/LoginPage'
import AddNewsForm from '@/views/AddNewsForm'
import News from '@/assets/News'

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
  {
    path: '/addnews',
    name: 'addnews',
    component: AddNewsForm
  },
  {
    path: '/news',
    name: 'news',
    component: News
  },
  
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
