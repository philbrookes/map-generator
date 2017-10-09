package config

import "fmt"

//Config stores configuration settings for the map generator
type Config struct {
	allowedOrigins []string
	allowedMethods []string
	port           int
	buffer         int
	viewport       int
}

// GetConfig returns an instance of a config
func GetConfig() *Config {
	config := Config{
		allowedOrigins: []string{"localhost", "127.0.0.1", "http://127.0.0.1:3000"},
		allowedMethods: []string{"GET", "PUT", "POST", "DELETE"},
		port:           8080,
		buffer:         15,
		viewport:       25,
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

//GetBufferSize returns the size of the buffer zone around the visible map
func (c *Config) GetBufferSize() int {
	return c.buffer
}

//GetViewportSize returns the size of the height and wifth of the visible map
func (c *Config) GetViewportSize() int {
	return c.viewport
}
