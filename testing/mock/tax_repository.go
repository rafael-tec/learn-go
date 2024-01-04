package mock

type Entity struct {
	ID  string
	Tax float64
}

type DatabaseClient interface {
	Save(statement string, entity Entity) error
}

type TaxRepository struct {
	dbClient DatabaseClient
}

func NewTaxRepository(dbClient DatabaseClient) TaxRepository {
	return TaxRepository{dbClient: dbClient}
}

func (r TaxRepository) SaveTax(tax float64) error {
	entity := Entity{ID: "a2b4c0g9", Tax: tax}

	err := r.dbClient.Save("", entity)
	if err != nil {
		return err
	}

	return nil
}
