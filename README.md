# Nitro
Nitro is a fast, request-based and asynchronous package manager based on GitHub repositories.

# Important Note
This README will soon be replaced by another.

![ghpages](https://cdn.jsdelivr.net/npm/@intergrav/devins-badges@3/assets/cozy/documentation/ghpages_vector.svg)
![discord-singular](https://cdn.jsdelivr.net/npm/@intergrav/devins-badges@3/assets/cozy/social/discord-singular_vector.svg)

## How does it work?
Nitro uses the [GitHub API](https://api.github.com) to fetch information about the contents of the repository. Then it launches a thread for every file, with the task, to fetch its content and write it to a file. You don't need to authenticate with the GitHub API because Nitro only needs to send one request per repository, and you get 60 free requests per hour.

## How do I install it?
[Download the repository contents](https://github.com/NoahOnFyre/Nitro/archive/refs/heads/master.zip) and extract the archive. Make sure you have [Golang](https://go.dev/dl) installed. Then run the following command for the OS you're using.

### Windows
```
setup
```

### Linux (not tested)
```
wineconsole setup.bat
```

## How do I use it?
Info: You can use Nitro from the command prompt, or as standalone CLI.

If you want to install the package "NoahOnFyre/FyUTILS" you can go ahead and run:
```
get noahonfyre/fyutils
```
If you want to install a repo created by me, you can just enter the repository name instead of the whole name. For example:
```
get fyutils
```
If you're using the application via your command prompt, you'll have to add the `nitro` prefix to every command.
```
nitro get fyutils
```

# Updating installation
You can easily update your installation by downloading the newest version and running the `setup` command.

# Supported platforms

| Platform | Tested             |
|----------|--------------------|
| Windows  | :white_check_mark: |
| Linux    | :x:                |
| macOS    | :x:                |
| Android  | :x:                |