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

    dbins.SetConnMaxLifetime(100)

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


// func MysqlInit() {
// 	var err error
// 	mysql_url := configs.ReadYaml().Mysql.Username + ":" + configs.ReadYaml().Mysql.Password + "@tcp(" + configs.ReadYaml().Mysql.Host + ":" +
// 		configs.ReadYaml().Mysql.Port + ")/" + configs.ReadYaml().Mysql.Database + "?charset=utf8mb4&parseTime=True"

// 	// fmt.Println("url", url)
// 	Db, err = sqlx.Open("mysql", mysql_url)
// 	if err != nil {
// 		// fmt.Println("数据库连接错误: ", err)
// 		panic("数据库连接错误: " + err.Error())
// 	}
// 	//设置连接池最大连接数
// 	Db.SetMaxOpenConns(100)
// 	//设置连接池最大空闲连接数
// 	Db.SetMaxIdleConns(20)

// 	// 验证是否连接
// 	ctx := context.Background()
// 	if err = Db.PingContext(ctx); err != nil {
// 		// fmt.Println("数据库未连接: ", err)
// 		panic("数据库未连接: " + err.Error())
// 	}

// 	// fmt.Println("数据库连接成功: ", Db.Ping())
// 	// Rdb, err = redis.Dial("tcp", Rdb_url, redis.DialDatabase(Rdb_db), redis.DialPassword(Rdb_pwd))
// 	// if err != nil {
// 	// 	panic("redis连接错误: " + err.Error())
// 	// }
// 	//StartRedis()
// }

func DbClose() {
	Db.Close()
}
