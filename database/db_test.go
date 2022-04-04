package database

import "testing"

func TestConnectDB(t *testing.T) {
	ConnectDB()
	defer CloseDB()
}
