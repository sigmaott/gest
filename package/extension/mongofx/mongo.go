package mongofx

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/fx"
)

func ForRoot(ctx context.Context, uri string, databaseName string) fx.Option {
	return fx.Module("mongofx", fx.Provide(func() *mongo.Client {
		return basicConnection(ctx, uri)
	}, func(client *mongo.Client) *mongo.Database {
		return ConnectToDatabase(ctx, client, databaseName)
	}))
}
func ConnectToDatabase(ctx context.Context, client *mongo.Client, databaseName string) (db *mongo.Database) {

	return client.Database(databaseName)
}

func basicConnection(ctx context.Context, uri string) (db *mongo.Client) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	return client
}
