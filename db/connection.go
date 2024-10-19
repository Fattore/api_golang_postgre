package db

import (
	"API/configs"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// função para apenas abrir a conexão, nada mais além disso
func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	//diz o tipo de banco de dados + as variaveis de conexão que conversão diretamente com as infos do package config
	conn, err := sql.Open("postgres", sc)

	//se der erro ele derruba tudo
	if err != nil {
		panic(err)
	}

	//faz uma pequena consulta no banco para verificar se está tudo certo
	err = conn.Ping()

	return conn, err
}
