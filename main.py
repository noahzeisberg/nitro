import os
import sys

from colorama import Fore, Back, init

init(convert=True)


def prefix(level: str = "INFO"):
    match level:
        case "INFO":
            level_color = Back.GREEN + Fore.BLACK
        case "WARN":
            level_color = Back.YELLOW + Fore.BLACK
        case "ERROR":
            level_color = Back.RED + Fore.BLACK
        case _:
            level_color = ""

    return level_color + " " + level + " " + Back.RESET + Fore.RESET + " "


def menu():
    print(Fore.GREEN + """   _____                       _           
  / ___/_________  _________  (_)___  ____ 
  \__ \/ ___/ __ \/ ___/ __ \/ / __ \/ __ \\
 ___/ / /__/ /_/ / /  / /_/ / / /_/ / / / /
/____/\___/\____/_/  / .___/_/\____/_/ /_/ 
                    /_/                                    """)
    print()
    print(Fore.GREEN + "SCORPION CLI " + Fore.LIGHTBLACK_EX + "-" + Fore.GREEN + " ASYNC GITHUB-BASED PACKAGE MANAGER")
    print()


menu()
path = os.path.dirname(os.path.abspath(sys.argv[0]))

while True:
    try:
        cmd = input(Fore.LIGHTBLACK_EX + "\\\\" + Fore.GREEN + "scorpion" + Fore.LIGHTBLACK_EX + " >>> " + Fore.RESET)
    except KeyboardInterrupt:
        sys.exit(0)
