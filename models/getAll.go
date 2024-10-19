package models

import "API/db"

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	//fechar a conexão
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}

	//vai andar por todos os items retornados e faz o scan de cada um deles
	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}

		//se não der error algum, vai juntar tudo em uma lista e vai retornar pro handler
		todos = append(todos, todo)
	}

	return
}
