@echo off

@REM Check if go is installed, if not, print error message and exit
where go >nul 2>nul
if errorlevel 1 (
    echo go is not installed or not in PATH
    exit /b 1
)

@REM Check if go is at least version 1.18, if not, print error message and exit
for /f "tokens=3" %%i in ('go version') do set version=%%i
set version=%version:~2,4%
if "%version%" LSS "1.18" (
    echo go version is too old, must be at least 1.18
    exit /b 1
)

@REM Check if the build directory exists, if it does, delete it
if exist build rmdir /s /q build

@REM Build the binary. if it fails, print error message and exit. if it succeeds, print success message
echo Building binary...
go build -v -o build\ .
if errorlevel 1 (
    echo Build failed
    exit /b 1
) else (
    echo Build succeeded
)