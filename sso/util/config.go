package util

type Redis_Config struct {
	Host string
	Port string
}

var Redis = new(Redis_Config)

const LOG_FILE = "D:/sso.log"

func init() {
	//Redis.Host = "10.211.55.3"
	//Redis.Port = "6379"

	Redis.Host = "127.0.0.1"
	Redis.Port = "6379"

	//Redis.Host = os.Getenv("Redis_Host")
	//Redis.Port = os.Getenv("Redis_Port")

}
