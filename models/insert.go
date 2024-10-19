package models

import "API/db"

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	//fechar a conexão
	defer conn.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
