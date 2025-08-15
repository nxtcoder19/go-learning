1. generate ssh key using: ssh-keygen -t ed25519 -C "your_email@example.com" OR ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
2. Add your public ssh key on github
3. make a config file inside .ssh
4. your config file should be like:
# Work GitHub account
Host work github.com
  HostName github.com
  PreferredAuthentications publickey
  IdentityFile ~/.ssh/id_rsa

# Second GitHub account
Host personal github.com
  HostName github.com
  PreferredAuthentications publickey
  IdentityFile ~/.ssh/id_ed25519
5. Now, clone repo using below commands:
 # Cloning a repository from the Work GitHub account
 git clone git@work:username/repository.git

 # Cloning a repository from the Personal GitHub account
 git clone git@personal:username/repository.git
 git clone git@personal:kloudlite/marketing_api.git
6.Now create repo using below commands:
 # git init
 # git add .
 # git commit -m "wip"
 # git remote add origin git@personal:nxtcoder19/deeplearning-ai.git(personal)
 # git remote add origin git@work:nxtcoder19/deeplearning-ai.git(work)
 # git push -u origin main
7. if you encounter below error while commiting and pushing:
 git commit -m "test http triggering with webhook and sensor with docker image"
Author identity unknown

*** Please tell me who you are.

Run

  git config --global user.email "you@example.com"
  git config --global user.name "Your Name"

to set your account's default identity.
Omit --global to set the identity only in this repository.

fatal: unable to auto-detect email address (got 'piyushkumar@pkumar_mac.(none)')
Then, solution will be:
git config user.email "piyush.acet@gmail.com"



# Setup

1) Create two SSH keys (one per GitHub account)
# For account abc
ssh-keygen -t ed25519 -C "abc-account" -f ~/.ssh/id_abc

# For account def
ssh-keygen -t ed25519 -C "def-account" -f ~/.ssh/id_def


(Press Enter for no passphrase, or set one if you prefer.)

2) Add the public keys to each GitHub account

Copy keys:

pbcopy < ~/.ssh/id_abc.pub   # macOS; on Linux use: xclip -sel clip < ~/.ssh/id_abc.pub
pbcopy < ~/.ssh/id_def.pub


In each GitHub account: Settings → SSH and GPG keys → New SSH key and paste the matching key.

3) Configure ~/.ssh/config

Create/edit this file:

# ~/.ssh/config

Host github-abc
  HostName github.com
  User git
  IdentityFile ~/.ssh/id_abc
  IdentitiesOnly yes

Host github-def
  HostName github.com
  User git
  IdentityFile ~/.ssh/id_def
  IdentitiesOnly yes

4) Test both connections
ssh -T github-abc   # should say you're authenticated as the abc account
ssh -T github-def   # should say you're authenticated as the def account

5) Point each repo to the right remote
Repo 1 → push to abc account
cd /path/to/repo1
# Use the repo SSH URL but replace 'github.com' with 'github-abc'
git remote set-url origin git@github-abc:abc-username/repo1.git
# (or add origin if it doesn't exist)

Repo 2 → push to def account
cd /path/to/repo2
git remote set-url origin git@github-def:def-username/repo2.git

6) Set identity per repo (recommended)

This keeps commit authorship correct:

# In repo1 (abc)
git config user.name  "ABC Name"
git config user.email "abc@email.com"

# In repo2 (def)
git config user.name  "DEF Name"
git config user.email "def@email.com"

7) Push as usual
git push -u origin main

If you prefer HTTPS instead of SSH (optional)

Use different tokens and embed host aliases so your credential manager can distinguish:

Create two classic tokens or fine-grained PATs (one per account).

Set remotes like:

git remote set-url origin https://github-abc/abc-username/repo1.git
git remote set-url origin https://github-def/def-username/repo2.git


Add entries to /etc/hosts (or use a proxy) mapping github-abc and github-def to github.com, then let your OS credential manager store different credentials per host alias. (SSH is simpler, so use SSH if possible.)

That’s it! Each repo will now push to the correct GitHub account from the same machine.
