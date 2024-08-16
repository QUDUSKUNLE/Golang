package integration

type ExternalRepository struct {}

func OpenExternalConnection() *ExternalRepository {
	return &ExternalRepository{}
}
