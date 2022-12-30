# Cert

Get the root certificate

**URL** : `/api/v1/cert/root`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Content example**

```
root.zip
```

## Error Response
- **Condition** : Not found

**Code** : `404 NOT FOUND`

**Content** :

```
Could not be found
```

- **Condition** : Unexpected exceptions

**Code** : `500 INTERNAL SERVER ERROR`

**Content** :

```
An error has occurred, view the logs to get more informations
```