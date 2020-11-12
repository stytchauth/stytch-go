package stytch

type BaseURI string
const (
	BaseURITest BaseURI = "https://test.stytch.com/v1"
	BaseURILive BaseURI = "https://api.stytch.com/v1"
)

type Env string
const (
	EnvTest Env = "lest"
	EnvLive Env = "live"
)

type config struct {
	env Env
	baseURI BaseURI
	/** The projectID to use with basic authentication */
	/** Replace the value of ProjectID with SetBasicAuthProjectID function */
	projectID string
	/** The secret to use with basic authentication */
	/** Replace the value of Secret with SetBasicAuthSecret function */
	secret string
}

//type Configuration interface {
//	BasicAuthProjectID() string
//	SetBasicAuthProjectID(projectID   string)
//	BasicAuthSecret() string
//	SetBasicAuthSecret(secret   string)
//	SetEnv(env Env)
//	GetEnv() Env
//	GetBaseURI() BaseURI
//}


func newConfig() *config{
	config := new(config)
	return config
}

/*
 * Getter function returning ProjectID
 */
func (c *config) BasicAuthProjectID() string{
	return c.projectID
}

/*
 * Setter function setting up ProjectID
 */
func (c *config) SetBasicAuthProjectID(projectID string) {
	c.projectID = projectID
}

/*
 * Getter function returning Secret
 */
func (c *config) BasicAuthSecret() string{
	return c.secret
}

/*
 * Setter function setting up Secret
 */
func (c *config) SetBasicAuthSecret(secret string) {
	c.secret = secret
}

/*
 * Getter function returning Env
 */
func (c *config) GetEnv() Env{
	return c.env
}

/*
 * Setter function setting up Env
 */
func (c *config) SetEnv(env Env) {
	c.env = env
	if env == EnvLive {
		c.baseURI = BaseURILive
	} else if env == EnvTest {
		c.baseURI = BaseURITest
	}
}

/*
 * Getter function returning BaseURI
 */
func (c *config) GetBaseURI() BaseURI{
	return c.baseURI
}
