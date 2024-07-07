import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate' 
let pinia = createPinia()
pinia.use(piniaPluginPersistedstate) 
export default pinia