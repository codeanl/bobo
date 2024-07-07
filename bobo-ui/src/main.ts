import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
const app = createApp(App)

//引入路由 
import router from './router'
app.use(router)

//引入pinia
import pinia from './store'
app.use(pinia)

app.mount('#app')
