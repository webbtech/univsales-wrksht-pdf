# if the ENV environment variable is not set to either stage or prod, makefile will fail
# ENV is confirmed below in the check_env directive
# example:
# for stage run: ENV=stage make
# for production run: ENV=prod make


include .env

default: check_env build awspackage awsdeploy

check_env:
	@echo -n "Your environment is $(ENV)? [y/N] " && read ans && [ $${ans:-N} = y ]

clean:
	@rm -rf dist
	@mkdir -p dist

build: clean
	@for dir in `ls handler`; do \
		GOOS=linux go build -o dist/$$dir github.com/pulpfree/univsales-wrksht-pdf/handler/$$dir; \
	done
	@GOOS=linux go build -o dist/authorizer github.com/pulpfree/univsales-wrksht-pdf/authorizer;
	@cp ./config/defaults.yml dist/
	@echo "build successful"

run: build
	sam local start-api -n env.json

samval:
	sam validate

awspackage:
	@aws cloudformation package \
	--template-file ${FILE_TEMPLATE} \
	--output-template-file ${FILE_PACKAGE} \
	--s3-bucket $(AWS_BUCKET_NAME) \
	--s3-prefix $(AWS_BUCKET_PREFIX) \
	--profile $(AWS_PROFILE) \
	--region $(AWS_REGION)

awsdeploy:
	@aws cloudformation deploy \
	--template-file ${FILE_PACKAGE} \
	--region $(AWS_REGION) \
	--stack-name $(PROJECT_NAME) \
	--capabilities CAPABILITY_IAM \
	--profile $(AWS_PROFILE) \
	--force-upload \
	--parameter-overrides \
	 	ParamCertificateArn=$(CERTIFICATE_ARN) \
		ParamCustomDomainName=$(CUSTOM_DOMAIN_NAME) \
		ParamENV=$(ENV) \
		ParamHostedZoneId=$(HOSTED_ZONE_ID) \
	 	ParamKMSKeyID=$(KMS_KEY_ID) \
		ParamProjectName=$(PROJECT_NAME) \
		ParamStorageBucket=$(AWS_STORAGE_BUCKET) \
		ParamSecurityGroupIds=$(SECURITY_GROUP_IDS) \
		ParamSubnetIds=$(SUBNET_IDS)

describe:
	@aws cloudformation describe-stacks \
		--region $(AWS_REGION) \
		--stack-name $(PROJECT_NAME)

outputs:
	@ make describe \
		| jq -r '.Stacks[0].Outputs'