package docker_applications

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	"os"
	"os/exec"
	"path"
	"strings"
)

type Application struct {
	File      string
	Dir       string
	DependsOn []string `mapstructure:"depends_on"`
}

type Config struct {
	Version      string
	Applications map[string]Application
}

func Execute(opts []string, config Config) error {
	args, err := config.Parse(funk.Head(opts).(string), funk.Tail(opts).([]string))
	if err != nil {
		return err
	}

	log.Infof("docker-compose %s", strings.Join(args, " "))

	cmd := exec.Command("docker-compose", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (c *Config) Parse(application string, opts []string) ([]string, error) {
	args, err := c.toArgs(application, []string{})
	if err != nil {
		return nil, err
	}

	return append(args, opts...), nil
}

func (c *Config) toArgs(application string, alreadyInArgs []string) ([]string, error) {
	configApplication, ok := c.Applications[application]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Undefined Application: %s", application))
	}

	dir := os.ExpandEnv(configApplication.Dir)
	if dir == "" {
		return nil, errors.New(fmt.Sprintf("Undefined directory in Application: %s", application))
	}

	file := os.ExpandEnv(configApplication.File)
	if file == "" {
		file = "docker-compose.yml"
	}

	var args []string
	args = append(args, "-f", path.Join(dir, file))

	alreadyInArgs = append(alreadyInArgs, application)

	for _, newApplication := range configApplication.DependsOn {
		if !funk.Contains(alreadyInArgs, newApplication) {
			newArgs, err := c.toArgs(newApplication, alreadyInArgs)
			if err != nil {
				return nil, err
			}

			args = append(args, newArgs...)
		}
	}

	return args, nil
}
