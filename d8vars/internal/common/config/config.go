package d8conf

import (
	"bytes"
	"log"

	"github.com/spf13/viper"
)

var (
	zeroBytes  = []byte{0x30}
	spaceBytes = []byte{0x20}
	lowBytes   = []byte{0x00}
)

// struct to map cobol copy values
type cobCopy struct {
	Copy []Var `mapstructure:"copy"`
}

type Var struct {
	Field   string `mapstructure:"field"`
	Type    int    `mapstructure:"type"`
	Length  int    `mapstructure:"length"`
	Decimal int    `mapstructure:"decimal"`
	Sign    bool   `mapstructure:"sign"`
}

// struct to map env values
type env struct {
	AppEnv       string `mapstructure:"APP_ENV"`
	CobolProgram string `mapstructure:"COBOL_PROGRAM"`
	CobolConfig  string `mapstructure:"COBOL_CONFIG"`
}
type cobol struct {
	Commarea bytes.Buffer
	Cvars    map[string][]int
}

// Initilize variables to access values
var Env *env
var Cobol cobol

func InitConfig() {
	var c cobCopy

	// Tell viper the path/location of your env file
	viper.AddConfigPath("../conf")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error reading env file: ", err)
	}
	// Viper unmarshals env variables into the struct
	if err := viper.Unmarshal(&Env); err != nil {
		log.Fatal("error reading env file: ", err)
	}

	// Load cobol copy values from yaml file (default: cobcopy.yaml)
	if Env.CobolConfig == "" {
		viper.SetConfigName("cobcopy")
	} else {
		viper.SetConfigName(Env.CobolConfig)
	}
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error reading cobol config file: ", err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatal("error reading cobol config file: ", err)
	}

	// Initialize Cvars
	Cobol.Cvars = make(map[string][]int)

	// load cobol copy values
	loadCobolConfig(&c)

	log.Println("loading variables Env:", Env.AppEnv)

}
func loadCobolConfig(config *cobCopy) {

	// loop through the copy values
	for i, counter := 0, 0; i < len(config.Copy); i++ {
		field := config.Copy[i].Field
		copyItem := config.Copy[i]
		fieldVars := Cobol.Cvars[field]
		// append values to map Cvars
		fieldVars = append(fieldVars, copyItem.Type)
		fieldVars = append(fieldVars, copyItem.Length)
		fieldVars = append(fieldVars, copyItem.Decimal)
		if copyItem.Sign {
			fieldVars = append(fieldVars, 1)
		} else {
			fieldVars = append(fieldVars, 0)
		}
		// add variable offset to map Cvars
		fieldVars = append(fieldVars, counter)
		Cobol.Cvars[field] = fieldVars

		// initialize commarea with default values
		switch copyItem.Type {
		// display initialized with zeroes
		case 0:
			counter = counter + copyItem.Length
			for i := 0; i < copyItem.Length; i++ {
				Cobol.Commarea.Write(zeroBytes[:])
			}
		// comp1 initialized with low values
		case 1:
			counter = counter + 4
			for i := 0; i < 4; i++ {
				Cobol.Commarea.Write(lowBytes[:])
			}
		// comp2 initialized with low values
		case 2:
			counter = counter + 8
			for i := 0; i < 8; i++ {
				Cobol.Commarea.Write(lowBytes[:])
			}
		// comp3 initialized with low values
		case 3:
			counter = counter + ((copyItem.Length / 2) + 1)
			for i := 0; i < ((copyItem.Length / 2) + 1); i++ {
				Cobol.Commarea.Write(lowBytes[:])
			}
		// comp, binary, comp4, comp5 initialized with low values
		case 4, 5:
			if copyItem.Length > 0 && copyItem.Length < 5 {
				counter = counter + 2
				for i := 0; i < 2; i++ {
					Cobol.Commarea.Write(lowBytes[:])
				}
			}
			if copyItem.Length > 4 && copyItem.Length < 10 {
				counter = counter + 4
				for i := 0; i < 4; i++ {
					Cobol.Commarea.Write(lowBytes[:])
				}
			}
			if copyItem.Length > 9 && copyItem.Length < 19 {
				counter = counter + 8
				for i := 0; i < 8; i++ {
					Cobol.Commarea.Write(lowBytes[:])
				}
			}
		// char initialized with space values
		case 9:
			counter = counter + copyItem.Length
			for i := 0; i < copyItem.Length; i++ {
				Cobol.Commarea.Write(spaceBytes[:])
			}
		default:
			log.Fatal("variable %s type %v not implemented", field, copyItem.Type)
		}
		// add variable offset to map Cvars
		Cobol.Cvars[field] = append(Cobol.Cvars[field], counter)
	}

}
