# go build -o ti main.go && go test ./test/... -v -count=1 && rm ti
go build -pgo ./cpu.pprof -o ti main.go && go test ./test/... -v -count=1 && rm ti
