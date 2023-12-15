module github.com/richelieu-yang/chimera/v2

go 1.21

require (
	fyne.io/fyne/v2 v2.4.2
	github.com/alwindoss/morse v1.0.1
	github.com/apache/pulsar-client-go v0.11.1
	github.com/apache/rocketmq-clients/golang/v5 v5.0.1-rc.5
	github.com/bytedance/sonic v1.10.2
	github.com/dablelv/cyan v0.0.52
	github.com/davidbyttow/govips/v2 v2.13.0
	github.com/deckarep/golang-set/v2 v2.5.0
	github.com/duke-git/lancet/v2 v2.2.8
	github.com/dustin/go-humanize v1.0.1
	github.com/emersion/go-imap v1.2.1
	github.com/emersion/go-imap-id v0.0.0-20190926060100-f94a56b9ecde
	github.com/emersion/go-message v0.17.0
	github.com/fatih/structs v1.1.0
	github.com/fsnotify/fsnotify v1.7.0
	github.com/gabriel-vasile/mimetype v1.4.3
	github.com/gin-contrib/cors v1.5.0
	github.com/gin-contrib/gzip v0.0.6
	github.com/gin-contrib/pprof v1.4.0
	github.com/gin-contrib/sse v0.1.0
	github.com/gin-gonic/gin v1.9.1
	github.com/go-kratos/kratos/v2 v2.7.2
	github.com/go-playground/validator/v10 v10.16.0
	github.com/go-redsync/redsync/v4 v4.11.0
	github.com/gogf/gf/v2 v2.5.7
	github.com/golang-jwt/jwt/v5 v5.2.0
	github.com/golang-module/carbon/v2 v2.2.14
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.5.0
	github.com/google/wire v0.5.0
	github.com/gorilla/securecookie v1.1.2
	github.com/gorilla/sessions v1.2.2
	github.com/gorilla/websocket v1.5.1
	github.com/h2non/bimg v1.1.9
	github.com/hashicorp/go-hclog v1.5.0
	github.com/hashicorp/raft v1.6.0
	github.com/hashicorp/raft-boltdb v0.0.0-20231211162105-6c830fa4535e
	github.com/hibiken/asynq v0.24.1
	github.com/imroc/req/v3 v3.42.2
	github.com/jedib0t/go-pretty/v6 v6.4.9
	github.com/jinzhu/copier v0.4.0
	github.com/joho/godotenv v1.5.1
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/json-iterator/go v1.1.12
	github.com/klauspost/compress v1.17.4
	github.com/klauspost/cpuid/v2 v2.2.6
	github.com/linxGnu/grocksdb v1.8.10
	github.com/lionsoul2014/ip2region/binding/golang v0.0.0-20231013030745-3066d243cd04
	github.com/matoous/go-nanoid/v2 v2.0.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/modern-go/reflect2 v1.0.2
	github.com/nacos-group/nacos-sdk-go/v2 v2.2.4
	github.com/natefinch/lumberjack/v3 v3.0.0-alpha
	github.com/oklog/ulid/v2 v2.1.0
	github.com/panjf2000/ants/v2 v2.9.0
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8
	github.com/redis/go-redis/v9 v9.3.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/rs/xid v1.5.0
	github.com/saintfish/chardet v0.0.0-20230101081208-5e3ef4b5456d
	github.com/samber/lo v1.39.0
	github.com/shirou/gopsutil/v3 v3.23.11
	github.com/shopspring/decimal v1.3.1
	github.com/sirupsen/logrus v1.9.3
	github.com/sony/sonyflake v1.2.0
	github.com/spf13/afero v1.11.0
	github.com/spf13/cast v1.6.0
	github.com/spf13/cobra v1.8.0
	github.com/spf13/viper v1.18.1
	github.com/swaggo/files v1.0.1
	github.com/swaggo/gin-swagger v1.6.0
	github.com/swaggo/swag v1.16.2
	github.com/tidwall/gjson v1.17.0
	github.com/xuri/excelize/v2 v2.8.0
	github.com/zeromicro/go-zero v1.6.0
	github.com/zeromicro/zero-contrib/logx/logrusx v0.0.0-20231030135404-af9ae855016f
	go.etcd.io/etcd/client/v3 v3.5.11
	go.opentelemetry.io/otel v1.21.0
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0
	go.opentelemetry.io/otel/metric v1.21.0
	go.opentelemetry.io/otel/sdk v1.21.0
	go.opentelemetry.io/otel/trace v1.21.0
	go.uber.org/automaxprocs v1.5.3
	go.uber.org/zap v1.26.0
	golang.org/x/exp v0.0.0-20231206192017-f3f8817b8deb
	golang.org/x/image v0.14.0
	golang.org/x/net v0.19.0
	golang.org/x/text v0.14.0
	google.golang.org/grpc v1.60.0
	google.golang.org/protobuf v1.31.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/mysql v1.5.2
	gorm.io/gorm v1.25.5
	sigs.k8s.io/yaml v1.4.0
)

