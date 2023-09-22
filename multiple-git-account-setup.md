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