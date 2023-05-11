package ids

import (
	uuid "github.com/rs/xid"
	"hash/crc32"
)

func New() int64 {
	return int64(crc32.ChecksumIEEE([]byte(getUUID())))
}

func getUUID() string {
	return uuid.New().String()
}
