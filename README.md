# gopool
go routines pool

## Demo
```go
func TestBad(t *testing.T) {
	GoPoolMaxNum = 2
	GoPoolDebug = true

	for i := 0; i < 10; i++ {
		Go(func() {
			time.Sleep(time.Second)
		})
		time.Sleep(time.Millisecond * 10)
	}
	time.Sleep(10*time.Second)
}
```