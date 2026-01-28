package config

type DbConfig struct {
	Host     string
	Port     string
	Db       string
	Username string
	Password string
}

func newDb() DbConfig {
	return DbConfig{
		Host:     getEnv("DB_HOST", "127.0.0.1"),
		Port:     getEnv("DB_PORT", "5432"),
		Db:       getEnv("DB_DATABASE", ""),
		Username: getEnv("DB_USERNAME", ""),
		Password: getEnv("DB_PASSWORD", ""),
	}
}
