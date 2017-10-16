# Copyright (C) 2017 David Capello

all:
	go build

package:
	env GOOS=darwin GOARCH=386 go build -v -o bin/findinfiles *.go
	cd bin && zip findinfiles-macosx.zip findinfiles && rm findinfiles && cd ..
	env GOOS=windows GOARCH=386 go build -v -o bin/findinfiles.exe *.go
	cd bin && zip findinfiles-windows.zip findinfiles.exe && rm findinfiles.exe && cd ..
	env GOOS=linux GOARCH=386 go build -v -o bin/findinfiles *.go
	cd bin && zip findinfiles-linux.zip findinfiles && rm findinfiles && cd ..
