package dbus

import "golang.org/x/sys/unix"

// Gandi variable defines whether we should use Gandi specific struct fields.
// When set to false, the Gandi specific fields will be empty
var Gandi = false

// BasicIO stores the statistics for NFS
// Each field is a counter
type BasicIO struct {
	Requested  uint64
	Transfered uint64
	Total      uint64
	Errors     uint64
	Latency    uint64
	QueueWait  uint64
}

// OperationStat stores statistics for 9P or NFS operations
type OperationStat struct {
	Total  uint64
	Errors uint64
}

// LayoutOperationStat stores statistics for pNFS operations
type LayoutOperationStat struct {
	Total  uint64
	Errors uint64
	Delays uint64
}

// StatsBaseAnswer is the base answer to stats requests, every
// statistics related answer begins with this
type StatsBaseAnswer struct {
	Status bool
	Error  string
	Time   unix.Timespec
}

// PNFSOperations is th response to Layouts stats call
// for NFSv4.1 depending of the status of the call, all
// fields may not be filled
type PNFSOperations struct {
	StatsBaseAnswer
	Getdevinfo   LayoutOperationStat
	LayoutGet    LayoutOperationStat
	LayoutCommit LayoutOperationStat
	LayoutReturn LayoutOperationStat
	LayoutRecall LayoutOperationStat
}

// BasicStats is the response to IO stats call,
// some of the fields may not be filled depending
// of the call type and status
type BasicStats struct {
	StatsBaseAnswer
	Read    BasicIO
	Write   BasicIO
	Open    OperationStat // Gandi specific
	Close   OperationStat // Gandi specific
	Getattr OperationStat // Gandi specific
	Lock    OperationStat // Gandi specific
}
