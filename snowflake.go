package snowflake

import (
	"fmt"
	"sync"
	"time"
)

// Constants definition
const (
	epoch          = int64(1609459200000)             // Custom epoch timestamp (2021-01-01 00:00:00)
	nodeBits       = uint(10)                         // Number of bits for machine ID
	sequenceBits   = uint(12)                         // Number of bits for sequence
	maxNodeID      = int64(-1 ^ (-1 << nodeBits))     // Maximum value for machine ID
	maxSequence    = int64(-1 ^ (-1 << sequenceBits)) // Maximum value for sequence
	nodeIDShift    = sequenceBits                     // Left shift for machine ID
	timestampShift = sequenceBits + nodeBits          // Left shift for timestamp
)

type Snowflake struct {
	mu        sync.Mutex
	timestamp int64
	nodeID    int64
	sequence  int64
}

func NewSnowflake(nodeID int64) (*Snowflake, error) {
	if nodeID < 0 || nodeID > maxNodeID {
		return nil, fmt.Errorf("nodeID must be between 0 and %d", maxNodeID)
	}
	return &Snowflake{
		timestamp: 0,
		nodeID:    nodeID,
		sequence:  0,
	}, nil
}

// NextID generates the next unique ID
func (s *Snowflake) NextID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixNano() / 1e6 // Current timestamp in milliseconds
	if s.timestamp == now {
		// If generating IDs within the same millisecond
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			// If sequence exceeds the max value, wait for the next millisecond
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		// Reset sequence when time changes
		s.sequence = 0
	}

	s.timestamp = now

	id := (now-epoch)<<timestampShift | (s.nodeID << nodeIDShift) | s.sequence
	return id
}
