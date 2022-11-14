# (Progress) Stick for Go Lang
Simple progress bar without any dependencies.

## Supports
- Percentage and Finished/Total task
- Elapsed time and estimated finish time
- Custom sizing
- Custom character for the finished section of the bar

## Usage
```golang
bar := stick.Create(100)
for i := 0; i < 100; i++ {
	time.Sleep(50 * time.Millisecond)
	bar.Add(1)
}
```

## Customize

Custom character for the finished section of the bar

```golang
bar := stick.Create(100)
bar.BarStar = "#"
for i := 0; i < 100; i++ {
	time.Sleep(50 * time.Millisecond)
	bar.Add(1)
}
```

Custom sizing

```golang
bar := stick.Create(100)
bar.Length = 50
for i := 0; i < 100; i++ {
	time.Sleep(50 * time.Millisecond)
	bar.Add(1)
}
```
