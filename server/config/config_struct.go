package config

type Config struct {
	Mysql Mysql
}

type Mysql struct {
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	DBname     string `mapstructure:"dbname"`
	Parameters string `mapstructure:"parameters"`
}

func DSN(m Mysql) string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DBname + "?" + m.Parameters
}
