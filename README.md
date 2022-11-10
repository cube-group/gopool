# gopool
go routines pool

## Demo
```go
func TestBad(t *testing.T) {
	GoPoolMaxNum = 2
	GoPoolDebug = true

	for i := 0; i < 100; i++ {
		Go(func() {
			time.Sleep(time.Millisecond*2)
		})
		time.Sleep(time.Millisecond)
	}
	time.Sleep(10*time.Second)
}
```