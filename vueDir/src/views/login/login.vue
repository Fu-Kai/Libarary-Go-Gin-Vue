<template>
  <div class="background"></div>

  <div class="front">
    <a-card style="opacity: 0.85;" hoverable="true" headStyle="font-size:19px;margin-left:0.9em" bordered="true"
            title="图书管理系统">
      <a-form
          layout="vertical"
          :model="formState"
          @finish="handleFinish"
          @finishFailed="handleFinishFailed"
      >
        <a-form-item style>
          <a-input v-model:value="formState.acc" placeholder="Username">
            <template #prefix>
              <UserOutlined style="color: rgba(0, 0, 0, 0.25)"/>
            </template>
          </a-input>
        </a-form-item>
        <a-form-item>
          <a-input v-model:value="formState.passwd" type="password" placeholder="Password">
            <template #prefix>
              <LockOutlined style="color: rgba(0, 0, 0, 0.25)"/>
            </template>
          </a-input>
        </a-form-item>
        <a-form-item>
          <a-radio-group v-model:value="formState.userType ">
            <a-radio value="1" name="type">学生</a-radio>
            <a-radio value="2" name="type">管理员</a-radio>
          </a-radio-group>
        </a-form-item>

        <a-form-item>
          <a-button
              style="width: 100%"
              type="primary"
              html-type="submit"
              :disabled="formState.acc === '' || formState.passwd === ''"
          >
            登录
          </a-button>
        </a-form-item>
      </a-form>
    </a-card>

  </div>
</template>
<script lang="ts">
import {UserOutlined, LockOutlined, SmileOutlined} from '@ant-design/icons-vue';
import {ValidateErrorEntity} from 'ant-design-vue/es/form/interface';
import {defineComponent, reactive, UnwrapRef, h} from 'vue';
import {useRouter} from 'vue-router'
import {message, notification} from "ant-design-vue";
import {login} from "@/api/ApiMethods";

interface FormState {
  acc: string;
  passwd: string;
  userType: string;
}

export default defineComponent({
  setup() {
    const router = useRouter();
    router.go(1)
    const formState: UnwrapRef<FormState> = reactive({
      acc: '',
      passwd: '',
      userType: '1',
    });
    const handleFinish = async (values: FormState) => {
      try {
        console.log(formState);

        let mid = await login(formState)
        console.log(mid)
        if (mid.data.code == 0 && formState.userType == '2') {
          openNotification(formState.acc, formState.userType);
          await router.push(`/admin`)
        } else if (mid.data.code == 0 && formState.userType == '1') {
          openNotification(formState.acc, formState.userType);
          let sid = mid.data.StuId;
          sessionStorage.setItem('StuId', sid)
          await router.push(`/student`)
        } else
          message.error('账号或密码错误')
      } catch (e) {
        console.log(e)
        message.error('失败')
      }
    };
    const handleFinishFailed = (errors: ValidateErrorEntity<FormState>) => {
      console.log(errors);
    };
    const openNotification = (name: any, userType: any) => {
      if (userType == '1') {
        notification.open({
          message: '欢迎来到在线图书馆',
          description:
              '你可以点击左侧菜单栏中的藏书阁浏览目前已上架图书进行借阅服务',
          icon: h(SmileOutlined, {style: 'color: #108ee9'}),
          onClick: () => {
            console.log('Notification Clicked!');
          },
        });
      } else {
        notification.open({
          message: '欢迎您，尊贵的管理员',
          description:
              '你可以点击左侧菜单栏对学生和馆内书籍进行管理',
          icon: h(SmileOutlined, {style: 'color: #108ee9'}),
          onClick: () => {
            console.log('Notification Clicked!');
          },
        });
      }
    };
    return {
      router,
      formState,
      handleFinish,
      handleFinishFailed,
      openNotification
    };
  },
  components: {
    UserOutlined,
    LockOutlined,
  },
});
</script>
<style>
.background {
  background: url("../../../public/src=http___img.zcool.cn_community_010f135bc40480a8012099c8401a6c.jpg&refer=http___img.zcool.jpeg");
  width: 100%;
  height: 120%; /**宽高100%是为了图片铺满屏幕 */
  z-index: -1;
  position: absolute;

}

.front {
  z-index: 1;
  position: absolute;
  top: 30%;
  left: 43%;
  width: 14%;
}

</style>
