<template>
    <div style="height: 100vh;width: 230px;position: relative;">
        <!--  -->
        <div style="padding: 10px 0 0 20px;">
            <n-avatar size="large"
                src="https://cdn.paopao.info/public/avatar/1f/28/84/e0/489b-40e4-948b-ac3c2f931546.jpeg" />
        </div>
        <!--  -->
        <div style="position: absolute;left: 0;bottom: 10px;width: 230px;display: grid;place-items: center;">
            <Login></Login>
        </div>
    </div>
</template>

<script setup lang="ts">
import Login from '@/components/login.vue'

</script>

<style lang="scss"></style>