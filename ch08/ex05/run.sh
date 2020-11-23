go build orig.go
go build main.go

echo "並列でない"
time ./orig > out.png
echo "並列(4つまでgoroutine)"
time ./main > out.png

#rm -f orig
#rm -f main
