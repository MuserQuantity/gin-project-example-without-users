package config

type Mysql struct {
	Ip           string `mapstructure:"ip" json:"ip" yaml:"ip"`                                     // ip address of server
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               // port
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // advanced configuration
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                      // database name
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // database username
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // database password
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // max idle connections
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // max open connections to database
	LogLevel     string `mapstructure:"log-level" json:"log-level" yaml:"log-level"`                // whether to open Gorm global log
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // whether to write log files through zap
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Ip + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogLevel
}
