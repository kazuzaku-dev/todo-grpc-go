package util

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// ULIDを生成する
func NewULID() string {
	t := time.Now().UTC()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
