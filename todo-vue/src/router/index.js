import { createRouter, createWebHistory } from 'vue-router'
import AddPage from '../pages/AddPage.vue'
import HomePage from '../pages/HomePage.vue'
import UpdatePage from '../pages/UpdatePage.vue'
const routes = [
    {
        path: '/new',
        name: 'AddPage',
        component: AddPage
    },
    {
        path: '/update/:id',
        name: 'UpdatePage',
        component: UpdatePage
    },
    {
        path: '/',
        name: 'HomePage',
        component: HomePage
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router