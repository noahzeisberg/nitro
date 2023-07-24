@echo off

rem Install requirements.
title Installing requirements...
echo Installing requirements...
pip install -r requirements.txt

rem Copy data to user directory.
title Installing Nitro...
echo Installing Nitro...
copy %cd%\nitro.py %userprofile%

title Setup finished!
echo Setup finished!

rem cd into user directory.
cd /d %userprofile%
echo Preparing for first use...

rem Start the CLI.
nitro