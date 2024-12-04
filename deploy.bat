@echo off

setlocal

set dirPath=C:\Users\Administrator\gocode\myads-upgrader
set binPath=%dirPath%\upgrader.exe
set uploadUrl=http://mkt.nextotp.com/admin/modules/myads/api/upgrader/binary

cd %dirPath%

git pull origin master

echo.
echo Golang build.
go build -o upgrader.exe main.go

echo.
echo Upload binary.
curl -F "file=@%binPath%" %uploadUrl%
echo.

echo.
echo Completed, please quit.

timeout /t 10 /nobreak >nul

endlocal
