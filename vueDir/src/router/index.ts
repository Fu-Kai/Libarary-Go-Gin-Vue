import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
const routes: Array<RouteRecordRaw> = [
  {
    path: '/admin',
    name: 'adminHome',
    component: () => import('../views/adminHome.vue')
  },
  {
    path: '/admin/book',
    name: 'adminBook',
    component: () => import('../views/adminBook.vue')
  },
  {
    path: '/admin/brrt',
    name: 'adminBrRt',
    component: () => import('../views/adminBrRt.vue')
  },
  {
    path: '/student',
    name: 'studentHome',
    component: () => import('../views/studentHome.vue')
  },
  {
    path: '/student/self',
    name: 'studentSelf',
    component: () => import('../views/studentSelf.vue')
  },
  {
    path: '/home',
    name: 'homes',
    component: () => import('../views/layout.vue')
  },
  {
    path: '/',
    name: 'home',
    component: () => import('../views/login/login.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

// router.beforeEach((to, from, next) => {
//   if (to.meta.requireAuth) {  // 判断该路由是否需要登录权限
//     //如果需要就执行下面的代码
//
//     if(sessionStorage.getItem("isLogin")==='1')
//     {
//       next({path:'/'});
//     }
//     else if(sessionStorage.getItem("isLogin")==='2')
//     {
//       // window.location.href = "/manager";
//       next({path:'/manager'});
//     }
//     else if(sessionStorage.getItem("isLogin")==='3')
//     {
//       next({path:'/admin'});
//     }
//         // 通过sessionStorage 获取当前的isLogin的值 将我们保存的isLogin的值赋给num,num是顺便取的名称，可以换
//         //我们在登录界面，我们使用请求，请求成功后，我们就使用sessionStorage为‘isLogin’保存一个值  1，如果请求失败，就不保存‘isLogin’的值
//     //如果请求成功，num的值就是1，请求失败就是null，所以下面进行判断
//     else{
//       next({
//         path: '/',//返回登录界面
//         // query: {redirect: to.fullPath}
//       })
//     }
//   }
//   else {
//     //如果不需要登录权限就直接跳转到该路由
//     next();
//   }
// })
export default router
