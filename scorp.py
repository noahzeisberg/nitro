import asyncio
import os
import shutil
import subprocess
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


def print_help_entry(command_name: str, description: str):
    first = Fore.GREEN + command_name.upper() + Fore.LIGHTBLACK_EX + ":" + Fore.RESET + " "*(12-len(command_name))
    second = Fore.WHITE + description.replace("(", Fore.LIGHTBLACK_EX + "(").replace(")", ")" + Fore.RESET)
    print(prefix() + first + second)


def valid_args(count: int):
    if len(args) < count:
        print(prefix("ERROR") + "Invalid arguments!")
        return False
    else:
        return True


def init(cmd, args):
    global package
    global target_path
    try:
        match cmd:
            case "install" | "get":
                if valid_args(1):
                    package = args[0]
                    check_package(package)
                    user = package.split("/")[0]
                    repository = package.split("/")[1]
                    target_path = out_dir + "\\" + repository
                    os.makedirs(target_path)
                    content = requests.get("https://api.github.com/repos/" + package + "/contents").json()
                    asyncio.run(start(content))

            case "list":
                for file in os.listdir(out_dir):
                    print(prefix() + Fore.GREEN + file.removeprefix(out_dir))

            case "run":
                if valid_args(1):
                    script = args[0]
                    if os.path.exists(out_dir + "\\" + script):
                        subprocess.call("python " + out_dir + "\\" + script + "\\main.py")
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
                    print(prefix() + "Successfully removed package!")

            case "help":
                print_help_entry("scorp", "Start the standalone CLI.")
                print_help_entry("get", "Fetch a package and save it to local storage.")
                print_help_entry("list", "List all installed packages.")
                print_help_entry("run", "Run a package. (only python supported)")
                print_help_entry("update", "Re-fetch a package and save it to local storage.")
                print_help_entry("remove", "Remove a package from local storage.")
                print_help_entry("help", "Get help about the commands.")
                print_help_entry("dir", "Open the location of your packages.")
                print_help_entry("clear", "Clear the terminal.")
                print_help_entry("restart", "Restart the application. (CLI only)")
                print_help_entry("exit", "Exit the application. (CLI only)")

            case "dir":
                os.system("start explorer.exe " + out_dir)

            case "clear" | "rl":
                os.system("cls")
                menu()

            case "restart" | "rs":
                os.system("start " + path + "\\scorp.py")
                sys.exit(0)

            case "exit":
                sys.exit(0)

            case _:
                if os.path.exists(out_dir + "\\" + cmd):
                    subprocess.call("python " + out_dir + "\\" + cmd + "\\main.py")
                else:
                    print(prefix("ERROR") + "Command not found.")

    except Exception as e:
        print(prefix("ERROR") + str(e))


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


path = os.path.dirname(os.path.abspath(sys.argv[0]))
out_dir = str(Path.home()) + "\\.scorpion"

if len(sys.argv) > 1:
    args = sys.argv
    cmd = args[1]
    args.remove(sys.argv[0])
    args.remove(cmd)
    print()
    init(cmd=cmd, args=args)
    sys.exit(0)

if not os.path.exists(out_dir):
    os.makedirs(out_dir)

menu()
while True:
    try:
        print()
        args = input(Fore.LIGHTBLACK_EX + "\\\\" + Fore.GREEN + "scorpion" + Fore.LIGHTBLACK_EX + " >>> " + Fore.RESET).split(" ")
        print()
        cmd = args[0]
        args.remove(cmd)
    except KeyboardInterrupt:
        sys.exit(0)
    init(cmd=cmd, args=args)
