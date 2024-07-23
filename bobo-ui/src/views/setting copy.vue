<template>
    <n-list hoverable clickable class="n-list">
        <n-list-item>
            <p class="listTitle">基本信息</p>
            <div class="profileInfo">
                <div>
                    <n-avatar style=" width: 80px;height: 80px;" size="medium" :src="userStore.avatar" />
                </div>
                <div class="profileInfoName">
                    <div style="margin-right: 5px">
                        <p style="font-size: 13px">昵称: {{ userStore.nickname }}</p>
                        <p style="font-size: 13px">用户名: @{{ userStore.username }}</p>
                    </div>
                    <!-- <n-button quaternary round type="success">修改</n-button> -->
                </div>
            </div>
        </n-list-item>
        <n-list-item>
            <p class="listTitle">邮箱号</p>
            <n-alert title="邮箱绑定提示" type="warning" v-if="userStore.email == ''">
                成功绑定邮箱后，才能进行换头像、发动态、回复等交互~
                <n-button quaternary round type="success"> 立即绑定 </n-button>
            </n-alert>
            <div class="youPhone" v-if="userStore.email != ''">
                <p>{{ userStore.email }}</p>
                <n-button quaternary round type="success"> 更换 </n-button>
            </div>
        </n-list-item>
        <n-list-item>
            <p class="listTitle">账号安全</p>
            <div class="youPhone">
                <p>您已设置密码</p>
                <n-button quaternary round type="success"> 重置 </n-button>
            </div>
            <div>
                <n-form ref="formRef" :label-width="80" :model="formValue" :rules="rules" size="small">
                    <n-form-item label="旧密码" path="old_password">
                        <n-input v-model:value="formValue.old_password" placeholder="旧密码" type="password" />
                    </n-form-item>
                    <n-form-item label="新密码" path="new_password">
                        <n-input v-model:value="formValue.new_password" placeholder="新密码" type="password" />
                    </n-form-item>
                    <n-form-item label="重复密码" path="second_new_password">
                        <n-input v-model:value="formValue.second_new_password" placeholder="重复密码" type="password" />
                    </n-form-item>
                </n-form>
                <n-row :gutter="[0, 24]">
                    <n-col :span="23">
                        <div style="display: flex; justify-content: flex-end">
                            <n-button round size="small">
                                取消
                            </n-button>
                            <n-button round type="primary" size="small" style="margin-left: 10px;"
                                @click="handleValidateClick">
                                验证
                            </n-button>
                        </div>
                    </n-col>
                </n-row>
            </div>
        </n-list-item>
    </n-list>
</template>

<script setup lang="ts">
import useUserStore from "@/store/user";
let userStore = useUserStore();
import { defineComponent, ref } from 'vue'
import type { FormInst } from 'naive-ui'
import { useMessage } from 'naive-ui'


const formRef = ref<FormInst | null>(null)
const message = useMessage()

const formValue = ref({
    old_password: '',
    new_password: '',
    second_new_password: ''
})
const rules = {
    old_password: {
        required: true,
        message: '请输入旧密码，长度为6-16位',
        trigger: 'blur',
        min: 6,
        max: 16
    },
    new_password: {
        required: true,
        message: '请输入新密码,长度为6-16位',
        trigger: ['input', 'blur'],
        min: 6,
        max: 16
    },
    second_new_password: {
        required: true,
        message: '请重复输入密码,长度为6-16位',
        trigger: ['blur', 'password-input'],
        min: 6,
        max: 16
    }
}
const handleValidateClick = (e: MouseEvent) => {
    e.preventDefault()
    formRef.value?.validate((errors) => {
        if (!errors) {
            message.success('Valid')
        }
        else {
            console.log(errors)
            message.error('Invalid')
        }
    })
}
</script>

<style scoped>
.listTitle {
    font-size: 16px;
    font-weight: bold;
    margin: 5px 0 15px 0;
}

.profileInfo {
    .profileInfoName {
        margin-top: 10px;
        display: flex;
    }
}

.youPhone {
    display: flex;
    align-items: center;
}
</style>