<template>
  <div>
    <main>
      <a-layout v-if="$route.path !== '/'">
        <a-layout-sider v-model:collapsed="collapsed" :trigger="null" collapsible>
          <div class="logo"/>
          <a-menu v-model:selectedKeys="selectedKeys" theme="dark" mode="inline">
            <a-menu-item key="1" v-if="$route.path.indexOf('admin') !== -1">
              <router-link to="/admin">
                <user-outlined/>
                <span>学生管理</span>
              </router-link>
            </a-menu-item>
            <a-menu-item key="2" v-if="$route.path.indexOf('admin') !== -1">
              <router-link to="/admin/book">
                <video-camera-outlined/>
                <span>书籍管理</span>
              </router-link>
            </a-menu-item>
            <a-menu-item key="3" v-if="$route.path.indexOf('admin') !== -1">
              <router-link to="/admin/brrt">
                <video-camera-outlined/>
                <span>借阅管理</span>
              </router-link>
            </a-menu-item>
            <a-menu-item key="4" v-if="$route.path.indexOf('student') !== -1">
              <router-link to="/student">
                <video-camera-outlined/>
                <span>藏书阁</span>
              </router-link>
            </a-menu-item>
            <a-menu-item key="5" v-if="$route.path.indexOf('student') !== -1">
              <router-link to="/student/self">
                <user-outlined/>
                <span>个人信息</span>
              </router-link>
            </a-menu-item>
          </a-menu>
        </a-layout-sider>
        <a-layout>
          <a-layout-header style="background: #fff; padding: 0">
            <menu-unfold-outlined
                v-if="collapsed"
                class="trigger"
                @click="() => (collapsed = !collapsed)"
            />
            <menu-fold-outlined v-else class="trigger" @click="() => (collapsed = !collapsed)"/>
          </a-layout-header>
          <a-layout-content
              :style="{ margin: '24px 16px', padding: '24px', background: '#fff', minHeight: '280px' }"
          >
            <!--            内嵌内容-->
            <router-view :key="$route.path"></router-view>
            <!--            内嵌内容-->
          </a-layout-content>
        </a-layout>
      </a-layout>
      <router-view :key="$route.path" v-if="$route.path === '/'">

      </router-view>
    </main>
  </div>
</template>

<script lang="ts">var collapsed;

import {defineComponent, ref} from "vue";
import {
  UserOutlined,
  VideoCameraOutlined,
  UploadOutlined,
  MenuUnfoldOutlined,
  MenuFoldOutlined,
} from '@ant-design/icons-vue';
import router from "@/router";

export default  defineComponent({
  components: {
    UserOutlined,
    VideoCameraOutlined,
    UploadOutlined,
    MenuUnfoldOutlined,
    MenuFoldOutlined,
  },
  setup() {
    return {
      selectedKeys: ref<string[]>(['1','4']),
      collapsed: ref<boolean>(false),
    };
  },
});
</script>

<style scoped>
main {
  min-height: 40em;
}

.trigger {
  font-size: 18px;
  line-height: 64px;
  padding: 0 24px;
  cursor: pointer;
  transition: color 0.3s;
}

.trigger:hover {
  color: #1890ff;
}

.logo {
  height: 32px;
  background: rgba(255, 255, 255, 0.3);
  margin: 16px;
}

.site-layout .site-layout-background {
  background: #fff;
}
</style>

