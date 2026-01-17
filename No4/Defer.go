package No4

import "fmt"

type DB struct {
	Name string
}

func (db *DB) Close() error {
	fmt.Println("Close DB: ", db.Name)
	return nil
}

func DeferTest() {
	db := DB{Name: "DB-1"}
	defer db.Close()
	defer func(db *DB) {
		err := db.Close()
		if err != nil {

		}
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}(&db)

	db.Name = "DB-2"
	fmt.Println("Use DB: ", db.Name)
	panic("DB panic")
}
