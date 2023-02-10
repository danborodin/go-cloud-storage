package mongo

import (
	"auth/src/configs"
	"auth/src/types"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DbMongo) AddUser(user types.User) error {
	user.Id = primitive.NewObjectID()
	coll := db.client.Database(configs.Conf.Mongo.Database).Collection(userCollection)

	cursor, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	db.l.InfoPrintln("user with ", cursor.InsertedID, "added")
	return nil
}

func (db *DbMongo) GetUserByUsername(username string) (*types.User, error) {
	coll := db.client.Database(configs.Conf.Mongo.Database).Collection(userCollection)

	filter := bson.D{{"username", username}}
	var user types.User
	err := coll.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DbMongo) UpdateUser(user types.User) error {
	coll := db.client.Database(configs.Conf.Mongo.Database).Collection(userCollection)

	filter := bson.D{{"username", user.Username}}
	_, err := coll.ReplaceOne(context.TODO(), filter, user)
	if err != nil {
		return err
	}
	return nil
}

func (db *DbMongo) DeleteUser(username string) error {
	coll := db.client.Database(configs.Conf.Mongo.Database).Collection(userCollection)

	filter := bson.D{{"username", username}}
	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (db *DbMongo) EmailExist(email string) (bool, error) {

	coll := db.client.Database(configs.Conf.Mongo.Database).Collection(userCollection)

	filter := bson.D{{"email", email}}
	var result bson.M
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (db *DbMongo) UsernameExist(username string) (bool, error) {
	coll := db.client.Database(configs.Conf.Mongo.Database).Collection(userCollection)

	filter := bson.D{{"username", username}}
	var result bson.M
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (db *DbMongo) CheckUserVerified(username string) (bool, error) {
	coll := db.client.Database(configs.Conf.Mongo.Database).Collection(userCollection)

	filter := bson.D{{"username", username}}
	var result types.User
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false, err
	}

	return result.Verified, nil
}

func (db *DbMongo) GetUnverifiedUsers() ([]types.User, error) {
	coll := db.client.Database(configs.Conf.Mongo.Database).Collection(userCollection)
	filter := bson.D{{"verified", false}}
	var users []types.User

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
