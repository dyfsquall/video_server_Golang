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
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("ADD", testAddUser)
	t.Run("GET", testGetUser)
	t.Run("DELETE", testDeleteUser)
	t.Run("REGET", testReGetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("squall", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("squall")
	// fmt.Println("PWD: ----", pwd)
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
	return

}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("squall", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testReGetUser(t *testing.T) {
	pwd, err := GetUserCredential("squall")
	if err != nil {
		t.Errorf("Error of ReGetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("ReGetUser: Deleteing user test failed!")
	}

}

func TestVideoWorkflow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
}

var tempvid string

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddNewVideo: %v", err)
	}
	tempvid = vi.ID
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video !"
	err := AddNewComments(vid, aid, content)
	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	//FormatInt: 将int 转化为字符串  //UnixNano : unix 时间： 1970年1月1日
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("error of ListComments: %v", err)
	}

	for i, ele := range res {
		// fmt.Printf("comment: %d, %v \n", i, ele)
		fmt.Println("comment: ", i, ele)
	}
}
