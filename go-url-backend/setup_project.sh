#!/bin/bash

# Define the project root directory
ROOT_DIR="go-url-backend"

# Create domain structure
mkdir -p $ROOT_DIR/domain/todo
mkdir -p $ROOT_DIR/domain/user

# Create application structure
mkdir -p $ROOT_DIR/application/dto
mkdir -p $ROOT_DIR/application/mapper
mkdir -p $ROOT_DIR/application/services

# Create infrastructure structure
mkdir -p $ROOT_DIR/infrastructure/repository
mkdir -p $ROOT_DIR/infrastructure/database
mkdir -p $ROOT_DIR/infrastructure/middleware

# Create interfaces structure
mkdir -p $ROOT_DIR/interfaces/controllers
mkdir -p $ROOT_DIR/interfaces/routers

# Create pkg structure
mkdir -p $ROOT_DIR/pkg/config
mkdir -p $ROOT_DIR/pkg/utils

# Create base files
touch $ROOT_DIR/domain/todo/todo.go
# touch $ROOT_DIR/domain/todo/todo_service.go  c'est pas outilis
touch $ROOT_DIR/domain/todo/todo_repository.go 
touch $ROOT_DIR/domain/user/user.go
# touch $ROOT_DIR/domain/user/user_service.go  c'est pas outilis
touch $ROOT_DIR/domain/user/user_repository.go

touch $ROOT_DIR/application/dto/todo_dto.go
touch $ROOT_DIR/application/dto/user_dto.go
touch $ROOT_DIR/application/mapper/todo_mapper.go
touch $ROOT_DIR/application/mapper/user_mapper.go
touch $ROOT_DIR/application/services/todo_service.go
touch $ROOT_DIR/application/services/user_service.go

touch $ROOT_DIR/infrastructure/repository/todo_repository.go
touch $ROOT_DIR/infrastructure/repository/user_repository.go
touch $ROOT_DIR/infrastructure/database/database.go
touch $ROOT_DIR/infrastructure/database/schema.sql
touch $ROOT_DIR/infrastructure/middleware/auth_middleware.go
touch $ROOT_DIR/infrastructure/middleware/logging_middleware.go

touch $ROOT_DIR/interfaces/controllers/todo_controller.go
touch $ROOT_DIR/interfaces/controllers/user_controller.go
touch $ROOT_DIR/interfaces/routers/router.go

touch $ROOT_DIR/pkg/config/config.go
touch $ROOT_DIR/pkg/utils/logger.go
touch $ROOT_DIR/pkg/utils/validator.go

echo "Project structure and base files created successfully."
