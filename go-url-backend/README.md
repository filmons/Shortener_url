# go-ingnerd-backend

go-ingnerd-backend/
│
├── domain/                          # Core domain logic
│   ├── url/
│   │   ├── url.go                  # Domain model for "Url"
│   │   └── todo_repository.go       # Repository interface for "Url"
│   └── user/
│       ├── user.go                  # Domain model for "User"
│       └── user_repository.go       # Repository interface for "User"
│
├── application/                     # Application services layer
│   ├── dto/                         # Data Transfer Objects
│   │   ├── url_dto.go              # DTO for "Url"
│   │   └── user_dto.go              # DTO for "User"
│   ├── mapper/                      # Mappers for converting domain models to DTOs
│   │   ├── url_mapper.go
│   │   └── user_mapper.go
│   └── services/
│       ├── url_service.go          # Application service for "Url"
│       └── user_service.go          # Application service for "User"
│
├── infrastructure/                  # Infrastructure specific implementations
│   ├── repository/
│   │   ├── url_repository.go       # Implementation of Url repository
│   │   └── user_repository.go       # Implementation of User repository
│   ├── database/
│   │   ├── database.go              # Database setup and connection utilities
│   │   └── schema.sql               # SQL schema definitions
│   └── middleware/
│       ├── auth_middleware.go       # Authentication middleware
│       └── logging_middleware.go    # Logging middleware
│
├── interfaces/                      # Interfaces to the outside world
│   ├── controllers/
│   │   ├── url_controller.go       # Controller for "Url" operations
│   │   └── user_controller.go       # Controller for "User" operations
│   └── routers/
│       └── router.go                # Setup of HTTP routing
│
├── pkg/                             # Shared packages across the application
│   ├── config/
│   │   └── config.go                # Configuration management
│   └── utils/
│       ├── logger.go                # Logger utility
│       └── validator.go             # Common validation utilities
├── test/                                # Dedicated test directory
│   ├── unit/                            # Unit tests
│   │   ├── domain/
│   │   │   ├── url_test.go
│   │   │   └── user_test.go
│   │   ├── application/
│   │   │   ├── url_service_test.go
│   │   │   └── user_service_test.go
│   │   └── utils/
│   │       └── utils_test.go
│   ├── integration/                      # Integration tests
│   │   ├── api/
│   │   │   ├── url_controller_test.go
│   │   │   └── user_controller_test.go
│   │   └── repository/
│   │       ├── url_repository_impl_test.go
│   │       └── user_repository_impl_test.go
│   └── mocks/                            # Mocks for testing
│       ├── url_repository_mock.go
│       └── user_repository_mock.go
├── .env                             # Environment variables
├── .gitignore                       # Specifies intentionally untracked files to ignore
├── go.mod                           # Go module definitions
├── go.sum                           # Go module checksums
├── main.go                          # Entry point of the application
└── README.md                        # Project overview and setup instructions