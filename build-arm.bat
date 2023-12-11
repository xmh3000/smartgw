@set GOARCH=arm
@set GOOS=linux
@set CGO_ENABLED=0
@set GOARM=7
@echo start build smartgw
@echo -------------------------------------------------------------------

@go build -o smartgw -ldflags "-w -s"

@upx.exe smartgw

@rmdir smartgw\smartgw /S /Q

@xcopy config\*.* smartgw\smartgw\config\*.* /y /s
@xcopy webroot\*.* smartgw\smartgw\webroot\*.* /y /s
@copy smartgw smartgw\smartgw\smartgw /y
@copy install.sh smartgw\smartgw\install.sh /y

@echo ok!