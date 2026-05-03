package external

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/examples/fullstack/emigo"
	"github.com/urfave/cli/v3"
)

/**
* Configuration generator
 */
type Config struct {
	// Sample string
	SampleString string `envconfig:"SAMPLE_STRING" description:"Sample string"`
}

func GetConfigCliFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "sample-string",
			Usage: "Sample string",
		},
	}
}
func CastConfigFromCli(config *Config, c emigo.CliCastable) {
	if c.IsSet("sample-string") {
		config.SampleString = c.String("sample-string")
	}
}
func GetConfigCli() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "sample-string",
			Usage: "Sample string (string)",
			Commands: []*cli.Command{
				{
					Name: "get",
					Action: func(ctx context.Context, c *cli.Command) error {
						fmt.Println(config.SampleString)
						return nil
					},
				},
				{
					Name: "set",
					Action: func(ctx context.Context, c *cli.Command) error {
						return emigo.ConfigSetString(c, config.SampleString, func(value string) {
							config.SampleString = value
							config.Save(".env")
						})
						return nil
					},
				},
			},
		},
	}
}

// The config is usually populated by env vars on LoadConfiguration
var config Config = Config{}

/*
*
You can call this function on first line of your main function.
This is different from fireback configuration (for now), you can
define config: in module3 file, similar to fields in entities,
and we generate the config struct and this function would read .env.local,
.env.prod, etc - depending on the ENV=xxx env variable.
*
*/
func LoadConfiguration() Config {
	emigo.HandleEnvVars(&config)
	return config
}
func (x *Config) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
func (x *Config) Save(filepath string) error {
	return emigo.SaveEnvFile(x, filepath)
}
