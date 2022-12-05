package config

type Timer struct {
	Start bool   `mapstructure:"start" json:"start" yaml:"start"` // whether to start the timer
	Spec  string `mapstructure:"spec" json:"spec" yaml:"spec"`    // cron spec
}
