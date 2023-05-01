package shared

type MD5Config struct {
	PrependSalt string `json:"prepend_salt,omitempty"`
	AppendSalt  string `json:"append_salt,omitempty"`
}

type Argon2Config struct {
	Salt            string `json:"salt"`
	IterationAmount int    `json:"iteration_amount"`
	Memory          int    `json:"memory"`
	Threads         int    `json:"threads"`
	KeyLength       int    `json:"key_length"`
}

type SHA1Config struct {
	PrependSalt string `json:"prepend_salt,omitempty"`
	AppendSalt  string `json:"append_salt,omitempty"`
}

type ScryptConfig struct {
	Salt       string `json:"salt"`
	NParameter int    `json:"n_parameter"`
	RParameter int    `json:"r_parameter"`
	PParameter int    `json:"p_parameter"`
	KeyLength  int    `json:"key_length"`
}

type HashType string

const (
	HashTypeBcrypt   HashType = "bcrypt"
	HashTypeMD5      HashType = "md_5"
	HashTypeArgon2I  HashType = "argon_2i"
	HashTypeArgon2ID HashType = "argon_2id"
	HashTypeSHA1     HashType = "sha_1"
	HashTypeScrypt   HashType = "scrypt"
	HashTypePHPass   HashType = "phpass"
)
