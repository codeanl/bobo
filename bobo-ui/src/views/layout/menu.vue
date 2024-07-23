<template>
    <div style="height: 100vh;width: 230px;position: relative;">
        <!--  -->
        <div style="padding: 10px 0 0 20px;">
            <n-avatar size="large"
                src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcREqDKQqOGiC-QLyaWSxqpAhe3FUdsYD27e1g&s" />
        </div>
        <!--  -->
        <div style="margin: 5px 0 0 0;">
            <n-menu v-model:value="activeKey" :options="MenuOptions" :on-update:value="changeMenu" />
        </div>
        <!--  -->
        <div style="position: absolute;left: 0;bottom: 20px;width: 230px;display: grid;place-items: center;"
            v-if="userStore.token != ''">
            <Logout></Logout>
        </div>
        <div style="position: absolute;left: 0;bottom: 20px;width: 230px;display: grid;place-items: center;"
            v-if="userStore.token == ''">
            <Login></Login>
        </div>
    </div>
</template>

<script setup lang="ts">
import Login from '@/components/login.vue'
import Logout from '@/components/logout.vue'

import useUserStore from "@/store/user";
let userStore = useUserStore();

import type { Component } from 'vue'
import { h, ref, computed } from 'vue'
import { NIcon } from 'naive-ui'
import type { MenuOption } from 'naive-ui'

import {
    BookOutline as BookIcon,
    PersonOutline as PersonIcon,
    WineOutline as WineIcon
} from '@vicons/ionicons5'

const activeKey = ref('home')
function renderIcon(icon: Component) {
    return () => h(NIcon, null, { default: () => h(icon) })
}
import { useRouter } from 'vue-router'
const router = useRouter()
const changeMenu = (key: string, item: MenuOption) => {
    console.log(key);
    activeKey.value = key
    router.push(key)
}
const MenuOptions = computed(() => [
    {
        label: '广场',
        key: 'home',
        icon: renderIcon(BookIcon)
    },
    {
        label: '设置',
        key: 'setting',
        icon: renderIcon(BookIcon),
        show: userStore.token != ''
    },
])
</script>

<style lang="scss"></style>
