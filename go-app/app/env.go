package main

import (
	"github.com/caarlos0/env/v6"
)

type Env struct {
	Addr     string `env:"ADDRESS,required"`
	MySQLURL string `env:"MYSQL_URL,required"`
}

func (e *Env) Parse() error {
	if err := env.Parse(e); err != nil {
		return err
	}
	return nil
}
