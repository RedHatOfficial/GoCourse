for i in `seq 1 5`
do
    nohup go run subscriber.go > subscriber${i}.log &
done

nohup go run subscriberX.go > subscriberX.log &
