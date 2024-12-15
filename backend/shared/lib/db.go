package lib

import (
	"context"
	ent "root/database"

	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
)

func ConnectDatabase(log *logrus.Logger) (*ent.Client, error) {

	dns := "postgresql://postgres.drrsenlocmidswbgjwzc:qwertyqwest7q8q1579@aws-0-eu-central-1.pooler.supabase.com:6543/postgres"
	client, err := ent.Open("postgres", dns)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
		return nil, err
	}

	log.Info("connected to the database successfully")

	autoMigrate(client, log)

	log.Info("successfully created schema resources")

	return client, nil
}

func autoMigrate(client *ent.Client, log *logrus.Logger) {
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
