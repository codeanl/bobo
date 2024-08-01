<template>
  <n-infinite-scroll style="max-height: 100vh" :distance="10" @load="handleLoad">
    <div v-for="(item, index) in dada" :key="index" class="message"
      style="height: 100px;background-color: aqua;margin: 5px;">
      <span> {{ item }} </span>
    </div>
    <div v-if="loading" class="text">
      加载中...
    </div>
    <div v-if="noMore" class="text">
      没有更多了 🤪
    </div>
  </n-infinite-scroll>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from 'vue'

export default defineComponent({
  setup() {
    const loading = ref(false)
    const pageSize = ref(10)
    const pageNum = ref(1)
    const total = ref(40)
    const dada = ref([1, 2, 3, 4, 5, 6, 7, 8, 9, 10])
    const dada1 = ref([1, 2, 3, 4, 5, 6, 7, 8, 9, 10])


    // const noMore = computed(() => items.value.length > 16)
    const noMore = ref(false)
    const handleLoad = async () => {
      if (loading.value || noMore.value)
        return
      if (pageNum.value >= Math.floor(total.value / pageSize.value)) {
        console.log("最后一页");
        return
      }
      loading.value = true
      await new Promise(resolve => setTimeout(resolve, 1000))
      pageNum.value++
      console.log("当前第" + pageNum.value);
      console.log("next");

      dada.value = [...dada.value, ...dada1.value]
      loading.value = false
    }

    return {
      dada1,
      noMore,
      loading,
      handleLoad,
      dada
    }
  }
})
</script>

<style scoped>
.message {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
  padding: 10px;
}

.message:last-child {
  margin-bottom: 0;
}

.reverse {
  flex-direction: row-reverse;
}

.text {
  text-align: center;
}

.reverse .avatar {
  margin-left: 10px;
}

.avatar {
  width: 28px;
  height: 28px;
  border-radius: 50%;
  margin-right: 10px;
}
</style>