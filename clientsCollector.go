package main

import (
	"github.com/Gandi/ganesha_exporter/dbus"
	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	clientsNfsV3RequestedDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v3_requested_bytes_total",
		"Number of requested bytes for NFSv3 operations",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV3TransferedDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v3_transfered_bytes_total",
		"Number of transfered bytes for NFSv3 operations",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV3OperationsDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v3_operations_total",
		"Number of operations for NFSv3",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV3ErrorsDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v3_operations_errors_total",
		"Number of operations in error for NFSv3",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV3LatencyDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v3_operations_latency_seconds_total",
		"Cumulative time consumed by operations for NFSv3",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV3QueueWaitDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v3_operations_queue_wait_seconds_total",
		"Cumulative time spent in rpc wait queue for NFSv3",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV40RequestedDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v40_requested_bytes_total",
		"Number of requested bytes for NFSv4.0 operations",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV40TransferedDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v40_transfered_bytes_total",
		"Number of transfered bytes for NFSv4.0 operations",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV40OperationsDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v40_operations_total",
		"Number of operations for NFSv4.0",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV40ErrorsDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v40_operations_errors_total",
		"Number of operations in error for NFSv4.0",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV40LatencyDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v40_operations_latency_seconds_total",
		"Cumulative time consumed by operations for NFSv4.0",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV40QueueWaitDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v40_operations_queue_wait_seconds_total",
		"Cumulative time spent in rpc wait queue for NFSv4.0",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV41RequestedDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v41_requested_bytes_total",
		"Number of requested bytes for NFSv4.1 operations",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV41TransferedDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v41_transfered_bytes_total",
		"Number of transfered bytes for NFSv4.1 operations",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV41OperationsDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v41_operations_total",
		"Number of operations for NFSv4.1",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV41ErrorsDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v41_operations_errors_total",
		"Number of operations in error for NFSv4.1",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV41LatencyDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v41_operations_latency_seconds_total",
		"Cumulative time consumed by operations for NFSv4.1",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV41QueueWaitDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v41_operations_queue_wait_seconds_total",
		"Cumulative time spent in rpc wait queue for NFSv4.1",
		[]string{"direction", "clientip"}, nil,
	)
	clientsPnfsLayoutOperationsDesc = prometheus.NewDesc(
		"ganesha_clients_pnfs_v41_layout_operations_total",
		"Numer of layout operations for pNFSv4.1",
		[]string{"type", "clientip"}, nil,
	)
	clientsPnfsLayoutErrorsDesc = prometheus.NewDesc(
		"ganesha_clients_pnfs_v41_layout_operations_errors_total",
		"Numer of layout operations in error for pNFSv4.1",
		[]string{"type", "clientip"}, nil,
	)
	clientsPnfsLayoutDelayDesc = prometheus.NewDesc(
		"ganesha_clients_pnfs_v41_layout_delay_seconds_total",
		"Cumulative delay time for pNFSv4.1",
		[]string{"direction", "clientip"}, nil,
	)
)

// ClientsCollector Collector for ganesha clients
type ClientsCollector struct {
	clientMgr                      dbus.ClientMgr
	nfsv3, nfsv40, nfsv41, pnfsv41 *bool
}

// NewClientsCollector creates a new collector
func NewClientsCollector() ClientsCollector {
	return ClientsCollector{
		clientMgr: dbus.NewClientMgr(),
		nfsv3:     kingpin.Flag("collector.clients.nfsv3", "Activate NFSv3 stats").Default("true").Bool(),
		nfsv40:    kingpin.Flag("collector.clients.nfsv40", "Activate NFSv4.0 stats").Default("true").Bool(),
		nfsv41:    kingpin.Flag("collector.clients.nfsv41", "Activate NFSv4.1 stats").Default("true").Bool(),
		pnfsv41:   kingpin.Flag("collector.clients.pnfsv41", "Activate pNFSv4.1 stats").Default("true").Bool(),
	}
}

// Describe prometheus description
func (ic ClientsCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(ic, ch)
}

