<template>
    <n-list clickable class="n-list" style="padding: 0 15px;">
        <n-list-item>
            <p class="listTitle">基本信息</p>
            <div>
                <div>
                    <n-avatar style=" width: 80px;height: 80px;" size="medium" :src="userStore.avatar" />
                </div>
                <div>
                    <n-upload action="http://localhost:4321/api/upload" @finish="handleFinish" :show-file-list=false>
                        <n-button quaternary round type="success">更换</n-button>
                    </n-upload>
                </div>
                <div style="margin-top: 5px">
                    <div style="display: flex;align-items: center;">
                        <p style=" font-size: 13px;margin-right: 5px">昵称</p>
                        <span v-if="!upNicknameStatus"> {{ userStore.nickname }}</span>
                        <n-input v-else v-model:value="nicknameInput" type="text" style="width: 100px;" />
                        <n-button quaternary round type="success" @click="upNicnNameInput" style="margin-left: 5px;"
                            size="small" v-if="!upNicknameStatus">修改</n-button>
                        <n-button quaternary round type="success" @click="upProfile" style="margin-left: 5px;"
                            size="small" v-else>保存</n-button>
                    </div>
                    <div style="display: flex;align-items: center;">
                        <p style="font-size: 13px">用户名</p>
                        <span style="margin-left: 5px">@{{ userStore.username }}</span>
                    </div>
                </div>
            </div>
        </n-list-item>
        <n-list-item>
            <p class="listTitle">邮箱号</p>
            <div class="youPhone">
                <p>{{ userStore.email }}</p>
                <n-button quaternary round type="success" v-if="!upEmailStatus" @click="upEmailStatus = !upEmailStatus">
                    修改 </n-button>
            </div>
            <div v-if="upEmailStatus">
                <n-form ref="formRef" :label-width="80" :model="upEmailFormValue" size="small">
                    <n-form-item label="新邮箱">
                        <n-input v-model:value="upEmailFormValue.email" placeholder="新邮箱" />
                        <n-button strong secondary type="success" @click="sendCode">
                            发送验证码
                        </n-button>
                    </n-form-item>
                    <n-form-item label="验证码">
                        <n-input v-model:value="upEmailFormValue.code" placeholder="验证码" />
                    </n-form-item>
                </n-form>
                <!--  -->
                <n-row :gutter="[0, 24]">
                    <n-col :span="23">
                        <div style="display: flex; justify-content: flex-end">
                            <n-button round size="small" @click="resetEmailForm">
                                取消
                            </n-button>
                            <n-button round type="primary" size="small" style="margin-left: 10px;" @click="upEmail">
                                修改
                            </n-button>
                        </div>
                    </n-col>
                </n-row>
            </div>
        </n-list-item>
        <n-list-item>
            <p class="listTitle">账号安全</p>
            <div class="youPhone">
                <p>您已设置密码</p>
                <n-button quaternary round type="success" v-if="!upPassStatus" @click="upPassStatus = !upPassStatus"> 修改
                </n-button>
            </div>
            <!--  -->
            <div v-if="upPassStatus">
                <n-form ref="formRef" :label-width="80" :model="formValue" :rules="upPassRules" size="small">
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
                <!--  -->
                <n-row :gutter="[0, 24]">
                    <n-col :span="23">
                        <div style="display: flex; justify-content: flex-end">
                            <n-button round size="small" @click="resetPassForm">
                                取消
                            </n-button>
                            <n-button round type="primary" size="small" style="margin-left: 10px;" @click="upPass">
                                修改
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
import { UpdatePassword, SendCode, UpdateEmail, UpdateProfile } from '@/api/user'

const formRef = ref<FormInst | null>(null)
const message = useMessage()

const formValue = ref({
    old_password: '',
    new_password: '',
    second_new_password: ''
})
const upPassRules = {
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
const upEmailFormValue = ref({
    code: '',
    email: '',
})

const sendCode = async () => {
    let res: any = await SendCode({
        params: { email: userStore.email }
    })
    if (res.code == 200) {
        message.success('邮件发送成功')
    } else {
        message.error(res.message)
    }
}


const upPassStatus = ref(false)
const upEmailStatus = ref(false)
const upNicknameStatus = ref(false)

const resetPassForm = () => {
    formRef.value?.restoreValidation()
    formValue.value = {
        old_password: '',
        new_password: '',
        second_new_password: ''
    }
    upPassStatus.value = false
}

const resetEmailForm = () => {
    upEmailStatus.value = false
}

import { useRouter } from "vue-router";
let router = useRouter();
// 修改密码
const upPass = (e: MouseEvent) => {
    e.preventDefault()
    formRef.value?.validate(async (errors) => {
        if (!errors) {
            let res = await UpdatePassword(formValue.value)
            if (res.code == 200) {
                message.success("修改密码成功")
                userStore.userLogout()
                router.push("/");
            } else {
                message.error(res.message)
            }
        }
        else {
            message.error('表单校验失败！')
        }
    })
}
// 修改邮箱
const upEmail = async () => {
    let res = await UpdateEmail(upEmailFormValue.value)
    if (res.code == 200) {
        message.success("修改邮箱成功")
        upEmailStatus.value = false
        userStore.email = upEmailFormValue.value.email
    } else {
        message.error(res.message)
    }
}

const nicknameInput = ref('')
// 切换nickname输入框
const upNicnNameInput = () => {
    nicknameInput.value = userStore.nickname
    upNicknameStatus.value = true
}
// 修改个人信息
const upProfile = async () => {
    console.log(userStore.nickname);
    let res = await UpdateProfile({
        nickname: nicknameInput.value,
        avatar: userStore.avatar
    })
    if (res.code == 200) {
        userStore.nickname = nicknameInput.value
        message.success("修改成功")
    } else {
        message.success(res.data)
    }
    upNicknameStatus.value = false
}

// 上传头像
import type { UploadFileInfo } from 'naive-ui'
const handleFinish = async ({
    file,
    event
}: {
    file: UploadFileInfo
    event?: ProgressEvent
}) => {
    let res = JSON.parse((event?.target as XMLHttpRequest).response);
    if (res.code == 200) {
        let ress = await UpdateProfile({
            nickname: userStore.nickname,
            avatar: res.data
        })
        if (ress.code == 200) {
            userStore.avatar = res.data
            message.success("修改成功")
        } else {
            message.success(ress.data)
        }
    } else {
        message.success("上传失败")
    }
}
</script>

<style scoped>
.listTitle {
    font-size: 16px;
    font-weight: bold;
    margin: 5px 0 15px 0;
}



.youPhone {
    display: flex;
    align-items: center;
}
</style>