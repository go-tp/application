package model

import (
	"github.com/gin-gonic/gin"
	"gtp/extend"
	"fmt"
	_"github.com/garyburd/redigo/redis"
)

func IndexM(c *gin.Context) interface{} {
	return "Welcome to go-tp.com."
}


type queryLogin struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}


func LoginM(c *gin.Context) interface{} {

	// 获取参数
	query := queryLogin{}
	c.ShouldBind(&query)
	username := query.Username
	password := query.Password

	// 连接 mysql
	mysqlCon,_ := extend.MysqlInit()

	if username == "" {
		return "用户名或密码错误(1)"
	}

	var password_ string
	var salt_ string
	var userid string
	// 取原值
	mysqlCon.QueryRow("SELECT id,password,salt FROM admin_user WHERE status = 1 and username=?", username).Scan(&userid,&password_,&salt_)

	pswd := extend.Md5(extend.Md5(password)+salt_)
	if(pswd != password_){
		return "密码错误(2)"
	}
	
	// 此方法需要根据用户登录情况 创建token
	// 生成jwt-token
	token,_ := extend.GenerateToken(userid)

	// todo 存入redis

	return token
}

// 上传图片
func UploadM(c *gin.Context) interface{} {
    file, errLoad := c.FormFile("file")
    if errLoad != nil {
        msg := "获取上传文件错误"
        return msg
    }
 
    // 上传文件到指定的路径
    ret := make(map[string]string)
	// 名称确保唯一
	uuid := extend.Uuid()
    ret["fileName"] = uuid + file.Filename
    ret["fileNameOrigin"] = file.Filename
	
    // public/upload 文件夹下
    filePath := "./public/upload/"+ ret["fileName"]

    err := c.SaveUploadedFile(file, filePath)
    if err != nil {
        return err
    }
    
    ret["picUrl"] = ret["fileName"]
	return ret["picUrl"]
}

func LogoutM(c *gin.Context) interface{} {
	// 清除缓存
	// 清除redis
	// 清除cookie
	c.SetCookie("vue_admin_template_token", "", -1, "/", "*", false, true)
	return "退出成功."
}

