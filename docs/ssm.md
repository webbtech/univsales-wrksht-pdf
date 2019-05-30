# Create SSM Params

## Parameters For Testing

* key alias: alias/Universal
* path: /test/univsales-wrksht-pdf

``` bash
# Create params
aws ssm put-parameter --name /test/univsales-wrksht-pdf/DBName \
  --value universal-sales --type String --overwrite
```

## Parameters for production

* key alias alias/Universal
* path: /prod/univsales-wrksht-pdf

``` bash
# Create String params
aws ssm put-parameter --name /prod/univsales-wrksht-pdf/DBHost \
  --value ca-central-1a.mongo,ca-central-1b.mongo --type String --overwrite

aws ssm put-parameter --name /prod/univsales-wrksht-pdf/DBName \
  --value universal-sales --type String --overwrite

# Create SecureString params
aws ssm put-parameter --name /prod/univsales-wrksht-pdf/DBPassword \
  --value siteUser --type SecureString \
  --key-id alias/Universal --overwrite

aws ssm put-parameter --name /prod/univsales-wrksht-pdf/DBUser \
  --value siteUser --type SecureString \
  --key-id alias/Universal --overwrite

# Get parameters
aws ssm get-parameters-by-path --path /prod/univsales-wrksht-pdf
aws ssm get-parameters-by-path --path /prod/univsales-wrksht-pdf --with-decryption
```
