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
                },
            ]
        },
        
    ],
})
export default router