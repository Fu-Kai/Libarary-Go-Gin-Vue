package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var Db *sql.DB

func Init() (err error) {
	Db, err = sql.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/LibraryManagement?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("数据库连接失败了")
		log.Fatalln(err)
	}

	Db.SetMaxIdleConns(20)
	Db.SetMaxOpenConns(20)

	if err := Db.Ping(); err != nil {
		log.Fatalln(err)
	}
	return
}

// Login ===============Public SQL Begin================
func Login(loginForm LoginForm_fukai) (exist int, err error) {
	fmt.Println(loginForm)
	var sqlStr = ``
	if loginForm.UserType == "1" {
		sqlStr = `SELECT stu_id, stu_acc,stu_passwd FROM student WHERE stu_acc = ? AND stu_passwd = ?`
	} else {
		sqlStr = `SELECT mana_id, mana_acc,mana_passwd FROM manager WHERE mana_acc = ? AND mana_passwd = ?`
	}
	lf := LoginForm_fukai{}
	err = Db.QueryRow(sqlStr, loginForm.Acc, loginForm.Passwd).Scan(&lf.Id, &lf.Acc, &lf.Passwd)
	if err != nil {
		return 0, err
	}
	fmt.Println(lf.Id)
	return lf.Id, nil
}

//===============Public SQL Begin================

//===============Student SQL Begin================

// GetOneStudents 获取某个学生信息
func GetOneStudents(sid string) (student Student_fukai, err error) {
	var sqlStr = `SELECT * FROM student where stu_id = ?`
	err = Db.QueryRow(sqlStr, sid).Scan(&student.StuId, &student.StuName, &student.StuGender, &student.StuAge, &student.StuMajor, &student.StuAble, &student.StuAcc, &student.StuPasswd)
	if err != nil {
		return student, err
	}
	return
}

// GetAllStudents 获取所有学生信息
func GetAllStudents() (students []Student_fukai, err error) {
	var sqlStr = `SELECT * FROM student;`
	rows, err := Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		student := Student_fukai{}
		err = rows.Scan(&student.StuId, &student.StuName, &student.StuGender, &student.StuAge, &student.StuMajor, &student.StuAble, &student.StuAcc, &student.StuPasswd)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	err = rows.Close()
	return
}

func AddStudent(student Student_fukai) (err error) {
	var sqlStr = `INSERT INTO student (stu_name, stu_gender, stu_age, stu_major, stu_acc, stu_passwd) VALUES (?,?,?,?,?,?)`

	_, err = Db.Exec(sqlStr, student.StuName, student.StuGender, student.StuAge, student.StuMajor, student.StuAcc, student.StuPasswd)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return
}

func ModifyStudent(student Student_fukai, sid string) (err error) {
	var sqlStr = `UPDATE student SET stu_name = ?, stu_gender = ?, stu_age = ?, stu_major = ?, stu_acc = ?, stu_passwd = ? WHERE stu_id = ?`
	_, err = Db.Exec(sqlStr, student.StuName, student.StuGender, student.StuAge, student.StuMajor, student.StuAcc, student.StuPasswd, sid)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return
}

func DeleteStudent(sid string) (err error) {
	var sqlStr = `DELETE FROM student where stu_id = ?`

	_, err = Db.Exec(sqlStr, sid)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return
}

//===============Student SQL End==================

//===============Book SQL Start==================

func AddBook(book Book_fukai) (err error) {
	var sqlStr = `INSERT INTO book (book_name, book_author, book_pub, book_num, book_record) VALUES (?,?,?,?,?)`

	_, err = Db.Exec(sqlStr, book.BookName, book.BookAuthor, book.BookPub, book.BookNum, book.BookRecord)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return
}

func ModifyBook(book Book_fukai, bid string) (err error) {
	//var sqlStr = `INSERT INTO student (stu_name, stu_gender, stu_age, stu_major, stu_acc, stu_passwd) VALUES (?,?,?,?,?,?)`
	var sqlStr = `UPDATE book SET book_name = ?, book_author = ?, book_pub = ? WHERE book_id = ?`
	_, err = Db.Exec(sqlStr, book.BookName, book.BookAuthor, book.BookPub, bid)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return
}

func GetOneBook(bid string) (book Book_fukai, err error) {
	var sqlStr = `SELECT * FROM book where book_id = ?`
	err = Db.QueryRow(sqlStr, bid).Scan(&book.BookId, &book.BookName, &book.BookAuthor, &book.BookPub, &book.BookNum, &book.BookRecord)
	if err != nil {
		return book, err
	}
	return
}

func GetAllBooks() (books []Book_fukai, err error) {
	var sqlStr = `SELECT * FROM book`
	rows, err := Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		book := Book_fukai{}
		err = rows.Scan(&book.BookId, &book.BookName, &book.BookAuthor, &book.BookPub, &book.BookNum, &book.BookRecord)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		books = append(books, book)
	}
	err = rows.Close()
	return
}
func DeleteBook(bid string) (err error) {
	var sqlStr = `DELETE FROM book where book_id = ?`

	_, err = Db.Exec(sqlStr, bid)
	if err != nil {

		fmt.Println(err)
		return err
	}
	return
}

