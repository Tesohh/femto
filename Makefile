debug:
	go build -gcflags="all=-N -l" ./cmd/femto
	./femto
	rm ./femto
