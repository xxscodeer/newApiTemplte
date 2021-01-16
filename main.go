package main

import (
	"XxxxMicroAPI/route"
	"XxxxMicroAPI/tools"
	"fmt"
	"github.com/kataras/iris/v12"
	"net"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"

	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"

	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/common/log"
)

var cfg *tools.Config

func init() {
	fmt.Println("init config ing...")
	cfg = tools.ParseConfig("./conf/config.yaml")
	fmt.Println("init config end ..")
}

func main() {
	addr := cfg.AppConfig.Host + ":" + cfg.AppConfig.Port
	etcdURL := cfg.EtcdConfig.Host + ":" + cfg.EtcdConfig.Port
	jaegerURL := cfg.JaegerConfig.Host + ":" + cfg.JaegerConfig.Port

	//链路追踪
	t, i, e := tools.NewTracer(cfg.JaegerConfig.Name, jaegerURL)
	if e != nil {
		logger.Fatal("jaeger init fail,", e)
	}
	defer i.Close()

	opentracing.SetGlobalTracer(t)

	//	//熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	//启动端口
	go func() {
		err := http.ListenAndServe(net.JoinHostPort("0.0.0.0", " "), hystrixStreamHandler)
		if err != nil {
			log.Error(err)
		}
	}()
	server:=micro.NewService(
		micro.Name(cfg.AppConfig.Name),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(registry.Addrs(etcdURL))),
		//添加链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		//添加熔断
		micro.WrapClient(tools.NewClientHystrixWrapper()),
		micro.WrapClient(roundrobin.NewClientWrapper()),
		)
	app := route.InitRouter(server)
	app.Logger().SetLevel(cfg.AppConfig.Mode)
	_ = app.Run(iris.Addr(addr))
}
