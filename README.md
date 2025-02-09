# HTTB

> [HTTB.dev](https://httb.dev) - Free and open-source HTTP bin.  
> Use HTTB endpoints to test your HTTP client and tools.

> [!WARNING]
> HTTB is work in progress. The API is not stable and some features are not yet implemented.

## Self host

HTTB can be self-hosted using [Docker](https://docker.com).

<table>
<tr>
<th>docker command</th>
<th>docker compose</th>
</tr>
<tr>
<td>

```sh
docker run -p 8080:8080 marvinjwendt/httb
```

</td>
<td>

```yml
services:
  httb:
    image: marvinjwendt/httb
    ports:
      - 8080:8080
```

</td>
</tr>
</table>

### Environment Variables

| Name                                 | Description                                                | Default                 |
|--------------------------------------|------------------------------------------------------------|-------------------------|
| `ADDR`                               | Address to listen to                                       | `0.0.0.0:8080`          |
| `LOG_LEVEL`                          | Log level (can be one of `debug`, `info`, `warn`, `error`) | `info`                  |
| `LOG_FORMAT`                         | Log format (can be one of `logfmt`, `json`)                | `logfmt`                |
| `TIMEOUT`                            | Timeout for responses                                      | `2m`                    |
| `SWAGGER_DEFAULT_SERVER`             | Default server for the Swagger UI (/docs)                  | `http://localhost:8080` |
| `SWAGGER_DEFAULT_SERVER_DESCRIPTION` | Description for the default Swagger UI server              | `for local testing`     |