// Collect do the actual job
func (ic ClientsCollector) Collect(ch chan<- prometheus.Metric) {
	var ()
	_, clients := ic.clientMgr.ShowClients()
	for _, client := range clients {
		clientip := client.Client
		if *ic.nfsv3 {
			var stats dbus.BasicStats
			if client.NFSv3 {
				stats = ic.clientMgr.GetNFSv3IO(client.Client)
			}
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV3RequestedDesc,
				prometheus.CounterValue,
				float64(stats.Read.Requested),
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV3TransferedDesc,
				prometheus.CounterValue,
				float64(stats.Read.Transfered),
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV3OperationsDesc,
				prometheus.CounterValue,
				float64(stats.Read.Total),
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV3ErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Read.Errors),
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV3LatencyDesc,
				prometheus.CounterValue,
				float64(stats.Read.Latency)/1e9,
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV3QueueWaitDesc,
				prometheus.CounterValue,
				float64(stats.Read.QueueWait)/1e9,
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV3RequestedDesc,
				prometheus.CounterValue,
				float64(stats.Write.Requested),
				"write", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV3TransferedDesc,
				prometheus.CounterValue,
				float64(stats.Write.Transfered),
				"write", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV3OperationsDesc,
				prometheus.CounterValue,
				float64(stats.Write.Total),
				"write", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV3ErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Write.Errors),
				"write", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV3LatencyDesc,
				prometheus.CounterValue,
				float64(stats.Write.Latency)/1e9,
				"write", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV3QueueWaitDesc,
				prometheus.CounterValue,
				float64(stats.Write.QueueWait)/1e9,
				"write", clientip)
		}
		if *ic.nfsv40 {
			stats := dbus.BasicStats{}
			if client.NFSv40 {
				stats = ic.clientMgr.GetNFSv40IO(client.Client)
			}
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV40RequestedDesc,
				prometheus.CounterValue,
				float64(stats.Read.Requested),
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV40TransferedDesc,
				prometheus.CounterValue,
				float64(stats.Read.Transfered),
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV40OperationsDesc,
				prometheus.CounterValue,
				float64(stats.Read.Total),
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV40ErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Read.Errors),
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV40LatencyDesc,
				prometheus.CounterValue,
				float64(stats.Read.Latency)/1e9,
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV40QueueWaitDesc,
				prometheus.CounterValue,
				float64(stats.Read.QueueWait)/1e9,
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV40RequestedDesc,
				prometheus.CounterValue,
				float64(stats.Write.Requested),
				"write", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV40TransferedDesc,
				prometheus.CounterValue,
				float64(stats.Write.Transfered),
				"write", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV40OperationsDesc,
				prometheus.CounterValue,
				float64(stats.Write.Total),
				"write", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV40ErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Write.Errors),
				"write", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV40LatencyDesc,
				prometheus.CounterValue,
				float64(stats.Write.Latency)/1e9,
				"write", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV40QueueWaitDesc,
				prometheus.CounterValue,
				float64(stats.Write.QueueWait)/1e9,
				"write", clientip)
		}
		if *ic.nfsv41 {
			stats := dbus.BasicStats{}
			if client.NFSv41 {
				stats = ic.clientMgr.GetNFSv41IO(client.Client)
			}
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV41RequestedDesc,
				prometheus.CounterValue,
				float64(stats.Read.Requested),
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV41TransferedDesc,
				prometheus.CounterValue,
				float64(stats.Read.Transfered),
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV41OperationsDesc,
				prometheus.CounterValue,
				float64(stats.Read.Total),
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV41ErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Read.Errors),
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV41LatencyDesc,
				prometheus.CounterValue,
				float64(stats.Read.Latency)/1e9,
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV41QueueWaitDesc,
				prometheus.CounterValue,
				float64(stats.Read.QueueWait)/1e9,
				"read", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV41RequestedDesc,
				prometheus.CounterValue,
				float64(stats.Write.Requested),
				"write", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV41TransferedDesc,
				prometheus.CounterValue,
				float64(stats.Write.Transfered),
				"write", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV41OperationsDesc,
				prometheus.CounterValue,
				float64(stats.Write.Total),
				"write", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV41ErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Write.Errors),
				"write", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV41LatencyDesc,
				prometheus.CounterValue,
				float64(stats.Write.Latency)/1e9,
				"write", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsNfsV41QueueWaitDesc,
				prometheus.CounterValue,
				float64(stats.Write.QueueWait)/1e9,
				"write", clientip)
		}
		if *ic.pnfsv41 {
			stats := dbus.PNFSOperations{}
			if client.NFSv41 {
				stats = ic.clientMgr.GetNFSv41Layouts(client.Client)
			}
			ch <- prometheus.MustNewConstMetric(
				clientsPnfsLayoutOperationsDesc,
				prometheus.CounterValue,
				float64(stats.Getdevinfo.Total),
				"getdevinfo", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsPnfsLayoutErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Getdevinfo.Errors),
				"getdevinfo", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsPnfsLayoutDelayDesc,
				prometheus.CounterValue,
				float64(stats.Getdevinfo.Delays)/1e9,
				"getdevinfo", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsPnfsLayoutOperationsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutGet.Total),
				"get", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsPnfsLayoutErrorsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutGet.Errors),
				"get", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsPnfsLayoutDelayDesc,
				prometheus.CounterValue,
				float64(stats.LayoutGet.Delays)/1e9,
				"get", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsPnfsLayoutOperationsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutCommit.Total),
				"commit", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsPnfsLayoutErrorsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutCommit.Errors),
				"commit", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsPnfsLayoutDelayDesc,
				prometheus.CounterValue,
				float64(stats.LayoutCommit.Delays)/1e9,
				"commit", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsPnfsLayoutOperationsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutReturn.Total),
				"return", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsPnfsLayoutErrorsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutReturn.Errors),
				"return", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsPnfsLayoutDelayDesc,
				prometheus.CounterValue,
				float64(stats.LayoutReturn.Delays)/1e9,
				"return", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsPnfsLayoutOperationsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutRecall.Total),
				"recall", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsPnfsLayoutErrorsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutRecall.Errors),
				"recall", clientip)
			ch <- prometheus.MustNewConstMetric(
				clientsPnfsLayoutDelayDesc,
				prometheus.CounterValue,
				float64(stats.LayoutRecall.Delays)/1e9,
				"recall", clientip)
		}
	}
}
