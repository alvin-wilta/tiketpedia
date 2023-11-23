.PHONY: proto

clean:
	rm proto/*.pb.go

proto:
	@echo " > Generating proto files"
	@protoc -I proto/ --go_out=. ./proto/*.proto

ls: 
	ls