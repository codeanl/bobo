<template>
  <!-- 列表内容 -->
  <n-list hoverable clickable>
    <n-list-item>
      <div style="display: grid;place-items: center;">
        <span>登录后，精彩更多</span>
        <Login></Login>
      </div>
    </n-list-item>
    <n-list-item v-for="i in test" :key="i">
      <!-- 列表项内容 -->
      <n-thing title="相见恨晚" content-style="margin-top: 10px;">
        <template #description>
          <n-space size="small" style="margin-top: 4px">
            <n-tag :bordered="false" type="info" size="small">
              暑夜
            </n-tag>
            <n-tag :bordered="false" type="info" size="small">
              晚春
            </n-tag>
          </n-space>
        </template>
        奋勇呀然后休息呀<br>
        完成你伟大的人生
      </n-thing>
    </n-list-item>
  </n-list>
</template>

<script setup lang="ts">
import { ref, onMounted, watchEffect } from 'vue';
const test = ref(20)
let isD = ref<Boolean>(false)
import Login from '@/components/login.vue'
const scrollContainer = ref<HTMLElement | null>(null);
function checkScrollPosition() {
  const container = scrollContainer.value;
  if (!container) return;

  const scrollHeight = container.scrollHeight;
  const scrollTop = container.scrollTop;
  const clientHeight = container.clientHeight;

  if (!isD.value) {
    // 检查是否滚动到底部
    if (scrollTop + clientHeight >= scrollHeight - 1) {
      console.log("gogo");
      test.value = test.value + 10
      if (test.value >= 100) {
        isD.value = true
        console.log("Stop");
        console.log(test.value);
      }
    }
  }
}

watchEffect(() => {
  const container = scrollContainer.value;
  if (container) {
    // 添加滚动事件监听器
    container.addEventListener('scroll', checkScrollPosition);
  }
});

onMounted(() => {
  // 组件卸载时移除事件监听器
  return () => {
    const container = scrollContainer.value;
    if (container) {
      container.removeEventListener('scroll', checkScrollPosition);
    }
  };
});
</script>