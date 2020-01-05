Read/write series of the data
=============================


Writer

```go

w, _ := NewFileWriter("test.out", 1024)
defer w.Close()
for i := uint64(0); i < 100; i++ {
    w.Write(i)
}
```

Reader

```go
r, _ := NewFileReader("test.out", 1024)
defer r.Close()

j := uint64(0)
sum := uint64(0)
for err := r.Read(&j); err == nil; err = r.Read(&j) {
    sum += j
}
```
