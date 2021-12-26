package db

import (
	"database/sql"
	"time"
)

type LoginForm_fukai struct {
	Id       int    `json:"id" form:"id"`
	Acc      string `json:"acc" form:"acc"`
	Passwd   string `json:"passwd" form:"passwd"`
	UserType string `json:"userType" form:"userType"`
}

// Student_fukai 定义学生表结构体
type Student_fukai struct {
	StuId     int    `sql:"stu_id" json:"StuId" form:"StuId"`
	StuName   string `sql:"stu_name" json:"StuName" form:"StuName"`
	StuGender string `sql:"stu_gender" json:"StuGender" form:"StuGender"`
	StuAge    string `sql:"stu_age" json:"StuAge" form:"StuAge"`
	StuMajor  string `sql:"stu_major" json:"StuMajor" form:"StuMajor"`
	StuAble   int    `sql:"stu_able" json:"StuAble" form:"StuAble"`
	StuAcc    string `sql:"stu_acc" json:"StuAcc" form:"StuAcc"`
	StuPasswd string `sql:"stu_passwd" json:"StuPasswd" form:"StuPasswd"`
}

// Book_fukai 定义书本结构体
type Book_fukai struct {
	BookId     int       `json:"BookId" form:"BookId"`
	BookName   string    `json:"BookName" form:"BookName"`
	BookAuthor string    `json:"BookAuthor" form:"BookAuthor"`
	BookPub    string    `json:"BookPub" form:"BookPub"`
	BookNum    int       `json:"BookNum" form:"BookNum"`
	BookRecord time.Time `json:"BookRecord" form:"BookRecord"`
}

// BorrowTab_fukai 定义借书记录结构体
type BorrowTab_fukai struct {
	SId              int       `json:"SId" form:"SId"`
	BId              int       `json:"BId" form:"BId"`
	BorrowDate       time.Time `json:"BorrowDate" form:"BorrowDate"`
	AgreedReturnDate time.Time `json:"AgreedReturnDate" form:"AgreedReturnDate"`
}

// BorrowReturnTab_fukai 定义借阅表结构
type BorrowReturnTab_fukai struct {
	BRId             int          `json:"BRId" form:"BRId"`
	SId              int          `json:"SId" form:"SId"`
	BId              int          `json:"BId" form:"BId"`
	BorrowDate       time.Time    `json:"BorrowDate" form:"BorrowDate"`
	AgreedReturnDate time.Time    `json:"AgreedReturnDate" form:"AgreedReturnDate"`
	ReturnDate       sql.NullTime `json:"ReturnDate" form:"ReturnDate"`
}

// Manager_fukai 定义管理员结构体
type Manager_fukai struct {
	ManaId     int    `json:"ManaId" form:"ManaId"`
	ManaName   string `json:"ManaName" form:"ManaName"`
	ManaPhone  string `json:"ManaPhone" form:"ManaPhone"`
	ManaAcc    string `json:"ManaAcc" form:"ManaAcc"`
	ManaPasswd string `json:"ManaPasswd" form:"ManaPasswd"`
}
