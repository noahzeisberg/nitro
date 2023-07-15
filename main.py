import asyncio
import os
import shutil
import sys
import threading
import time

import requests
from colorama import Fore, Back, init
from requests_futures.sessions import FuturesSession

init(convert=True)


def fetch(dataset):
    if dataset["type"] != "file":
        return
    content = requests.get(dataset["download_url"]).content
    with open(output_path + "\\" + dataset["name"], mode="wb") as file:
        file.write(content)
    print(prefix() + "Fetched " + Fore.GREEN + dataset["name"] + Fore.RESET + f"!" + Fore.LIGHTBLACK_EX + " (" + str(dataset["size"]) + " Bytes)")


async def start(content: str):
    for entry in content:
        await asyncio.to_thread(fetch, entry)


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


def valid_args(count: int):
    if len(args) < count:
        print(prefix("ERROR") + "Invalid arguments!")
        return False
    else:
        return True


def menu():
    print(Fore.GREEN + """   _____                       _           
  / ___/_________  _________  (_)___  ____ 
  \__ \/ ___/ __ \/ ___/ __ \/ / __ \/ __ \\
 ___/ / /__/ /_/ / /  / /_/ / / /_/ / / / /
/____/\___/\____/_/  / .___/_/\____/_/ /_/ 
                    /_/                                    """)
    print()
    print(Fore.GREEN + "SCORPION CLI " + Fore.LIGHTBLACK_EX + "-" + Fore.GREEN + " ASYNC GITHUB-BASED PACKAGE MANAGER")


menu()
path = os.path.dirname(os.path.abspath(sys.argv[0]))

while True:
    try:
        print()
        args = input(Fore.LIGHTBLACK_EX + "\\\\" + Fore.GREEN + "scorpion" + Fore.LIGHTBLACK_EX + " >>> " + Fore.RESET).split(" ")
        print()
        cmd = args[0]
        args.remove(cmd)
    except KeyboardInterrupt:
        sys.exit(0)

    match cmd:
        case "install":
            if valid_args(1):
                package = args[0]
                user = package.split("/")[0]
                repository = package.split("/")[1]
                output_path = path + "\\" + user + "_" + repository
                os.makedirs(output_path)
                content = requests.get("https://api.github.com/repos/" + package + "/contents").json()
                asyncio.run(start(content))

        case "uninstall" | "remove":
            if valid_args(1):
                package = args[0]
                user = package.split("/")[0]
                repository = package.split("/")[1]
                shutil.rmtree(path + "\\" + user + "_" + repository)

        case "exit":
            sys.exit(0)

        case "clear" | "rl":
            os.system("cls")
            menu()

        case "restart" | "rs":
            os.system("start " + path + "\\main.py")
            sys.exit(0)
        case _:
            print(prefix("ERROR") + "Unknown command.")
