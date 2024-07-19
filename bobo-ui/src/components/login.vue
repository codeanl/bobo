<template>
    <div
        style="height: 50px;width: 100%;text-align: center;display: flex;align-items: center;justify-content: center; ">
        <n-button strong secondary round type="primary" class="btnbtn" @click="toLogin">
            登录
        </n-button>
        <n-button strong secondary round type="info" class="btnbtn" @click="toRegister">
            注册
        </n-button>
    </div>

    <!--  -->
    <n-modal v-model:show="showModal">
        <n-card style="width: 350px" :bordered="false" size="huge" aria-modal="true">
            <n-tabs class="card-tabs" :default-value="tabName" size="large" animated pane-wrapper-style="margin: 0 -4px"
                pane-style="padding-left: 4px; padding-right: 4px; box-sizing: border-box;">
                <!--  -->
                <n-tab-pane name="login" tab="登录">
                    <n-form>
                        <n-form :model:value="loginForm">
                            <n-form-item-row label="用户名">
                                <n-input v-model:value="loginForm.username" />
                            </n-form-item-row>
                            <n-form-item-row label="密码">
                                <n-input v-model:value="loginForm.password" />
                            </n-form-item-row>
                        </n-form>
                    </n-form>
                    <n-button type="primary" block secondary strong @click="login">
                        登录
                    </n-button>
                </n-tab-pane>
                <!--  -->
                <n-tab-pane name="register" tab="注册">
                    <n-form :model:value="registerForm">
                        <n-form-item-row label="用户名">
                            <n-input v-model:value="registerForm.username" placeholder="请输入用户名" />
                        </n-form-item-row>
                        <n-form-item-row label="邮箱">
                            <n-input v-model:value="registerForm.email" placeholder="请输入邮箱" />
                            <n-button strong secondary type="success" @click="sendCode">
                                获取
                            </n-button>
                        </n-form-item-row>
                        <n-form-item-row label="验证码">
                            <n-input v-model:value="registerForm.code" placeholder="请输入验证码" />
                        </n-form-item-row>
                        <n-form-item-row label="密码">
                            <n-input type="password" show-password-on="mousedown" placeholder="请输入密码" :maxlength="16"
                                v-model:value="registerForm.password" />
                        </n-form-item-row>
                    </n-form>
                    <n-button type="primary" block secondary strong @click=register>
                        注册
                    </n-button>
                </n-tab-pane>
            </n-tabs>
        </n-card>
    </n-modal>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";
import { Component, h } from "vue";
import { NIcon } from "naive-ui";
function renderIcon(icon: Component) {
    return () => h(NIcon, null, { default: () => h(icon) });
}
import { useMessage } from "naive-ui";
const message = useMessage();
import { SendCode, Register, Login, Profile } from '@/api/user'
import { log } from "console";

// 注册表单
let registerForm = reactive({
    username: "",
    password: "",
    email: "",
    code: "",
});
//表单数据
let loginForm = reactive({
    username: "codeanl",
    password: "123456",
});

const tabName = ref("login")
const showModal = ref(false);
//点击登录按钮
const toLogin = () => {
    showModal.value = true;
    tabName.value = "login";
};
//点击注册按钮
const toRegister = () => {
    registerForm = reactive({
        username: "",
        password: "",
        email: "",
        code: "",
    });
    showModal.value = true;
    tabName.value = "register";
};

//发送验证码
const sendCode = async () => {
    let res: any = await SendCode({
        params: { email: registerForm.email }
    })
    if (res.code == 200) {
        message.success('邮件发送成功')
    } else {
        message.error(res.message)
    }
}
//注册
const register = async () => {
    let res = await Register(registerForm)
    if (res.code == 200) {
        showModal.value = false;
        message.success('注册成功')
    } else {
        message.error(res.message)
    }
}

import useUserStore from "@/store/user";
let userStore = useUserStore();

const login = async () => {
    tabName.value = "login";
    // userStore.token = ''
    // console.log(userStore.token);
    let res = await Login(loginForm)
    if (res.code == 200) {
        userStore.token = res.data
        let profileReq = await Profile()
        if (profileReq.code == 200) {
            userStore.created_at = profileReq.data.created_at
            userStore.email = profileReq.data.email
            userStore.id = profileReq.data.id
            userStore.ip_address = profileReq.data.ip_address
            userStore.ip_source = profileReq.data.ip_source
            userStore.last_login_time = profileReq.data.last_login_time
            userStore.nickname = profileReq.data.nickname
            userStore.role = profileReq.data.role
            userStore.status = profileReq.data.status
            userStore.username = profileReq.data.username
            userStore.avatar = profileReq.data.avatar
            showModal.value = false;
            message.success('登陆成功')
            console.log(userStore.token);
        } else {
            message.error(profileReq.message)
        }
    } else {
        message.error(res.message)
    }
}


</script>

<style scoped>
.btnbtn {
    margin: 0 5px;
}
</style>