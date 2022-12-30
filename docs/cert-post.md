# Cert

Create a new cert

**URL** : `/api/v1/cert`

**Method** : `POST` ; `PATCH`

**Data constraints**

```json
{
    "name": "[identifier]",
    "ip": "[ip addr]",
    "dns": "[dns(.)domain.tld]"
}
```

## Success Response

**Code** : `200 OK`

**JSON Content example**

```json
{
    "name": "test",
    "dns": "test.tld",
    "ip": "10.0.0.1"
}
```

**File Content example**

```
test.zip
```

## Error Response

- **Condition** : Missing json key/value

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
    "key": "is missing, but required"
}
```

- **Condition** : Invalid response type

**Code** : `406 NOT ACCEPTABLE`

**Content** :

```
Invalid response type requested - 'json' (default) or 'file' are allowed
```

- **Condition** : Invalid or missing body

**Code** : `409 CONFLICT`

**Content** :

```
Cert already exists, if you want to recreate use PATCH request
```

- **Condition** : Invalid or missing body

**Code** : `422 UNPROCESSABLE ENTITY`

**Content** :

```
Body couldn't be parsed, check the documentations for the correct request body
```

- **Condition** : Unexpected exceptions

**Code** : `500 INTERNAL SERVER ERROR`

**Content** :

```
An error has occurred, view the logs to get more informations
```