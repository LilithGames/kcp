package kcp

const (
	IKCP_RTO_NDL       = 30  // no delay min rto
	IKCP_RTO_MIN       = 100 // normal min rto
	IKCP_RTO_DEF       = 200
	IKCP_RTO_MAX       = 60000
	IKCP_CMD_PUSH      = 81 // cmd: push data
	IKCP_CMD_ACK       = 82 // cmd: ack
	IKCP_CMD_WASK      = 83 // cmd: window probe (ask)
	IKCP_CMD_WINS      = 84 // cmd: window size (tell)
	IKCP_ASK_SEND      = 1  // need to send IKCP_CMD_WASK
	IKCP_ASK_TELL      = 2  // need to send IKCP_CMD_WINS
	IKCP_WND_SND       = 32
	IKCP_WND_RCV       = 32
	IKCP_MTU_DEF       = 1400
	IKCP_ACK_FAST      = 3
	IKCP_INTERVAL      = 100
	IKCP_OVERHEAD      = 24
	IKCP_DEADLINK      = 10
	IKCP_THRESH_INIT   = 2
	IKCP_THRESH_MIN    = 2
	IKCP_PROBE_INIT    = 7000   // 7 secs to probe window size
	IKCP_PROBE_LIMIT   = 120000 // up to 120 secs to probe window
	IKCP_FASTACK_LIMIT = 5
)

const (
	IKCP_LOG_OUTPUT    = 1
	IKCP_LOG_INPUT     = 2
	IKCP_LOG_SEND      = 4
	IKCP_LOG_RECV      = 8
	IKCP_LOG_IN_DATA   = 16
	IKCP_LOG_IN_ACK    = 32
	IKCP_LOG_IN_PROBE  = 64
	IKCP_LOG_IN_WIN    = 128
	IKCP_LOG_OUT_DATA  = 256
	IKCP_LOG_OUT_ACK   = 512
	IKCP_LOG_OUT_PROBE = 1024
	IKCP_LOG_OUT_WINS  = 2048
)
