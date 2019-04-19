# Dymo

[![OpenFaaS](https://img.shields.io/badge/openfaas-cloud-blue.svg)](https://www.openfaas.com)

A simple domain name reversal function.

# Invoke the function

```sh
$ curl https://rgee0.o6s.io/dymo -d 'rgee0.o6s.io'
io.o6s.rgee0
```

It will also accommodate http:// and https:// schemes:

```sh
$ curl https://rgee0.o6s.io/dymo -d 'https://rgee0.o6s.io'
io.o6s.rgee0
```

And strip tailing paths:

```sh
$ curl https://rgee0.o6s.io/dymo -d 'https://rgee0.o6s.io/dymo/makes/labels'
io.o6s.rgee0
```
