package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/gocql/gocql"
)

var Pdb *sql.DB

var Cdb *gocql.ClusterConfig

func init() {
	var err error

	//Postgres Connection
	Pdb, err = sql.Open("postgres", "postgres://user007:yV2kYebKVEX7bi637GOU@localhost/gusers?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = Pdb.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your PostgreSQL database.")

	//Cassandra Cluster Config

	// connect to the cluster
	Cdb = gocql.NewCluster("127.0.0.1")
	Cdb.Keyspace = "explorer"

}
