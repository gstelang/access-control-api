# Role-Based Domain Name Management using OpenFGA:
* Verify your ability to perform actions (such as managing or transferring) a domain name based on your role (e.g., owner, delegate).

# Install 
```
brew install openfga/tap/fga
```

# Validate model
```
fga model validate --file stores/user/folder_service.fga     
```

# Test model
```
fga model test --tests stores/user/test.folder_service.fga.yaml
```

# Docker setup
```
docker pull openfga/openfga
docker run -p 8080:8080 -p 8081:8081 -p 3000:3000 openfga/openfga run 
```

# Run the service 
```
go run cmd/service/main.go
```
