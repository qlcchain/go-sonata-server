# Go Sonata API

[![Actions Status](https://github.com/qlcchain/go-sonata-server/workflows/Main%20workflow/badge.svg)](https://github.com/qlcchain/go-sonata-server/actions)

## Key Features

- [x] support Serviceability(Address Validation, Site Queries, Product Offering Qualification) API with mock data
- [x] support Quote API
- [x] support ProductOrder API
- [x] support Inventory API with mock data
- [ ] Trouble-ticketing Billing
- [ ] Contract & Catalog
- [x] support JWT authentication

For more details, please check [swagger spec](spec/all.yaml)

## Build and Run
```bash
make clean build
./build/gsonata
```

## CLI
```bash
gsonata 0.0.1-6e3eb21.2020-05-25T11:45:37Z+0800Usage:
  gsonata [OPTIONS]

A set of APIs based on the LSO Reference Architecture for
Serviceability (Address Validation, Site Queries, Product Offering Qualification) |
Quoting | Product Inventory | Ordering | Trouble-ticketing Billing | Contract & Catalog



Application Options:
      --scheme=            the listeners to enable, this can be repeated and defaults to the schemes in the swagger spec
      --cleanup-timeout=   grace period for which to wait before killing idle connections (default: 10s)
      --graceful-timeout=  grace period for which to wait before shutting down the server (default: 15s)
      --max-header-size=   controls the maximum number of bytes the server will read parsing the request header's keys and values, including the
                           request line. It does not limit the size of the request body. (default: 1MiB)
      --socket-path=       the unix socket to listen on (default: /var/run/sonata.sock)
      --host=              the IP to listen on (default: localhost) [$HOST]
      --port=              the port to listen on for insecure connections, defaults to a random value [$PORT]
      --listen-limit=      limit the number of outstanding requests
      --keep-alive=        sets the TCP keep-alive timeouts on accepted connections. It prunes dead TCP connections ( e.g. closing laptop
                           mid-download) (default: 3m)
      --read-timeout=      maximum duration before timing out read of the request (default: 30s)
      --write-timeout=     maximum duration before timing out write of the response (default: 60s)
      --tls-host=          the IP to listen on for tls, when not specified it's the same as --host [$TLS_HOST]
      --tls-port=          the port to listen on for secure connections, defaults to a random value [$TLS_PORT]
      --tls-certificate=   the certificate to use for secure connections [$TLS_CERTIFICATE]
      --tls-key=           the private key to use for secure connections [$TLS_PRIVATE_KEY]
      --tls-ca=            the certificate authority file to be used with mutual tls auth [$TLS_CA_CERTIFICATE]
      --tls-listen-limit=  limit the number of outstanding requests
      --tls-keep-alive=    sets the TCP keep-alive timeouts on accepted connections. It prunes dead TCP connections ( e.g. closing laptop
                           mid-download)
      --tls-read-timeout=  maximum duration before timing out read of the request
      --tls-write-timeout= maximum duration before timing out write of the response

jwt:
  -K, --key=               private key(elliptic P521) to sign JWT token

debug:
  -V, --verbose            show verbose debug information
  -m, --mock               mock sample data
  -f, --file               set database as file mode

server:
      --domain=            bind domain (default: localhost)
      --allowedOrigins=    AllowedOrigins of CORS (default: *)

Help Options:
  -h, --help               Show this help message

```

## Docker
You can build the docker image yourself or download it from docker hub
### Build docker images

```bash
 docker build -t qlcchain/go-sonata-server .
```

### Download docker images from docker hub

```bash
docker pull qlcchain/go-sonata-server:latest
```

### Start docker container
You can choose to run a normal node without an account or run an account node.

#### Run a normal node without an account

```bash
docker container run -d --name go-sonata-server \
    -p 127.0.0.1:9999:9999 \
    qlcchain/go-sonata-server:latest
```
#### Run node by Docker Compose

```bash
docker-compose down -v && docker-compose up -d
```

For more infomartion, please check [docker-compose](docker-compose.yml)

## Links & Resources
* [Yellow Paper](https://github.com/qlcchain/YellowPaper)
* [QLC Website](https://qlcchain.org)
* [Reddit](https://www.reddit.com/r/QLCChain/)
* [Medium](https://medium.com/qlc-chain)
* [Twitter](https://twitter.com/QLCchain)
* [Telegram](https://t.me/qlinkmobile)
