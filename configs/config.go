package configs

import "github.com/spf13/viper"

//cfg é um ponteiro
var cfg *config

//uma struct é similar com uma classe
type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

//função sempre será chamada no star das aplicações
//viper permite setar valores padrões, logo, se ocê não as informar, os valores serão os setados no init()
func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

//função para carregar as configurações
func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	//se o erro for diferente de não encontrei o arquivo, não tem problema, pq as informações foram setadas no init()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	//vai criar um ponteiro da struct
	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

//função para returnar o banco
func GetDB() DBConfig {
	return cfg.DB
}

//função para returnar só a porta
func GetServerPort() string {
	return cfg.API.Port
}
