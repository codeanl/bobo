import { createRouter, createWebHashHistory } from 'vue-router'
let router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: '/',
            component: () => import('@/views/layout/index.vue'),
            name: 'layout',
            redirect: '/home',  
            children:[
                {
                    path: '/home',
                    component: () => import('@/views/home.vue'),
                    meta: {
                        keepAlive: true,
                    },
                },
                {
                    path: '/setting',
                    component: () => import('@/views/setting.vue'),
                    meta: {
                        keepAlive: false,
                    },
                },
                {
                    path: '/daily',
                    component: () => import('@/views/daily.vue'),
                    meta: {
                        keepAlive: false,
                    },
                },
            ]
        },
        
    ],
})
export default router