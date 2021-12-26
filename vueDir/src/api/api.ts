export const api = {
    'login': 'http://localhost:7070/login',
    //用户
    'addStudent': 'http://localhost:7070/student/add',//插入学生
    'modifyStudent': 'http://localhost:7070/student/modify',//修改学生信息
    'getAllStudents': 'http://localhost:7070/student',//获取所有学生
    'getOneStudent': 'http://localhost:7070/student',//获取单个学生
    'deleteStudent': 'http://localhost:7070/student/delete',//删除学生
    //书本
    'addBook': 'http://localhost:7070/book/add',
    'modifyBook': 'http://localhost:7070/book/modify',//修改学生信息
    'getAllBooks': 'http://localhost:7070/book',
    'getOneBook': 'http://localhost:7070/book',//获取单个学生
    'deleteBook': 'http://localhost:7070/book/delete',
    //借还
    'getAllBorrows': 'http://localhost:7070/borrows',
    'getOneBorrow': 'http://localhost:7070/borrow',
    'borrow': 'http://localhost:7070/student/borrow',
    'return': 'http://localhost:7070/student/return',
    'deleteBorrow': 'http://localhost:7070/delete/borrow',
};
