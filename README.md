
# before you run go run ./cmd/main.go make sure to do the following :

#### tip for new kafka4.1.1 
docker-compose pull (to pull the image)
then generate the uuid with: 
$ docker run --rm apache/kafka:4.1.1 /opt/kafka/bin/kafka-storage.sh random-uuid
you will get a uuid like this : JYBMCjxoT9yrH_xj2_WHFw save it and put it in your docker-compose.yml

then run this 

docker run --rm \
  -v kafka-data:/var/lib/kafka/data \
  -e KAFKA_CLUSTER_ID="JYBMCjxoT9yrH_xj2_WHFw" \
  apache/kafka:4.1.1 \
  /opt/kafka/bin/kafka-storage.sh format --cluster-id JYBMCjxoT9yrH_xj2_WHFw --ignore-formatted --config /opt/kafka/config/server.properties --standalone

in this new version of kafka you are also expected to create events before hand ( no longer created automatically)


docker exec -it kafka-jar-service /opt/kafka/bin/kafka-topics.sh --create --topic jar-events --partitions 1 --replication-factor 1 --bootstrap-server localhost:9092
