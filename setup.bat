@echo off

title Installing requirements...
echo Installing requirements...
pip install -r requirements.txt

title Installing Nitro...
echo Installing Nitro...
copy %cd%\nitro.py %userprofile%

title Setup finished!
echo Setup finished!

cd /d %userprofile%
echo Preparing for first use...

nitro