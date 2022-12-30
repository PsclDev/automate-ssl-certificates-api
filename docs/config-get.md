# Config

Get the make scripts with the current configuration

**URL** : `/api/v1/config/root` ; `/api/v1/config/certificate`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Content example**

```
#!/bin/bash

# !!! Update this variable !!!
yourPath=_data/cert/root
# -------

mkdir -p  $yourPath
cd $yourPath

cat > root.conf <<EOF
[req]
prompt = no
default_bits = 4096
default_md = sha512
distinguished_name = dn
x509_extensions = my_extensions

# !!! Update this section !!!
[dn] 
C=US
ST=NY
L=NY
O=crazy
CN=crazy.tld
# -------

[my_extensions]
basicConstraints = critical, CA:TRUE
subjectKeyIdentifier = hash
authorityKeyIdentifier = keyid:always, issuer:always
keyUsage = critical,  cRLSign, digitalSignature, keyCertSign

EOF

openssl genrsa -out root.key 4096
openssl req -x509 -new -nodes -key root.key -sha512 -days 3650 -out root.pem -config root.conf

openssl x509 -in root.pem -outform der -out root.crt

echo "[*] Created root certificates"
```

## Error Response

- **Condition** : Unexpected exceptions

**Code** : `500 INTERNAL SERVER ERROR`

**Content** :

```
An error has occurred, view the logs to get more informations
```