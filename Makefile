GOOS=linux
GOARCH=amd64
BUILDPATH=$(CURDIR)
BINPATH=$(BUILDPATH)/bin
EXENAME=report


clean:
	@rm $(BINPATH)/${EXENAME}_server || true
	
protoc:
	@protoc -I proto/ -I${GOPATH}/src --go_out=plugins=grpc:grpc/ proto/model/output.proto
	@protoc -I proto/ -I${GOPATH}/src --go_out=plugins=grpc:grpc/ proto/service/service.proto
	@sed -i "s/\/proto\//\/grpc\//g" grpc/service/service.pb.go
	@echo "Go code generated from proto files..."
	
	
build: clean
	@if [ ! -d $(BINPATH) ] ; then mkdir -p $(BINPATH) ; fi
	@GOOS=$(GOOS) GOARCH=$(GOARCH) cd server && go build -o $(BINPATH)/$(EXENAME)_server . || (echo "build failed $$?"; exit 1)
	@echo 'Build suceeded... done'
	