package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nickel-dime/user-management/config"
	"github.com/nickel-dime/user-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// gets all the users in the db
func GetAllUsers(c *fiber.Ctx) error {
	userCollection := config.MI.DB.Collection(os.Getenv("USER_COLLECTION"))

	query := bson.D{{}}

	cursor, err := userCollection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	var users []models.User = make([]models.User, 0)

	err = cursor.All(c.Context(), &users)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"users": users,
		},
	})
}

// creates a user in the db with unique id
func CreateUser(c *fiber.Ctx) error {
	userCollection := config.MI.DB.Collection(os.Getenv("USER_COLLECTION"))

	data := new(models.User)

	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	data.RegisteredAt = time.Now()

	result, err := userCollection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert user",
			"error":   err,
		})
	}

	user := &models.User{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	userCollection.FindOne(c.Context(), query).Decode(user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user": user,
		},
	})
}

// updates a user given an id, will return error if id not found
func UpdateUser(c *fiber.Ctx) error {
	userCollection := config.MI.DB.Collection(os.Getenv("USER_COLLECTION"))

	paramID := c.Params("id")

	id, err := primitive.ObjectIDFromHex(paramID)

	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
			"error":   err,
		})
	}

	data := new(models.User)
	err = c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	query := bson.D{{Key: "_id", Value: id}}

	var userToUpdate bson.D

	// If the field's type is a pointer, you can use the x != nil check,
	// when using fields with non-pointer types

	// handle errors, keep along the happy path
	if data.Age < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Age is invalid",
			"error":   nil,
		})
	}

	if data.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Name is an empty string",
			"error":   nil,
		})
	}

	userToUpdate = append(userToUpdate, bson.E{Key: "name", Value: data.Name})
	userToUpdate = append(userToUpdate, bson.E{Key: "age", Value: data.Age})
	update := bson.D{
		{Key: "$set", Value: userToUpdate},
	}

	if err := userCollection.FindOneAndUpdate(c.Context(), query, update).Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "User Not found",
				"error":   err,
			})
		}

		// given that the outer condition isn't false (meaning the FindOneAndUpdate call failed)
		// and the error is not the same as ErrNoDocuments, we can safely assume that this error
		// falls outside of the expected edge-cases for the function
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot update user",
			"error":   err,
		})
	}

	// in the case of MongoDB and how multi-tenant services can run at scale
	// (meaning things will run under much higher, but distributed workloads)
	// Golang can, in some instances, execute the function faster than MongoDB.
	// When creating a new resource/document in a Mongo collection, opt for returning

	user := &models.User{}
	userCollection.FindOne(c.Context(), query).Decode(user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user": user,
		},
	})
}

// deletes user of given id, will return error if not found
func DeleteUser(c *fiber.Ctx) error {
	userCollection := config.MI.DB.Collection(os.Getenv("USER_COLLECTION"))

	paramID := c.Params("id")

	id, err := primitive.ObjectIDFromHex(paramID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
			"error":   err,
		})
	}

	query := bson.D{{Key: "_id", Value: id}}

	err = userCollection.FindOneAndDelete(c.Context(), query).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "User Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete user",
			"error":   err,
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
