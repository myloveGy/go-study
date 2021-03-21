package config

type MySqlConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	DbName   string `ini:"db"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Dsn      string `ini:"dsn"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
}

type Config struct {
	AppName      string `ini:"app_name"`
	Debug        bool   `ini:"debug"`
	*MySqlConfig `ini:"mysql"`
	*RedisConfig `ini:"redis"`
}
