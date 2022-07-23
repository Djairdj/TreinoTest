package product

//Implementação do Repository

import (
	"TreinoTest/internal/entity"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(db *mongo.Database) Repository {
	return &mongoRepository{
		db: db,
	}
}

func (r mongoRepository) Create(ctx context.Context, name string) (*entity.Product, error) {
	collection := r.db.Collection("product")
	data := bson.D{{"name", name}, {"votes", 0}}
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	id, ok := res.InsertedID.(primitive.ObjectID) //feitura de cast
	if !ok {
		return nil, fmt.Errorf("nao foi possivel converter o id")
	}
	p := entity.Product{
		ID:    id.Hex(),
		Name:  name,
		Votes: 0,
	}
	return &p, nil
}

func (r mongoRepository) List(ctx context.Context) ([]entity.Product, error) {
	collection := r.db.Collection("product")
	findOptions := options.Find()
	//Set the limit of the number of record to find
	findOptions.SetLimit(5)
	//Define an array in which you can store the decoded documents
	results := []entity.Product{}

	//Passing the bson.D{{}} as the filter matches  documents in the collection
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("could not find the products: %v", err)
	}

	for cur.Next(ctx) {
		//Create a value into which the single document can be decoded
		elem := entity.Product{}
		err = cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	err = cur.Close(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}
