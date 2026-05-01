package emigo

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

func structToEnvMap(config interface{}) (map[string]string, error) {
	envMap := make(map[string]string)
	val := reflect.ValueOf(config).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		envKey := field.Tag.Get("envconfig")

		if envKey != "" {
			value := val.Field(i).Interface()

			switch v := value.(type) {
			case string:
				envMap[envKey] = v
			case bool:
				envMap[envKey] = strconv.FormatBool(v)
			case int, int64:
				envMap[envKey] = strconv.FormatInt(reflect.ValueOf(v).Int(), 10)
			case float64:
				envMap[envKey] = strconv.FormatFloat(reflect.ValueOf(v).Float(), 'f', -1, 64)
			default:
				return nil, fmt.Errorf("unsupported type: %s", reflect.TypeOf(v))
			}
		}
	}

	return envMap, nil
}

func SaveEnvFile(config interface{}, filename string) error {

	envMap, err := structToEnvMap(config)

	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for key, value := range envMap {
		if value == "" {
			continue
		}

		_, err := file.WriteString(fmt.Sprintf("%s=%s\n", key, value))
		if err != nil {
			return err
		}
	}

	return nil
}

func HandleEnvVars(spec interface{}) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	filename := ".env." + env
	err := godotenv.Load(filename)
	if err != nil {
		log.Printf("environment variable file expected: %s was not loaded. Error: %v", filename, err)
	}

	envconfig.MustProcess("", spec)
}

func ConfigSetBoolean(c *cli.Context, currentValue bool, setValue func(value bool)) error {
	if len(c.Args()) > 0 {
		var value bool = false
		read := c.Args()[0]
		if read == "true" || read == "1" || read == "yes" {
			value = true
		} else if read == "false" || read == "0" || read == "no" {
			value = false
		} else {
			return errors.New("the value for boolean needs to be true, false, 0, 1, yes, no")
		}

		setValue(value)
	} else {
		curr := "unknown"
		if currentValue {
			curr = "true"
		} else {
			curr = "false"
		}
		result := AskForSelect("Set the value to? Current value: "+curr, []string{"true", "false"})

		if result == "true" {
			setValue(true)
		}
		if result == "false" {
			setValue(false)
		}
	}

	return nil
}

func ConfigSetString(c *cli.Context, currentValue string, setValue func(value string)) error {
	if len(c.Args()) > 0 {
		var value string = c.Args()[0]
		setValue(value)
	} else {
		result := AskForInput("Set the value to?", currentValue)
		setValue(result)
	}

	return nil
}

func ConfigSetInt64(c *cli.Context, currentValue int64, setValue func(value int64)) error {
	if len(c.Args()) > 0 {
		var value string = c.Args()[0]

		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			setValue(intValue)
		}

	} else {
		result := AskForInput("Set the value to?", strconv.FormatInt(currentValue, 10))
		intValue, err := strconv.ParseInt(result, 10, 64)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			setValue(intValue)
		}
	}

	return nil
}

func ConfigSetInt(c *cli.Context, currentValue int, setValue func(value int)) error {
	if len(c.Args()) > 0 {
		var value string = c.Args()[0]

		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			setValue(int(intValue))
		}

	} else {
		result := AskForInput("Set the value to?", strconv.FormatInt(int64(currentValue), 10))
		intValue, err := strconv.ParseInt(result, 10, 64)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			setValue(int(intValue))
		}
	}

	return nil
}

func ConfigSetFloat64(c *cli.Context, currentValue float64, setValue func(value float64)) error {
	if len(c.Args()) > 0 {
		var value string = c.Args()[0]

		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			setValue(floatValue)
		}

	} else {
		result := AskForInput("Set the value to?", fmt.Sprintf("%f", currentValue))
		floatValue, err := strconv.ParseFloat(result, 64)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			setValue(floatValue)
		}
	}

	return nil
}

func AskForSelect(label string, items []string) string {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		if err.Error() == "^C" {
			os.Exit(35)
			return ""
		}
		return ""
	}

	index := strings.Index(result, ">>>")
	if index <= 0 {
		return result
	}
	return strings.Trim(result[0:index], " ")

}

func AskForInput(label string, defaultV string) string {
	validate := func(input string) error {
		if input == "" {
			return errors.New("this is necessary")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    label,
		Validate: validate,
		Default:  defaultV,
	}

	value, err := promptVariable.Run()
	if err != nil {
		if err.Error() == "^C" {
			os.Exit(35)
			return ""
		}
		return ""
	}

	return value
}
