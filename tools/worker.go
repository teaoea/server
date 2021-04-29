package tools

import (
	"Server/config"
	"sync"
	"time"
)

var (
	conf     config.Config
	c        = conf.Yaml()
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
		admin := c.Mail.Admin
		subject := "时钟回拨"
		body := []byte("<h2>服务器时间回调,请同步服务器时间</h2>")
		SendMail(admin, subject, body, "")
		future := time.Now().AddDate(1, 0, 0).UnixNano() / 1e6 // 获取一年后的时间戳
		return future
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
