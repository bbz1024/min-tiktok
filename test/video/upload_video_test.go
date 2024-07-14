package video

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// UploadVideoToServer uploads a video file to a server.
// It returns an error if the upload fails.
func UploadVideoToServer(url, filePath, title, token string) error {
	// Open the file to be uploaded
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a multipart writer to build the request body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the file field to the multipart request body
	part, err := writer.CreateFormFile("data", filepath.Base(filePath))
	if err != nil {
		return fmt.Errorf("failed to create form file: %w", err)
	}

	// Write the file content into the multipart request body's file field
	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("failed to write file to body: %w", err)
	}

	// Add other fields to the request body
	err = writer.WriteField("title", title)
	if err != nil {
		return fmt.Errorf("failed to add title field: %w", err)
	}
	err = writer.WriteField("token", token)
	if err != nil {
		return fmt.Errorf("failed to add token field: %w", err)
	}

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set the Content-Type header to multipart/form-data format with boundary
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status code for success
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(all))

	return nil
}
func TestUploadVideo(t *testing.T) {
	token := "aea5c417-cbab-4e32-9d8c-43f2da3d0635"
	url := "http://localhost:8891/douyin/publish/action/"
	videosPath := "C:\\Users\\lzb\\Downloads\\Video3"
	dir, err := os.ReadDir(videosPath)
	if err != nil {
		t.Errorf("ReadDir() error = %v", err)
		return
	}
	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		path := filepath.Join(videosPath, file.Name())
		title := file.Name()
		err := UploadVideoToServer(url, path, title, token)
		if err != nil {
			t.Errorf("UploadVideoToServer() error = %v", err)
		}
		time.Sleep(time.Second * 1)
	}
}
