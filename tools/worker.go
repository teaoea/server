package tools

import (
	"sync"
	"time"

	"server/config"
)

var (
	conf     config.Config
	c        = conf.Conf()
	workerId = c.Worker.WorkerId
	centerId = c.Worker.CenterId
	sequence = c.Worker.Sequence
	epoch    = c.Worker.Epoch
)

const (
	sequenceBits   uint64 = 12
	sequenceMax    int64  = -1 ^ (-1 << sequenceBits)
	timestampShift uint8  = 22
	centerIdShift  uint8  = 17
	workerIdShift  uint8  = 12
)

type Worker struct {
	sync.Mutex
	lastTimestamp int64
	workerId      int64
	centerID      int64
	sequence      int64
}

func (*Worker) NewWorker() *Worker {
	return &Worker{
		workerId:      workerId,
		centerID:      centerId,
		sequence:      sequence,
		lastTimestamp: 0,
	}
}

func NewId() int64 {
	var w Worker
	w.Lock()
	timestamp := time.Now().UnixNano() / 1e6
	if timestamp < w.lastTimestamp {
		subject := "Inaccurate system time,level:0"
		content := `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
</head>
<body style="font-family: sans-serif">
<h1 style="text-align: center">Inaccurate system time,level:0</h1>
<h2 style="text-align: left"><strong>Inaccurate system time,please synchronize time</strong></h2>
</body>
</html>
`
		SendAdmin(subject, content)
		panic("Inaccurate system time")
	}
	if timestamp == w.lastTimestamp {
		w.sequence = (w.sequence + 1) & sequenceMax
		if w.sequence == 0 {
			for timestamp <= w.lastTimestamp {
				timestamp = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.sequence = 0
	}
	w.lastTimestamp = timestamp
	id := ((timestamp - epoch) << timestampShift) | (w.centerID << centerIdShift) | (w.workerId << workerIdShift) | w.sequence
	defer w.Unlock()
	return id
}
