<template>
    <div style="height: 100vh;width: 200px;position: relative;">
        <!-- 按钮 -->
        <div style="height: 50px;width: 200px;position: absolute;left: 0;bottom: 0;
            text-align: center;display: flex;align-items: center;justify-content: center; ">
            <n-button strong secondary round type="primary" class="btnbtn" @click="toLogin">
                登录
            </n-button>
            <n-button strong secondary round type="info" class="btnbtn" @click="toRegister">
                注册
            </n-button>
        </div>
    </div>
</template>

<script setup lang="ts">

</script>

<style scoped>
.btnbtn {
    margin: 0 5px;
}
</style>