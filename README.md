# Features:
* Send HTTP requests (GET, POST, PUT, DELETE) with role-based access.


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