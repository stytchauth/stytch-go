package shared

type Key struct {
	Typ     string   `json:"kty"`
	Use     string   `json:"use"`
	KeyOps  []string `json:"key_ops"`
	Alg     string   `json:"alg"`
	KeyID   string   `json:"kid"`
	X5C     []string `json:"x5c"`
	X5TS256 string   `json:"x5tS256"`
	N       string   `json:"n"`
	E       string   `json:"e"`
}
