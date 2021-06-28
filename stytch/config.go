package stytch

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

type config struct {
	env     Env
	baseURI BaseURI
	/** The projectID to use with basic authentication */
	/** Replace the value of ProjectID with SetBasicAuthProjectID function */
	projectID string
	/** The secret to use with basic authentication */
	/** Replace the value of Secret with SetBasicAuthSecret function */
	secret string
}

func newConfig() *config {
	config := new(config)
	return config
}

func (c *config) BasicAuthProjectID() string {
	return c.projectID
}

func (c *config) SetBasicAuthProjectID(projectID string) {
	c.projectID = projectID
}

func (c *config) BasicAuthSecret() string {
	return c.secret
}

func (c *config) SetBasicAuthSecret(secret string) {
	c.secret = secret
}

func (c *config) GetEnv() Env {
	return c.env
}

func (c *config) SetEnv(env Env) {
	c.env = env
	if env == EnvLive {
		c.baseURI = BaseURILive
	} else if env == EnvTest {
		c.baseURI = BaseURITest
	}
}

func (c *config) GetBaseURI() BaseURI {
	return c.baseURI
}
