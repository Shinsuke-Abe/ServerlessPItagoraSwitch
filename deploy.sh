#! /bin/sh

rm -fr ./build

mkdir ./build

GOOS=linux go build -o ./build/pitagoraOn
cp ./template.yml ./build/template.yml

cd ./build

zip handler.zip ./pitagoraOn

aws cloudformation package \
    --template-file ./template.yml \
    --output-template-file serverless-output.yml \
    --s3-bucket my-pitagora-switch &&
aws cloudformation deploy \
    --template-file serverless-output.yml \
    --stack-name pitagora-switch \
    --capabilities CAPABILITY_IAM
