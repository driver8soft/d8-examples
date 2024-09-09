package d8conf

import (
	"github.com/spf13/viper"
)

// struct to map cobol copy values
type cobC struct {
	Copy []CobCopy `mapstructure:"copy"`
}

type CobCopy struct {
	Field   string `mapstructure:"field"`
	Type    int    `mapstructure:"type"`
	Length  int    `mapstructure:"length"`
	Decimal int    `mapstructure:"decimal"`
	Sign    bool   `mapstructure:"sign"`
}

// struct to map env values
type envC struct {
	AppEnv       string `mapstructure:"APP_ENV"`
	CobolProgram string `mapstructure:"COBOL_PROGRAM"`
}

// Initilize variables to access values
var EnvC *envC
var CobC *cobC

func InitConfig() (err error) {

	// Tell viper the path/location of your env file
	viper.AddConfigPath("../conf")

	// Viper unmarshals env variables into the struct
	if err = loadConfig("app", "env"); err != nil {
		return
	}
	if err = viper.Unmarshal(&EnvC); err != nil {
		return
	}
	// Viper unmarshals cobol variables into the struct
	if err = loadConfig("cobcopy", "yaml"); err != nil {
		return
	}
	if err = viper.Unmarshal(&CobC); err != nil {
		return
	}
	return nil

}
func loadConfig(s, t string) error {
	viper.SetConfigName(s)
	viper.SetConfigType(t)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
