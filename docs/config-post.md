# Config

Set your domain names

**URL** : `/api/v1/config`

**Method** : `POST`

**Data constraints**

```json
{
    "country": "[iso 3166 alpha-2]",
    "state": "[iso 3166 alpha-2]",
    "location": "[iso 3166 alpha-2]",
    "domain": "[name]",
    "tld": "[tld]",
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```
Config was set
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