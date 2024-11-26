@echo off

setlocal

set dirPath=C:\Users\Administrator\gocode\myads-upgrader
set binPath=%dirPath%\upgrader.exe
set uploadUrl=http://mkt.nextotp.com/admin/modules/myads/upgrader/binary

cd %dirPath%

git pull origin master

go build -o upgrader.exe main.go

curl -F "file=@%binPath%" %uploadUrl%

endlocal
