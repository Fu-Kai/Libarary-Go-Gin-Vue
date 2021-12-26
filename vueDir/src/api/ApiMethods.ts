import {request} from './http'
import {api} from './api'

export declare interface Student {
    StuId: number,
    StuName: string,
    StuGender: string,
    StuAge: string,
    StuMajor: string,
    StuAble: number,
    StuAcc: string,
    StuPasswd: string,
}

export declare interface Book {
    BookId: number,
    BookName: string,
    BookAuthor: string,
    BookPub: string,
    BookNum: number,
    BookRecord: string,
}

export declare interface BorrowTab {
    SId: number,
    BId: number,
    BorrowDate: string,
    AgreedReturnDate: string,
}

function login(formState: any) {
    return request({
        method: 'post',
        herders: {},
        data: formState,
        url: api.login,
    })
}

function addStudent(studentFormState: any) {
    return request({
        method: 'post',
        herders: {},
        data: studentFormState,
        url: api.addStudent,
    })
}

function addBook(bookFormState: any) {
    return request({
        method: 'post',
        herders: {},
        data: bookFormState,
        url: api.addBook,
    })
}

function ModifyStudent(studentFormState: any, sid: any) {
    return request({
        method: 'post',
        herders: {},
        data: studentFormState,
        url: api.modifyStudent + '/' + sid,
    })
}

function ModifyBook(bookFormState: any, bid: any) {
    return request({
        method: 'post',
        herders: {},
        data: bookFormState,
        url: api.modifyBook + '/' + bid,
    })
}

function getOneStudentInfo(sid: any) {
    return request({
        method: 'get',
        herders: {},
        url: api.getOneStudent + '/' + sid,
    })
}

function getAllStudentInfo() {
    return request({
        method: 'get',
        herders: {},
        url: api.getAllStudents,
    })
}

function getAllBookInfo() {
    return request({
        method: 'get',
        herders: {},
        url: api.getAllBooks,
    })
}

function DeleteStudent(sid: any) {
    return request({
        method: 'post',
        herders: {
            'Content-Type': 'application/json; charset=UTF-8'
        },
        url: api.deleteStudent + '/' + sid,
    })
}

function DeleteBook(bid: any) {
    return request({
        method: 'post',
        herders: {
            'Content-Type': 'application/json; charset=UTF-8'
        },
        url: api.deleteBook + '/' + bid,
    })
}

function DeleteBorrow(BRId: any) {
    return request({
        method: 'post',
        herders: {
            'Content-Type': 'application/json; charset=UTF-8'
        },
        url: api.deleteBorrow + '/' + BRId,
    })
}

function Borrow(borrowState: any) {
    borrowState.SId = parseInt(borrowState.SId);
    return request({
        method: 'post',
        herders: {},
        data: borrowState,
        url: api.borrow,
    })
}

function Return(BRId: any, bid: any, sid: any) {
    return request({
        method: 'post',
        herders: {},
        url: api.return + '/' + BRId + '/' + bid + '/' + sid,
    })
}

function getOneBorrowInfo(BRId: any) {
    return request({
        method: 'get',
        herders: {},
        url: api.getOneBorrow + '/' + BRId,
    })
}

function getAllBorrows() {
    return request({
        method: 'get',
        herders: {},
        url: api.getAllBorrows,
    })
}

function getOneBookInfo(bid: any) {
    return request({
        method: 'get',
        herders: {},
        url: api.getOneBook + '/' + bid,
    })
}

function nowTime() {
    let date = new Date();
    let year = date.getFullYear();
    let month = addZero(date.getMonth() + 1);
    let day = addZero(date.getDate());
    let Hours = addZero(date.getHours());
    let Minutes = addZero(date.getMinutes());
    let Seconds = addZero(date.getSeconds());

    let s_createtime = year + '-' + month + '-' + day + 'T' + Hours + ':' + Minutes + ':' + Seconds + 'Z';
    return s_createtime.toString()
}

function addZero(date: number) {
    let newdate = date.toString();
    if (date < 10) {
        newdate = "0" + newdate;
    }
    return newdate;
}

function toMysqlTime(dates: any) {
    let date = new Date(dates)
    let year = date.getFullYear();
    let month = addZero(date.getMonth() + 1);
    let day = addZero(date.getDate());
    let Hours = addZero(date.getHours());
    let Minutes = addZero(date.getMinutes());
    let Seconds = addZero(date.getSeconds());

    let s_createtime = year + '-' + month + '-' + day + 'T' + Hours + ':' + Minutes + ':' + Seconds + 'Z';
    return s_createtime.toString()
}

export {
    login,
    getOneStudentInfo,
    getAllStudentInfo,
    getOneBookInfo,
    getAllBookInfo,
    addStudent,
    addBook,
    ModifyStudent,
    ModifyBook,
    DeleteStudent,
    DeleteBook,
    nowTime,
    toMysqlTime,
    Borrow,
    Return,
    getAllBorrows,
    getOneBorrowInfo,
    DeleteBorrow,
}
