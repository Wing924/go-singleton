package singleton_test

import (
	"github.com/Wing924/go-singleton"
	"sync"
	"testing"
)

type MySingleton struct {
	value int
}

func TestSingleton_GetOrInit(t *testing.T) {
	var instance singleton.Singleton[*MySingleton]
	getInstance := func(i int) *MySingleton {
		return instance.GetOrInit(func() *MySingleton {
			return &MySingleton{value: i}
		})
	}

	for i := 100; i < 105; i++ {
		s := getInstance(i)
		if s.value != 100 {
			t.Errorf("expected %d, got %d", 100, s.value)
		}
	}
}

func TestSingleton_GetOrInit_thread_safe(t *testing.T) {
	var instance singleton.Singleton[*MySingleton]
	getInstance := func(i int) *MySingleton {
		return instance.GetOrInit(func() *MySingleton {
			return &MySingleton{value: i}
		})
	}

	var wg sync.WaitGroup
	result := make(chan int, 100)

	wg.Add(100)
	for i := range 100 {
		go func() {
			defer wg.Done()
			s := getInstance(i)
			result <- s.value
		}()
	}
	wg.Wait()
	close(result)
	first := <-result
	for v := range result {
		if v != first {
			t.Errorf("expected %d, got %d", first, v)
		}
	}
}

func BenchmarkSingleton_GetOrInit(b *testing.B) {
	var instance singleton.Singleton[*MySingleton]
	getInstance := func() *MySingleton {
		return instance.GetOrInit(func() *MySingleton {
			return &MySingleton{value: 1}
		})
	}

	b.ResetTimer()
	for range b.N {
		getInstance()
	}
}
