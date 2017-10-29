go build recursive.go
go build parallel.go

DIR=~

echo "parellelized recursive (goroutines)"
time ./parallel $DIR | wc -l

echo "find command"
time find $DIR | wc -l

echo "naive recursive"
time ./recursive $DIR | wc -l
