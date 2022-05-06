# tcld (`Beta`)
A cli tool for managing Temporal Cloud namespaces.

> This cli tool is currently in `beta` and access to Temporal Cloud via the cli is restricted. Please reach out to temporal-cloud support for more information.

# Installation
## Build from source
1. Verify that you have Go 1.18+ installed. If `go` is not installed, follow instructions on [the Go website](https://golang.org/doc/install).
```
go version
```
2. Clone the `tcld` repository and run `make`.
```
git clone https://github.com/temporalio/tcld.git
cd tcld
make
```
3. Copy the tcld executable to any directory that appears in the PATH environment variable; for example, `/usr/local/bin/`.
```
cp tcld /usr/local/bin/tcld
```
4. Run `tcld version` to check if it worked.
```
tcld version
```

# Authentication and Login
In order to use the cli you must first login by running the following command:
```
tcld login
```
You will be sent a link to confirm your device code and login. After logging in, you are now authenticated and can make requests with this cli.

# Namespace Management

### List namespaces user has access to:
```
tcld namespace list
```

### Get namespace information:
```
tcld namespace get -n <namespace>
```

### Update the CA certificate:
```
tcld namespace update accepted-client-ca set -n <namespace> --ca-certificate-file <ca-pem-filepath>
```
> :warning: If the update removes a certificate, any clients (tctl/workers) still using the removed certificate will fail to connect to the namespace after the update completes.

### Add new search attributes:
```
tcld namespace update search-attributes add -n <namespace> --sa "<attribute-name>=<search-attribute-type>" --sa "<attribute-name>=<search-attribute-type>"
```
Supported search attribute types: `SearchAttributeTypeKeyword SearchAttributeTypeText SearchAttributeTypeInt SearchAttributeTypeDouble SearchAttributeTypeDatetime SearchAttributeTypeBool`

### Rename existing search attribute:
```
tcld namespace update search-attributes rename -n <namespace> --existing-name <existing-attribute-name> --new-name <new-attribute-name>
```
> :warning: Any workflows that are using the old search attribute name will fail after the update.

# Asynchronous Operations
Any update operations making changes to the namespaces hosted on Temporal Cloud are asynchronous. Such operations are tracked using a `request-id` that can be passed in when invoking the update operation or will be auto-generated by the server if one is not specified. Once an asynchronous request is initiated, a `request-id` is returned. Use the `request get` command to query the status of an asynchronous request.
```
tcld request get -r <request-id> -n <namespace>
```

# License

MIT License, please see [LICENSE](https://github.com/temporalio/tcld/blob/master/LICENSE) for details.
