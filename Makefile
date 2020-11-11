test:
	go test -race -cover -coverprofile=cover.out $$(go list ./... | grep -Ev "coronator$$|testutil|mocks")