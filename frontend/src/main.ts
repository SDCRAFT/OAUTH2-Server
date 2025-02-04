import './assets/main.css'
import 'element-plus/theme-chalk/dark/css-vars.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

import App from './App.vue'
import { routes } from 'vue-router/auto-routes'
import { createMemoryHistory, createRouter, createWebHistory } from 'vue-router'

const app = createApp(App)

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

app.use(pinia)
console.log(routes)

const router = createRouter({
  history: createWebHistory(),
  routes,
})
app.use(router)

app.mount('#app')


