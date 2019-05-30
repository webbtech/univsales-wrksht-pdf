package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/pulpfree/univsales-wrksht-pdf/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	log "github.com/sirupsen/logrus"
)

// DB and Table constants
const (
	colAddress    = "addresses"
	colCustomer   = "customers"
	colGroupTypes = "group-types"
	colJS         = "jobsheets"
	colJSGroups   = "jobsheet-win-grps"
	colJSOther    = "jobsheet-other"
	colJSWindows  = "jobsheet-wins"
	colPayments   = "payments"
	colProducts   = "products"
	colQuotes     = "quotes"
	Users         = "users"
)

// MDB struct
type MDB struct {
	client *mongo.Client
	dbName string
	db     *mongo.Database
}

// NewDB sets up new MDB struct
func NewDB(connection string, dbNm string) (model.DBHandler, error) {

	clientOptions := options.Client().ApplyURI(connection)
	err := clientOptions.Validate()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	// defer suite.db.Close()

	return &MDB{
		client: client,
		dbName: dbNm,
		db:     client.Database(dbNm),
	}, err
}

// FetchQuote method
func (db *MDB) FetchQuote(quoteID string) (*model.Quote, error) {

	// Initialize
	q := &model.Quote{}

	// Fetch quote
	err := db.getQuote(q, quoteID)
	if err != nil {
		return q, err
	}

	// Fetch Group items
	err = db.getGroupItems(q)
	if err != nil {
		return q, err
	}

	// Fetch Windows
	err = db.getWindowItems(q)
	if err != nil {
		return q, err
	}

	// Fetch Other items
	err = db.getOtherItems(q)
	if err != nil {
		return q, err
	}

	// Fetch Jobsheet features
	err = db.getJobsheetFeatures(q)
	if err != nil {
		return q, err
	}

	// Fetch customer data
	err = db.getCustomer(q)
	if err != nil {
		return q, err
	}

	return q, nil
}

func (db *MDB) getQuote(q *model.Quote, quoteID string) error {

	if quoteID == "" {
		return errors.New("Missing quoteID string")
	}

	col := db.db.Collection(colQuotes)
	objectIDS, err := primitive.ObjectIDFromHex(quoteID)
	// filter := bson.D{{"_id", objectIDS}}
	// found answer to go-vet issue in above filter here: https://stackoverflow.com/questions/54548441/composite-literal-uses-unkeyed-fields#answer-54548495
	filter := bson.D{primitive.E{Key: "_id", Value: objectIDS}}
	err = col.FindOne(context.Background(), filter).Decode(&q)
	if err != nil {
		log.Fatalf("quote table error: %s", err)
		return err
	}

	return nil
}

func (db *MDB) getGroupItems(q *model.Quote) error {

	// fixme: this should be in FetchQuote
	q.Items = &model.Items{}

	col := db.db.Collection(colJSGroups)
	for _, i := range q.ItemIds.Group {
		item := &model.Group{}
		// these IDs were stored as strings, so now need to convert
		objectIDS, err := primitive.ObjectIDFromHex(i)
		if err != nil {
			return errors.New("Invalid groupID")
		}

		filter := bson.D{primitive.E{Key: "_id", Value: objectIDS}}
		if err := col.FindOne(context.Background(), filter).Decode(&item); err != nil {
			log.Fatalf("Group not found. Error: %s, group id: %s", err, i)
			return err
		}
		// Fetch group type
		// colGT := db.db.Collection(colGroupTypes)
		// ret := bson.M{}
		// gtFilter := bson.D{primitive.E{Key: "_id", Value: item.Specs["groupType"]}}
		// if err := colGT.FindOne(context.Background(), gtFilter).Decode(&ret); err != nil {
		// 	log.Fatalf("Group Type not found. Error: %s", err)
		// 	return err
		// }
		// // Set group type name
		// item.Specs["groupTypeName"] = ret["name"]
		q.Items.Group = append(q.Items.Group, item)
	}

	return nil
}

func (db *MDB) getWindowItems(q *model.Quote) error {

	col := db.db.Collection(colJSWindows)
	for _, i := range q.ItemIds.Window {
		item := &model.Window{}
		objectIDS, err := primitive.ObjectIDFromHex(i)
		if err != nil {
			return errors.New("Invalid windowID")
		}

		filter := bson.D{primitive.E{Key: "_id", Value: objectIDS}}
		if err := col.FindOne(context.Background(), filter).Decode(&item); err != nil {
			log.Fatalf("Window error for id: %s. Error: %s", i, err)
			return err
		}
		// fetch product info
		prod, err := db.getProductName(item.ProductID)
		if err != nil {
			return err
		}
		item.ProductName = prod.Name
		q.Items.Window = append(q.Items.Window, item)
	}

	return nil
}

func (db *MDB) getOtherItems(q *model.Quote) error {

	col := db.db.Collection(colJSOther)
	for _, i := range q.ItemIds.Other {
		item := &model.Other{}
		objectIDS, err := primitive.ObjectIDFromHex(i)
		if err != nil {
			return errors.New("Invalid otherID")
		}

		filter := bson.D{primitive.E{Key: "_id", Value: objectIDS}}
		if err := col.FindOne(context.Background(), filter).Decode(&item); err != nil {
			log.Fatalf("Other not found. Error: %s", err)
			return err
		}
		q.Items.Other = append(q.Items.Other, item)
	}

	return nil
}

func (db *MDB) getProductName(productID primitive.ObjectID) (*model.Product, error) {

	col := db.db.Collection(colProducts)
	p := &model.Product{}
	filter := bson.D{primitive.E{Key: "_id", Value: productID}}
	if err := col.FindOne(context.Background(), filter).Decode(&p); err != nil {
		return p, err
	}
	return p, nil
}

func (db *MDB) getCustomer(q *model.Quote) error {

	col := db.db.Collection(colCustomer)
	filter := bson.D{primitive.E{Key: "_id", Value: q.CustomerID}}
	err := col.FindOne(context.Background(), filter).Decode(&q.Customer)
	if err != nil {
		return err
	}
	phoneMap := map[string]string{}
	for _, v := range q.Customer.Phones {
		phoneMap[v.Type] = v.Number
	}
	q.Customer.PhoneMap = phoneMap

	// // Fetch customer address data
	col = db.db.Collection(colAddress)
	adFilter := bson.D{primitive.E{Key: "customerID", Value: q.CustomerID}, primitive.E{Key: "associate", Value: "customer"}}
	// err = col.Find(bson.M{"customerID": q.CustomerID, "associate": "customer"}).One(&q.Customer.Address)
	err = col.FindOne(context.Background(), adFilter).Decode(&q.Customer.Address)
	if err != nil {
		return err
	}

	return nil
}

func (db *MDB) getJobsheetFeatures(q *model.Quote) error {

	col := db.db.Collection(colJS)
	jobSheet := &model.JobSheet{}

	// see: https://stackoverflow.com/questions/53120116/how-to-filter-fields-from-a-mongo-document-with-the-official-mongo-go-driver
	/* projection := model.JobSheet{
		// ID: nil,
		Features: nil,
	} */

	filter := bson.D{primitive.E{Key: "_id", Value: q.JobsheetID}}
	if err := col.FindOne(context.Background(), filter).Decode(&jobSheet); err != nil {
		log.Fatalf("Jobsheet query error: %s", err)
		return err
	}
	q.Features = jobSheet.Features

	return nil
}

// Close method
func (db *MDB) Close() {
	err := db.client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
