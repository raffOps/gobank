package clients

type Titular struct {
	name       string
	cpf        string
	profession string
}

func NewTitular(name, cpf, profession string) (*Titular, error) {
	titular := Titular{name: name, cpf: cpf, profession: profession}
	return &titular, nil
}
