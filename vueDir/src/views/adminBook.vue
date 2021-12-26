<template>
  <a-card>
    <a-button style="margin-bottom: 10px;" type="primary" @click="showModal(1)">
      <PlusCircleTwoTone/>
      添加书籍
    </a-button>
    <a-table :columns="columns" :loading="stateBook.loading" :data-source="stateBook.data">
      <template #filterDropdown="{ setSelectedKeys, selectedKeys, confirm, clearFilters, column }">
        <div style="width:163px;padding: 8px;">
          <a-input
              ref="searchInput"
              :placeholder="`搜索学生`"
              :value="selectedKeys[0]"
              style="width: 150px; margin-bottom: 8px; display: block"
              @change="e => setSelectedKeys(e.target.value ? [e.target.value] : [])"
              @pressEnter="handleSearch(selectedKeys, confirm, column.dataIndex)"
          />
          <a-button
              type="primary"
              size="small"
              style="width: 68px; margin-right: 10px"
              @click="handleSearch(selectedKeys, confirm, column.dataIndex)"
          >
            <template #icon>
              <SearchOutlined/>
            </template>
            搜索
          </a-button>
          <a-button size="small" style="width: 68px" @click="handleReset(clearFilters)">
            取消
          </a-button>
        </div>
      </template>
      <template #filterIcon="filtered">
        <search-outlined :style="{ color: filtered ? '#108ee9' : undefined }"/>
      </template>
      <template #status="{text, record}">
              <span v-if="record.BookNum == 1">
                否
              </span>
        <span v-else>
                是
              </span>
      </template>
      <template #action="{text, record}">
        <a @click="deleteBook(record.BookId)">
          <DeleteTwoTone twoToneColor="red"/>
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        </a>
        <a @click="modifyBook(record)">
          <EditTwoTone/>
        </a>
      </template>
    </a-table>
  </a-card>
  <a-modal v-model:visible="visible" title="书本信息" :footer="null" html-type="submit">
    <a-form
        layout="vertical"
        :model="bookFormState"
    >
      <a-form-item style>
        <a-input v-model:value="bookFormState.BookName" placeholder="书名">
          <template #prefix>
            <ProjectOutlined style="color: rgba(0, 0, 0, 0.25)"/>
          </template>
        </a-input>
      </a-form-item>
      <a-form-item style>
        <a-input v-model:value="bookFormState.BookAuthor" placeholder="作者">
          <template #prefix>
            <UserOutlined style="color: rgba(0, 0, 0, 0.25)"/>
          </template>
        </a-input>
      </a-form-item>
      <a-form-item style>
        <a-input v-model:value="bookFormState.BookPub" placeholder="出版社">
          <template #prefix>
            <MenuFoldOutlined style="color: rgba(0, 0, 0, 0.25)"/>
          </template>
        </a-input>
      </a-form-item>


      <a-form-item>
        <a-button type="primary" style="float: right" @click="handleFinish" v-if="which === 1">
          添加
        </a-button>
        <a-button type="primary" style="float: right" @click="modifyFinish" v-else>
          修改
        </a-button>
        <a-button style="float: right; margin-right: 5px" @click="closeModal">
          取消
        </a-button>
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script lang="ts">
import {defineComponent, reactive, ref, toRefs, UnwrapRef} from 'vue';
import {useRouter} from 'vue-router'
import {
  getAllStudentInfo,
  DeleteBook,
  Book,
  ModifyStudent,
  addBook,
  getAllBookInfo,
  nowTime,
  ModifyBook
} from "@/api/ApiMethods";
import {message} from "ant-design-vue";
import {
  DeleteTwoTone,
  CheckCircleOutlined,
  CloseCircleOutlined,
  QuestionCircleOutlined,
  SearchOutlined,
  SyncOutlined,
  UserOutlined,
  VideoCameraOutlined,
  UploadOutlined,
  MenuUnfoldOutlined,
  MenuFoldOutlined,
  PlusCircleTwoTone,
  LockOutlined,
  ManOutlined,
  ProjectOutlined,
  NumberOutlined,
  EditTwoTone
} from '@ant-design/icons-vue'
import {ValidateErrorEntity} from "ant-design-vue/es/form/interface";
import router from "@/router";

interface BookForm {
  BookName: string,
  BookAuthor: string,
  BookPub: string,
  BookNum: number,
  BookRecord: string,
}

