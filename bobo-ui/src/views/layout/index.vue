<template>
  <n-config-provider :theme="theme">
    <div style="margin: 0 auto;display: flex;">
      <!-- 左侧菜单 -->
      <Menu></Menu>
      <!-- 内容-->
      <div style="display: grid; grid-template-columns: auto 1fr;height: 100vh;">
        <!--  -->
        <div style="grid-column: 1 / -1;display: flex;justify-content: space-between;height: 60px;width: 610px;
       align-items: center;padding: 0 10px;" :class="{ 'dark-border': active, 'no-dark-border': !active }">
          <span style="font-size: 15px;font-weight: bold;">bobo社区</span>
          <n-switch v-model:value="active" size="small" @change="changetheme">
            <template #checked-icon>
              <n-icon :component="Sunny" />
            </template>
            <template #unchecked-icon>
              <n-icon :component="Moon" />
            </template>
          </n-switch>
        </div>
        <!--  -->
        <div style="grid-column: 2;overflow-y: auto;width: 630px; border-top: 0;border-bottom: 0;"
          :class="{ 'dark-border ': active, 'no-dark-border ': !active }">
          <router-view></router-view>
        </div>
      </div>
      <!-- 右侧 -->
      <Advertise></Advertise>
    </div>
    <n-global-style />
  </n-config-provider>
</template>

<script setup lang="ts">
import Menu from '@/views/layout/menu.vue'
import Advertise from '@/views/layout/advertise.vue'
import { Moon, Sunny } from '@vicons/ionicons5'
import { ref } from 'vue';
import { darkTheme } from 'naive-ui'
import type { GlobalTheme } from 'naive-ui'
import { defineComponent } from 'vue'

const active = ref(false)
const theme = ref<GlobalTheme | null>(null)
const changetheme = () => {
  theme.value = active.value == true ? darkTheme : null;
};

</script>
<style scoped>
.no-dark-border {
  border: 1px solid #efeff6;
}

.dark-border {
  border: 1px solid rgb(43, 43, 46);
}
</style>
