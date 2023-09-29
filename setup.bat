@echo off
echo Installing Nitro...

REM Setting up Go in current directory and building sources will skip automatically,
REM if there are no Go module files and source files found.

echo Installing Go dependencies and building binaries...
go install
go build -v -o nitro.exe

REM Here are the installation steps, if there is the "nitro.exe" file in the same directory
REM as the setup file is. This will also run, if the Go build process completes,
REM because they're using the same output file.

echo Copying output files...
copy %cd%\nitro.exe %userprofile%

echo Nitro installation complete! Check out further instructions in the repository's README.md!
echo.
echo https://github.com/NoahOnFyre/Nitro
echo.
echo Also, please consider starring the repository on GitHub, because it would support me a lot. :)
pause

exit