func MenuM(c *gin.Context) interface{} {
	username,_ := c.Get("username")
	fmt.Println("username:",username)

	// 连接 mysql
	//mysqlCon,_ := extend.MysqlInit()
	mysqlCon,_ := extend.MysqlInit()

	type admin_user struct{
		id int
		username string
		name string
		email *string `db:"email" json:",omitempty"`
		phone *string `db:"phone" json:",omitempty"`
		role_id int
		status int
		created_at int
		updated_at int
	}

	var adminUser admin_user
	mysqlCon.QueryRow("SELECT id,username,name,email,phone,role_id,status,created_at,updated_at FROM admin_user WHERE status = 1 and id=?", username).Scan(
		&adminUser.id,
		&adminUser.username,
		&adminUser.name,
		&adminUser.email,
		&adminUser.phone,
		&adminUser.role_id,
		&adminUser.status,
		&adminUser.created_at,
		&adminUser.updated_at)
	

	t1 := make(map[string]interface{})
	t1["id"] = adminUser.id
	t1["username"] = adminUser.username
	t1["name"] = adminUser.name
	t1["email"] = adminUser.email
	t1["phone"] = adminUser.phone
	t1["role_id"] = adminUser.role_id
	t1["status"] = adminUser.status
	t1["created_at"] = adminUser.created_at
	t1["updated_at"] = adminUser.updated_at

	t_ := make(map[int]interface{})
	t_[0] = adminUser.username

	t1["roles"] = t_

	type admin_role struct{
		id int
		title string
		description string
		created_at int
		updated_at int
		status int
		sort int
		config string
	}
	var adminRole admin_role
	mysqlCon.QueryRow("SELECT id,title,description,created_at,updated_at,status,sort,config FROM admin_role WHERE status = 1 and id=?",adminUser.role_id).Scan(
		&adminRole.id,
		&adminRole.title,
		&adminRole.description,
		&adminRole.created_at,
		&adminRole.updated_at,
		&adminRole.status,
		&adminRole.sort,
		&adminRole.config)
	t2 := make(map[string]interface{})

	t2["id"] = adminRole.id
	t2["title"] = adminRole.title
	t2["description"] = adminRole.description
	t2["created_at"] = adminRole.created_at
	t2["updated_at"] = adminRole.updated_at
	t2["status"] = adminRole.status
	t2["sort"] = adminRole.sort
	if adminRole.config == ""{
		t2["config"] = 0
	}else{
		t2["config"] = adminRole.config
	}
	
	type admin_power struct{
		id int
		title *string `db:"title" json:",omitempty"`
		controller *string `db:"controller" json:",omitempty"`
		action *string `db:"action" json:",omitempty"`
		active *string `db:"active" json:",omitempty"`
		level *int `db:"level" json:",omitempty"`
		sort *int `db:"sort" json:",omitempty"`
		seo_description *string `db:"seo_description" json:",omitempty"`
		seo_title *string `db:"seo_title" json:",omitempty"`
		seo_keyword *string `db:"seo_keyword" json:",omitempty"`
		is_show *int `db:"is_show" json:",omitempty"`
		ico *string `db:"ico" json:",omitempty"`
		created_at *int `db:"created_at" json:",omitempty"`
		updated_at *int `db:"updated_at" json:",omitempty"`
		status *int `db:"status" json:",omitempty"`
	}
	var adminPower admin_power
	if t2["config"] == 0{
		sqlStr := "SELECT id,title,controller,action,active,level,sort,seo_description,seo_title,seo_keyword,is_show,ico,created_at,updated_at,status FROM admin_power WHERE status > 0 and level = 0 order by id asc"
		rows,_ := mysqlCon.Query(sqlStr)
		defer rows.Close()

		result_data := make([]map[string]interface{}, 0)
		for rows.Next() {
			t1 := make(map[string]interface{})
			rows.Scan(
				&adminPower.id,
				&adminPower.title,
				&adminPower.controller,
				&adminPower.action,
				&adminPower.active,
				&adminPower.level,
				&adminPower.sort,
				&adminPower.seo_description,
				&adminPower.seo_title,
				&adminPower.seo_keyword,
				&adminPower.is_show,
				&adminPower.ico,
				&adminPower.created_at,
				&adminPower.updated_at,
				&adminPower.status)

			t1["id"] = adminPower.id
			t1["controller"] = *adminPower.controller
			if *adminPower.is_show == 1{
				t1["hidden"] = false
			}else{
				t1["hidden"] = true
			}
			t1["name"] = adminPower.controller
			t1["status"] = true

			meta := make(map[string]interface{})
			meta["icon"] = "fa "+*adminPower.ico
			meta["title"] = adminPower.title
			t1["meta"] = meta

			// todo 愚蠢的多级嵌套 目前只支持二级
			sqlStr := "SELECT id,title,controller,action,active,level,sort,seo_description,seo_title,seo_keyword,is_show,ico,created_at,updated_at,status FROM admin_power WHERE status > 0 and level = ? order by id asc"
			rows,_ := mysqlCon.Query(sqlStr,adminPower.id)
			result_child := make([]map[string]interface{}, 0)
			for rows.Next() {
				t2 := make(map[string]interface{})
				rows.Scan(
					&adminPower.id,
					&adminPower.title,
					&adminPower.controller,
					&adminPower.action,
					&adminPower.active,
					&adminPower.level,
					&adminPower.sort,
					&adminPower.seo_description,
					&adminPower.seo_title,
					&adminPower.seo_keyword,
					&adminPower.is_show,
					&adminPower.ico,
					&adminPower.created_at,
					&adminPower.updated_at,
					&adminPower.status)
				t2["id"] = adminPower.id
				t2["controller"] = adminPower.controller
				if *adminPower.is_show == 1{
					t1["hidden"] = false
				}else{
					t1["hidden"] = true
				}
				t2["action"] = adminPower.action
				t2["name"] = *adminPower.controller + *adminPower.action
				t2["status"] = true

				meta1 := make(map[string]interface{})
				meta1["icon"] = "fa "
				meta1["title"] = adminPower.title
				t2["meta"] = meta1
				result_child = append(result_child, t2)
			}
			t1["children"] = result_child
			result_data = append(result_data, t1)

		}
		t1["menu"] = result_data
	}

	t1["role"] = t2
	
	defer mysqlCon.Close()
	return t1
}