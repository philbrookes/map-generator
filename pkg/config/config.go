package config

import "fmt"

//Config stores configuration settings for the map generator
type Config struct {
	allowedOrigins []string
	allowedMethods []string
	port           int
}

// GetConfig returns an instance of a config
func GetConfig() *Config {
	config := Config{
		allowedOrigins: []string{"localhost", "127.0.0.1"},
		allowedMethods: []string{"GET", "PUT", "POST", "DELETE"},
		port:           8080,
	}

	return &config
}

//GetAllowedOrigins returns an array of allowed origins
func (c *Config) GetAllowedOrigins() []string {
	return c.allowedOrigins
}

//GetAllowedMethods returns an array of allowed methods
func (c *Config) GetAllowedMethods() []string {
	return c.allowedMethods
}

//GetPortListenerStr returns the port for listening on, formatted as ":<port>"
func (c *Config) GetPortListenerStr() string {
	return fmt.Sprintf(":%d", c.port)
}
