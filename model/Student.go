package model

import (
	"time"
)

//id, first_name, last_name, birthday
type Student struct {
	Id int
	FirstName	string
	LastName	string
	Birthday time.Time
}

func FindAll() ([]Student, error) {
	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]Student, 0)
	for rows.Next() {
		item := Student{}
		err := rows.Scan(&item.Id, &item.FirstName, &item.LastName, &item.Birthday)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}