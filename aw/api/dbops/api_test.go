package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate sessions")
	dbConn.Exec("truncate comments")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	//subTest方法
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Delete", testDeleteUser)
	//t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("avenssi", "123")
	if err != nil {
		t.Errorf("Error of AddUser : %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if err != nil {
		t.Errorf("Error of GetUser : %v", err)
	}

	if pwd != "123" {
		t.Errorf("Password mismatch, expected '123', got '%s'", pwd)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("avenssi", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser : %v", err)
	}
}

var temvid string

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddVideo", testAddNewVideo)
	t.Run("GetVideo", testGetVideo)
	t.Run("DeleteVideo", testDeleteVideo)

}

func testAddNewVideo(t *testing.T) {
	vi, err := AddNewVideo(1, "my_video")
	if err != nil {
		t.Errorf("Error of AddNewVideo : %v", err)
	}
	temvid = vi.Id
}

func testGetVideo(t *testing.T) {
	_, err := GetVideoInfo(temvid)
	if err != nil {
		t.Errorf("Error of GetVideo : %v", err)
	}
}

func testDeleteVideo(t *testing.T) {
	err := DeleteVideoInfo(temvid)
	if err != nil {
		t.Errorf("Error of DeleteVideo : %v", err)
	}
}

// comments的test
func TestCommentWorkFlow(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddComment)
	t.Run("ListComments", testListComments)

}
func testAddComment(t *testing.T) {
	//func AddNewComments(aid int, content string, vid int) error {
	err := AddNewComments(1, "I like this video", "12345")
	if err != nil {
		t.Errorf("Error of AddNewComments : %v", err)
	}
}
func testListComments(t *testing.T) {
	//func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments : %v", err)
	}

	for id, elseInfo := range res {
		fmt.Printf("comment: %d , %v \n", id, elseInfo)
	}
}
