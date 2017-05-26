package util

import "log"

type DB_Config struct {
	Db_host     string
	Db_port     string
	Db_user     string
	Db_password string
	Db_database string
	Db_max_open int
	Db_max_idle int
}

type Server_Config struct {
	SSO                  string
	SSO_Login            string
	SSO_Service_Validate string
	OAuth                string
	OAuth_CAS_Check      string
}

type Redis_Config struct {
	Host string
	Port string
}

var DB = new(DB_Config)

var Server = new(Server_Config)

var Redis = new(Redis_Config)

const LOG_FILE = "D:/oauth.log"

func init() {
	DB.Db_host = "127.0.0.1"
	DB.Db_port = "3306"
	DB.Db_user = "root"
	DB.Db_password = "1234"
	DB.Db_database = "oauth"

	DB.Db_max_open = 100
	DB.Db_max_idle = 20




	Server.SSO = "http://test.yourhost.com:5000"
	Server.SSO_Login = "http://test.yourhost.com:5000/login"
	Server.SSO_Service_Validate = "http://test.yourhost.com:5000/serviceValidate"
	Server.OAuth = "http://test.yourhost.com:3000"
	Server.OAuth_CAS_Check = "http://test.yourhost.com:3000/cas_check"

	Redis.Host = "127.0.0.1"
	Redis.Port = "6379"

	log.Println("DB_CONFIG:"+DB.Db_host)
	log.Println("DB_CONFIG:"+DB.Db_port)
	log.Println("DB_CONFIG:"+DB.Db_user)
	log.Println("DB_CONFIG:"+DB.Db_password)
	log.Println("DB_CONFIG:"+DB.Db_database)
	log.Println("DB_CONFIG:"+string(DB.Db_max_open))
	log.Println("DB_CONFIG:"+string(DB.Db_max_idle))

	log.Println("SERVER_CONFIG:"+Server.SSO)
	log.Println("SERVER_CONFIG:"+Server.SSO_Login)
	log.Println("SERVER_CONFIG:"+Server.SSO_Service_Validate)
	log.Println("SERVER_CONFIG:"+Server.OAuth)
	log.Println("SERVER_CONFIG:"+Server.OAuth_CAS_Check)

	log.Println("REDIS_CONFIG:"+Redis.Host)
	log.Println("REDIS_CONFIG:"+Redis.Port)

}
