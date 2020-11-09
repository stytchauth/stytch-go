# Getting started

This is the Stytch api.  You can find out more about Stytch at 
[stytch.com](https://stytch.com).


## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
    * **Go**  (Visit [https://golang.org/doc/install](https://golang.org/doc/install) for more details on how to install Go)
    * **Java VM** Version 8 or later
    * **Eclipse 4.6 (Neon)** or later ([http://www.eclipse.org/neon/](http://www.eclipse.org/neon/))
    * **GoClipse** setup within above installed Eclipse (Follow the instructions at [this link](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions) to setup GoClipse)
* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.
* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:

```Go
go get github.com/apimatic/unirest-go
```

This will install unirest-go in the ```GOPATH``` you specified in the system variables.

Now follow the steps mentioned below to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

![Importing SDK into Eclipse - Step 1](https://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](https://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=Stytch-GoLang&projectName=stytch_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=stytch_lib)

## How to Use

The following section explains how to use the StytchLib library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=stytch_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=stytch_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=stytch_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=stytch_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=stytch_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=Stytch-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=stytch_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=stytch_lib)

## Initialization

### Authentication
In order to setup authentication of the API client, you need the following information.

| Parameter | Description |
|-----------|-------------|
| basicAuthUserName | The username to use with basic authentication |
| basicAuthPassword | The password to use with basic authentication |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


# Class Reference

## <a name="list_of_controllers"></a>List of Controllers

* [users_pkg](#users_pkg)
* [magiclinks_pkg](#magiclinks_pkg)
* [emails_pkg](#emails_pkg)

## <a name="users_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".users_pkg") users_pkg

### Get instance

Factory for the ``` USERS ``` interface can be accessed from the package users_pkg.

```go
users := users_pkg.NewUSERS()
```

### <a name="create_user"></a>![Method: ](https://apidocs.io/img/method.png ".users_pkg.CreateUser") CreateUser

> Add a user to Stytch. A user_id is returned in the response that can then be used to perform other operations within Stytch.


```go
func (me *USERS_IMPL) CreateUser(body *models_pkg.UserCreate)(*models_pkg.UserCreateResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | Created user object |


#### Example Usage

```go
var body *models_pkg.UserCreate

var result *models_pkg.UserCreateResponse
result,_ = users.CreateUser(body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not found |
| 429 | Too many requests |
| 500 | Internal server error |



### <a name="get_user_by_id"></a>![Method: ](https://apidocs.io/img/method.png ".users_pkg.GetUserByID") GetUserByID

> Fetch a given user to see what their various attributes are.


```go
func (me *USERS_IMPL) GetUserByID(userId string)(*models_pkg.UserGetResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| userId |  ``` Required ```  | The user_id for the user to fetch. |


#### Example Usage

```go
userId := "user-test-b8797f2c-a93c-11ea-bb37-0242ac130002"

var result *models_pkg.UserGetResponse
result,_ = users.GetUserByID(userId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not found |
| 429 | Too many requests |
| 500 | Internal server error |



### <a name="update_user"></a>![Method: ](https://apidocs.io/img/method.png ".users_pkg.UpdateUser") UpdateUser

> Update a user's attributes. For example, you can add additional emails or change the user's primary email.


```go
func (me *USERS_IMPL) UpdateUser(
            userId string,
            body *models_pkg.UserUpdate)(*models_pkg.UserUpdateResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| userId |  ``` Required ```  | The user_id to update. |
| body |  ``` Required ```  | Updated user object |


#### Example Usage

```go
userId := "user_id"
var body *models_pkg.UserUpdate

var result *models_pkg.UserUpdateResponse
result,_ = users.UpdateUser(userId, body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not found |
| 429 | Too many requests |
| 500 | Internal server error |



### <a name="delete_user"></a>![Method: ](https://apidocs.io/img/method.png ".users_pkg.DeleteUser") DeleteUser

> Remove a user from Stytch.


```go
func (me *USERS_IMPL) DeleteUser(userId string)(*models_pkg.UserDeleteResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| userId |  ``` Required ```  | The user_id to be deleted. |


#### Example Usage

```go
userId := "user-test-b8797f2c-a93c-11ea-bb37-0242ac130002"

var result *models_pkg.UserDeleteResponse
result,_ = users.DeleteUser(userId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not found |
| 429 | Too many requests |
| 500 | Internal server error |



[Back to List of Controllers](#list_of_controllers)

## <a name="magiclinks_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".magiclinks_pkg") magiclinks_pkg

### Get instance

Factory for the ``` MAGICLINKS ``` interface can be accessed from the package magiclinks_pkg.

```go
magicLinks := magiclinks_pkg.NewMAGICLINKS()
```

### <a name="create_send_magic_link"></a>![Method: ](https://apidocs.io/img/method.png ".magiclinks_pkg.CreateSendMagicLink") CreateSendMagicLink

> Send a magic link to the user. You can optionally include additional security measures such as requiring the ip address the link is requested from match the one it's clicked from.


```go
func (me *MAGICLINKS_IMPL) CreateSendMagicLink(body *models_pkg.MagicLinkSend)(*models_pkg.MagicLinkSendResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.MagicLinkSend

var result *models_pkg.MagicLinkSendResponse
result,_ = magicLinks.CreateSendMagicLink(body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not found |
| 429 | Too many requests |
| 500 | Internal server error |



### <a name="create_send_email_magic_link"></a>![Method: ](https://apidocs.io/img/method.png ".magiclinks_pkg.CreateSendEmailMagicLink") CreateSendEmailMagicLink

> Send a magic link to the user. You can optionally include additional security measures such as requiring the ip address the link is requested from match the one it's clicked from.


```go
func (me *MAGICLINKS_IMPL) CreateSendEmailMagicLink(body *models_pkg.MagicLinkSendByEmail)(*models_pkg.MagicLinkSendByEmailResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.MagicLinkSendByEmail

var result *models_pkg.MagicLinkSendByEmailResponse
result,_ = magicLinks.CreateSendEmailMagicLink(body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not found |
| 429 | Too many requests |
| 500 | Internal server error |



### <a name="post_user_magic_link_authenticate"></a>![Method: ](https://apidocs.io/img/method.png ".magiclinks_pkg.PostUserMagicLinkAuthenticate") PostUserMagicLinkAuthenticate

> Authenticate a user given a magic link. This endpoint verifies that the link is valid, hasn't expired, and any optional security settings such as ip match or user agent match are satisfied. Not to be confused with the emails verify endpoint meant for initial, one time verification that the correct email was supplied during sign up.


```go
func (me *MAGICLINKS_IMPL) PostUserMagicLinkAuthenticate(
            token string,
            body *models_pkg.MagicLinkAuthenticate)(*models_pkg.MagicLinkAuthenticateResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| token |  ``` Required ```  | TODO: Add a parameter description |
| body |  ``` Required ```  | Magic link object |


#### Example Usage

```go
token := "token"
var body *models_pkg.MagicLinkAuthenticate

var result *models_pkg.MagicLinkAuthenticateResponse
result,_ = magicLinks.PostUserMagicLinkAuthenticate(token, body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not found |
| 429 | Too many requests |
| 500 | Internal server error |



[Back to List of Controllers](#list_of_controllers)

## <a name="emails_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".emails_pkg") emails_pkg

### Get instance

Factory for the ``` EMAILS ``` interface can be accessed from the package emails_pkg.

```go
emails := emails_pkg.NewEMAILS()
```

### <a name="delete_email"></a>![Method: ](https://apidocs.io/img/method.png ".emails_pkg.DeleteEmail") DeleteEmail

> Remove an email from a given user.


```go
func (me *EMAILS_IMPL) DeleteEmail(
            emailId string,
            userId string)(*models_pkg.UserEmailDeleteResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| emailId |  ``` Required ```  | The email_id to be deleted. |
| userId |  ``` Required ```  | The user_id to delete an email from. |


#### Example Usage

```go
emailId := "email-test-c1a1d554-a93c-11ea-bb37-0242ac130002"
userId := "user-test-b8797f2c-a93c-11ea-bb37-0242ac130002"

var result *models_pkg.UserEmailDeleteResponse
result,_ = emails.DeleteEmail(emailId, userId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not found |
| 429 | Too many requests |
| 500 | Internal server error |



### <a name="create_send_email_verification"></a>![Method: ](https://apidocs.io/img/method.png ".emails_pkg.CreateSendEmailVerification") CreateSendEmailVerification

> Prompt for a verification email to be sent to the user to confirm the correct email was entered. The email must be verified before the user needs to login next.


```go
func (me *EMAILS_IMPL) CreateSendEmailVerification(
            userId string,
            emailId string)(*models_pkg.SendVerificationResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| userId |  ``` Required ```  | The user_id for the user to fetch. |
| emailId |  ``` Required ```  | The email_id for the given user to verify. |


#### Example Usage

```go
userId := "user-test-b8797f2c-a93c-11ea-bb37-0242ac130002"
emailId := "email-test-c1a1d554-a93c-11ea-bb37-0242ac130002"

var result *models_pkg.SendVerificationResponse
result,_ = emails.CreateSendEmailVerification(userId, emailId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not found |
| 429 | Too many requests |
| 500 | Internal server error |



### <a name="create_verify_email"></a>![Method: ](https://apidocs.io/img/method.png ".emails_pkg.CreateVerifyEmail") CreateVerifyEmail

> Verify that a user supplied the correct email during signup.


```go
func (me *EMAILS_IMPL) CreateVerifyEmail(token string)(*models_pkg.VerifyEmailResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| token |  ``` Required ```  | The token used to verify user's email. |


#### Example Usage

```go
token := "KKFa7u0KgAgHGXkZ77gOEd8YjyzzcC1rvMINgsZuIxM"

var result *models_pkg.VerifyEmailResponse
result,_ = emails.CreateVerifyEmail(token)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not found |
| 429 | Too many requests |
| 500 | Internal server error |



[Back to List of Controllers](#list_of_controllers)



