package methodoptions

type Authorization struct {
	// A secret token for a given Stytch Session
	SessionToken string
	// The JSON Web Token (JWT) for a given Stytch Session.
	SessionJWT string
}

func (a *Authorization) AddHeaders(headers map[string][]string) map[string][]string {
	if headers == nil {
		headers = make(map[string][]string)
	}
	if a.SessionToken != "" {
		headers["X-Stytch-Member-Session"] = []string{a.SessionToken}
	}
	if a.SessionJWT != "" {
		headers["X-Stytch-Member-SessionJWT"] = []string{a.SessionJWT}
	}
	return headers
}
