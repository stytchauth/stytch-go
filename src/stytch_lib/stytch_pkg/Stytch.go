/*
 * stytch_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package StytchClient

import(
	"stytch_lib/configuration_pkg"
	"stytch_lib/users_pkg"
	"stytch_lib/magiclinks_pkg"
	"stytch_lib/emails_pkg"
)


/*
 * Interface for the STYTCH_IMPL
 */
type STYTCH interface {
    Users()                 users_pkg.USERS
    MagicLinks()            magiclinks_pkg.MAGICLINKS
    Emails()                emails_pkg.EMAILS
    Configuration()         configuration_pkg.CONFIGURATION
}

/*
 * Factory for the STYTCH interface returning STYTCH_IMPL
 */
func NewSTYTCH(basicAuthUserName string, basicAuthPassword string) STYTCH {
    stytchClient := new(STYTCH_IMPL)
    stytchClient.config = configuration_pkg.NewCONFIGURATION()

    stytchClient.config.SetBasicAuthUserName(basicAuthUserName)
    stytchClient.config.SetBasicAuthPassword(basicAuthPassword)

    return stytchClient
}
