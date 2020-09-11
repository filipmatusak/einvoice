package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/slovak-egov/einvoice/common"
	"io/ioutil"
)

type GSC struct {
	bkt *storage.BucketHandle
	ctx context.Context
}

func (storage *GSC) SaveObject(path, value string) error {
	obj := storage.bkt.Object(path)
	w := obj.NewWriter(storage.ctx)

	if _, err := fmt.Fprintf(w, value); err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}

func (storage *GSC) ReadObject(path string) (string, error) {
	obj := storage.bkt.Object(path)
	r, err := obj.NewReader(storage.ctx)
	if err != nil {
		return "", nil
	}
	defer r.Close()

	res, err := ioutil.ReadAll(r)
	if err != nil {
		return "", nil
	}

	return string(res), err
}

func NewGSC() *GSC {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		panic(err)
	}

	bktName := common.GetRequiredEnvVariable("GCS_BUCKET")
	bkt := client.Bucket(bktName)

	return &GSC{bkt, ctx}
}
