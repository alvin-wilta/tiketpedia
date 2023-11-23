.PHONY: proto

OUTDIR = model/

clean:
	rm $(OUTDIR)/*.pb.go

proto:
	@echo " > Generating proto files"
	@protoc -I $(OUTDIR)/ --go_out=. ./$(OUTDIR)/*.proto

ls: 
	ls