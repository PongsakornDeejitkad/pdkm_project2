package utils

import (
	"log"
	"strings"
	"unicode"

	"github.com/spf13/viper"
)

func InitViper() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read in viper config:%s", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

}

func ViperGetString(path string) string {
	return viper.GetString(path)
}

func ViperGetInt(path string) int {
	return viper.GetInt(path)
}

func ViperGetFloat(path string) float64 {
	return viper.GetFloat64(path)
}

func CheckLanguage(text string) bool {
	for _, char := range text {
		if !unicode.IsOneOf([]*unicode.RangeTable{unicode.Thai, unicode.Latin, unicode.Space, unicode.Number}, char) {
			return false
		}
	}
	return true
}
