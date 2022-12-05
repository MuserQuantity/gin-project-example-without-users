package config

type Server struct {
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Timer  Timer  `mapstructure:"timer" json:"timer" yaml:"timer"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
}
