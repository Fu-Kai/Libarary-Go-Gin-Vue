<template>
  <a-card>
    <a-button style="margin-bottom: 10px;" type="primary" @click="showModal(1)">
      <PlusCircleTwoTone/>
      添加学生
    </a-button>
    <a-table :columns="columns" :loading="stateStudent.loading" :data-source="stateStudent.data">
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
              <span v-if="record.StuAble == 1">
                未借书
              </span>
        <span v-else>
                已借书
              </span>
      </template>
      <template #action="{text, record}">
        <a @click="deleteStudent(record.StuId)">
          <DeleteTwoTone twoToneColor="red"/>
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        </a>
        <a @click="modifyStudent(record)">
          <EditTwoTone/>
        </a>
      </template>
    </a-table>
  </a-card>
  <a-modal v-model:visible="visible" title="学生信息" :footer="null" html-type="submit">
    <a-form
        layout="vertical"
        :model="studentFormState"
    >
      <a-form-item style>
        <a-input v-model:value="studentFormState.StuName" placeholder="姓名">

          <template #prefix>
            <UserOutlined style="color: rgba(0, 0, 0, 0.25)"/>
          </template>
        </a-input>
      </a-form-item>
      <a-form-item style>
        <a-input v-model:value="studentFormState.StuGender" placeholder="性别">
          <template #prefix>
            <ManOutlined style="color: rgba(0, 0, 0, 0.25)"/>
          </template>
        </a-input>
      </a-form-item>
      <a-form-item style>
        <a-input v-model:value="studentFormState.StuAge" placeholder="年龄">
          <template #prefix>
            <UserOutlined style="color: rgba(0, 0, 0, 0.25)"/>
          </template>
        </a-input>
      </a-form-item>
      <a-form-item style>
        <a-input v-model:value="studentFormState.StuMajor" placeholder="专业">
          <template #prefix>
            <ProjectOutlined style="color: rgba(0, 0, 0, 0.25)"/>
          </template>
        </a-input>
      </a-form-item>
      <a-form-item style>
        <a-input v-model:value="studentFormState.StuAcc" placeholder="账号">
          <template #prefix>
            <NumberOutlined style="color: rgba(0, 0, 0, 0.25)"/>
          </template>
        </a-input>
      </a-form-item>
      <a-form-item>
        <a-input v-model:value="studentFormState.StuPasswd" placeholder="密码">
          <template #prefix>
            <LockOutlined style="color: rgba(0, 0, 0, 0.25)"/>
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
import {getAllStudentInfo, DeleteStudent, Student, addStudent, ModifyStudent} from "@/api/ApiMethods";
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

interface StudentForm {
  StuName: string,
  StuGender: string,
  StuAge: string,
  StuMajor: string,
  StuAcc: string,
  StuPasswd: string,
}

export default defineComponent({
  name: "adminHome",
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
    const studentId = ref<number>(0);
    const studentFormState: UnwrapRef<StudentForm> = reactive({
      StuName: '',
      StuGender: '',
      StuAge: '',
      StuMajor: '',
      StuAcc: '',
      StuPasswd: '',
    });

    function refreshStudentFormState() {
      studentId.value = 0;
      studentFormState.StuName = '';
      studentFormState.StuGender = '';
      studentFormState.StuAge = '';
      studentFormState.StuMajor = '';
      studentFormState.StuAcc = '';
      studentFormState.StuPasswd = '';
    }

    const handleFinish = async (values: StudentForm) => {
      try {
        console.log(studentFormState);
        let mid = await addStudent(studentFormState);
        console.log('mid', mid)
        if (mid.data.code == 0) {
          message.success('学生添加成功 姓名：' + studentFormState.StuName);
          await refreshStudentFormState();
        } else {
          message.info('该账号已存在，请重新选择');
          studentFormState.StuAcc = '';
          return
        }
      } catch (e) {
        console.log(e)
        message.error('失败')
      }
      visible.value = false;
      await fetchData();
    };
    const handleFinishFailed = (errors: ValidateErrorEntity<StudentForm>) => {
      console.log(errors);
    };
    const showModal = (swhich: number) => {
      which.value = swhich;
      visible.value = true;
    };
    const closeModal = async () => {
      visible.value = false;
      await refreshStudentFormState();
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
        title: '学生姓名',
        dataIndex: 'StuName',
        key: 'StuName',
        align: 'center',
        slots: {
          filterDropdown: 'filterDropdown',
          filterIcon: 'filterIcon',
        },
        onFilter: (value: string, record: { StuName: { toString: () => string; }; }) =>
            record.StuName.toString().toLowerCase().includes(value.toLowerCase()),
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
        title: '性别',
        align: 'center',
        dataIndex: 'StuGender',
        key: 'StuGender',
      },
      {
        title: '年龄',
        align: 'center',
        dataIndex: 'StuAge',
        key: 'StuAge',
      },
      {
        title: '专业',
        align: 'center',
        dataIndex: 'StuMajor',
        key: 'StuMajor',
      },
      {
        title: '是否已借书',
        align: 'center',
        dataIndex: 'StuAble',
        key: 'StuAble',
        slots: {customRender: 'status'}
      },
      {
        title: '账号',
        align: 'center',
        dataIndex: 'StuAcc',
        key: 'StuAcc',
      },
      {
        title: '密码',
        align: 'center',
        dataIndex: 'StuPasswd',
        key: 'StuPasswd',
      },
      {
        title: '删除 /  修改',
        dataIndex: 'action',
        key: 'action',
        align: 'center',
        slots: {customRender: 'action'}
      }
    ]

    let stateStudent = reactive<{ loading: boolean, data: Student[] }>({
      loading: true,
      data: []
    })

    async function fetchData() {
      try {
        stateStudent.loading = true;
        let mid = await getAllStudentInfo()
        let studentsJSON = mid.data.data;
        console.log(studentsJSON)
        if (studentsJSON == null) {
          stateStudent.loading = false;
          stateStudent.data = [];
          return
        }

        const result = []
        for (let student of studentsJSON) {
          result.push(student);
        }
        stateStudent.data = result;
        stateStudent.loading = false;
      } catch (e) {
        message.error('获取学生数据失败')
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
    const deleteStudent = async (index: number) => {
      console.log(index)
      await DeleteStudent(index);
      message.success('删除成功')
      await fetchData();
    }
    const modifyFinish = async () => {
      try {
        await console.log(studentFormState.StuAge)
        await ModifyStudent(studentFormState, studentId.value);
        message.success('修改成功')
      } catch (e) {
        console.log(e);
      }
      await closeModal();
      await refreshStudentFormState();
      await fetchData();
    }
    const modifyStudent = async (record: any) => {
      studentFormState.StuName = record.StuName;
      studentFormState.StuGender = record.StuGender;
      studentFormState.StuMajor = record.StuMajor;
      studentFormState.StuAge = record.StuAge;
      studentFormState.StuAcc = record.StuAcc;
      studentFormState.StuPasswd = record.StuPasswd;
      studentId.value = record.StuId;
      showModal(2);
    }
    fetchData();
    return {
      studentFormState,
      handleFinish,
      handleFinishFailed,
      visible,
      showModal,
      closeModal,
      handleOk,
      stateStudent,
      columns,
      handleSearch,
      handleReset,
      searchInput,
      ...toRefs(searchState),
      deleteStudent,
      modifyStudent,
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
