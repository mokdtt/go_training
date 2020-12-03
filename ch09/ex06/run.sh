go build main.go

echo "GOMAXPROCS=1"
GOMAXPROCS=1 time ./main > out.png
echo "GOMAXPROCS=2"
GOMAXPROCS=2 time ./main > out.png
echo "GOMAXPROCS=3"
GOMAXPROCS=3 time ./main > out.png
echo "GOMAXPROCS=4"
GOMAXPROCS=4 time ./main > out.png
echo "GOMAXPROCS=5"
GOMAXPROCS=5 time ./main > out.png
echo "GOMAXPROCS=6"
GOMAXPROCS=6 time ./main > out.png
echo "GOMAXPROCS=7"
GOMAXPROCS=7 time ./main > out.png
