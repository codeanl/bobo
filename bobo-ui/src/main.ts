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

// 导入视频播放组件
import VueVideoPlayer from '@videojs-player/vue'
import 'video.js/dist/video-js.css'
app.use(VueVideoPlayer)

//无限滚动
import infiniteScroll from 'vue3-infinite-scroll-better'
// import infiniteScroll from 'vue3-infinite-scroll-better'
app.use(infiniteScroll)


// import useUserStore from './store/user'
// let userStore = useUserStore(pinia)
// router.beforeEach(async (to:any, from:any, next:any) => {
//     let token = userStore.token
//     let username = userStore.username
//     if (token) {
//             if (username) {
//                 next()
//             } else {
//                 console.log("1111377777");
//                 try {
//                     await userStore.userInfo()
//                     next({ ...to, replace: true })
//                 } catch (error) {
//                     console.error(error)
//                     next({ path: '/' })
//                 }
//             }
//         } else {
//             next();
//             console.log("nono");
//             next({ path: '/home' })
//             try {
//                 await userStore.userLogout()
//             } catch (error) {
//                 console.error(error)
//             }
//         }
// })

app.mount('#app')
