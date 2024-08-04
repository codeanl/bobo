<template>
  <n-infinite-scroll style="max-height:  calc(100vh - 80px)" :distance="2" @load="handleLoad">
    <n-list hoverable clickable show-divider>
      <n-list bordered>
        <Compose @submit-success="composeSuccess"></Compose>
      </n-list>
      <n-list-item class="list" v-for="(i, index) in dailyList" :key="index" @click="lookInfo(i.id)">
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
          <div>
            <n-image-group>
              <n-space>
                <n-image-group>
                  <n-grid :x-gap="4" :y-gap="4" :cols="3">
                    <template v-for="(img, index) in i?.DailyAttachment?.filter(item1 => item1.type === '1')"
                      :key="index">
                      <n-gi>
                        <n-image @click.stop class="post-img" object-fit="cover" :src="img.url" />
                      </n-gi>
                    </template>
                  </n-grid>
                </n-image-group>
                <!-- <n-image 
                  v-for="item in i?.DailyAttachment?.filter(item1 => item1.type === '1')" :src="item.url"
                  style="max-height: 158px;object-fit: cover;" width="158" @click.stop />
               -->
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
            <p>0</p>
          </div>
          <div class="commit btn">
            <n-icon size="18">
              <ChatbubbleEllipsesOutline />
            </n-icon>
            <p v-if="i.comment_count >= 100">{{ (parseInt(i.comment_count, 10) / 1000).toFixed(1) }}k</p>
            <p v-else>{{ i.comment_count ? i.comment_count : 0 }}</p>
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
            <p>0</p>
          </div>
        </div>
      </n-list-item>
    </n-list>
    <div style="padding: 20px 0">
      <div v-if="loading" class="text" style="display: flex;align-items: baseline;justify-content: center;">
        加载中...
      </div>
      <div v-if="noMore" class="text" style="display: flex;align-items: baseline;justify-content: center;">
        没有更多了 🤪
      </div>
    </div>
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

onMounted(() => {
  GetData()
});

const message = useMessage();

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
    total.value = res.count
  } else {
    message.error(res.message)
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


import { useRouter } from 'vue-router';
import { useMessage } from 'naive-ui';
const router = useRouter();
const lookInfo = (id) => {
  router.push({ path: '/daily', query: { id: id } });
}

const loading = ref(false)
const noMore = ref(false)
const total = ref(0)

const handleLoad = async () => {
  console.log(111);

  if (loading.value || noMore.value)
    return
  if (page_num.value > Math.floor(total.value / page_size.value)) {
    console.log("最后一页");
    noMore.value = true
    return
  }
  loading.value = true
  page_num.value++
  console.log("获取第" + page_num.value + "页数据");
  let res: any = await DailyList({
    params: {
      page_size: page_size.value, page_num: page_num.value
    }
  })
  if (res.code == 200) {
    dailyList.value.push(...res.data)
    loading.value = false
  } else {
    page_num.value--
    loading.value = false
  }
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

.post-img {
  img {
    width: 100%;
    height: 100%;
  }
}
</style>