module user

go 1.15

require (
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef // indirect
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/micro/micro/v3 v3.0.0-beta.7
	github.com/satori/go.uuid v1.2.0
	gorm.io/gorm v1.20.6
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
