<template>
  <a-card>
    <a-space>
      <a-form
          layout="inline"
          :model="formState"
          @finish="handleFinish"
      >
        <a-form-item>
          <a-select
              v-model:value="formState.SId"
              style="width: 120px"
              placeholder="用户名"
              @change="handleChange1"
              :options="state.studentData.map(stu => ({ value: stu.StuName + '&emsp;SID:' + stu.StuId}))"
          >
          </a-select>
        </a-form-item>
        <a-form-item>
          <a-select
              v-model:value="formState.BId"
              style="width: 140px"
              placeholder="书名"
              @change="handleChange2"
              :options="state.bookData.map(book => ({ value: book.BookName + '&emsp;BID:' + book.BookId }))"
          >
          </a-select>
        </a-form-item>
        <a-form-item>
          <a-date-picker
              v-model:value="formState.AgreedReturnDate"
              @change="handleChange2"
              show-time
              placeholder="选择最晚归还时间"/>
        </a-form-item>
        <a-form-item>
          <a-button type="primary" html-type="submit" :disabled="disabled < 3">
            录入借阅关系
          </a-button>
        </a-form-item>
      </a-form>
    </a-space>
    <a-divider/>
    <div style="text-align: center">
      <span style="font-size: 20px">借阅&归还</span>
      <a-table :columns="columns" :data-source="state.borrowsData">
        <template #action="{text, record}">
          <a @click="deleteBorrow(record.BRId)" v-if="record.ReturnDate !== ''">
            <DeleteTwoTone twoToneColor="red"/>
          </a>
          <a @click="returnBook(record)" v-else>
            <LoginOutlined/>
          </a>
        </template>
      </a-table>
    </div>

  </a-card>
</template>
<script lang="ts">
import {
  Book,
  getAllBookInfo,
  getAllStudentInfo,
  nowTime,
  Student,
  toMysqlTime,
  Borrow,
  BorrowTab, getAllBorrows, getOneStudentInfo, getOneBookInfo, Return, DeleteBorrow
} from "@/api/ApiMethods";


import {defineComponent, reactive, UnwrapRef, ref} from 'vue';
import {message} from "ant-design-vue";
import {DeleteTwoTone, EditTwoTone, LoginOutlined} from '@ant-design/icons-vue'

interface FormState {
  BRId: string;
  SId: string;
  BId: string;
  BorrowDate: string;
  AgreedReturnDate: string;
  ReturnDate: string;
}

export default defineComponent({
  components: {
    EditTwoTone,
    DeleteTwoTone,
    LoginOutlined,
  },
  setup() {
    const columns = [
      {
        title: '姓名',
        align: 'center',
        dataIndex: 'SName',
        key: 'SName',
      },
      {
        title: '书名',
        align: 'center',
        dataIndex: 'BName',
        key: 'BName',
      },
      {
        title: '借书时间',
        align: 'center',
        dataIndex: 'BorrowDate',
        key: 'BorrowDate',
      },
      {
        title: '应归还时间',
        align: 'center',
        dataIndex: 'AgreedReturnDate',
        key: 'AgreedReturnDate',
      },
      {
        title: '实际归还时间',
        align: 'center',
        dataIndex: 'ReturnDate',
        key: 'ReturnDate',
      },
      {
        title: '删除 / 还书',
        dataIndex: 'action',
        key: 'action',
        align: 'center',
        slots: {customRender: 'action'}
      }
    ]
    const formState: UnwrapRef<FormState> = reactive({
      BRId: '',
      SId: '',
      BId: '',
      BorrowDate: '',
      AgreedReturnDate: '',
      ReturnDate: '',
    });
    const disabled = ref(0);
    const handleChange1 = (value: string) => {
      console.log(`selected ${value}`);
      disabled.value++;
    };
    const handleChange2 = (value: string) => {
      console.log(`selected ${value}`);
      disabled.value++;
    };
    const handleFinish = async (values: FormState) => {
      formState.BorrowDate = nowTime();
      let SId = '';
      let BId = '';
      for (let i = 0; i < formState.SId.length; i++) {
        if (formState.SId[i] == ':') {
          for (let j = i + 1; j < formState.SId.length; j++) {
            SId += formState.SId[j];
          }
          break;
        }
      }
      for (let i = 0; i < formState.BId.length; i++) {
        if (formState.BId[i] == ':') {
          for (let j = i + 1; j < formState.BId.length; j++) {
            BId += formState.BId[j];
          }
          break;
        }
      }
      let borrowState = {
        SId: parseInt(SId),
        BId: parseInt(BId),
        BorrowDate: formState.BorrowDate,
        AgreedReturnDate: toMysqlTime(formState.AgreedReturnDate),
      };
      try {
        await Borrow(borrowState);
        message.success('添加借阅关系成功')
      } catch (e) {
        console.log(e)
      }
      console.log(values, borrowState);
      await fetchData();
    };

    const state = reactive<{ studentData: Student[], bookData: Book[], borrowsData: BorrowTab[] }>({
      studentData: [],
      bookData: [],
      borrowsData: [],
    })

    async function refreshState() {
      state.studentData = [];
      state.bookData = [];
      formState.BRId = '';
      formState.SId = '';
      formState.BId = '';
      formState.BorrowDate = '';
      formState.AgreedReturnDate = '';
      formState.ReturnDate = '';
    }

    async function fetchData() {
      try {
        await refreshState();
        let mid = await getAllStudentInfo()
        let studentsJSON = mid.data.data;
        mid = await getAllBookInfo()
        let booksJSON = mid.data.data;
        console.log(studentsJSON)
        console.log(booksJSON)
        for (let student of studentsJSON) {
          if (student.StuAble == 1)
            state.studentData.push(student);
        }
        for (let book of booksJSON) {
          if (book.BookNum > 0) {
            state.bookData.push(book)
          }
        }
        let mid2 = await getAllBorrows()
        if(mid2.data.data == null){
          state.borrowsData = [];
          return
        }
        let borrowsJSON = mid2.data.data;
        console.log(borrowsJSON)
        const result = []
        for (let borrows of borrowsJSON) {
          let midS = await getOneStudentInfo(borrows.SId);
          let midB = await getOneBookInfo(borrows.BId);
          borrows.SName = midS.data.data.StuName;
          borrows.BName = midB.data.data.BookName;
          borrows.BorrowDate = borrows.BorrowDate.replace(/T/g, ' ').replace(/Z/g, '')
          borrows.AgreedReturnDate = borrows.AgreedReturnDate.replace(/T/g, ' ').replace(/Z/g, '')
          if (borrows.ReturnDate.Valid)
            borrows.ReturnDate = borrows.ReturnDate.Time.replace(/T/g, ' ').replace(/Z/g, '')
          else
            borrows.ReturnDate = '';
          result.push(borrows);
        }
        state.borrowsData = result;
      } catch (e) {
        //message.error('获取学生数据失败')
      }
    }

    fetchData()

    async function returnBook(record: any) {
      console.log('record', record)
      try {
        await Return(record.BRId, record.SId, record.BId);
        await refreshState();
        await fetchData();
        message.success('还书成功');
      } catch (e) {
        console.log(e)
      }
    }

    async function deleteBorrow(BRId: any) {
      console.log(BRId)
      try {
        await DeleteBorrow(BRId)
        await refreshState();
        message.success('删除借阅记录成功');
        await fetchData();
      } catch (e) {
        console.log(e)
      }
    }

    return {
      state,
      handleChange1,
      handleChange2,
      handleFinish,
      formState,
      disabled,
      columns,
      returnBook,
      deleteBorrow,
    };
  },
});
</script>

