build:
	protoc -I. --go_out=plugins=micro:. \
    proto/policy/policy.proto
	docker build -t policy-service .

run:
	docker run -d --net="host" \
		-p 50054 \
		-e MICRO_SERVER_ADDRESS=:50054 \
		-e MICRO_REGISTRY=mdns \
		policy-service
