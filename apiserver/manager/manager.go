package manager

import (
	"encoding/json"
	"github.com/slovak-egov/einvoice/apiserver/db"
	"github.com/slovak-egov/einvoice/apiserver/invoice"
	"github.com/slovak-egov/einvoice/apiserver/storage"
	"github.com/slovak-egov/einvoice/apiserver/xml/d16b"
	"github.com/slovak-egov/einvoice/apiserver/xml/ubl21"
)

type Manager interface {
	Create(invoice *invoice.Invoice) (*invoice.Meta, error)
	CreateUBL(value string) (*invoice.Meta, error)
	CreateD16B(value string) (*invoice.Meta, error)
	CreateJSON(value string) (*invoice.Meta, error)

	GetFull(id string, format string) (string, error)
	GetMeta(id string) (*invoice.Meta, error)
	GetAllInvoiceMeta() ([]invoice.Meta, error)
}

type managerImpl struct {
	db      db.DBConnector
	storage storage.Storage
}

func NewManager(db db.DBConnector, storage storage.Storage) Manager {
	return &managerImpl{db, storage}
}

func (manager *managerImpl) Create(invoice *invoice.Invoice) (*invoice.Meta, error) {
	meta := invoice.GetMeta()

	if err := manager.db.CreateInvoice(meta); err != nil {
		return nil, err
	}

	jsonString, err := json.Marshal(invoice)
	if err != nil {
		return nil, err
	}

	if err = manager.storage.SaveObject("invoice-"+meta.Id+".json", string(jsonString)); err != nil {
		return nil, err
	}

	return meta, nil
}

func fileNameFromMeta(meta *invoice.Meta) string {
	return fileName(meta.Id, meta.Format)
}

func fileName(id, format string) string {
	extension := "json"
	if format != invoice.JsonFormat {
		extension = "xml"
	}

	return "invoice-" + id + "." + extension
}

func (manager *managerImpl) CreateJSON(value string) (*invoice.Meta, error) {
	var inv = new(invoice.Invoice)

	if err := json.Unmarshal([]byte(value), &inv); err != nil {
		return nil, err
	}

	meta := inv.GetMeta()

	if err := manager.db.CreateInvoice(meta); err != nil {
		return nil, err
	}

	if err := manager.storage.SaveObject(fileNameFromMeta(meta), value); err != nil {
		return nil, err
	}
	return meta, nil
}

func (manager *managerImpl) CreateUBL(value string) (*invoice.Meta, error) {
	err, meta := ubl21.Create(value)
	if err != nil {
		return nil, err
	}

	if err = manager.db.CreateInvoice(meta); err != nil {
		return nil, err
	}

	err = manager.storage.SaveObject(fileNameFromMeta(meta), value)
	if err != nil {
		return nil, err
	}
	return meta, nil
}

func (manager *managerImpl) CreateD16B(value string) (*invoice.Meta, error) {
	err, meta := d16b.Create(value)
	if err != nil {
		return nil, err
	}

	if err = manager.db.CreateInvoice(meta); err != nil {
		return nil, err
	}

	err = manager.storage.SaveObject(fileNameFromMeta(meta), value)
	if err != nil {
		return nil, err
	}
	return meta, err
}

func (manager *managerImpl) GetMeta(id string) (*invoice.Meta, error) {
	return manager.db.GetInvoiceMeta(id)
}

func (manager *managerImpl) GetFull(id string, format string) (string, error) {
	invoiceStr, err := manager.storage.ReadObject(fileName(id, format))
	if err != nil {
		return "", err
	}
	return invoiceStr, nil
}

func (manager *managerImpl) GetAllInvoiceMeta() ([]invoice.Meta, error) {
	return manager.db.GetAllInvoice()
}
