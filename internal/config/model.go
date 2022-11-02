package config

type Config struct {
	Threads  int    `yaml:"Threads"`
	LogLevel int    `yaml:"LogLevel"`
	User     []User `yaml:"User"`
	Cron     string `yaml:"Cron"`
}

type User struct {
	BookList []int  `yaml:"BookList"`
	InfoAPI  string `yaml:"InfoAPI"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
}
