debug:
	rm femto.log
	go build -gcflags="all=-N -l" ./cmd/femto
	./femto
	rm ./femto
	cat femto.log
