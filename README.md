# OKTA TOKEN APIs Simple

This is a test application which is able to generate token when you pass the authcode in the body.

```json
{
  "code": "abcdefghijk"
}
```

Above Code is a OKTA Auth Code

## How to get AUTH Code

Use the below URL To get the login page of your okta domain.

```
https://{domain-name}/oauth2/default/v1/authorize?client_id={app-clinet-id}&response_type=code&scope={scope-name}&redirect_uri={app-redirect-uri}&state=1234
```

# Postmen Collection

```
https://www.getpostman.com/collections/9d5df68098927fdae3c3
```
