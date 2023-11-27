package dbops

import (
	"testing"
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
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Delete", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("avenss", "123")
	if err != nil {
		t.Errorf("error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("avenss")
	if pwd != "123" || err != nil {
		t.Errorf("error of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("avenss", "123")
	if err != nil {
		t.Errorf("error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("avenss")
	if err != nil {
		t.Errorf("error of RegetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

func AddNewVideo(aid int, name string) error
