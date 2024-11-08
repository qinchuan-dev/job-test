buf generate
swagger-combine ./config.json -o ./swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true
rm ./job-test.swagger.json
