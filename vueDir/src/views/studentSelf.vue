<template>
  <a-card style="width: 330px; font-size: 17px; float:left;" title="个人信息">
    <a-avatar style="float: right; color: #f56a00; background-color: #fde3cf">{{ studentInfoState.data.StuName }}
    </a-avatar>
    <p>姓名：{{ studentInfoState.data.StuName }}</p>
    <p>性别：{{ studentInfoState.data.StuGender }}</p>
    <p>年龄：{{ studentInfoState.data.StuAge }}</p>
    <p>专业：{{ studentInfoState.data.StuMajor }}</p>
    <p>借阅状态：
      <span v-if="studentInfoState.data.StuAble">可借</span>
      <span v-else>已借</span>
    </p>
    <p>
      账号：{{ studentInfoState.data.StuAcc }}
      <a-button type="danger" style="float: right" @click="logout">退出登录</a-button>
    </p>
  </a-card>
  <a-card style="margin-left: 20em; width: 240px; font-size: 17px; float:left;" title="正在看的">
    <a-empty v-if="studentInfoState.data.StuAble">
      <template #description>
      <span>
        还没有借书，去
        <a href="http://localhost:8080/student">藏书阁</a>挑选
      </span>
      </template>
    </a-empty>
    <template #cover>
      <img
          v-if="!studentInfoState.data.StuAble"
          alt="example"
          src="https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fbkimg.cdn.bcebos.com%2Fpic%2F4d086e061d950a7be3613b8e0ad162d9f2d3c920&refer=http%3A%2F%2Fbkimg.cdn.bcebos.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1642592826&t=baa46a2600922bd6c7771e00d838aa5d"/>
    </template>
    <a-card-meta :title="borrowInfoState.dataBook.BookName" v-if="!studentInfoState.data.StuAble">
      <template #description>
        {{ '作者：' + borrowInfoState.dataBook.BookAuthor }}
        <br>
        {{ borrowInfoState.dataBook.BookPub }}
        <a style="float: right" @click="returnBook">
          <logout-outlined/>
          归还
        </a>
      </template>
    </a-card-meta>
  </a-card>
  <a-card style="float:right; width: 25em" v-if="!studentInfoState.data.StuAble">
    <a-row :gutter="1">
      <a-col :span="24">
        <a-statistic
            title="应归还日期"
            :value="borrowInfoState.dataBorrow.AgreedReturnDate"

            @finish="onFinish"
        />
      </a-col>
      <a-col :span="24">
        <a-statistic-countdown title="还书日倒计时" :value="borrowInfoState.dataBorrow.AgreedReturnDate"
                               format="D 天 H 时 m 分 s 秒"/>
      </a-col>
    </a-row>
  </a-card>
</template>


<script lang="ts">
import {
  UserOutlined,
  VideoCameraOutlined,
  UploadOutlined,
  MenuUnfoldOutlined,
  MenuFoldOutlined,
  LogoutOutlined,
  EditOutlined,
} from '@ant-design/icons-vue';
import {defineComponent, reactive, ref, UnwrapRef} from 'vue';
import {useRouter} from 'vue-router'
import {Moment} from 'moment';
import {Book, Student, getAllBookInfo, getOneStudentInfo, BorrowTab, getOneBorrowInfo, Return} from "@/api/ApiMethods";
import {message} from "ant-design-vue";
interface BorrowInfo {
  BRId: string,
  SId: string,
  BId: string,
  BorrowDate: string,
  AgreedReturnDate: string,
  ReturnDate: string
}
export default defineComponent({
  components: {
    UserOutlined,
    VideoCameraOutlined,
    UploadOutlined,
    MenuUnfoldOutlined,
    MenuFoldOutlined,
    LogoutOutlined,
    EditOutlined,
  },
  setup() {
    const sid = sessionStorage.getItem('StuId')
    const studentInfoState = reactive<{ loading: boolean, data: Student }>({
      loading: true,
      data: {
        StuId: 0,
        StuName: '',
        StuGender: '',
        StuAge: '',
        StuMajor: '',
        StuAble: 1,
        StuAcc: '',
        StuPasswd: '',
      }
    })
    const borrowInfoState = reactive<{ loading: boolean, dataBorrow: BorrowInfo, dataBook: Book }>({
      loading: true,
      dataBorrow: {
        BRId: '',
        SId: '',
        BId: '',
        BorrowDate: '',
        AgreedReturnDate: '',
        ReturnDate: ''
      },
      dataBook: {
        BookId: 0,
        BookName: '',
        BookAuthor: '',
        BookPub: '',
        BookNum: 1,
        BookRecord: '',
      }
    })

    async function fetchData() {
      let mid = await getOneStudentInfo(sid)
      let midBorrow = await getOneBorrowInfo(sid)
      studentInfoState.data = mid.data.data

      midBorrow.data.borrowData.AgreedReturnDate = midBorrow.data.borrowData.AgreedReturnDate.replace(/T/g, ' ').replace(/Z/g, '')
      if (midBorrow.data.borrowData.ReturnDate.Valid)
        midBorrow.data.borrowData.ReturnDate = midBorrow.data.borrowData.ReturnDate.Time.replace(/T/g, ' ').replace(/Z/g, '')
      else
        midBorrow.data.borrowData.ReturnDate = '';
      borrowInfoState.dataBorrow = midBorrow.data.borrowData

      borrowInfoState.dataBook = midBorrow.data.bookData
      console.log('123', midBorrow)
      console.log('1234', borrowInfoState)
      borrowInfoState.loading = false;
      studentInfoState.loading = false;
    }

    fetchData();

    async function returnBook() {
      try {
        await Return(borrowInfoState.dataBorrow.BRId, borrowInfoState.dataBorrow.SId, borrowInfoState.dataBorrow.BId);
        message.success('还书成功');
        await fetchData();
      } catch (e) {
        console.log(e)
      }
    }
    const router = useRouter();

    async function logout() {
      message.success('退出登录成功，缓存已清理')
      sessionStorage.clear();
      await router.push(`/`)
    }

    const onFinish = () => {
      console.log('finished!');
    };
    return {
      sid,
      studentInfoState,
      logout,
      deadline: Date.now() + 1000 * 60 * 60 * 24 * 2 + 1000 * 30,
      onFinish,
      borrowInfoState,
      returnBook,
    };
  },
});
</script>
