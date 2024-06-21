package initialize

import (
	"fmt"
	"go-ecommerce/global"

	"github.com/spf13/viper"
)

// LoadConfig load configuration
func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config") // path to the config
	viper.SetConfigName("local")    // file name
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}
	// read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Server Port::", viper.GetString("security.jwt.key"))

	// configuration structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Println("Unable to decode configuration %v", err)
	}
}
