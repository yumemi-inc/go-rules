//go:build ruleguard

package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

func UseJoinHostPort(m dsl.Matcher) {
	m.Import("fmt")
	m.
		Match("fmt.Sprintf(\"%s:%d\", $host, $port)").
		Report("use net.JoinHostPort instead of fmt.Sprintf").
		Suggest("net.JoinHostPort($host, $port)")
}
