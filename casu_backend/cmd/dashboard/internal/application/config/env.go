package config

import (
	"github.com/caarlos0/env/v6"
	"runtime"
	"time"
)

type Config struct {
	App struct {
		MyPodIP         string        `env:"MY_POD_IP"            envDefault:"127.0.0.1"`
		Domain          string        `env:"APP_DOMAIN"           envDefault:"http://localhost:3000"`
		Environment     string        `env:"APP_ENV"              envDefault:"development"`
		ShutdownTimeout time.Duration `env:"APP_SHUTDOWN_TIMEOUT" envDefault:"5s"`
		Secret          string        `env:"USER_SECRET"          envDefault:"secret"`
		ApiBaseURL      string        `env:"USER_BASE_URL"        envDefault:"http://localhost:3000"`
	}
	Debug struct {
		Host string `env:"DEBUG_HOST" envDefault:"0.0.0.0"`
		Port int    `env:"DEBUG_PORT" envDefault:"4001"`
	}
	Mailer struct {
		Host     string `env:"MAILER_HOST"     envDefault:"smtp.gmail.com"`
		Port     int    `env:"MAILER_PORT"     envDefault:"587"`
		User     string `env:"MAILER_USER"     envDefault:"seroka207@gmail.com"`
		Password string `env:"MAILER_PASSWORD" envDefault:"nuhn lyev wqht ezia"`
	}

	HTTP struct {
		Host string `env:"HOST"      envDefault:"0.0.0.0"`
		Port int    `env:"HTTP_PORT" envDefault:"5002"`
		// Origins should follow format: scheme "://" host [ ":" port ]
		Origins []string `env:"HTTP_ORIGINS" envSeparator:"|" envDefault:"*"`

		ReadTimeout  time.Duration `env:"HTTP_SERVER_READ_TIMEOUT"     envDefault:"5s"`
		WriteTimeout time.Duration `env:"HTTP_SERVER_WRITE_TIMEOUT"    envDefault:"10s"`
		IdleTimeout  time.Duration `env:"HTTP_SERVER_SHUTDOWN_TIMEOUT" envDefault:"120s"`
	}

	GRPC struct {
		Host string `env:"HOST"      envDefault:"0.0.0.0"`
		Port int    `env:"GRPC_PORT" envDefault:"6002"`

		LogName       string        `env:"LOG_NAME"      envDefault:"user"`
		ServerMinTime time.Duration `env:"GRPC_SERVER_MIN_TIME" envDefault:"5m"` // if a client pings more than once every 5 minutes (default), terminate the connection
		ServerTime    time.Duration `env:"GRPC_SERVER_TIME" envDefault:"2h"`     // ping the client if it is idle for 2 hours (default) to ensure the connection is still active
		ServerTimeout time.Duration `env:"GRPC_SERVER_TIMEOUT" envDefault:"20s"` // wait 20 second (default) for the ping ack before assuming the connection is dead
		ConnTime      time.Duration `env:"GRPC_CONN_TIME" envDefault:"10s"`      // send pings every 10 seconds if there is no activity
		ConnTimeout   time.Duration `env:"GRPC_CONN_TIMEOUT" envDefault:"20s"`   // wait 20 second for ping ack before considering the connection dead
	}
	MongoDB struct {
		User     string `env:"MONGO_USER"     envDefault:"admin"`
		Pass     string `env:"MONGO_PASS"     envDefault:"admin"`
		Host     string `env:"MONGO_HOST"     envDefault:"mongo"`
		Port     int    `env:"MONGO_PORT"     envDefault:"27017"`
		Database string `env:"MONGO_DATABASE" envDefault:"admin"`
	}

	AuthClient struct {
		Host   string `env:"AUTH_HOST" envDefault:"0.0.0.0"` // Auth service host
		Secret string `env:"AUTH_SECRET"                envDefault:"secret"`
		Port   int    `env:"AUTH_PORT" envDefault:"6000"`
	}
	Facebook struct {
		ClientID     string `env:"FACEBOOK_CLIENT_ID"`
		ClientSecret string `env:"FACEBOOK_CLIENT_SECRET"`
	}
	Google struct {
		ClientID     string `env:"GOOGLE_CLIENT_ID"`
		ClientSecret string `env:"GOOGLE_CLIENT_SECRET"`
	}
	CommandBus struct {
		QueueSize int `env:"COMMAND_BUS_BUFFER" envDefault:"100"`
	}
	EventBus struct {
		QueueSize int `env:"COMMAND_BUS_BUFFER" envDefault:"100"`
	}
}

func FromEnv() *Config {
	var c Config

	if err := env.Parse(&c.App); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.Debug); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.Mailer); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.HTTP); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.GRPC); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.MongoDB); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.AuthClient); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.CommandBus); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.EventBus); err != nil {
		panic(err)
	}

	if c.CommandBus.QueueSize == 0 {
		c.CommandBus.QueueSize = runtime.NumCPU()
	}
	if c.EventBus.QueueSize == 0 {
		c.EventBus.QueueSize = runtime.NumCPU()
	}

	return &c
}
