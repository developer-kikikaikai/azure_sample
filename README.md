# azure_sample

Install azure cli

```
curl -sL https://aka.ms/InstallAzureCLIDeb | sudo bash
```

Coding

```
func init
```

Run local

```
func host start
```

Deploy

```
# login
az login
# deploy
func azure functionapp publish xxx --force
```
