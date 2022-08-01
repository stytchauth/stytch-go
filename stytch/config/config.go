package config

const APIVersion = "5.12.0"

type BaseURI string

const (
	BaseURITest BaseURI = "https://test.stytch.com/v1"
	BaseURILive BaseURI = "https://api.stytch.com/v1"
)

type Env string

const (
	EnvTest Env = "test"
	EnvLive Env = "live"
)

type Config struct {
	Env     Env
	BaseURI BaseURI
	/** The ProjectID to use with basic authentication */
	/** Replace the value of ProjectID with SetBasicAuthProjectID function */
	ProjectID string
	/** The Secret to use with basic authentication */
	/** Replace the value of Secret with SetBasicAuthSecret function */
	Secret string
}

func New() *Config {
	config := new(Config)
	return config
}

func (c *Config) BasicAuthProjectID() string {
	return c.ProjectID
}

func (c *Config) SetBasicAuthProjectID(projectID string) {
	c.ProjectID = projectID
}

func (c *Config) BasicAuthSecret() string {
	return c.Secret
}

func (c *Config) SetBasicAuthSecret(secret string) {
	c.Secret = secret
}

func (c *Config) GetEnv() Env {
	return c.Env
}

func (c *Config) SetEnv(env Env) {
	c.Env = env
	if env == EnvLive {
		c.BaseURI = BaseURILive
	} else if env == EnvTest {
		c.BaseURI = BaseURITest
	}
}

func (c *Config) GetBaseURI() BaseURI {
	return c.BaseURI
}
