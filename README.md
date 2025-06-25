## My tool to backup my repositories on Github
It is hardcoded to only backup repositories under my username "kivattt", don't use this yourself

This takes up ~631M of disk space as of 2025-06-25.

```
Size history:

~550 MB of disk space as of 2025-03-03
631 M of disk space as of 2025-06-25.
```

Steps:
```
sudo apt install gh
gh auth login
gh api /users/kivattt/repos -F per_page=6969 --method GET > repos.json
go build
./github-backup-user
```
