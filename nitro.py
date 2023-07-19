import asyncio
import json
import os
import shutil
import subprocess
import sys
from pathlib import Path
import platform

import packaging.version as version_parser
import requests
import requests_futures.sessions
from colorama import Fore, Back, init

init(convert=True)


def fetch(dataset):
    if dataset["type"] != "file":
        return
    content = requests.get(dataset["download_url"]).content
    with open(target_path + "\\" + dataset["name"], mode="wb") as file:
        file.write(content)


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


def valid_args(args, count: int):
    if len(args) < count:
        print(prefix("ERROR") + "Invalid arguments!")
        return False
    else:
        return True


async def handle_command(cmd, args):
    global package
    global target_path
    try:
        match cmd:
            case "install" | "get":
                if valid_args(args, 1):
                    package = args[0].lower()
                    package = check_package(package)
                    user = package.split("/")[0]
                    repository = package.split("/")[1]
                    target_path = out_dir + "\\" + repository
                    os.makedirs(target_path)
                    content = requests.get("https://api.github.com/repos/" + package + "/contents").json()
                    print(prefix() + "Fetching content of " + Fore.GREEN + package + Fore.RESET + "...")
                    tasks = []
                    for entry in content:
                        tasks.append(asyncio.to_thread(fetch, entry))
                    await asyncio.gather(*tasks)
                    print(prefix() + "All artifacts collected!")
                    print(prefix() + "Verifying...")
                    artifacts = os.listdir(target_path)
                    includes_manifest = False
                    for artifact in artifacts:
                        artifact.removeprefix(target_path)
                        if artifact == "manifest.nitro":
                            print(prefix() + "Found project manifests!")
                            includes_manifest = True

                    if includes_manifest:
                        print(prefix() + "Using project manifests.")
                    else:
                        print(prefix() + "Creating manifest...")

                        with open(target_path + "\\manifest.nitro", "x") as file:
                            file.write(json.dumps({
                                "package": package,
                                "manifest_version": 1,
                                "main": "main.py"
                            }, indent=4))
                    print(prefix() + "Successfully collected package " + Fore.GREEN + package + Fore.RESET + "!")

            case "list":
                for file in os.listdir(out_dir):
                    print(prefix() + Fore.GREEN + file.removeprefix(out_dir))

            case "run":
                if valid_args(args, 1):
                    script = args[0]
                    if os.path.exists(out_dir + "\\" + script):
                        subprocess.call("python " + out_dir + "\\" + script + "\\main.py")
                    else:
                        os.system(script)

            case "update":
                current_version = version_parser.parse(version)
                target_version = version_parser.parse("")
                if current_version < target_version:
                    print(prefix() + "Downloading update...")

            case "uninstall" | "remove":
                if valid_args(args, 1):
                    package = args[0].lower()
                    package = check_package(package)
                    user = package.split("/")[0]
                    repository = package.split("/")[1]
                    target_path = out_dir + "\\" + repository
                    await asyncio.to_thread(lambda: shutil.rmtree(target_path))
                    print(prefix() + "Successfully removed package!")

            case "help":
                print_help_entry("nitro", "Start the standalone CLI.")
                print_help_entry("get", "Fetch a package and save it to local storage.")
                print_help_entry("list", "List all installed packages.")
                print_help_entry("run", "Run a package. (only python supported)")
                print_help_entry("update", "Update your Nitro instance to the newest version.")
                print_help_entry("remove", "Remove a package from local storage.")
                print_help_entry("help", "Get help about the commands.")
                print_help_entry("dir", "Open the location of your packages.")
                print_help_entry("clear", "Clear the terminal.")
                print_help_entry("restart", "Restart the application. (CLI only)")
                print_help_entry("exit", "Exit the application. (CLI only)")

            case "dir":
                os.system("start explorer.exe " + out_dir)

            case "clear" | "rl":
                menu()

            case "restart" | "rs":
                os.system("start " + path + "\\nitro.py")
                sys.exit(0)

            case "exit":
                os.system("cls")
                sys.exit(0)

            case _:
                print(prefix("ERROR") + "Command not found.")

    except Exception as e:
        print(prefix("ERROR") + str(e))


def check_package(pkg: str):
    if not pkg.__contains__("/"):
        return "noahonfyre/" + pkg
    else:
        return pkg


def menu():
    os.system("cls")
    os.system("title Nitro - Async Package Manager")
    print(Fore.GREEN + "     _   __ _  __")
    print(Fore.GREEN + "    / | / /(_)/ /_ _____ ____" + " "*4 + Fore.LIGHTBLACK_EX + "Python: " + Fore.GREEN + platform.python_version())
    print(Fore.GREEN + "   /  |/ // // __// ___// __ \\" + " "*3 + Fore.LIGHTBLACK_EX + "Packages: " + Fore.GREEN + str(len(os.listdir(out_dir))) + " package" + Fore.LIGHTBLACK_EX + "(s)" + Fore.GREEN + ".")
    print(Fore.GREEN + "  / /|  // // /_ / /   / /_/ /" + " "*3 + Fore.LIGHTBLACK_EX + "Version: " + Fore.GREEN + version.replace(".", Fore.LIGHTBLACK_EX + "." + Fore.GREEN))
    print(Fore.GREEN + " /_/ |_//_/ \\__//_/    \\____/" + " "*4 + Fore.LIGHTBLACK_EX + "Made by: " + Fore.GREEN + "NoahOnFyre")


async def main():
    if not os.path.exists(out_dir):
        os.makedirs(out_dir)

    if len(sys.argv) > 1:
        args = sys.argv
        cmd = args[1]
        args.remove(sys.argv[0])
        args.remove(cmd)
        print()
        await handle_command(cmd=cmd, args=args)
        sys.exit(0)

    menu()
    while True:
        try:
            print()
            args = input(Fore.LIGHTBLACK_EX + "\\\\" + Fore.GREEN + "nitro" + Fore.LIGHTBLACK_EX + " ~ " + Fore.RESET + "nitro ").split(" ")
            print()
            cmd = args[0]
            args.remove(cmd)
        except KeyboardInterrupt:
            sys.exit(0)
        await handle_command(cmd=cmd, args=args)


path = os.path.dirname(os.path.abspath(sys.argv[0]))
version = "1.0.0"
out_dir = str(Path.home()) + "\\.nitro"

asyncio.run(main())
