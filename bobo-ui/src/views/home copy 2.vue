<template>
  <n-infinite-scroll style="max-height: 100vh" :distance="10" @load="handleLoad">
    <n-list hoverable clickable show-divider>

      <n-list bordered>
        <Compose @submit-success="composeSuccess"></Compose>
      </n-list>
      <n-list-item ref="scrollContainer" class="list" v-for="(i, index) in dailyList" :key="index"
        @click="lookInfo(i.id)">
        <div class="info">
          <div class="profile">
            <n-avatar round size="medium" :src="i.user.avatar" />
            <p>{{ i.user?.nickname }}</p>
            <p>@{{ i.user?.username }}</p>
            <n-tag type="warning" size="small" v-if="i.is_top == '1'"> 置顶</n-tag>
          </div>
          <div class="time">
            <p>{{ i.ip_loc }} · {{ formatDate(i.created_at) }}</p>
          </div>
        </div>
        <div class="content">
          <p v-html="returnBr(i.content)"></p>
          <div v-if="i?.DailyAttachment?.filter(item1 => item1.type === '3').length > 0">
            <n-button type="success" dashed size="small" style="margin: 0 0 5px 0;"
              @click.stop="downloadZip(i?.DailyAttachment?.filter(item1 => item1.type === '3')[0].url)">
              附件
            </n-button>
          </div>
          <!--  -->
          <!-- 图片 -->
          <div class="images">
            <n-image-group>
              <n-space>
                <n-image v-for="item in i?.DailyAttachment?.filter(item1 => item1.type === '1')" :src="item.url"
                  style="max-height: 158px;object-fit: cover;" width="158" @click.stop />
              </n-space>
            </n-image-group>
          </div>
          <!-- 视频 -->
          <div v-if="i?.DailyAttachment?.filter(item1 => item1.type === '2').length > 0" style="width: 70%;">
            <Video :url="i?.DailyAttachment?.filter(item1 => item1.type === '2')[0].url" @click.stop></Video>
          </div>
          <!-- 链接 -->
          <div class="link">
            <div class="linkto" v-for="item in i?.DailyAttachment?.filter(item1 => item1.type === '4')">
              <n-icon size="14" color="green">
                <Link />
              </n-icon>
              <a :href="item.url">{{ item.url }}</a>
            </div>
          </div>
        </div>


        <div class="start">
          <div class="lick btn">
            <n-icon size="18">
              <HeartOutline />
            </n-icon>
            <p>11</p>
          </div>
          <div class="commit btn">
            <n-icon size="18">
              <ChatbubbleEllipsesOutline />
            </n-icon>
            <p>11</p>
          </div>
          <div class="eye btn">
            <n-icon size="18">
              <EyeOutline />
            </n-icon>
            <p v-if="i.view_count >= 100">{{ (parseInt(i.view_count, 10) / 1000).toFixed(1) }}k</p>
            <p v-else>{{ i.view_count ? i.view_count : 0 }}</p>
          </div>
          <div class="share btn">
            <n-icon size="18">
              <ShareSocialOutline />
            </n-icon>
            <p>11</p>
          </div>
        </div>
      </n-list-item>
    </n-list>
  </n-infinite-scroll>
</template>

<script setup lang="ts">
import Compose from '@/components/compose.vue'
import Video from '@/components/video.vue'
// 监听子组件的事件
const composeSuccess = () => {
  GetData()
}

import {
  ShareSocialOutline,
  ChatbubbleEllipsesOutline,
  HeartOutline,
  Link,
  EyeOutline
} from "@vicons/ionicons5";
import { onMounted, onUnmounted, ref } from "vue";
import { DailyList, SaveOrUpdate } from '@/api/daily'

const returnBr = (i: any) => {
  return i.replace(/\n/g, "<br />")
}
import useUserStore from "@/store/user";
let userStore = useUserStore();
import { useMessage } from 'naive-ui'
import { download } from 'naive-ui/es/_utils';

onMounted(() => {
  window.addEventListener('scroll', handleScroll);
  GetData()
});

const scrollContainerRef = ref(null);

// 定义一个方法来处理滚动事件
const handleScroll = () => {
  const scrollElement = scrollContainerRef.value;
  if (!scrollElement) return;

  // 计算是否滚动到底部
  const isAtBottom = scrollElement.scrollHeight - scrollElement.scrollTop === scrollElement.clientHeight;
  if (isAtBottom) {
    console.log('滚动到底部了！');
    // 执行加载更多数据的逻辑
  }
};
onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll);
});


const page_num = ref(1)
const page_size = ref(10)
const dailyList = ref([])
// 获取动态数据
const GetData = async () => {
  let res: any = await DailyList({
    params: {
      page_size: page_size.value, page_num: page_num.value
    }
  })
  if (res.code == 200) {
    dailyList.value = res.data
  }
}

// 时间转化
const formatDate = (dateString: string) => {
  const date = new Date(dateString);
  const year = date.getFullYear();
  const month = (date.getMonth() + 1).toString().padStart(2, '0');
  const day = date.getDate().toString().padStart(2, '0');
  const hours = date.getHours().toString().padStart(2, '0');
  const minutes = date.getMinutes().toString().padStart(2, '0');
  return `${year}-${month}-${day} ${hours}:${minutes}`;
}

const downloadZip = (url) => {
  window.open(url)
}

const lookInfo = (id) => {
  console.log(id);
}

const loading = ref(false)
const noMore = ref(false)
const total = ref(0)

const handleLoad = async () => {
  if (loading.value || noMore.value)
    return
  if (page_num.value > Math.floor(total.value / page_size.value)) {
    console.log("最后一页");
    noMore.value = true
    return
  }
  loading.value = true
  const data = ref([])
  data.value = dailyList.value

  // GetData()
  page_num.value++
  console.log("获取第" + page_num.value + "页数据");
  let res: any = await DailyList({
    params: {
      page_size: page_size.value, page_num: page_num.value
    }
  })
  if (res.code == 200) {
    dailyList.value.push(...res.data)
  } else {
    page_num.value--
  }
  loading.value = false
}
</script>

<style lang="scss">
.list {
  .info {
    display: flex;
    justify-content: space-between;

    .profile {
      display: flex;
      align-items: center;

      p:first-of-type {
        font-size: 14px;
        font-weight: bold;
        margin: 0 5px;
      }

      p:nth-of-type(2) {
        font-size: 12px;
        color: #625f5f;
        margin-right: 5px;
      }
    }

    .time {
      display: flex;
      align-items: center;

      p {
        font-size: 12px;
        margin: 0 5px;
      }
    }
  }

  .content {
    font-size: 14px;
    margin: 0px 20px;

    .link {
      .linkto {
        display: flex;
        flex-wrap: wrap;
        align-items: center;
      }

      a {
        padding-left: 3px;
        color: green;
        text-decoration: none;
      }
    }
  }

  .start {
    margin: 10px 30px;
    display: flex;
    justify-content: space-between;
    align-items: center;

    .btn {
      display: flex;
      align-items: center;
    }

    .btn p {
      margin-left: 10px;
    }
  }
}
</style>