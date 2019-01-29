BIN_DIR := /usr/local/bin
SYSTEMD_DIR := /etc/systemd/system

all:
	cd dev; go install
	go build cluster-df-node.go config.go data.go
	go build cluster-df-router.go config.go data.go
	go build cluster-df.go config.go data.go
	go build cluster-df-local.go config.go data.go
clean:
	cd dev; go clean
	go clean
install:
	install -v cluster-df-local cluster-df-node cluster-df-router cluster-df $(BIN_DIR)
	install docs/cluster-df-node.service $(SYSTEMD_DIR)
	systemctl enable cluster-df-node.service
	systemctl start cluster-df-node.service
uninstall:
	systemctl disable cluster-df-node.service
	systemctl stop cluster-df-node.service
	rm -f $(SYSTEMD_DIR)/cluster-df-node.service 
	rm -f $(BIN_DIR)/cluster-df-local
	rm -f $(BIN_DIR)/cluster-df-node
	rm -f $(BIN_DIR)/cluster-df-router
	rm -f $(BIN_DIR)/cluster-df


install-router:
	install docs/cluster-df-router.service $(SYSTEMD_DIR)
	systemctl enable cluster-df-router.service
	systemctl start cluster-df-router.service

uninstall-router:
	systemctl disable cluster-df-router.service
	systemctl stop cluster-df-router.service
	rm -f $(SYSTEMD_DIR)/cluster-df-router.service 
