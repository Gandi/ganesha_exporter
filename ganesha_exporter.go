package main

import (
	"net/http"

	"github.com/Gandi/ganesha_exporter/dbus"
	"github.com/alecthomas/kingpin/v2"
	"github.com/davecgh/go-spew/spew"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		listenAddress     = kingpin.Flag("web.listen-address", "Address on which to expose metrics and web interface.").Default(":9587").String()
		metricsPath       = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").String()
		gandi             = kingpin.Flag("gandi", "Activate Gandi specific fields").Default("false").Bool()
		exporterCollector = kingpin.Flag("collector.exports", "Activate exports collector").Default("true").Bool()
	)
	ec := NewExportsCollector()
	var clientCollector = kingpin.Flag("collector.clients", "Activate clients collector").Default("true").Bool()
	cc := NewClientsCollector()

	kingpin.Version(version.Print("ctld_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	dbus.Gandi = *gandi

	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		collectors.NewGoCollector(),
	)
	if *exporterCollector {
		reg.MustRegister(ec)
	}
	if *clientCollector {
		reg.MustRegister(cc)
	}
	http.Handle(*metricsPath, promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`<html>
			<head><title>ctld Exporter</title></head>
			<body>
			<h1>ctld Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
		if err != nil {
			log.Errorln(err)
		}
	})

	log.Infoln("Listening on", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
	mgr := dbus.NewExportMgr()
	time, exports := mgr.ShowExports()
	spew.Dump(time, exports)
}
