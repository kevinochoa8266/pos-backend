package store_test

import (
	"github.com/kevinochoa8266/pos-backend/store"
	"testing"
)
var db, _ = store.GetConnection(dbUrl)

// func init() {
// 	name := "John"
// 	number := 9417160432
// 	address := "123 abc"
// 	db.Exec("INSERT INTO employee VALUES(?, ?, ?);")
// }

func TestGet(t *testing.T) {

	employeeStore := store.NewEmployeeStore(db)
	
	_, err := employeeStore.Get(1); if err != nil {
		t.Error("could not retrieve a employee with the given id.")
	}

}