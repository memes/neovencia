build:
	protoc -I. --go_out=plugins=micro:. \
    proto/billing/billing.proto
	docker build -t billing-service .

run:
	docker run -d --net="host" \
		-p 50055 \
		-e MICRO_SERVER_ADDRESS=:50055 \
		-e MICRO_REGISTRY=mdns \
		billing-service
