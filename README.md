# Scorpion
Scorpion is a fast, request-based and asynchronous package manager based on GitHub repositories.

![ghpages](https://cdn.jsdelivr.net/npm/@intergrav/devins-badges@3/assets/cozy/documentation/ghpages_vector.svg)
![discord-singular](https://cdn.jsdelivr.net/npm/@intergrav/devins-badges@3/assets/cozy/social/discord-singular_vector.svg)

## How does it work?
Scorpion uses the [GitHub API](https://api.github.com) to fetch information about the contents of the repository. Then it launches a thread for every file, with the task, to fetch its content and write it to a file. You don't need to authenticate with the GitHub API because Scorpion only needs to send one request per repository, and you get 60 free requests per hour.

## How do I install it?
[Download the repository contents](https://github.com/NoahOnFyre/Scorpion/archive/refs/heads/master.zip) and extract the archive. Make sure you have [Python](https://python.org) installed. Open up a command prompt window in your current directory and run:
```
setup
```
This should install all the dependencies you need, will copy the main file to your user directory and execute some post-install steps, so you can instantly start using the tool.

## How do I use it?
Info: You can use Scorpion from the command prompt, or as standalone CLI.

If you want to install the package "Eltotiz/StreamHunter" you can go ahead and run:
```
scorp get eltotiz/streamhunter
```
If you want to install a repo created by me, you can just enter the repository name instead of the whole name. For example:
```
scorp get scorpion
```