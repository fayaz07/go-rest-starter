# the Authentication layer

1. [API Key](#api-key)
2. [Init authentication](#init-authentication)
3. [Registration](#registration)
4. [Login](#login)
5. [Logout](#logout)
6. [Refresh token](#refresh-token)
7. [Reset password](#reset-password)
8. [Change password](#change-password)
9. [Delete account](#delete-account)

### API Key

1. There will be a super-user account, which is used to handle the creation of API Keys, which will be shipped over to the client.

> Note: Only one/two super-user accounts are allowed(based on requirement).

### Init authentication

1. The API Key will be used to authenticate the client. (First handshake)
2. Client sends the API Key to the server along with information of the host, device details, etc (encrypted).
3. Server checks if the API Key is valid.
4. If API Key is valid, then store the client information in the database.
5. The server sends the client a response with the authentication token. (Client uses that auth-token to hit register and login endpoints).
6. The auth-token comes with an expiry time of 30 minutes.

Prototype of request:

```http
POST http://localhost:8080/api/v1/auth/init

Headers:
x-api-token: <token>
```

Body (for browser):

```json
{
  "host": "localhost",
  "device": "desktop",
  "os": "windows",
  "browser": "chrome",
  "browser_version": "80.0.3987.149"
}
```

Body (for mobile):

```json
{
  "device": "mobile",
  "os": "android",
  "version": "7.0",
  "model": "Samsung Galaxy S7",
  "manufacturer": "Samsung",
  "connection": "wifi"
}
```

### Registration

1. Client sends the auth-token to the server along with the registration details.
2. Server checks if the auth-token is valid.
3. A max of 10 requests are allowed per auth-token and it's subsequent IP address.
4. Once all the information is properly in place and registration is successful, the server should close any registration requests from the same IP address for one hour.

Prototype of request:

```http
POST http://localhost:8080/api/v1/auth/register

Headers:
x-auth-token: <token>
```

body:

```json
{
  "email": "user@domain.com",
  "password": "Password@1"
}
```

### Login

1. Client sends the auth-token to the server along with the login details.
2. Server checks if the auth-token is valid.
3. A max of 10 requests are allowed per auth-token and it's subsequent IP address.
4. If the user has not registered yet, then the server should send a response with the error code `REGISTRATION_REQUIRED`.
5. If the user sends invalid password for 5 attempts with the same auth-token, then the server should send a response with the error code `INVALID_PASSWORD` and locks the user's account until one hour.

Prototype of request:

```http
POST http://localhost:8080/api/v1/auth/login

Headers:
x-auth-token: <token>
```

body:

```json
{
  "email": "user@domain.com",
  "password": "Password@1"
}
```
