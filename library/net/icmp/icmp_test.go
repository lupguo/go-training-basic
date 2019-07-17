package icmp

import "testing"

func BenchmarkSend(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		Send()
	}
}


