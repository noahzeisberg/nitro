import asyncio
import os
import shutil
import sys
from pathlib import Path

import requests
from colorama import Fore, Back, init

init(convert=True)


def fetch(dataset):
    if dataset["type"] != "file":
        return
    content = requests.get(dataset["download_url"]).content
    with open(target_path + "\\" + dataset["name"], mode="wb") as file:
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


def check_package(pkg: str):
    if not pkg.__contains__("/"):
        global package
        package = "NoahOnFyre/" + pkg


def menu():
    os.system("title Scorpion - Async Package Manager")
    print(Fore.GREEN + """   _____                       _           
  / ___/_________  _________  (_)___  ____ 
  \__ \/ ___/ __ \/ ___/ __ \/ / __ \/ __ \\
 ___/ / /__/ /_/ / /  / /_/ / / /_/ / / / /
/____/\___/\____/_/  / .___/_/\____/_/ /_/ 
                    /_/                                    """)


menu()
path = os.path.dirname(os.path.abspath(sys.argv[0]))
out_dir = str(Path.home()) + "\\.scorpion\\"

if len(sys.argv) > 1:
    script = sys.argv[1]
    os.system("title Scorpion - Running: " + script)
    print()
    print(Fore.LIGHTBLACK_EX + "\\\\" + Fore.GREEN + "scorpion" + Fore.LIGHTBLACK_EX + " >>> " + Fore.RESET + "run " + script)
    print()
    if os.path.exists(out_dir + script):
        with open(out_dir + script + "\\main.py", "rb") as file:
            exec(file.read())
    else:
        os.system(script)
    sys.exit()

if not os.path.exists(out_dir):
    os.makedirs(out_dir)

while True:
    try:
        print()
        args = input(Fore.LIGHTBLACK_EX + "\\\\" + Fore.GREEN + "scorpion" + Fore.LIGHTBLACK_EX + " >>> " + Fore.RESET).split(" ")
        print()
        cmd = args[0]
        args.remove(cmd)
    except KeyboardInterrupt:
        sys.exit(0)

    try:
        match cmd:
            case "install" | "get" | "add":
                if valid_args(1):
                    package = args[0]
                    check_package(package)
                    user = package.split("/")[0]
                    repository = package.split("/")[1]
                    target_path = out_dir + "\\" + repository
                    os.makedirs(target_path)
                    content = requests.get("https://api.github.com/repos/" + package + "/contents").json()
                    asyncio.run(start(content))

            case "run":
                if valid_args(1):
                    script = args[0]
                    if os.path.exists(out_dir + script):
                        with open(out_dir + script + "\\main.py", "rb") as file:
                            exec(file.read())
                    else:
                        os.system(script)

            case "update":
                if valid_args(1):
                    package = args[0]
                    check_package(package)
                    user = package.split("/")[0]
                    repository = package.split("/")[1]
                    target_path = out_dir + "\\" + repository
                    shutil.rmtree(target_path)
                    os.makedirs(target_path)
                    content = requests.get("https://api.github.com/repos/" + package + "/contents").json()
                    asyncio.run(start(content))

            case "uninstall" | "remove":
                if valid_args(1):
                    package = args[0]
                    check_package(package)
                    user = package.split("/")[0]
                    repository = package.split("/")[1]
                    target_path = out_dir + "\\" + repository
                    shutil.rmtree(target_path)

            case "exit":
                sys.exit(0)

            case "clear" | "rl":
                os.system("cls")
                menu()

            case "restart" | "rs":
                os.system("start " + path + "\\scorp.py")
                sys.exit(0)
            case _:
                print(prefix("ERROR") + "Unknown command.")
    except Exception as e:
        print(prefix("ERROR") + str(e))
        continue
