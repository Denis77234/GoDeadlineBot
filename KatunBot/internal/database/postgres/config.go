package postgres

import (
	"fmt"
)

type Config struct {
	Login string
	Password string
	Host string
	Database string
	Ssl string
}



func (c Config) ConnString() string {
	return fmt.Sprintf("user=%v password=%v host=%v database=%v sslmode=%v", c.Login,c.Password,c.Host,c.Database,c.Ssl)
}