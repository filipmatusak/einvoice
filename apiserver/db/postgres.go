package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/slovak-egov/einvoice/apiserver/invoice"
	"io/ioutil"
	"strconv"
)

type dbConnector struct {
	db *pg.DB
}

func NewDBConnector() DBConnector {
	return &dbConnector{}
}

func (connector *dbConnector) Connect(config ConnectionConfig) {
	connector.db = pg.Connect(&pg.Options{
		Addr:     config.Host + ":" + strconv.Itoa(config.Port),
		User:     config.User,
		Password: config.Password,
		Database: config.Database,
	})
}

func (connector *dbConnector) Close() {
	connector.db.Close()
}

func (connector *dbConnector) InitDB() error {

	query, err := ioutil.ReadFile("sql/setup.sql")
	if err != nil {
		return err
	}

	_, err = connector.db.Exec(string(query))

	return err
}

func (connector *dbConnector) GetAllInvoice() ([]invoice.Meta, error) {
	var invoices []invoice.Meta
	err := connector.db.Model(&invoices).Select()
	if err != nil {
		return nil, err
	}
	return invoices, nil
}

func (connector *dbConnector) GetInvoiceMeta(id string) (*invoice.Meta, error) {
	inv := &invoice.Meta{}
	err := connector.db.Model(inv).Where("id = ?", id).Select(inv)
	if err != nil {
		return nil, err
	}
	return inv, nil
}

func (connector *dbConnector) CreateInvoice(invoice *invoice.Meta) error {
	_, err := connector.db.Model(invoice).Insert(invoice)

	return err
}
