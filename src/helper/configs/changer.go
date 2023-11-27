package configs

import "strings"

func SetParam(file, name, value string) {
	confList := GetAllLines(file)
	config := new(ConfigLines)
	config.EnvFile = file
	isSet := false

	for _, line := range confList {
		if strings.TrimSpace(line) == "" || strings.TrimSpace(line)[:1] == "#" {
			config.AddRawLine(line)
		} else {
			opt := strings.Split(strings.TrimSpace(line), "=")
			if opt[0] == name {
				config.AddLine(opt[0], value)
				isSet = true
			} else {
				config.AddRawLine(line)
			}
		}
	}

	if !isSet {
		config.AddEmptyLine()
		config.AddLine(name, value)
	}

	if len(config.Lines) > 0 {
		config.SaveLines()
	}
}

func ChangeParamName(file string, names map[string]string) {
	confList := GetAllLines(file)
	config := new(ConfigLines)
	config.EnvFile = file

	for _, line := range confList {
		if strings.TrimSpace(line) == "" || strings.TrimSpace(line)[:1] == "#" {
			config.AddRawLine(line)
		} else {
			opt := strings.Split(strings.TrimSpace(line), "=")
			if newName, ok := names[opt[0]]; ok {
				config.AddLine(newName, opt[1])
			} else {
				config.AddRawLine(line)
			}
		}
	}

	if len(config.Lines) > 0 {
		config.SaveLines()
	}
}
