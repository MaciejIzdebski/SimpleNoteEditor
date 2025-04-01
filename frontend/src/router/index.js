import { createRouter, createWebHistory } from 'vue-router'
import NoteView from '../views/NoteView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: NoteView,
    }
  ],
})

export default router
