# Planon API

API for the Planon BookMySpace system.

## Usage

* Copy `config.toml.sample` to `config.toml` and adjust the values.
* `go build && ./planon` or use the prebuilt Linux x64 binary at the releases page
* See [API.md](https://github.com/mhvis/planon/blob/master/API.md) for the server endpoints

## Reverse engineering notes

Implementation is based on reverse engineering of the Planon Android app using
a HTTP proxy. By using Charles Proxy I had problems somehow with the SSL
certificate and couldn't get it to work, but using mitmproxy worked fine.

## License

Unlicensed for now, send me a mail if you want to use it.
