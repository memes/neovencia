build:
	protoc -I. --go_out=plugins=micro:. \
    proto/settlement/settlement.proto
	docker build -t settlement-service .

run:
	docker run -d --net="host" \
		-p 50053 \
		-e MICRO_SERVER_ADDRESS=:50053 \
		-e MICRO_REGISTRY=mdns \
		settlement-service