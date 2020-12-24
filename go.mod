module github.com/logzio/jaeger-logzio

go 1.12

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/avast/retry-go v2.6.0+incompatible
	github.com/beeker1121/goque v2.0.1+incompatible // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/hashicorp/go-hclog v0.14.0
	github.com/jaegertracing/jaeger v1.21.0
	github.com/logzio/logzio-go v0.0.0-20200316143903-ac8fc0e2910e
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/olivere/elastic v6.2.34+incompatible
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/shirou/gopsutil v2.18.12+incompatible // indirect
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.5.1
	github.com/syndtr/goleveldb v1.0.0 // indirect
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/logzio/logzio-go => github.com/albertteoh/logzio-go d8c39718d666153ebe82d651fa9f2437c9c7fdd9
