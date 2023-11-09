CONFIG_PATH="configs/config.yaml" 
# в случае чего копировать и вставить в терминал export=CONFIG_PATH="configs/config.yaml" 

compose:
		docker compose up -d

upmigrate:
		migrate -path internal/db/migration -database 'postgres://wb:password@localhost:5040/orderdb?sslmode=disable' up

downmigrate:
		migrate -path internal/db/migration -database 'postgres://wb:password@localhost:5040/orderdb?sslmode=disable' down	

# sub:
# 	go run cmd/main.go 

# pub:
# 	go run cmd/publisher/publisher.go 

stop:
	docker stop orders-service && docker stop nats-stream 	

clean:
	docker rm orders-service && docker rm nats-stream


.PHONY: compose upmigrate downmigrate sub pub stop clean