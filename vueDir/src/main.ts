import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import Antd from 'ant-design-vue';
import './antd.css';
//import 'ant-design-vue/dist/antd.css';
//import 'ant-design-vue/dist/antd.dark.css';
import './index.css';

// @ts-ignore
createApp(App).use(router).use(Antd).mount('#app')
