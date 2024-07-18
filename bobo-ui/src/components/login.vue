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
                        <n-form-item-row label="用户名">
                            <n-input />
                        </n-form-item-row>
                        <n-form-item-row label="密码">
                            <n-input />
                        </n-form-item-row>
                    </n-form>
                    <n-button type="primary" block secondary strong>
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
                    <n-button type="primary" block secondary strong>
                        注册
                    </n-button>
                </n-tab-pane>
            </n-tabs>
        </n-card>
    </n-modal>
</template>

<script setup lang="ts">
import { NIcon } from "naive-ui";
import { ref, reactive } from "vue";
import { Component, h } from "vue";
function renderIcon(icon: Component) {
    return () => h(NIcon, null, { default: () => h(icon) });
}
const tabName = ref("login")
const showModal = ref(false);
import { useMessage, useDialog } from "naive-ui";
import { SendCode } from '@/api/user'
const message = useMessage();
const dialog = useDialog();


let registerForm = reactive({
    username: "",
    password: "",
    email: "",
    code: "",
});
const toLogin = () => {
    showModal.value = true;
    tabName.value = "login";
};
const toRegister = () => {
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
</script>

<style scoped>
.btnbtn {
    margin: 0 5px;
}
</style>