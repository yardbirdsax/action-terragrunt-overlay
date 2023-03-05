/*
Package config is used to store configuration data for the Action.
*/
package config

import (
	"github.com/sethvargo/go-githubactions"
	"github.com/yardbirdsax/action-terragrunt/internal/interfaces/github"
)

const (
	ActionInputBaseDirectory    string = "base-directory"
	ActionInputTerraformCommand string = "terraform-command"
	ActionInputToken            string = "token"
)

// Config is a struct that contains the required elements for configuring the Action.
type Config struct {
	// The GitHub Action Context associated with the run
	gitHubContext *githubactions.GitHubContext
	// The path where the base configuration resides
	baseDirectory string
	// The Terragrunt command to run
	command string
	// The GitHub token to use when interacting with the API
	token string
}

// configOptsFn is used for functional options operating on the Config struct.
type configOptsFn func(*Config)

// NewConfig initializes a new Config object from an Action struct.
func NewConfig(action github.Action, optFns ...configOptsFn) (*Config, error) {
	config := &Config{}
	context, _ := action.Context()
	config.gitHubContext = context
	config.token = action.GetInput(ActionInputToken)
	config.baseDirectory = action.GetInput(ActionInputBaseDirectory)
	config.command = action.GetInput(ActionInputTerraformCommand)
	for _, f := range optFns {
		f(config)
	}
	return config, nil
}

// BaseDirectory gets the base directory for the struct.
func (c *Config) BaseDirectory() string {
	return c.baseDirectory
}

func (c *Config) Command() string {
	return c.command
}

func (c *Config) GitHubContext() githubactions.GitHubContext {
	return *c.gitHubContext
}
