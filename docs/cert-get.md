# Cert

Get the certificates

**URL** : `/api/v1/cert` ; `/api/v1/cert/:name`

**Method** : `GET`

**Query Params** : `res` 

**Params Options** json (default), file

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
name.zip
```

## Error Response
- **Condition** : Not found

**Code** : `404 NOT FOUND`

**Content** :

```
Could not be found
```

- **Condition** : Invalid response type

**Code** : `406 NOT ACCEPTABLE`

**Content** :

```
Invalid response type requested - 'json' (default) or 'file' are allowed
```

- **Condition** : Unexpected exceptions

**Code** : `500 INTERNAL SERVER ERROR`

**Content** :

```
An error has occurred, view the logs to get more informations
```