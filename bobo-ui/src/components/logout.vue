<template>
    <div style="display: flex;width: 170px;;">
        <div style="display: flex; ">
            <n-avatar round size="medium" :src="userStore.avatar" />
        </div>
        <div style="margin-left: 10px;">
            <p style="font-size: 14px;line-height: 14px;margin: 0;padding: 0;">
                {{ userStore.nickname }}
            </p>
            <p style="font-size: 13px;line-height: 13px;margin: 0;padding: 0;margin-top: 5px;">@{{ userStore.username }}
            </p>
        </div>
        <div @click="logout" style="margin-left: 10px;">
            <n-icon>
                <MdExit />
            </n-icon>
        </div>
    </div>
</template>

<script setup lang="ts">
import { MdExit } from "@vicons/ionicons4";
import { NIcon } from "naive-ui";
import { ref, onMounted, reactive } from "vue";
import { Component, h } from "vue";
function renderIcon(icon: Component) {
    return () => h(NIcon, null, { default: () => h(icon) });
}
import useUserStore from "@/store/user";
let userStore = useUserStore();

import { useRouter } from "vue-router";
let router = useRouter();

const logout = () => {
    userStore.userLogout()
    router.push("/");
}

</script>

<style lang="scss"></style>