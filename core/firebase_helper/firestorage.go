package firebase_helper

import (
	"context"
	"encoding/json"
	"fmt"

	"example.com/test/core/env"
	firebase "firebase.google.com/go"
	"github.com/google/uuid"

	"google.golang.org/api/option"
)

func UploadAsJson(data any) {
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("error initializing app: %v\n", err)
		return
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		fmt.Printf("error getting storage client: %v\n", err)
		return
	}

	FIRE_STORAGE_BUCKET := env.Get("FIRE_STORAGE_BUCKET")
	bucket, err := client.Bucket(FIRE_STORAGE_BUCKET)
	if err != nil {
		fmt.Printf("error getting bucket: %v\n", err)
		return
	}

	tokenId := uuid.New()
	fileName := fmt.Sprintf("%s.json", uuid.New())
	obj := bucket.Object(fileName)
	wc := obj.NewWriter(context.Background())
	wc.ContentType = "application/json"
	wc.ObjectAttrs.Metadata = map[string]string{
		"firebaseStorageDownloadTokens": tokenId.String(),
	}

	dataToaBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error marshalling data: %v\n", err)
		return
	}

	if _, err := wc.Write(dataToaBytes); err != nil {
		fmt.Printf("error writing data to bucket: %v\n", err)
		return
	}

	errSaving := wc.Close()

	if errSaving != nil {
		fmt.Printf("error closing writer: %v\n", errSaving)
		return
	}
}
