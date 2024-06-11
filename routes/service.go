package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ImArnav19/mongo/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var entryCollection *mongo.Collection = OpenCollection(Client, "customerdb")

func Create(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var customer models.Customer

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"BindError": err.Error()})
		fmt.Println(err)
		return
	}
	validateErr := validate.Struct(customer)

	if validateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"BindError": validateErr.Error()})
		fmt.Println(validateErr)
		return
	}

	customer.ID = primitive.NewObjectID()

	results, insertErr := entryCollection.InsertOne(ctx, customer)
	if insertErr != nil {
		msg := "Insert Problem in DB!"
		c.JSON(http.StatusInternalServerError, gin.H{"Inserror": msg})
		fmt.Println(insertErr)
		return
	}

	defer cancel()

	fmt.Println(results)
	c.JSON(http.StatusOK, results)

}

func Read(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var customers []bson.M
	cursor, err := entryCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}

	if curErr := cursor.All(ctx, &customers); curErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message Cursor Error": curErr.Error(),
		})
		fmt.Println(curErr.Error())
		return
	}

	defer cancel()
	fmt.Println(customers)
	c.JSON(http.StatusOK, customers)

}
func Update(c *gin.Context) {
	//binding validate both very important
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
	custID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(custID)

	var customer models.Customer
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"BindErr": err.Error()})
		fmt.Println(err)
		return
	}

	validateErr := validate.Struct(customer)

	if validateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"BindError": validateErr.Error()})
		fmt.Println(validateErr)
		return
	}

	results, Updateerr := entryCollection.ReplaceOne(ctx,

		bson.M{"_id": docID},
		bson.M{
			"Name":  customer.Name,
			"Email": customer.Email,
			"Age":   customer.Age,
		},
	)
	if Updateerr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"UpdateError": Updateerr.Error()})
		fmt.Println(Updateerr)
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, results)

}
func Delete(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
	custID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(custID)

	result, delErr := entryCollection.DeleteOne(ctx, bson.M{"_id": docID})

	if delErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"UpdateError": delErr.Error()})
		fmt.Println(delErr)
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result.DeletedCount)

}
