<template>
  <div ref="scrollContainer" style="overflow-y: auto; "> <!-- 添加 ref 到容器 -->
    <n-list hoverable clickable>
      <n-list-item v-for="i in test" :key="i">
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
      <!-- 添加一个隐藏的锚点元素 -->
      <div ref="sentinel" style="height: 1px;"></div>
    </n-list>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';
import { NList, NListItem, NThing, NSpace, NTag } from 'naive-ui';

const test = ref(20);
const scrollContainer = ref<HTMLElement | null>(null);
const sentinel = ref<HTMLElement | null>(null);

function handleScrollBottom() {
  console.log('666');
}

onMounted(() => {
  const observer = new IntersectionObserver((entries) => {
    if (entries[0].isIntersecting) {
      handleScrollBottom();
    }
  }, {
    root: scrollContainer.value,
    threshold: 1.0
  });

  if (sentinel.value) {
    observer.observe(sentinel.value);
  }

  onUnmounted(() => {
    if (sentinel.value) {
      observer.unobserve(sentinel.value);
    }
  });
});
</script>
