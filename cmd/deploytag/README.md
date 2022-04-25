# Tag Maker
This script creates `tag string` for deployment according to the rules defined in `lab-development` team in order to deploy.

## Usage

```shell
go run cmd/deploytag/main.go --env=prod
# prod-20220424t150839
```
### Options
* `env`: Target environment to deploy.


## Format
`{env}-{yyyymmdd}t{hhmmss}`


# Reference
* [lab-development Tag Based Deployment](https://www.notion.so/mathpresso/Tag-Based-Deployment-e1391c05e65c4350b8126d6aa79093d2)
