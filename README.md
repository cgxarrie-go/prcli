# prq
Interaction with pull requests from command line
List and create PRs in Azure or Github repositories


## Installation

- Use the binaries the [latest release](https://github.com/cgxarrie-go/prq/releases/latest)

Allocate the binay in a folder accessible from command line
- prq.exe for Windows
- prq for Mac
    

## Usage

### Config commands
- prq config : display config

#### modifiers
- prq config az -pat : set PAT in Azure config
- prq config az -branch : set the default PRs target branch in Azure config
- prq config gh -pat : set PAT in Github config
- prq config gh -branch : set the default PRs target branch in Github config
- prq config remotes -a **remote** : Add a remote to config
- prq config remotes -r **remote** : Remove remote from config

### List PR commands 
List will list all active PRs in the remote of the current folder's local git 
repository (Azure ot Github)

- prq list : Lists all PR in status Active in the repository in the current directory

#### modifiers
-o : select repositories to get PRs from
    - d : Lists all PR in status Active in all the repositories found in the current directory tree
    - c : Lists all PR in status Active in all the repositories in config remotes

-f : filter PRs by title, author or status

### Create PR commands 
- prq create : creates a draft PR from current branch to default destination 
brnach with default title

default destination branch is **master** in Azure and **main** in Github
deafult title is **PR from spurce-branch-name to destination-branch-name**

#### modifiers
-g : specify target brnach of the PR
-t : set the title of the PR
-f : specify if PR is draft. Default is true
-d : PR decription
-m: template file to be added to PR description

- prq create -d **branchname** : creates a draft PR from current branch to **branchname** with default title
- prq create -t **pr-title** : crecreates a draft PR from current branch to default destination branch with title **pr-title**
- prq create  -d **branchname** -t **pr-title** : crecreates a draft PR from current branch to **branchname** with title **pr-title**
