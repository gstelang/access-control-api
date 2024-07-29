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