export default defineComponent({
  name: "adminBook",
  components: {
    DeleteTwoTone,
    CheckCircleOutlined,
    SyncOutlined,
    CloseCircleOutlined,
    SearchOutlined,
    QuestionCircleOutlined,
    UserOutlined,
    VideoCameraOutlined,
    UploadOutlined,
    MenuUnfoldOutlined,
    MenuFoldOutlined,
    PlusCircleTwoTone,
    LockOutlined,
    ManOutlined,
    ProjectOutlined,
    NumberOutlined,
    EditTwoTone
  },
  setup() {
    const visible = ref<boolean>(false);
    const which = ref<number>(2);
    const BookId = ref<number>(0);
    const bookFormState: UnwrapRef<BookForm> = reactive({
      BookName: '',
      BookAuthor: '',
      BookPub: '',
      BookNum: 1,
      BookRecord: '',
    });

    function refreshBookFormState() {
      BookId.value = 0;
      bookFormState.BookName = '';
      bookFormState.BookAuthor = '';
      bookFormState.BookPub = '';
      bookFormState.BookNum = 1;
      bookFormState.BookRecord = '';
    }

    const handleFinish = async (values: BookForm) => {
      try {
        console.log(bookFormState);
        bookFormState.BookRecord = nowTime().toString()
        console.log(bookFormState);
        let mid = await addBook(bookFormState);
        console.log('mid', mid)
        if (mid.data.code == 0) {
          message.success('书本添加成功  ' + bookFormState.BookName);
          await refreshBookFormState();
        } else {
          message.info('添加失败');
          return
        }
      } catch (e) {
        console.log(e)
        message.error('失败')
      }
      visible.value = false;
      await fetchData();
    };
    const handleFinishFailed = (errors: ValidateErrorEntity<BookForm>) => {
      console.log(errors);
    };
    const showModal = (swhich: number) => {
      which.value = swhich;
      visible.value = true;
    };
    const closeModal = async () => {
      visible.value = false;
      await refreshBookFormState();
    };

    const handleOk = (e: MouseEvent) => {
      console.log(e);
      visible.value = false;
    };
    const searchState = reactive({
      searchText: '',
      searchedColumn: '',
    });
    const searchInput = ref();

    const columns = [
      {
        title: '书本名称',
        dataIndex: 'BookName',
        key: 'BookName',
        align: 'center',
        slots: {
          filterDropdown: 'filterDropdown',
          filterIcon: 'filterIcon',
        },
        onFilter: (value: string, record: { BookName: { toString: () => string; }; }) =>
            record.BookName.toString().toLowerCase().includes(value.toLowerCase()),
        onFilterDropdownVisibleChange: (visible: any) => {
          if (visible) {
            setTimeout(() => {
              console.log(searchInput.value);
              searchInput.value.focus();
            }, 100);
          }
        },
      },
      {
        title: '作者',
        align: 'center',
        dataIndex: 'BookAuthor',
        key: 'BookAuthor',
      },
      {
        title: '出版社',
        align: 'center',
        dataIndex: 'BookPub',
        key: 'BookPub',
      },
      {
        title: '是否被借',
        align: 'center',
        dataIndex: 'BookNum',
        key: 'BookNum',
        slots: {customRender: 'status'}
      },
      {
        title: '上架时间',
        align: 'center',
        dataIndex: 'BookRecord',
        key: 'BookRecord',
      },
      {
        title: '删除 /  修改',
        dataIndex: 'action',
        key: 'action',
        align: 'center',
        slots: {customRender: 'action'}
      }
    ]

    let stateBook = reactive<{ loading: boolean, data: Book[] }>({
      loading: true,
      data: []
    })

    async function fetchData() {
      try {
        stateBook.loading = true;
        let mid = await getAllBookInfo()
        let booksJSON = mid.data.data;
        console.log(booksJSON)
        if (booksJSON == null) {
          stateBook.loading = false;
          stateBook.data = [];
          return
        }

        const result = []
        for (let book of booksJSON) {
          book.BookRecord = book.BookRecord.replace(/T/g, ' ').replace(/Z/g, '')
          result.push(book);
        }
        stateBook.data = result;
        stateBook.loading = false;
      } catch (e) {
        message.error('获取书籍数据失败')
      }
    }

    const handleSearch = (selectedKeys: string[], confirm: () => void, dataIndex: string) => {
      confirm();
      searchState.searchText = selectedKeys[0];
      searchState.searchedColumn = dataIndex;
    };

    const handleReset = (clearFilters: () => void) => {
      clearFilters();
      searchState.searchText = '';
    };
    const deleteBook = async (index: number) => {
      console.log(index)
      await DeleteBook(index);
      message.success('删除成功')
      await fetchData();
    }
    const modifyBook = async (record: any) => {
      bookFormState.BookName = record.BookName;
      bookFormState.BookAuthor = record.BookAuthor;
      bookFormState.BookPub = record.BookPub;
      bookFormState.BookRecord = nowTime().toString();
      BookId.value = record.BookId;
      showModal(2);
    }
    const modifyFinish = async () => {
      try {
        await ModifyBook(bookFormState, BookId.value);
        message.success('修改成功')
      } catch (e) {
        console.log(e);
      }
      await closeModal();
      await refreshBookFormState();
      await fetchData();
    }
    fetchData();
    return {
      bookFormState,
      handleFinish,
      handleFinishFailed,
      visible,
      showModal,
      closeModal,
      handleOk,
      stateBook,
      columns,
      handleSearch,
      handleReset,
      searchInput,
      ...toRefs(searchState),
      deleteBook,
      modifyBook,
      modifyFinish,
      which,
    }
  }
})
</script>
<style>
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
