C_DIR= ./c
GO_DIR = ./srv

.PHONY: all

all: vpp_lib vppapi stats

vpp_lib:
	$(MAKE) -C $(C_DIR)
	sudo $(MAKE) -C $(C_DIR) install

vppapi: stats
	go build -o $(GO_DIR)/vppapi $(GO_DIR)/vppapi.go

stats:
	go build -o $(GO_DIR)/stats/stats $(GO_DIR)/stats/stats.go

testclient: vpp_lib vppapi
	go build -o testclient testclient.go

clean:
	-rm $(C_DIR)/*.la
	-rm $(C_DIR)/*.lo
	-rm $(C_DIR)/*.o
	-rm $(GO_DIR)/vppapi
	-rm testclient
