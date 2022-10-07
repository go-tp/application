package extend

import (
	"gtp/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
	"database/sql"
)

var Db *sqlx.DB

func MysqlInit() (*sql.DB, error) {

    //构建连接信息
	mysql_url := configs.ReadYaml().Mysql.Username + ":" + configs.ReadYaml().Mysql.Password + "@tcp(" + configs.ReadYaml().Mysql.Host + ":" +
		configs.ReadYaml().Mysql.Port + ")/" + configs.ReadYaml().Mysql.Database + "?charset=utf8mb4&parseTime=True"


    //打开数据库，前面是驱动名称，所以要导入:mysql驱动github.com/go-sql-driver/mysql

    dbins, err := sql.Open("mysql", mysql_url)

    if nil != err {

        fmt.Println("Open Database Error:", err)

        return nil, err

    }

    // 设置数据库的最大连接数
    dbins.SetConnMaxLifetime(2000)

    // 设置最大打开连接数
    dbins.SetMaxOpenConns(2000)

    // 设置数据库最大的闲置连接数
    dbins.SetMaxIdleConns(10)

    // 验证连接

    if err = dbins.Ping(); nil != err {

        fmt.Println("Open Database Fail,Error:", err)

        return nil, err

    }

    fmt.Println("Connect Success!!!")

    return dbins, nil

}

func DbClose() {
	Db.Close()
}

// demo

// 连接 mysql
// mysqlCon,_ := extend.MysqlInit()

// result, _ := mysqlCon.Exec("INSERT INTO article (title, content, status,created_at,updated_at) VALUES (?,?,?, ?, ?)","文章1","正文1","1","1664378735","1664378735")

// fmt.Println("result:",  result)


// // 插入单条
// stmt, _ := mysqlCon.Prepare("INSERT article SET title=?,content=?,status=?,created_at=?,updated_at=?")
// res, _ := stmt.Exec("文章1","正文1","1","1664378735","1664378735")

// id, _ := res.LastInsertId()

// fmt.Println("id:",      id)

// // 查询单条
// var title, content string
// var status int
// mysqlCon.QueryRow("SELECT title,content,status FROM article WHERE id=?", 1).Scan(&title, &content, &status)

// fmt.Println("title:",   title)
// fmt.Println("content:", content)
// fmt.Println("status:",  status)

// // 查询多条

// // 事务
// tx, err := mysqlCon.Begin()
// fmt.Println("err:",     err)

// err1 := tx.Commit()
// fmt.Println("err:",     err1)

// err2 := tx.Rollback()
// fmt.Println("err:",     err2)
