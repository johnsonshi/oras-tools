# ORAS Tools

Helper tools for [ORAS](https://oras.land).

## Quick Start

### Install

To install, run the following commands.

```bash
curl -LO https://github.com/johnsonshi/oras-tools/releases/download/v0.0.1/oras-tools
chmod +x oras-tools
sudo mv oras-tools /usr/local/bin
```

## Commands

### Delete

Delete [ORAS artifacts](https://oras.land/#how-oras-works) from a registry.

To delete an ORAS artifact (referenced by its hash `$digest`) in `$registry_url` and `$repository_name`:

```bash
./bin/oras-tools delete \
  --username "$registry_username" \
  --password "$registry_password" \
  --registry "$registry_url" \
  --repository "$repository_name" \
  --digest "$digest"
```
