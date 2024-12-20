package mongodb

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(
	ctx context.Context,
	uri string,
	username string,
	password string,
) (*mongo.Client, error) {
	const mongoPrefix = "mongodb://"
	const mongoClusterPrefix = "mongodb+srv://"

	cred := fmt.Sprintf("%s:%s@", username, password)
	if strings.HasPrefix(uri, mongoPrefix) {
		if username != "" && !regexp.MustCompile(`^mongodb:\/\/.*:.*@`).MatchString(uri) {
			index := strings.Index(uri, mongoPrefix) + len(mongoPrefix)
			uri = uri[:index] + cred + uri[index:]
		}
	} else if strings.HasPrefix(uri, mongoClusterPrefix) {
		if username != "" && !regexp.MustCompile(`^mongodb+srv:\/\/.*:.*@`).MatchString(uri) {
			index := strings.Index(uri, mongoClusterPrefix) + len(mongoClusterPrefix)
			uri = uri[:index] + cred + uri[index:]
		}
	}

	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongodb: %w", err)
	}

	return client, nil
}
