package db

import (
	"log"
	"time"

	"github.com/yashdiq/nitip_user-service/pkg/config"

	"github.com/gocql/gocql"
)

type Config struct {
	Hosts        []string
	Port         int
	Username     string
	Password     string
	ProtoVersion int
	Consistency  string
	Keyspace     string
	Timeout      time.Duration
}

func NewConfig(cfg config.Config) Config {
	return Config{
		Hosts:        []string{cfg.CASSANDRA_URL},
		Port:         cfg.CASSANDRA_PORT,
		Username:     cfg.CASSANDRA_USER,
		Password:     cfg.CASSANDRA_PASS,
		ProtoVersion: 4,
		Consistency:  "Quorum",
		Keyspace:     cfg.CASSANDRA_KEYSPACE,
		Timeout:      time.Second * 5,
	}
}

func DBConnect() (*gocql.Session, error) {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load the Congiguration File: ", err)
	}

	var cfg = NewConfig(c)

	cluster := gocql.NewCluster(cfg.Hosts...)
	
	cluster.Port = cfg.Port
	cluster.ProtoVersion = cfg.ProtoVersion
	cluster.Keyspace = cfg.Keyspace
	cluster.Consistency = gocql.ParseConsistency(cfg.Consistency)
	cluster.Timeout = cfg.Timeout
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cfg.Username,
		Password: cfg.Password,
	}

	return cluster.CreateSession()
}
