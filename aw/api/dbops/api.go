package dbops

import (
	"awesomeProject4/api/defs"
	"awesomeProject4/api/utils"
	"database/sql"
	_ "database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func AddUserCredential(loginName string, password string) error {
	// 记录日志，表示开始执行插入操作
	log.Printf("Adding user: %s with password: %s", loginName, password)

	// 预编译 SQL 语句
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		log.Printf("Error preparing SQL statement: %v", err) // 如果预编译失败，记录错误
		return err
	}
	defer stmtIns.Close() // defer 在函数结束时自动关闭

	// 执行 SQL 语句插入用户数据
	_, err = stmtIns.Exec(loginName, password)
	if err != nil {
		log.Printf("Error executing SQL statement: %v", err) // 如果执行插入失败，记录错误
		return err
	}

	log.Printf("Successfully added user: %s", loginName)
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	// 记录日志，表示开始获取用户密码
	log.Printf("Fetching password for user: %s", loginName)

	// 预编译 SQL 语句
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %v", err) // 记录 SQL 预编译错误
		return "", err
	}
	defer stmtOut.Close()

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error querying password: %v", err) // 记录查询错误
		return "", err
	}

	log.Printf("Successfully fetched password for user: %s, password: %s", loginName, pwd)
	return pwd, nil
}

func DeleteUser(login_name string, pwd string) error { //删
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd =?")
	if err != nil {
		log.Printf("DeleteUser error :  %s", err)
		return err
	}

	defer stmtDel.Close()
	//执行

	_, err = stmtDel.Exec(login_name, pwd)
	if err != nil {
		log.Printf("DeleteUser error :  %s", err)
		return err
	}
	return nil
}

// 添加新视频的函数，输入aid,name ，返回一个object和err
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	//create uuid
	//utils包下的type,全局的辨识符号
	vid, err := utils.NewUUID()
	if err != nil {
		log.Printf("AddNewVideo error :  %s", err)
		return nil, err
	}
	//视频创建的时间
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	//添加到数据库里面,预编译
	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info 
		(id,author_id,name,display_ctime) VALUES (?,?,?,?)`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %v", err)
		return nil, err
	}
	//执行sql语句
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		log.Printf("Error executing SQL statement: %v", err)
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: ctime}
	defer stmtIns.Close()
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	//根据authorIde find the videos
	stmtOut, err := dbConn.Prepare("SELECT author_id,name,display_ctime FROM video_info WHERE id = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %v", err)
		return nil, err
	}
	defer stmtOut.Close()
	var author_id string
	var name string
	var display_ctime string
	err = stmtOut.QueryRow(vid).Scan(&author_id, &name, &display_ctime) //这里还有no row的错误
	if err != nil && err == sql.ErrNoRows {
		log.Printf("Error querying video_info: %v", err)
	}
	res := &defs.VideoInfo{
		Id:           author_id,
		Name:         name,
		DisplayCtime: display_ctime,
	}
	return res, nil
}

// 删除视频
func DeleteVideoInfo(vid string) error {
	stmtDele, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %v", err)
	}
	defer stmtDele.Close()
	_, err = stmtDele.Exec(vid)
	if err != nil {
		log.Printf("Error executing SQL statement: %v", err)
		return err
	}
	return nil
}

// 添加新评论
func AddNewComments(aid int, content string, vid string) error {
	id, err := utils.NewUUID()
	if err != nil {
		log.Printf("AddNewComments error :  %s", err)
		return err
	}
	stmtIns, err := dbConn.Prepare("INSERT INTO comments(id,video_id,author_id,content) VALUES (?,?,?,?)")
	if err != nil {
		log.Printf("Error preparing SQL statement: %v", err)
		return err
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		log.Printf("Error executing SQL statement: %v", err)
		return err
	}
	return nil
}

// 展示comments
func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(` SELECT comments.id, users.Login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil
}
