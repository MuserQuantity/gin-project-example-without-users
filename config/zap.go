package config

type Zap struct {
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // prefix
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // log level
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // log format
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                  // log file save path
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // encode level
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // stacktrace key
	MaxAge        int    `mapstructure:"mintage" json:"max-age" yaml:"max-age"`                      // save file max age
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // show line
	LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // log in console
}
