！！开始配置之前，请保证电脑7070与8080端口未被占用！！

依赖包均已打包，无需重新上网拉

go依赖包已打包在`Libarary `--> `LibararySystem `--> `vender`下

vue依赖包已打包在`Libarary `--> `vueDir `--> `node_modules.zip`

## 后端

#### 1.在mysql中创建名为`LibraryManagement`的Database

#### 2.将dump.sql脚本文件在该Database下执行，会完成自动建表与测试数据的添加

![image-20211226000001826](https://gitee.com/mekaifu/u-pic/raw/master/uPic/image-20211226000001826.png)

#### 3.用编辑器打开后端数据库配置文件并修改 `Libarary `--> `LibararySystem `--> `db` --> `db.go`

a：替换为您的mysql用户名

b:  替换为您的mysql密码

c:  替换为您的mysql服务运行地址

<img src="https://gitee.com/mekaifu/u-pic/raw/master/uPic/image-20211225233829688.png" alt="image-20211225233829688" style="zoom: 33%;" />

4.回到`LibararySystem`目录编译运行main.go

`go build main.go`      `./main或者./main.exe(windows)`

至此后端运行完毕

<img src="https://gitee.com/mekaifu/u-pic/raw/master/uPic/image-20211225235831874.png" alt="image-20211225235831874" style="zoom: 33%;" />

## 前端

#### 1.解压依赖包 `Libarary `--> `vueDir `--> `node_modules.zip`

#### 2.在vueDir目录下运行`npm run serve` 或 `yarn serve`

<img src="/Users/kai/Library/Application Support/typora-user-images/image-20211226000400023.png" alt="image-20211226000400023" style="zoom: 33%;" />

#### 3.在浏览器中打开localhost:8080开始使用

<img src="/Users/kai/Library/Application Support/typora-user-images/image-20211226000506174.png" alt="image-20211226000506174" style="zoom:50%;" />