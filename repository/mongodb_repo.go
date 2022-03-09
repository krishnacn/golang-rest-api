package repository

import (
	"context"

	"golang-rest-api/model"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repo struct{}

func NewMongoDBRepository() PostRepository {
	return &repo{}
}

const (
	// TODO: use secrets manager to retrieve password
	mongodbUri     = "mongodb+srv://admin:csXl105sGQgtUl6g@cluster0.q2w6w.mongodb.net/dev?retryWrites=true&w=majority"
	databaseName   = "dev"
	collectionName = "posts"
)

func (*repo) FindAll() ([]model.Post, error) {

	// connect to mongodb
	ctx, client, err := connectToMongoDb(mongodbUri)
	if err != nil {
		logrus.Errorf("unable to connect to mongodb. error: %v", err)
		return nil, err
	}
	logrus.Info("successfully connected to mongodb")

	defer client.Disconnect(ctx)

	// find all documents in posts collection
	var posts []model.Post
	filter := bson.D{}
	cursor, err := client.Database(databaseName).Collection(collectionName).Find(context.TODO(), filter)
	if err != nil {
		logrus.Errorf("unable to retrieve data from mongodb. error: %v", err)
		return nil, err
	}
	logrus.Infof("successfully retrieved data from mongodb")

	// use cursor to iterate over all the retrieved documents
	for cursor.Next(ctx) {
		var post model.Post
		err := cursor.Decode(&post)
		if err != nil {
			logrus.Errorf("unable to decode data from cursor. error: %v", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	// check if cursor encountered any error
	err = cursor.Err()
	if err != nil {
		logrus.Errorf("error while iterating over retrieved documents. error: %v", err)
		return nil, err
	}
	logrus.Infof("successfully iterated over retrieved documents")

	cursor.Close(ctx)

	return posts, nil
}

func (*repo) Save(post *model.Post) (*model.Post, error) {

	// connect to mongodb
	ctx, client, err := connectToMongoDb(mongodbUri)
	if err != nil {
		logrus.Errorf("unable to connect to mongodb. error: %v", err)
		return nil, err
	}
	logrus.Info("successfully connected to mongodb")

	defer client.Disconnect(ctx)

	// insert in posts collection
	_, err = client.Database(databaseName).Collection(collectionName).InsertOne(context.TODO(), &post)
	if err != nil {
		logrus.Errorf("unable to insert data into mongodb. error: %v", err)
		return nil, err
	}
	logrus.Info("successfully inserted 1 document into mongodb")

	return post, nil
}

func connectToMongoDb(mongoDbUri string) (context.Context, *mongo.Client, error) {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(mongodbUri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logrus.Errorf("unable to create mongo client. error: %v", err)
		return nil, nil, err
	}
	return ctx, client, nil
}
