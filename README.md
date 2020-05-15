
# gorm tools

gorm tools is a set of  tools for database operation. 

*This tool uses the xorm-cmd toolkit, thanks to the xorm tool author's open source*

## Source Install

`git clone https://github.com/liangyt123/gorm`


## Commands

All the commands below.

* **reverse**     reverse a db to codes
* **shell**       a general shell to operate all kinds of database
* **dump**        dump database all table struct's and data to standard output
* **source**      execute a sql from std in
* **driver**      list all supported drivers

## Reverse

Reverse command is a tool to convert your database struct to all kinds languages of structs or classes. After you installed the tool, you can type 

`gorm help reverse`

to get help

example:

`cd $GOPATH/src/github.com/liangyt123/gorm`

sqlite:
`gorm reverse sqite3 test.db templates/gogorm`

mysql:
`gorm reverse mysql 'root:123456@tcp(127.0.0.1:3306)/gorm_test'?charset=utf8 templates/gogorm`

mymysql:
`gorm reverse mymysql gorm_test2/root/ templates/gogorm`

postgres:
`gorm reverse postgres "dbname=gorm_test sslmode=disable" templates/gogorm`

mssql:
`gorm reverse mssql "server=test;user id=testid;password=testpwd;database=testdb" templates/gogorm`

will generated go files in `./model` directory

### Template and Config

Now, gorm tool supports go and c++ two languages and have go, gogorm, c++ three of default templates. In template directory, we can put a config file to control how to generating.

```
lang=go
genJson=1
```

lang must be go or c++ now.
genJson can be 1 or 0, if 1 then the struct will have json tag.

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
