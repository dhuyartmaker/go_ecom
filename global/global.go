package global

type ConfigStruct struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"databases"`
}

var (
	Config ConfigStruct
)