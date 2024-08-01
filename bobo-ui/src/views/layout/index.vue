<template>
  <n-modal-provider>
    <n-dialog-provider>
      <n-config-provider :theme="theme">
        <n-message-provider>
          <!--  -->
          <div style="width: 230px;height: 100vh; position: fixed;left:calc(50% - 550px); " class="media">
            <Menu></Menu>
          </div>

          <!--  -->
          <div style="width: 230px;height: 100vh; position: fixed;right: calc(50% - 550px ); " class="media">
            <Advertise></Advertise>
          </div>

          <!--  -->
          <div style="width: 600px;" class="home">
            <!-- header -->
            <div style="position: sticky ;z-index: 99;left: 0; top: 0;display: flex;justify-content: space-between;height: 60px;
              align-items: center;padding: 0 10px;"
              :class="{ 'dark-border': active, 'no-dark-border': !active, 'bg-black': active, 'bg-white': !active }">
              <div style="display: flex;align-items: center;">
                <n-icon :component="MenuSharp" size="20" style="margin-right: 10px;display: none;" class="sizeBtn"
                  @click="isMedia = !isMedia" />
                <span style="font-size: 15px;font-weight: bold;">BOBO社区</span>
              </div>
              <n-switch v-model:value="active" size="small" @change="changetheme">
                <template #checked-icon>
                  <n-icon :component="Sunny" />
                </template>
                <template #unchecked-icon>
                  <n-icon :component="Moon" />
                </template>
              </n-switch>
            </div>

            <div>
              <!-- 登陆 -->
              <n-list bordered>
                <div style=" display: grid;place-items: center;padding: 20px 0;" v-if="userStore.token == ''">
                  <span>登录后，精彩更多</span>
                  <Login></Login>
                </div>
              </n-list>
              <!--  -->
              <div :class="{ 'dark-border-3': active, 'no-dark-border-3': !active }">
                <router-view></router-view>
              </div>

            </div>
          </div>
          <!--  -->
          <n-drawer v-model:show="isMedia" width="40%" placement="left">
            <n-drawer-content>
              <template #header>
                Header
              </template>
              <template #footer>
                <n-button>Footer</n-button>
              </template>
            </n-drawer-content>
          </n-drawer>
          <!--  -->

          <n-global-style />
        </n-message-provider>
      </n-config-provider>
    </n-dialog-provider>
  </n-modal-provider>
</template>

<script setup lang="ts">
import Menu from '@/views/layout/menu.vue'
import Advertise from '@/views/layout/advertise.vue'
import Login from '@/components/login.vue'
import { Moon, Sunny, MenuSharp } from '@vicons/ionicons5'
import { ref } from 'vue';
import { darkTheme } from 'naive-ui'
import type { GlobalTheme } from 'naive-ui'
import useUserStore from "@/store/user";
let userStore = useUserStore();
const active = ref(false)
const isMedia = ref(false)
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

.no-dark-border-3 {
  border-left: 1px solid #efeff6;
  border-right: 1px solid #efeff6;
  border-bottom: 1px solid #efeff6;
}

.dark-border-3 {
  border-left: 1px solid rgb(43, 43, 46);
  border-right: 1px solid rgb(43, 43, 46);
  border-bottom: 1px solid rgb(43, 43, 46);
}

.bg-black {
  background-color: rgb(24, 24, 28);
}

.bg-white {
  background-color: rgb(255, 255, 255);
}

.media {
  display: block;
}

/* 当屏幕宽度小于400px时，div元素消失 */
@media (max-width: 820px) {
  .media {
    display: none;
  }

  .sizeBtn {
    display: block !important;
  }

  .home {
    max-width: 100vw;
  }
}
</style>
