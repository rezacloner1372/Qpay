package bank

type Config struct {
	MerchantID string `koanf:"merchantId"`
	BaseURL    string `koanf:"baseurl"`
}
