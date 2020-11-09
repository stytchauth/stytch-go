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
 * Client structure as interface implementation
 */
type STYTCH_IMPL struct {
     users users_pkg.USERS
     magiclinks magiclinks_pkg.MAGICLINKS
     emails emails_pkg.EMAILS
     config  configuration_pkg.CONFIGURATION
}

/**
     * Access to Configuration
     * @return Returns the Configuration instance
*/
func (me *STYTCH_IMPL) Configuration() configuration_pkg.CONFIGURATION {
    return me.config
}
/**
     * Access to Users controller
     * @return Returns the Users() instance
*/
func (me *STYTCH_IMPL) Users() users_pkg.USERS {
    if(me.users) == nil {
        me.users = users_pkg.NewUSERS(me.config)
    }
    return me.users
}
/**
     * Access to MagicLinks controller
     * @return Returns the MagicLinks() instance
*/
func (me *STYTCH_IMPL) MagicLinks() magiclinks_pkg.MAGICLINKS {
    if(me.magiclinks) == nil {
        me.magiclinks = magiclinks_pkg.NewMAGICLINKS(me.config)
    }
    return me.magiclinks
}
/**
     * Access to Emails controller
     * @return Returns the Emails() instance
*/
func (me *STYTCH_IMPL) Emails() emails_pkg.EMAILS {
    if(me.emails) == nil {
        me.emails = emails_pkg.NewEMAILS(me.config)
    }
    return me.emails
}

