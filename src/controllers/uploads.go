package controllers

import (
	"context"
	"fmt"
	"go-template/src/config"
	"go-template/src/services"
	"go-template/src/types"
	"go-template/src/utils"
	"path"

	"github.com/gofiber/fiber/v2"
)

func HandleUpload(c *fiber.Ctx) error {
	// Fetch the R2Service from the context that we previously saved in the context in main.go
	R2Service, ok := c.Locals("R2Service").(*services.R2Service)
	if !ok || R2Service == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "R2Service not found in context",
		})
	}

	// Since this endpoint expects a file the POST request should be a multipart form so we need to parse that
	form, err := c.MultipartForm()
	if err != nil {
		return utils.HandleError(c, err, fiber.StatusBadRequest, "Cannot parse form data")
	}

	// Check if the file is present in the form data (you could accept multiple files but to keep this simple we only accept one)
	files := form.File["file"]
	if len(files) != 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Please upload a single file",
		})
	}

	file := files[0]
	errorMessage, validationResult := utils.ValidateFile(file)
	if !validationResult {
		return c.Status(fiber.StatusBadRequest).JSON(errorMessage)
	}

	// Right now just generate a random file name and save the file in the uploads folder
	fileName := fmt.Sprintf("%v%v", utils.GenerateFileName(10), path.Ext(file.Filename))
	filePath := fmt.Sprintf("uploads/%v", fileName)

	openFile, err := file.Open()
	if err != nil {
		return utils.HandleError(c, err, fiber.StatusInternalServerError, "Failed to open file")
	}
	defer openFile.Close()

	// Sometimes if the client doesn't send the content type it's usually octet-stream so you could try to guess it based on the file extension
	contentType := file.Header.Get("Content-Type")

	if err := R2Service.UploadFile(context.TODO(), filePath, openFile, int64(file.Size), contentType); err != nil {
		return utils.HandleError(c, err, fiber.StatusInternalServerError, "Failed to upload file")
	}

	// Construct the response
	response := types.File{
		Name: fileName,
		URL:  fmt.Sprintf("%v/%v", config.GetR2URL(), filePath),
	}

	return c.JSON(types.JSONResponse{
		Success: true,
		Files:   []types.File{response},
	})
}
