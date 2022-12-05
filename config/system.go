package config

import "strconv"

type System struct {
	Env          string `mapstructure:"env" json:"env" yaml:"env"`                               // environment value
	Mode         string `mapstructure:"mode" json:"mode" yaml:"mode"`                            // run mode
	Ip           string `mapstructure:"ip" json:"ip" yaml:"ip"`                                  // ip
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`                            // port
	OssType      string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`                // oss type
	UseRedis     bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`             // whether to use redis
	IpLimitCount int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"` // ip limit count
	IpLimitTime  int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`    // ip limit time
}

func (s *System) GetAddr() string {
	return s.Ip + ":" + strconv.Itoa(s.Port)
}
