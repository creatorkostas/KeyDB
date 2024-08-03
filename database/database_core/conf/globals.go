package conf

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Tools settings
var DB_filename string = "db.gob"
var Accounts_filename string = "accounts.gob"

// API settings
var Send_all_errors_in_requests bool = true

// DB setting
var Append_only_in_file bool = false
var Append_file string = "aof.txt"

const Append_size int = 64

type configs struct {
	DB_filename                 string `yaml:"DB_filename"`
	Accounts_filename           string `yaml:"Accounts_filename"`
	Send_all_errors_in_requests bool   `yaml:"Send_all_errors_in_requests"`
	Append_file                 string `yaml:"Append_file"`
	Append_only_in_file         bool   `yaml:"Append_only_in_file"`
	// Append_size string  `yaml:"MakefileName"`
}

func load_yaml_to_struct(yaml_path string, struct_data *configs) *configs {
	data, err := os.ReadFile(yaml_path)

	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, struct_data); err != nil {
		panic(err)
	}

	// print the fields to the console
	return struct_data
}

func Load_configs(path string) {
	var conf configs
	load_yaml_to_struct(path, &conf)
	DB_filename = conf.DB_filename
	Accounts_filename = conf.Accounts_filename
	Send_all_errors_in_requests = conf.Send_all_errors_in_requests
	Append_only_in_file = conf.Append_only_in_file
	Append_file = conf.Append_file
	// Append_size = conf.Append_size
}
