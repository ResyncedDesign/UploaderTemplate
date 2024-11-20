package utils

import (
	"fmt"
	"mime/multipart"
	"time"

	"math/rand"

	"github.com/gofiber/fiber/v2"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// Simple error handler to keep the code "cleaner"
func HandleError(c *fiber.Ctx, err error, status int, message string) error {
	fmt.Println(message, ":", err)
	return c.Status(status).JSON(fiber.Map{
		"error": message,
	})
}

func ValidateFile(file *multipart.FileHeader) (string, bool) {
	if file == nil {
		return "No file uploaded", false
	}

	fileSize := file.Size
	maxFileSize := int64(100 * 1024 * 1024) // 100MB

	if fileSize > maxFileSize {
		return "File is too large", false
	}

	fileContent, err := file.Open()
	if err != nil {
		return "Failed to open the file", false
	}
	defer fileContent.Close()

	buffer := make([]byte, 512)
	if _, err := fileContent.Read(buffer); err != nil {
		return "Failed to read the file", false
	}

	// Here you could add more checks like the file type by checking the buffer content etc.

	return "", true
}

// Pasted straight outta stackoverflow 🖕🤬🖕
func GenerateFileName(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
