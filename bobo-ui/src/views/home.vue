<template>
  <n-list hoverable clickable show-divider>

    <n-list bordered>
      <Compose @submit-success="composeSuccess"></Compose>
    </n-list>

    <n-list-item class="list" v-for="(i, index) in dailyList" :key="index">
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
        <!-- 图片 -->
        <div class="images">
          <n-image-group>
            <n-space>
              <n-image v-for="item in i?.DailyAttachment?.filter(item1 => item1.type === '1')" :src="item.url"
                style="max-height: 158px;object-fit: cover;" width="158" />
            </n-space>
          </n-image-group>
        </div>
        <!-- 视频 -->
        <div v-if="i?.DailyAttachment?.filter(item1 => item1.type === '2').length > 0">
          <Video :url="i?.DailyAttachment?.filter(item1 => item1.type === '2')[0].url"></Video>
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
import { onMounted, ref } from "vue";
import { DailyList, SaveOrUpdate } from '@/api/daily'

const returnBr = (i: any) => {
  return i.replace(/\\n/g, "<br />")
}
import useUserStore from "@/store/user";
let userStore = useUserStore();
import { useMessage } from 'naive-ui'

onMounted(() => {
  GetData()
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

const calculateWidth = (len) => {
  if (len >= 3) {
    return '33%'
  }
  if (len == 2) {
    return '30%'
  }
  return '50%'
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