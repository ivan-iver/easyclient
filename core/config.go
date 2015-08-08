package core

import (
	"path"

	"os"

	"github.com/nubleer/config"
)

type Config struct {
	*config.Config
	Settings map[string]string
}

func (self *Config) Init() (err error) {
	var pwd string
	if pwd, err = os.Getwd(); err != nil {
		panic(err)
	}
	file := path.Join(pwd, "app.conf")
	if self.Config, err = config.ReadDefault(file); err != nil {
		panic(err)
	}
	return self.loadSettings()
}

func (self *Config) loadSettings() (err error) {
	self.Settings = make(map[string]string)
	if self.Settings["url"], err = self.Config.String("paypal", "url"); err != nil {
		return err
	}
	if self.Settings["api"], err = self.Config.String("paypal", "api"); err != nil {
		return err
	}
	if self.Settings["user"], err = self.Config.String("paypal", "user"); err != nil {
		return err
	}
	if self.Settings["pwd"], err = self.Config.String("paypal", "pwd"); err != nil {
		return err
	}
	if self.Settings["signature"], err = self.Config.String("paypal", "signature"); err != nil {
		return err
	}
	return
}

func (self Config) Url() string {
	return self.Settings["url"]
}

func (self Config) Api() string {
	return self.Settings["api"]
}
func (self Config) User() string {
	return self.Settings["user"]
}
func (self Config) Pwd() string {
	return self.Settings["pwd"]
}

func (self Config) Signature() string {
	return self.Settings["signature"]
}
