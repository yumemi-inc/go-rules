//go:build ruleguard

package gorules

import (
	"fmt"
	"testing"
)

func TestUseJoinHostPort(t *testing.T) {
	_ = fmt.Sprintf("%s:%d", "localhost", 3306)
}
