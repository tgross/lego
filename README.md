<div align="center">
  <img alt="lego logo" src="./docs/static/images/lego-logo.min.svg">
  <p>Automatic Certificates and HTTPS for everyone.</p>
</div>

# Lego

Let's Encrypt client and ACME library written in Go.

[![Go Reference](https://pkg.go.dev/badge/github.com/go-acme/lego/v4.svg)](https://pkg.go.dev/github.com/go-acme/lego/v4)
[![Build Status](https://github.com//go-acme/lego/workflows/Main/badge.svg?branch=master)](https://github.com//go-acme/lego/actions)
[![Docker Pulls](https://img.shields.io/docker/pulls/goacme/lego.svg)](https://hub.docker.com/r/goacme/lego/)

## Features

- ACME v2 [RFC 8555](https://www.rfc-editor.org/rfc/rfc8555.html)
  - Support [RFC 8737](https://www.rfc-editor.org/rfc/rfc8737.html): TLS Application‑Layer Protocol Negotiation (ALPN) Challenge Extension
  - Support [RFC 8738](https://www.rfc-editor.org/rfc/rfc8738.html): certificates for IP addresses
  - Support [draft-ietf-acme-ari-03](https://datatracker.ietf.org/doc/draft-ietf-acme-ari/): Renewal Information (ARI) Extension
- Register with CA
- Obtain certificates, both from scratch or with an existing CSR
- Renew certificates
- Revoke certificates
- Robust implementation of all ACME challenges
  - HTTP (http-01)
  - DNS (dns-01)
  - TLS (tls-alpn-01)
- SAN certificate support
- [CNAME support](https://letsencrypt.org/2019/10/09/onboarding-your-customers-with-lets-encrypt-and-acme.html) by default
- Comes with multiple optional [DNS providers](https://go-acme.github.io/lego/dns)
- [Custom challenge solvers](https://go-acme.github.io/lego/usage/library/writing-a-challenge-solver/)
- Certificate bundling
- OCSP helper function

## Installation

How to [install](https://go-acme.github.io/lego/installation/).

## Usage

- as a [CLI](https://go-acme.github.io/lego/usage/cli)
- as a [library](https://go-acme.github.io/lego/usage/library)

## Documentation

Documentation is hosted live at https://go-acme.github.io/lego/.

## DNS providers

Detailed documentation is available [here](https://go-acme.github.io/lego/dns).

<!-- START DNS PROVIDERS LIST -->

<table><tr>
  <td><a href="https://go-acme.github.io/lego/dns/route53/">Amazon Route 53</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/bunny/">Bunny</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/exec/">External program</a></td>
  <td><a href="https://go-acme.github.io/lego/dns/acme-dns/">Joohoi&#39;s ACME-DNS</a></td>
</tr><tr>
  <td><a href="https://go-acme.github.io/lego/dns/manual/">Manual</a></td>
  <td></td>
  <td></td>
  <td></td>
</tr></table>

<!-- END DNS PROVIDERS LIST -->

If your DNS provider is not supported, please open an [issue](https://github.com/go-acme/lego/issues/new?assignees=&labels=enhancement%2C+new-provider&template=new_dns_provider.md).
