package config

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // which database of redis
	Ip       string `mapstructure:"ip" json:"ip" yaml:"ip"`                   // ip address of server
	Port     string `mapstructure:"port" json:"port" yaml:"port"`             // port
	Password string `mapstructure:"password" json:"password" yaml:"password"` // password
}

func (r *Redis) GetRedisAddr() string {
	return r.Ip + ":" + r.Port
}
