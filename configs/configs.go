package configs

import "github.com/spf13/viper"

const CONFIG_PATH string = "configs"
const CONFIG_NAME string = "app"
const CONFIG_TYPE string = "env"

type Config struct {
	AppHost string `mapstructure:"APP_HOST"`
	AppPort string `mapstructure:"APP_PORT"`

	DBDriver string `mapstructure:"DB_DRIVER"`
	DBHost   string `mapstructure:"DB_HOST"`
	DBPort   string `mapstructure:"DB_PORT"`
	DBUser   string `mapstructure:"DB_USER"`
	DBPass   string `mapstructure:"DB_PASSWORD"`
	DBNAME   string `mapstructure:"DB_NAME"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(CONFIG_PATH)
	viper.SetConfigName(CONFIG_NAME)
	viper.SetConfigType(CONFIG_TYPE)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
