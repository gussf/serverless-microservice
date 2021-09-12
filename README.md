# serverless-microservice
Go service using AWS's API Gateway, Lambda

``` bash
# Test using AWS SAM

# Requirements: Docker, aws-sam-cli
make sam-run
# Access localhost:3000/cars/available afterwards
``` 

# To Do
* Access DynamoDB
* Unit test lambdas package
* Create a CarRequest and CarResponse, in order to abstract other models from http io layer