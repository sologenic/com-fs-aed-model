module github.com/sologenic/com-fs-aed-model

go 1.23.8

require (
	github.com/google/go-cmp v0.7.0
	github.com/shopspring/decimal v1.4.0
	github.com/sologenic/com-fs-asset-model v0.0.0-20250430175832-81d20906dca7
	github.com/sologenic/com-fs-utils-lib/go/cache v0.0.0-20250502172059-baf86ddc0d0a
	github.com/sologenic/com-fs-utils-lib/go/grpc-client v0.0.0-20250417194121-8d91c0b8405b
	github.com/sologenic/com-fs-utils-lib/go/logger v0.0.0-20240906214824-8118c42c5b35
	github.com/sologenic/com-fs-utils-lib/models/metadata v0.0.0-20250417194121-8d91c0b8405b
	google.golang.org/grpc v1.71.1
	google.golang.org/protobuf v1.36.6
)

// TODO: Remove the replace directive when the com-fs-asset-model module is published
replace github.com/sologenic/com-fs-asset-model => ../com-fs-asset-model

require (
	cloud.google.com/go/compute/metadata v0.6.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	github.com/sologenic/com-fs-utils-lib/models/audit v0.0.0-20250131210530-7306ae3e5762 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/oauth2 v0.25.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/api v0.149.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
)