//===============Book SQL End==================

//===============Borrow && Return SQL Start==================

func BorrowBook(borrowTab BorrowTab_fukai) (err error) {
	//开启事务
	tx, err := Db.Begin()
	if err != nil {
		return
	}
	var sqlStr1 = `INSERT INTO borrow_return_tab (sid, bid, borrow_date, agreed_return_date) VALUES (?,?,?,?);`
	var sqlStr2 = `UPDATE student SET stu_able = 0 WHERE stu_id = ?;`
	var sqlStr3 = `UPDATE book SET book_num = 0 WHERE book_id = ?;`
	//任务1插入借阅表
	insertBorrow, err := tx.Prepare(sqlStr1)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func(insertBorrow *sql.Stmt) {
		err := insertBorrow.Close()
		if err != nil {

		}
	}(insertBorrow)
	//任务2设置学生借阅状态
	setAble, err := tx.Prepare(sqlStr2)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func(setAble *sql.Stmt) {
		err := setAble.Close()
		if err != nil {

		}
	}(setAble)
	//任务3设置书本下架
	setNum, err := tx.Prepare(sqlStr3)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func(setNum *sql.Stmt) {
		err := setNum.Close()
		if err != nil {

		}
	}(setNum)

	if _, err := insertBorrow.Exec(borrowTab.SId, borrowTab.BId, borrowTab.BorrowDate, borrowTab.AgreedReturnDate); err != nil {

		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	if _, err := setAble.Exec(borrowTab.SId); err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	if _, err := setNum.Exec(borrowTab.BId); err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		fmt.Println(err)
		return err
	}
	err = tx.Commit()
	return
}
func ReturnBook(BRId, sid, bid string) (err error) {
	//开启事务
	tx, err := Db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	var sqlStr1 = `UPDATE borrow_return_tab SET return_date = ? WHERE br_id = ?`
	var sqlStr2 = `UPDATE student SET stu_able = 1 WHERE stu_id = ?;`
	var sqlStr3 = `UPDATE book SET book_num = 1 WHERE book_id = ?;`
	//任务1插入借阅表
	insertReturn, err := tx.Prepare(sqlStr1)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func(insertReturn *sql.Stmt) {
		err := insertReturn.Close()
		if err != nil {

		}
	}(insertReturn)
	//任务2设置学生借阅状态
	setAble, err := tx.Prepare(sqlStr2)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func(setAble *sql.Stmt) {
		err := setAble.Close()
		if err != nil {

		}
	}(setAble)
	//任务3设置书本下架
	setNum, err := tx.Prepare(sqlStr3)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func(setNum *sql.Stmt) {
		err := setNum.Close()
		if err != nil {

		}
	}(setNum)

	h, _ := time.ParseDuration("-1h")
	if _, err := insertReturn.Exec(time.Now().Add(-8*h), BRId); err != nil {

		err := tx.Rollback()
		if err != nil {
			fmt.Println(err)
			return err
		}
		return err
	}
	if _, err := setAble.Exec(sid); err != nil {
		err := tx.Rollback()
		if err != nil {
			fmt.Println(err)
			return err
		}
		return err
	}
	if _, err := setNum.Exec(bid); err != nil {
		err := tx.Rollback()
		if err != nil {
			fmt.Println(err)
			return err
		}
		return err
	}
	err = tx.Commit()
	return
}
func GetAllBorrows() (borrows []BorrowReturnTab_fukai, err error) {
	var sqlStr = `SELECT * FROM borrow_return_tab`
	rows, err := Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		borrow := BorrowReturnTab_fukai{}
		err = rows.Scan(&borrow.BRId, &borrow.SId, &borrow.BId, &borrow.BorrowDate, &borrow.AgreedReturnDate, &borrow.ReturnDate)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		borrows = append(borrows, borrow)
	}
	err = rows.Close()
	return
}
func GetOneBorrow(sid string) (borrow BorrowReturnTab_fukai, book Book_fukai, err error) {
	fmt.Println(sid)
	var sqlStr = `select * from borrow_return_tab inner join book where bid = book_id and sid = ? and return_date IS NULL;`
	err = Db.QueryRow(sqlStr, sid).Scan(&borrow.BRId, &borrow.SId, &borrow.BId, &borrow.BorrowDate, &borrow.AgreedReturnDate, &borrow.ReturnDate, &book.BookId, &book.BookName, &book.BookAuthor, &book.BookPub, &book.BookNum, &book.BookRecord)
	if err != nil {
		return borrow, book, err
	}
	return
}
func DeleteBorrow(BRId string) (err error) {
	var sqlStr = `DELETE FROM borrow_return_tab where br_id = ?`
	_, err = Db.Exec(sqlStr, BRId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return
}

//===============Borrow && Return  SQL  End==================
