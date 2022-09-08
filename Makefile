# Copyright (C) 2017 David Capello

all:
	go build -o findinfiles *.go

package:
	env GOOS=darwin GOARCH=amd64 go build -v -o bin/findinfiles *.go
	cd bin && zip findinfiles-macosx.zip findinfiles && rm findinfiles && cd ..
	env GOOS=windows GOARCH=amd64 go build -v -o bin/findinfiles.exe *.go
	cd bin && zip findinfiles-windows.zip findinfiles.exe && rm findinfiles.exe && cd ..
	env GOOS=linux GOARCH=amd64 go build -v -o bin/findinfiles *.go
	cd bin && zip findinfiles-linux.zip findinfiles && rm findinfiles && cd ..
