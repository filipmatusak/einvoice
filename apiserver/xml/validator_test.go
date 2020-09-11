package xml

import (
	"github.com/lestrrat-go/libxml2/xsd"
	"io/ioutil"
	"testing"
)

func TestD16BValidation(t *testing.T) {
	validator := NewValidator("../../xml/d16b/xsd", "../../xml/ubl21/xsd")

	bytes, err := ioutil.ReadFile("../../xml/d16b/example/d16b_invoice.xml")
	if err != nil {
		panic(err)
	}

	if err = validator.ValidateD16B(bytes); err != nil {
		switch v := err.(type) {
		case xsd.SchemaValidationError:
			t.Error(v.Errors())
		default:
			t.Error(err)
		}
	}
}

func TestUBL21Validation(t *testing.T) {
	validator := NewValidator("../../xml/d16b/xsd", "../../xml/ubl21/xsd")

	bytes, err := ioutil.ReadFile("../../xml/ubl21/example/ubl21_invoice.xml")
	if err != nil {
		panic(err)
	}

	if err = validator.ValidateUBL21(bytes); err != nil {
		switch v := err.(type) {
		case xsd.SchemaValidationError:
			t.Error(v.Errors())
		default:
			t.Error(err)
		}
	}
}
