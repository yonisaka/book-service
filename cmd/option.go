package cmd

import (
	"github.com/yonisaka/book-service/config"
	"github.com/yonisaka/book-service/domain/service"
)

// CommandOption is an option type
type CommandOption func(c *Command)

// WithConfig is a function of command option
func WithConfig(conf *config.Config) CommandOption {
	return func(c *Command) {
		c.conf = conf
	}
}

// WithRepo is a function of command option
func WithRepo(repo *service.Repositories) CommandOption {
	return func(c *Command) {
		c.repo = repo
	}
}