require (
	contrib.go.opencensus.io/exporter/ocagent v0.6.0 // indirect
	fyne.io/systray v1.10.1-0.20231115130155-104f5ef7839e // indirect
	github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // indirect
	github.com/99designs/keyring v1.2.1 // indirect
	github.com/AthenZ/athenz v1.10.39 // indirect
	github.com/DataDog/zstd v1.5.2 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/alibabacloud-go/debug v0.0.0-20190504072949-9472017b5c68 // indirect
	github.com/alibabacloud-go/tea v1.1.17 // indirect
	github.com/alibabacloud-go/tea-utils v1.4.4 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1800 // indirect
	github.com/aliyun/alibabacloud-dkms-gcs-go-sdk v0.2.2 // indirect
	github.com/aliyun/alibabacloud-dkms-transfer-go-sdk v0.1.7 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/ardielle/ardielle-go v1.5.2 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.4.0 // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.0 // indirect
	github.com/cloudflare/circl v1.3.3 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/danieljoos/wincred v1.1.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dchest/siphash v1.2.3 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dvsekhvalnov/jose2go v1.5.0 // indirect
	github.com/emersion/go-sasl v0.0.0-20200509203442-7bfe0ed36a21 // indirect
	github.com/emersion/go-textwrapper v0.0.0-20200911093747-65d896831594 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/fredbi/uri v1.0.0 // indirect
	github.com/fyne-io/gl-js v0.0.0-20220119005834-d2da28d9ccfe // indirect
	github.com/fyne-io/glfw-js v0.0.0-20220120001248-ee7290d23504 // indirect
	github.com/fyne-io/image v0.0.0-20220602074514-4956b0afb3d2 // indirect
	github.com/gaukas/godicttls v0.0.4 // indirect
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20221017161538-93cebf72946b // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/go-text/render v0.0.0-20230619120952-35bccb6164b8 // indirect
	github.com/go-text/typesetting v0.0.0-20230616162802-9c17dd34aa4a // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt v3.2.1+incompatible // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/pprof v0.0.0-20230901174712-0191c66da455 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.0 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/go-msgpack/v2 v2.1.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.3.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jsummers/gobmp v0.0.0-20151104160322-e2ba15ffa76e // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/linkedin/goavro/v2 v2.9.8 // indirect
	github.com/lufia/plan9stats v0.0.0-20230326075908-cb1d2100619a // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/onsi/ginkgo/v2 v2.12.0 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20221212215047-62379fc7944b // indirect
	github.com/prometheus/client_golang v1.17.0 // indirect
	github.com/prometheus/client_model v0.4.1-0.20230718164431-9a2bf3000d16 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.11.1 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/quic-go/qtls-go1-20 v0.3.3 // indirect
	github.com/quic-go/quic-go v0.38.1 // indirect
	github.com/refraction-networking/utls v1.5.3 // indirect
	github.com/richardlehane/mscfb v1.0.4 // indirect
	github.com/richardlehane/msoleps v1.0.3 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/srwiley/oksvg v0.0.0-20221011165216-be6e8873101c // indirect
	github.com/srwiley/rasterx v0.0.0-20220730225603-2ab79fcdd4ef // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/tevino/abool v1.2.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/valyala/fastrand v1.1.0 // indirect
	github.com/xuri/efp v0.0.0-20230802181842-ad255f2331ca // indirect
	github.com/xuri/nfp v0.0.0-20230819163627-dc951e3ffe1a // indirect
	github.com/yuin/goldmark v1.5.5 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	go.etcd.io/etcd/api/v3 v3.5.11 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.11 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/arch v0.5.0 // indirect
	golang.org/x/crypto v0.16.0 // indirect
	golang.org/x/mobile v0.0.0-20230531173138-3c911d8e3eda // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/oauth2 v0.15.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/term v0.15.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.16.0 // indirect
	google.golang.org/api v0.153.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231120223509-83a465c0220f // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	honnef.co/go/js/dom v0.0.0-20210725211120-f030747120f2 // indirect
)
