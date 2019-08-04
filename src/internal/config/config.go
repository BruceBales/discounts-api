package config

import "os"

//Config defines variables required for Database connections
type Config struct {
	MysqlUser   string
	MysqlPass   string
	MysqlHost   string
	MysqlPort   string
	MysqlDriver string

	RedisPass string
	RedisHost string
	RedisPort string
}

//GetConfig fetches values from important environment variables and returns a populate config struct
func GetConfig() Config {
	//Each value in the Config will populate with a default if no environment variable is found
	mysqlDriver := os.Getenv("DB_DRIVER")
	if mysqlDriver == "" {
		mysqlDriver = "mysql"
	}
	mysqlUser := os.Getenv("MYSQL_USER")
	if mysqlUser == "" {
		mysqlUser = "root"
	}
	mysqlPass := os.Getenv("MYSQL_PASSWORD")
	if mysqlPass == "" {
		mysqlPass = "localtesting"
	}
	//If running local from outside docker-compose, host will be localhost.
	mysqlHost := os.Getenv("MYSQL_HOST")
	if mysqlHost == "" {
		mysqlHost = "localhost"
	}
	mysqlPort := os.Getenv("MYSQL_PORT")
	if mysqlPort == "" {
		mysqlPort = "3306"
	}

	//Redis Username and Password are intentionally blank on Localhost
	redisPass := os.Getenv("REDIS_PASS")
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}
	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}

	return Config{
		MysqlUser:   mysqlUser,
		MysqlPass:   mysqlPass,
		MysqlHost:   mysqlHost,
		MysqlPort:   mysqlPort,
		MysqlDriver: mysqlDriver,

		RedisPass: redisPass,
		RedisHost: redisHost,
		RedisPort: redisPort,
	}
}
