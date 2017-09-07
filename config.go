package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

//Conf variable fog start config params init
var Conf GPIONames

//GPIONames global config string
type GPIONames struct {
	In  string
	Out string
}

//Init config params initialization
func (c *GPIONames) Init() {
	fmt.Println("Configuration variables initialization...")

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error: %s \n", err)
		os.Exit(1)
	}
	c.In = viper.GetString("gpio.in")
	c.Out = viper.GetString("gpio.out")

	fmt.Println("[Tournament.Config] Initializad")
}
