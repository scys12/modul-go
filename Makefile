.PHONY : create_cover_test,func_cover_test,html_cover_test,create_cover_project,func_cover_project,html_cover_project

create_cover_test:
	go test ./... -coverprofile=cover.out

func_cover_test:
	go tool cover -func cover.out

html_cover_test:
	go tool cover -html cover.out

create_cover_project:
	go test -v -coverpkg=./... -coverprofile=profile.out ./...

func_cover_project:
	go tool cover -func profile.out

html_cover_project:
	go tool cover -html profile.out