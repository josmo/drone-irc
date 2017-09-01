package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

var build string

func main() {
	app := cli.NewApp()
	app.Name = "irc plugin"
	app.Usage = "irc plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.0+%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "prefix",
			Usage:  "prefix for notification",
			EnvVar: "PLUGIN_PREFIX",
			Value:  "build",
		},
		cli.StringFlag{
			Name:   "nick",
			Usage:  "nickname used by bot",
			EnvVar: "PLUGIN_NICK,",
		},
		cli.StringFlag{
			Name:   "channel",
			Usage:  "channel to post message in",
			EnvVar: "PLUGIN_CHANNEL,",
		},
		cli.StringFlag{
			Name:   "recipient",
			Usage:  "recipient",
			EnvVar: "PLUGIN_RECIPIENT",
		},
		cli.IntFlag{
			Name:   "irc-host",
			Usage:  "irc-host",
			EnvVar: "PLUGIN_IRC_HOST",
			Value:  6667,
		},
		cli.StringFlag{
			Name:   "irc-port",
			Usage:  "irc-port",
			EnvVar: "PLUGIN_IRC_PORT",
		},
		cli.StringFlag{
			Name:   "irc-password",
			Usage:  "irc-password",
			EnvVar: "PLUGIN_IRC_PASSWORD",
		},
		cli.BoolFlag{
			Name:   "irc-enable-tls",
			Usage:  "enable-tls",
			EnvVar: "PLUGIN_IRC_ENABLE_TLS",
		},
		cli.StringFlag{
			Name:   "template",
			Usage:  "Custom template for webhook",
			EnvVar: "PLUGIN_TEMPLATE",
		},
		cli.StringSliceFlag{
			Name:   "headers",
			Usage:  "Custom headers key map",
			EnvVar: "PLUGIN_HEADERS",
		},
		cli.StringSliceFlag{
			Name:   "urls",
			Usage:  "List of urls to perform the action on",
			EnvVar: "PLUGIN_URLS",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "For debug information",
			EnvVar: "PLUGIN_DEBUG",
		},
		cli.BoolFlag{
			Name:   "skip-verify",
			Usage:  "Skip ssl verification",
			EnvVar: "PLUGIN_SKIP_VERIFY",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.Int64Flag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Int64Flag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.Int64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_JOB_STARTED",
		},
		cli.StringFlag{
			Name:  "env-file",
			Usage: "source env file",
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {

	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Build: Build{
			Tag:     c.String("build.tag"),
			Number:  c.Int("build.number"),
			Event:   c.String("build.event"),
			Status:  c.String("build.status"),
			Commit:  c.String("commit.sha"),
			Ref:     c.String("commit.ref"),
			Branch:  c.String("commit.branch"),
			Author:  c.String("commit.author"),
			Message: c.String("commit.message"),
			Link:    c.String("build.link"),
			Started: c.Int64("build.started"),
			Created: c.Int64("build.created"),
		},
		Job: Job{
			Started: c.Int64("job.started"),
		},
		Config: Config{
			Prefix:       c.String("prefix"),
			Nick:         c.String("nick"),
			Channel:      c.String("channel"),
			Recipient:    c.String("recipient"),
			IRCHost:      c.String("irc-host"),
			IRCPort:      c.Int("irc-port"),
			IRCPassword:  c.String("irc-password"),
			IRCEnableTLS: c.Bool("irc-enable-tls"),
		},
	}
	return plugin.Exec()
}
