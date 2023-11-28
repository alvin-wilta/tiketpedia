.PHONY: proto

OUTDIR = proto

proto: clean proto-generate

tidy:
	@go mod tidy -compat=1.17
	@echo " > Go package tidy completed"

clean:
	@rm -f $(OUTDIR)/*.pb.go || true

proto-generate:
	@echo " > Generating proto files"
	@protoc -I $(OUTDIR)/ --go_out=. --go-grpc_out=. ./$(OUTDIR)/*.proto
	@echo " > Done generating proto files"
