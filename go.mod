module fleet_management

go 1.22.5

require (
	github.com/golang-migrate/migrate/v4 v4.15.2
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.10.9
	github.com/stretchr/testify v1.9.0
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
)

replace (
	github.com/golang-migrate/migrate/v4/database/postgres => github.com/golang-migrate/migrate/database/postgres v1.2.2
	github.com/golang-migrate/migrate/v4/source/file => github.com/golang-migrate/migrate/source/file v1.4.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
