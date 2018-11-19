package viper

type Facade struct {
	instance *Viper
}

func (c *Facade) LoadFile(path, filename, extension string) error {
	c.instance = New()
	c.instance.AddConfigPath(path)
	c.instance.SetConfigName(filename)
	c.instance.SetConfigType(extension)
	return c.instance.ReadInConfig()
}

func (c *Facade) LoadEnvVars(vars map[string]string) error {
	var err error
	for k, v := range vars {
		err = c.instance.BindEnv(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Facade) Get(key string) interface{} {
	return c.instance.Get(key)
}
