# create ECR repository
aws ecr create-repository --repository-name go-ecommerce-product --region ap-southeast-1


# setup repository variable on github Actions secrets and variables

https://github.com/<username>>/<app-name>/settings/secrets/actions


AWS_ACCESS_KEY_ID = value ( ERC push IAM)
AWS_REGION = value
AWS_SECRET_ACCESS_KEY = value ( ERC push IAM)
REPO_NAME = <app-name>
REPO_TOKEN = <monorepo-gitops-repo-token>

# setup deploy keys
https://github.com/<username>>/<app-name>/settings/keys

# generate ssh-key

<path-tempat-aman> contoh
 mkdir -p /Users/kowlon/.ssh_repo_keys ( sesuaikan path nya)

 ```
$ ssh-keygen

Generating public/private ed25519 key pair.
Enter file in which to save the key (/Users/kowlon/.ssh/id_ed25519): /Users/kowlon/.ssh_repo_keys/user_service
Enter passphrase (empty for no passphrase): 
Enter same passphrase again: 
Your identification has been saved in /Users/kowlon/.ssh_repo_keys/user_service
Your public key has been saved in /Users/kowlon/.ssh_repo_keys/user_service.pub
```

# add deploy key

https://github.com/<username>>/<app-name>/settings/keys/new
paste public key
contoh copy -> /Users/kowlon/.ssh_repo_keys/user_service.pub



