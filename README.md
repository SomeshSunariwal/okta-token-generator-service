# OKTA TOKEN APIs Simple

This is a test application which is able to generate token when you pass the authcode in the body.

```json
{
  "code": "abcdefghijk"
}
```

Above "code" is a OKTA Auth Code

## How to get Auth Code

Use the below URL To get the login page of your okta domain.

```
https://{domain-name}/oauth2/default/v1/authorize?client_id={app-clinet-id}&response_type=code&scope={scope-name}&redirect_uri={app-redirect-uri}&state=1234
```

After Login will be redirected back to HTML page with the Auth Code in the URL.

## Setup
update the variables in this file for local testing : https://github.com/SomeshSunariwal/okta-token-generator-service/blob/master/config/config.go

*** be carefull with ClientId and Client Secret in case of multiple okta applications

# Postmen Collection

```
https://www.getpostman.com/collections/9d5df68098927fdae3c3
```
