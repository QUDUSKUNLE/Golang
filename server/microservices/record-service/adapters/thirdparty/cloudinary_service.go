package thirdparty

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func CloudinaryUploader(file string) (string, error) {
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
	if err != nil {
		fmt.Printf("Error handshaking cloudinary %s", err.Error())
		return "", err
	}
	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{Folder: "halalmeat"})
	if err != nil {
		fmt.Printf("Error uploading %s to cloudinary", err.Error())
		return "", err
	}
	fmt.Println("File uploaded successfully.")
	return uploadResult.SecureURL, nil
}
