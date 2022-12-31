# Automate SSL Certificates API
The api is based on the [automate-ssl-certificates repo](https://github.com/PsclDev/automate-ssl-certificates) from a friend and myself. I just wanted to do something with go and decided to wrap those two scripts with an api for an easier and faster use.

## Image
```
docker pull ghcr.io/pscldev/automate-ssl-certificates-api:latest
```

## Envs
| Env | Type | Required |
|---|---|---|
| PORT | int | ✅ |
| PROD | bool | ❌ |
| NETLOG_URL | string | ❌ |
| SENTRY_DSN | string | ❌ |

## Routes
ℹ️ The project contains a postman collection which you can import to have an overview and example for each request.

⚠️ The `config` and `cert` routes contain a prefix: `/api/v1`

### **Config**
* [Get makeRoot.sh](docs/config-get.md) : `GET /config/root`
* [Get makeCertificate.sh](docs/config-get.md) : `GET /certificate`
* [Post your domain name settings](docs/config-post.md) : `POST /config`

### **Cert**
* [Get all certs](docs/cert-get.md) : `GET /cert`
* [Get root cert archive](docs/cert-get-root.md) : `GET /cert/root`
* [Get cert by name](docs/cert-get.md) : `GET /cert/:name`
* [Create a new certificate](docs/config-get.md) : `POST /cert`
* [Recreate a existing certificate](docs/config-get.md) : `PATCH /cert`
* [Delete a cert by name](docs/cert-delete.md) : `DELETE /cert/:name`

### **Health w/o prefix**
* [Check API health](docs/health.md) : `GET /health`