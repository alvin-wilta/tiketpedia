.PHONY: proto

OUTDIR = proto

proto: clean proto-generate

clean:
	@rm -f $(OUTDIR)/*.pb.go || true

proto-generate:
	@echo " > Generating proto files"
	@protoc -I $(OUTDIR)/ --go_out=. ./$(OUTDIR)/*.proto
	@echo " > Done generating proto files"
