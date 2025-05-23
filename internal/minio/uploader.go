package minio

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func UploadParts(folderPath, format string) ([]string, error) {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	bucketName := os.Getenv("MINIO_BUCKET")

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	var urls []string

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasSuffix(name, "."+format) {
			path := filepath.Join(folderPath, name)
			_, err := client.FPutObject(context.Background(), bucketName, name, path, minio.PutObjectOptions{ContentType: "video/mp4"})
			if err != nil {
				log.Println("Upload error:", name, err)
				continue
			}
			urls = append(urls, fmt.Sprintf("http://%s/%s/%s", endpoint, bucketName, name))
		}
	}
	return urls, nil
}
