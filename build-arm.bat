@set GOARCH=arm
@set GOOS=linux
@set CGO_ENABLED=0
@set GOARM=7
@echo start build smartgw
@echo -------------------------------------------------------------------

@go build -o smartgw -ldflags "-w -s"

@upx.exe smartgw

@rmdir release\smartgw /S /Q

@xcopy config\*.* release\smartgw\config\*.* /y /s
@xcopy webroot\*.* release\smartgw\webroot\*.* /y /s
@copy smartgw release\smartgw\smartgw /y

@echo ok!