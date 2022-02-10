/**
* @Author:Tristan
* @Date: 2022/1/4 5:08 下午
 */

package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Conf Config

type LogName struct {
	StatLogFile        string `yaml:"stat_log_file"`
	ErrorLogFile       string `yaml:"error_log_file"`
	PanicLogFile       string `yaml:"panic_log_file"`
	StderrPanicLogFile string `yaml:"stderr_panic_log_file"`
}
type Config struct {
	LogDir        string   `yaml:"log_dir"`
	LogName       LogName  `yaml:"log_name"`
	GrpcPort      string   `yaml:"grpc_port"`
}

func NewConfig(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, &Conf)
	if err != nil {
		return fmt.Errorf("fileName %q err: %v", filename, err)
	}
	return nil
}

func init() {
	err := NewConfig("./config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
}
