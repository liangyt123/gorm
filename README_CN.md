# gorm 工具

gorm 是一组数据库操作命令行工具。 

*该工具用到了xorm-cmd的工具包,感谢xorm的工具作者的开源*

## 源码安装

`git clone https://github.com/liangyt123/gorm`

## 命令列表

有如下可用的命令：

* **reverse**     反转一个数据库结构，生成代码
* **shell**       通用的数据库操作客户端，可对数据库结构和数据操作
* **dump**        Dump数据库中所有结构和数据到标准输出
* **source**      从标注输入中执行SQL文件
* **driver**      列出所有支持的数据库驱动

## 反转

Reverse 命令让你根据数据库的表来生成结构体或者类代码文件。安装好工具之后，可以通过

`gorm help reverse`

获得帮助。

例子:

首先要进入到当前项目的目录下，主要是后面的命令最后一个参数中用到的模版存放在当前项目的目录下

`cd $GOPATH/github.com/liangyt123/gorm`

sqlite:
`gorm reverse sqite3 test.db templates/gogorm`

mysql:
`gorm reverse mysql 'root:123456@tcp(127.0.0.1:3306)/gorm_test'?charset=utf8 templates/gogorm`

mymysql:
`gorm reverse mymysql gorm_test2/root/ templates/gogorm`

postgres:
`gorm reverse postgres "dbname=gorm_test sslmode=disable" templates/gogorm`

之后将会生成代码 generated go files in `./model` directory

### 模版和配置

当前，默认支持Go，C++ 和 objc 代码的生成。具体可以查看源码下的 templates 目录。在每个模版目录中，需要放置一个配置文件来控制代码的生成。如下：

```
lang=go
genJson=1
```

`lang` 目前支持 go， c++ 和 objc。
`genJson` 可以为0或者1，如果是1则结构会包含json的tag，此项配置目前仅支持Go语言。

## Shell

Shell command provides a tool to operate database. For example, you can create table, alter table, insert data, delete data and etc.

`gorm shell sqlite3 test.db` will connect to the sqlite3 database and you can type `help` to list all the shell commands.

## Dump

Dump command provides a tool to dump all database structs and data as SQL to your standard output.

`gorm dump sqlite3 test.db` could dump sqlite3 database test.db to standard output. If you want to save to file, just
type `gorm dump sqlite3 test.db > test.sql`.

## Source

`gorm source sqlite3 test.db < test.sql` will execute sql file on the test.db.

## Driver

List all supported drivers since default build will not include sqlite3.

## LICENSE

 BSD License
 [http://creativecommons.org/licenses/BSD/](http://creativecommons.org/licenses/BSD/)
