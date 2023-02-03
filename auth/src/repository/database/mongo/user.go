package mongo

import (
	"auth/src/configs"
	"auth/src/types"
	"context"
	"time"

	"github.com/danborodin/go-logd"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	userCollection = "users"
)

type DbMongo struct {
	l      *logd.Logger
	client *mongo.Client
}

func NewDB(l *logd.Logger) *DbMongo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.Conf.Mongo.Uri))
	if err != nil {
		l.ErrPrintln(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		l.ErrPrintln(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		l.ErrPrintln(err)
	}

	return &DbMongo{
		l:      l,
		client: client,
	}
}

func (db *DbMongo) Disconnect() {
	err := db.client.Disconnect(context.TODO())
	if err != nil {
		db.l.ErrPrintln(err) // maybe change this
	}
}

func (db *DbMongo) AddUser(user *types.User) error {
	user.Id = primitive.NewObjectID()
	coll := db.client.Database(configs.Conf.Mongo.Database).Collection(userCollection)

	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	db.l.InfoPrintln("user with ", result.InsertedID, "added")
	return nil
}
