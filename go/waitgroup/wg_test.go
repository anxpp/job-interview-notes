package waitgroup

import (
	"sync"
	"testing"
)

func TestWG(t *testing.T) {
	var wg sync.WaitGroup
	wg.Wait()
}
