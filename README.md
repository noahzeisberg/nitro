# Scorpion
Scorpion is a fast, request-based and asynchronous package manager based on GitHub repositories.

![ghpages](https://cdn.jsdelivr.net/npm/@intergrav/devins-badges@3/assets/cozy/documentation/ghpages_vector.svg)
![discord-singular](https://cdn.jsdelivr.net/npm/@intergrav/devins-badges@3/assets/cozy/social/discord-singular_vector.svg)

## How does it work?
Scorpion uses the [GitHub API](https://api.github.com) to fetch information about the contents of the repository. Then it launches a thread for every file, with the task, to fetch its content and write it to a file. You don't need to authenticate with the GitHub API because Scorpion only needs to send one request per repository, and you get 60 free requests per hour.

## How do I install it?
[Download the repository contents](https://github.com/NoahOnFyre/Scorpion/archive/refs/heads/master.zip) and extract the archive. Make sure you have [Python](https://python.org) installed. Open up a command prompt window in your current directory and run:
```
pip install -r requirements.txt
```
This should install all the dependencies you need. Then, start Scorpion by double-clicking on the file.

## How do I use it?
If you want to install the package "Eltotiz/StreamHunter" you can go ahead and run:
```
get eltotiz/streamhunter
```
If you want to install a package globally, add the "-g" flag to it. For example:
```
get eltotiz/streamhunter -g
```
If you want to install a repo created by me, you can just enter the repository name instead of the whole name. For example:
```
get scorpion
```
This works with the "global" flag too:
```
get scorpion -g
```