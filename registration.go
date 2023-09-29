package main

func CommandRegistration() {
	RegisterCommand("get", "Get a package from github.com.", Arguments{"pkg"}, GetCommand)
	RegisterCommand("remove", "Remove a package from local storage.", Arguments{"pkg"}, RemoveCommand)
	RegisterCommand("help", "List all commands you need to operate Nitro.", Arguments{}, HelpCommand)
	RegisterCommand("list", "List all of your installed packages in your local storage.", Arguments{}, ListCommand)
	RegisterCommand("exit", "Exit the Nitro CLI.", Arguments{}, ExitCommand)
}
