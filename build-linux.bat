@set GOARCH=amd64
@set GOOS=linux
@echo start build smartgw
@echo -------------------------------------------------------------------

@go build -o smartgw -ldflags "-w -s"

@upx.exe smartgw

@copy smartgw release\smartgw /y
@xcopy config\*.* release\config\*.* /y /s
@xcopy webroot\*.* release\webroot\*.* /y /s
