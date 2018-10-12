build:
	protoc -I. --go_out=plugins=micro:. \
    proto/claim/claim.proto
	docker build -t claim-service .

run:
	docker run -d --net="host" \
		-p 50053 \
		-e MICRO_SERVER_ADDRESS=:50053 \
		-e MICRO_REGISTRY=mdns \
		claim-service
