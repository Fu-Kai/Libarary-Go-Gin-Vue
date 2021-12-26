<template>
  <a-list :grid="{ gutter: 16, column: 4 }" :data-source="stateBook.data">
    <template #renderItem="{ item }">
      <a-list-item>
        <a-card hoverable style="width: 240px">
          <template #cover>
            <img alt="example"
                 src="https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fbkimg.cdn.bcebos.com%2Fpic%2F4d086e061d950a7be3613b8e0ad162d9f2d3c920&refer=http%3A%2F%2Fbkimg.cdn.bcebos.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1642592826&t=baa46a2600922bd6c7771e00d838aa5d"/>
          </template>
          <a-card-meta :title="item.BookName">
            <template #description>
              {{ '作者：' + item.BookAuthor }}
              <br>
              {{ item.BookPub }}
              <a style="float: right" @click="showModal(item)" v-if="Able">
                <logout-outlined/>
                借阅
              </a>
            </template>
          </a-card-meta>

        </a-card>
      </a-list-item>
    </template>
  </a-list>
  <a-modal v-model:visible="visible" :title=title :footer="null" html-type="submit">
    <a-form
        layout="vertical"
        :model="borrowFormState"
    >

      <span style="float: left; margin-top: 5px">借阅人：</span>
      <a-form-item style>
        <a-input :value=SName placeholder="借阅人" disabled="true">
          <template #prefix>
          </template>
        </a-input>
      </a-form-item>

      <span style="float: left; margin-top: 5px">书本ID：</span>
      <a-form-item style>
        <a-input v-model:value=borrowFormState.Bid placeholder="书本 id" disabled="true">
          <template #prefix>
          </template>
        </a-input>
      </a-form-item>

      <span style="float: left; margin-top: 5px">还书日期：</span>
      <a-form-item>
        <a-date-picker
            v-model:value="borrowFormState.AgreedReturnDate"
            show-time
            placeholder="选择最晚归还时间"
        />
      </a-form-item>

      <a-form-item>
        <a-button type="primary" style="float: right" @click="finish">
          借阅
        </a-button>
        <a-button style="float: right; margin-right: 5px" @click="closeModal">
          取消
        </a-button>
      </a-form-item>
    </a-form>
  </a-modal>
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
import {Book, Borrow, getAllBookInfo, getOneStudentInfo, nowTime, toMysqlTime} from "@/api/ApiMethods";
import {message} from "ant-design-vue";

interface BorrowForm {
  Bid: string,
  BorrowDate: string,
  AgreedReturnDate: string,
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
    const SName = ref<any>('');
    const Able = ref('');
    const visible = ref<boolean>(false);
    const title = ref<string>('');
    const borrowFormState: UnwrapRef<BorrowForm> = reactive({
      Bid: '',
      BorrowDate: '',
      AgreedReturnDate: '',
    });
    const stateBook = reactive<{ loading: boolean, data: Book[] }>({
      loading: true,
      data: []
    })
    const showModal = (item: any) => {
      borrowFormState.Bid = item.BookId
      title.value = item.BookName;
      visible.value = true;
    };
    const closeModal = async () => {
      visible.value = false;
    };

    async function fetchData() {
      try {
        stateBook.loading = true;
        let midName = await getOneStudentInfo(sid);
        SName.value = midName.data.data.StuName;
        Able.value = midName.data.data.StuAble;
        let mid = await getAllBookInfo()
        let booksJSON = mid.data.data;
        console.log(booksJSON)
        if (booksJSON == null) {
          stateBook.loading = false;
          return
        }
        const result = []
        for (let book of booksJSON) {
          book.BookRecord = book.BookRecord.replace(/T/g, ' ').replace(/Z/g, '')
          if (book.BookNum != 0)
            result.push(book);
        }
        stateBook.data = result;
        stateBook.loading = false;
      } catch (e) {
        message.error('获取书籍数据失败')
      }
    }

    fetchData();

    const finish = async () => {
      if (toMysqlTime(borrowFormState.AgreedReturnDate) <= nowTime()) {
        message.info('还书时间不能早于现在时间')
        return;
      }
      let borrowState = {
        SId: sid,
        BId: parseInt(borrowFormState.Bid),
        BorrowDate: nowTime(),
        AgreedReturnDate: toMysqlTime(borrowFormState.AgreedReturnDate),
      };
      try {
        await Borrow(borrowState);
        message.success('借阅成功')
        await closeModal();
      } catch (e) {
        console.log(e)
      }
      borrowFormState.AgreedReturnDate = '';
      await fetchData();
    };
    return {
      stateBook,
      sid,
      SName,
      Able,
      borrowFormState,
      visible,
      title,
      showModal,
      closeModal,
      finish,
    };
  },
});
</script>

<style>

</style>
