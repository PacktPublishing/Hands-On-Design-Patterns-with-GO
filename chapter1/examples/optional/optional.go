package chapter1

type config struct {
	level       string
	environment string
	command     string
}

type option func(c *config) error

func optionLevel(level string) option {
	return func(c *config) error {
		c.level = level
		return nil
	}
}

func optionEnvironment(environment string) option {
	return func(c *config) error {
		c.environment = environment
		return nil
	}
}

func optionCommand(command string) option {
	return func(c *config) error {
		c.command = command
		return nil
	}
}

func parseOptions(options ...option) (*config, error) {
	conf := &config{
		level:       "low",
		environment: "home",
		command:     "none",
	}

	for _, option := range options {
		if err := option(conf); err != nil {
			return nil, err
		}
	}
	return conf, nil
}
