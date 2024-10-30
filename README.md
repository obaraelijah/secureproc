## SecureProc
A secure job worker service that enables authorized users to run arbitrary Linux processes on remote hosts with resource constraints and output streaming capabilities.

### Features
* Secure process execution with resource isolation
* Process output streaming
* Resource control via cgroups (CPU, Memory, Block I/O)
* mTLS authentication and authorization
* Multiple concurrent client support
* Process namespace isolation

### Components
* **Library**: Core Go library for job management
* **API**: gRPC API for secure client-server communication
* **Client**: Command-line tool for interacting with the service

### Requirements
* Go 1.19 or higher
* Linux kernel with cgroup support
* OpenSSL for certificate generation

### Building 
`make all`

> Building individual components
`make client`
`make server`

### Development
`make deps`

> Run tests
`make test`
