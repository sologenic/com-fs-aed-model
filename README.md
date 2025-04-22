# AED model

The Asset Epoch Data model proto provides all the functionality required to interact with the aed service.

## Client

The client for golang and typescript are provided in the `/client` directory.

## Start the service

The client is self initializing. For this it requires the following environment variable:

- `AED_STORE` - The host of the service. Host is in the format `host:port`
- `AED__STORE_TESTING` - Starts the GRPC service in testmode with an in memory buffer. Presence of the variable with the value `TRUE` is required to put the client in test mode.

If the client is included and the non of the env variable is set, the client panics.

## Building the required files

Once the proto file is updated, the following files need to be generated:

* go - `aed.pb.go`, `aed_grpc.pb.go`
* typescript - `aed.ts`

Generating these files can be done by running a build script with following command, which will also generate a build dir and add newly generated files and dir to git commit.

```sh
./bin/build.sh
```
