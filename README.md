Revoker
=======

Small Go service for revoking JWT tokens with [KrakenD](https://www.krakend.io/). To learn more about KrakenD's JWT revocation with bloom filters, check out their [documentation](https://www.krakend.io/docs/authorization/revoking-tokens/).

## Work In Progress

**This project is a work in progress and should not be considered complete, tested, functional or even remotely close to production ready.**

## Configuration

The service requires one environment variable, `BLOOM_SERVER`, which should contain the address of your KrakenD bloom filter RPC server.

```bash
BLOOM_SERVER=localhost:1234
```

### Routes

The service maintains three (3) routes:

#### `/ (GET)`

The base route returns a 200 response to any GET request. This request exists primarily for health checks.

**Response:**

```json
{
  "status": 200,
  "message": "OK"
}
```

#### `/add (POST)`

The add route is the primary service route, it allows you to add revocations to the bloom filter.

**Payload:**

```json
{
  "key": "key",
  "subject": "subject"
}
```

`key`: Any JWT claim you've configured in your bloom filter config's `token_key` property (refer to the [documentation](https://www.krakend.io/docs/authorization/revoking-tokens/) for more details).

`subject`: The data associated to the key.

**Response:**

```json
{
  "status": 201,
  "message": "Added"
}
```


#### `/check (POST)`

The check route allows you to check if a revocation exists in the bloom filter.

**Payload:**

```json
{
  "key": "key",
  "subject": "subject"
}
```

`key`: Any JWT claim you've configured in your bloom filter config's `token_key` property (refer to the [documentation](https://www.krakend.io/docs/authorization/revoking-tokens/) for more details).

`subject`: The data associated to the key.

**Response:**

```json
{
    "status": 200,
    "subject": "key-subject",
    "exists": true
}
```

## Development Environment

During active development, the `docker-compose.yml` network is configured for development with our own projects. This will be removed once ready for general release.
