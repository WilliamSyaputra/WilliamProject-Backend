package controllers

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

// get list images
func GetImages(c *fiber.Ctx) error {
	files, err := os.ReadDir("images")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membaca folder"})
	}

	var images []string
	for _, file := range files {
		images = append(images, file.Name())
	}

	return c.JSON(images)
}

// Get Image by name
func GetImage(c *fiber.Ctx) error {
	name := c.Params("name")
	path := filepath.Join("images", name)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return c.Status(404).JSON(fiber.Map{"error": "Gambar tidak ditemukan"})
	}

	c.Set("Content-Disposition", "attachment; filename="+name)
	c.Set("Content-Type", "application/octet-stream")
	return c.SendFile(path)
}

// Upload Image
func UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "File tidak ditemukan"})
	}

	path := filepath.Join("images", file.Filename)
	if err := c.SaveFile(file, path); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan file"})
	}

	return c.JSON(fiber.Map{"message": "Upload sukses", "filename": file.Filename})
}

// upload multiple
func UploadImages(c *fiber.Ctx) error {
	// Ambil semua file yang diunggah
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal membaca form data"})
	}

	files := form.File["files"] // Mengambil semua file dari input name="files"

	if len(files) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Tidak ada file yang diunggah"})
	}

	var uploadedFiles []string

	for _, file := range files {
		path := filepath.Join("images", file.Filename)

		if err := c.SaveFile(file, path); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan file", "file": file.Filename})
		}

		uploadedFiles = append(uploadedFiles, file.Filename)
	}

	// Berikan respons sukses dengan daftar file yang telah diunggah
	return c.JSON(fiber.Map{"message": "Upload sukses", "files": uploadedFiles})
}

// Delete Image by name
func DeleteImage(c *fiber.Ctx) error {
	name := c.Params("name")
	path := filepath.Join("images", name)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return c.Status(404).JSON(fiber.Map{"error": "Gambar tidak ditemukan"})
	}

	if err := os.Remove(path); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus file"})
	}

	return c.JSON(fiber.Map{"message": "Gambar berhasil dihapus"})
}
