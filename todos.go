package main

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

var DB []TodoItem = make([]TodoItem, 0)
var PrimaryKey int = 1

func GetTodoItems() ([]TodoItem, error) {
	result := make([]TodoItem, 0)

	rows, err := db.Query("SELECT * FROM TODOS")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var item TodoItem

	for rows.Next() {
		rows.Scan(&item.TodoId, &item.Text, &item.Done)
		result = append(result, item)
	}

	return result, nil
}

func GetTodoItem(id int) (*TodoEditor, error) {
	result := &TodoEditor{}

	rows, err := db.Query("SELECT `Text`, `Done` FROM TODOS WHERE `TodoId`=?", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		rows.Scan(&result.Text, &result.Done)
	} else {
		return nil, errors.New("Item not found")
	}

	return result, nil
}

func CreateTodoItem(editor TodoEditor) error {
	_, err := db.Exec("INSERT INTO TODOS(`Text`, `Done`) VALUES(?, ?)", editor.Text, editor.Done)
	return err
}

func UpdateTodoItem(editor TodoEditor, id int) error {
	//UPDATE TODOS SET `Text`=:text, `Done`=:done WHERE `TodoId`=:id
	_, err := db.Exec("UPDATE TODOS SET `Text` = ?, `Done` = ? WHERE `TodoId`=?", editor.Text, editor.Done, id)
	return err
}

func DeleteTodoItem(id int) error {
	//DELETE FROM TODOS WHERE `TodoId`=:id
	_, err := db.Exec("DELETE FROM TODOS WHERE `TodoId`=?", id)
	return err
}
