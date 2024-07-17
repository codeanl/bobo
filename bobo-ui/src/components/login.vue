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
            <n-tabs class="card-tabs" default-value="signin" size="large" animated pane-wrapper-style="margin: 0 -4px"
                pane-style="padding-left: 4px; padding-right: 4px; box-sizing: border-box;">
                <n-tab-pane name="signin" tab="登录">
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
                <n-tab-pane name="signup" tab="注册">
                    <n-form>
                        <n-form-item-row label="用户名">
                            <n-input />
                        </n-form-item-row>
                        <n-form-item-row label="密码">
                            <n-input />
                        </n-form-item-row>
                        <n-form-item-row label="重复密码">
                            <n-input />
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
import { ref, onMounted, reactive } from "vue";
import { Component, h } from "vue";
function renderIcon(icon: Component) {
    return () => h(NIcon, null, { default: () => h(icon) });
}

const showModal = ref(false);
import { useMessage, useDialog } from "naive-ui";
const message = useMessage();
const dialog = useDialog();

const toLogin = () => {
    showModal.value = true;
};
const toRegister = () => {
    showModal.value = true;
};
</script>

<style scoped>
.btnbtn {
    margin: 0 5px;
}
</style>