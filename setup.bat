@echo off

title Installing requirements...
echo Installing requirements...
pip install -r requirements.txt

title Installing Scorpion...
echo Installing Scorpion...
copy %cd%\scorp.py %userprofile%

title Setup finished!
echo Setup finished!

C:
cd %userprofile%
scorp clear
scorp help