package invoice

const (
	JsonFormat = "json"
	UblFormat  = "ubl2.1"
	D16bFormat = "d16b"
)

type Invoice struct {
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Price    float64 `json:"price"`
}

type Meta struct {
	Id       string  `json:"id"`
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Format   string  `json:"format"`
	Price    float64 `json:"price"`
}

func (invoice *Invoice) GetMeta() *Meta {
	return &Meta{
		Sender:   invoice.Sender,
		Receiver: invoice.Receiver,
		Format:   JsonFormat,
		Price:    invoice.Price,
	}
}
