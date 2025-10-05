# Tasks: Build a REST API for Social Media Application

**Input**: Design documents from `/specs/002-build-an-rest/`
**Prerequisites**: plan.md, research.md, data-model.md, contracts/openapi.yaml, quickstart.md

## Current Progress Summary
- **Total Tasks**: 63 tasks
- **Completed**: 28/63 (44.4%)
- **Current Phase**: Moving to Phase 3.5 (Application Layer)
- **Last Completed**: Phase 3.4 (Infrastructure Layer) ✅

### Phase Completion Status:
- ✅ Phase 3.1 (Setup): 4/4 tasks completed
- ✅ Phase 3.2 (Tests First): 7/7 tasks completed
- ✅ Phase 3.3 (Core Domain Layer): 7/7 tasks completed
- ✅ Phase 3.4 (Infrastructure Layer): 10/10 tasks completed
- 🔄 Phase 3.5 (Application Layer): Ready to begin (0/16 tasks)
- ⏳ Phase 3.6 (Presentation Layer): Pending (0/9 tasks)
- ⏳ Phase 3.7 (Integration & Polish): Pending (0/9 tasks)

### Test Coverage Achieved:
- ✅ 17 contract tests covering all API endpoints
- ✅ 1 integration test with 10 sub-tests (complete user journey)
- ✅ 2 k6 performance test scripts ready for load testing
- ✅ All tests properly fail with connection errors (correct TDD behavior)

### Domain Layer Implementation:
- ✅ 5 core entities with comprehensive validation (User, Post, Connection, Like, Comment)
- ✅ 5 repository interfaces with complete CRUD and business operations
- ✅ Domain error types with structured validation
- ✅ Business logic encapsulated in entities following SOLID principles
- ✅ UUID-based primary keys and proper relationship modeling

### Infrastructure Layer Implementation:
- ✅ PostgreSQL database connection with connection pooling and health checks
- ✅ 5 database migration files with embedded filesystem support
- ✅ 5 PostgreSQL repository implementations with comprehensive CRUD operations
- ✅ Configuration loader with environment variable support and validation
- ✅ JWT token service with secure signing and validation
- ✅ Password hashing service with bcrypt and strength validation
- ✅ Migration runner for automated database schema management

## Execution Flow (main)
```
1. Load plan.md from feature directory
   → Tech stack: Go, PostgreSQL, standard library + database driver
   → Structure: Clean Architecture with cmd/, internal/, pkg/, tests/
2. Load design documents:
   → data-model.md: 5 entities (User, Post, Connection, Like, Comment)
   → contracts/openapi.yaml: 15+ API endpoints with authentication
   → research.md: JWT authentication, immediate consistency, k6 testing
3. Generate tasks by category:
   → Setup: Go project init, dependencies, linting
   → Tests: contract tests, integration tests, performance tests
   → Core: domain entities, application services, infrastructure
   → Integration: database, middleware, routing
   → Polish: unit tests, performance, documentation
4. Apply task rules:
   → Different files = mark [P] for parallel
   → Same file = sequential (no [P])
   → Tests before implementation (TDD)
5. Number tasks sequentially (T001, T002...)
6. Generate dependency graph
7. Create parallel execution examples
8. Validate task completeness:
   → All contracts have tests? ✓
   → All entities have models? ✓
   → All endpoints implemented? ✓
9. Return: SUCCESS (tasks ready for execution)
```

## Format: `[ID] [P?] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- Include exact file paths in descriptions

## Phase 3.1: Setup
- [x] T001 Create Clean Architecture project structure (cmd/, internal/domain/, internal/application/, internal/infrastructure/, internal/presentation/, pkg/, tests/)
- [x] T002 Initialize Go module with PostgreSQL driver (lib/pq) and JWT library
- [x] T003 [P] Configure Go linting (golangci-lint) and formatting (gofmt)
- [x] T004 [P] Set up testing framework (Go testing + testify) and k6 for performance testing

## Phase 3.2: Tests First (TDD) ✅ COMPLETED
- [x] T005 [P] Create failing contract tests for authentication endpoints (/auth/register, /auth/login) - tests/contract/auth_test.go
- [x] T006 [P] Create failing contract tests for user endpoints (/users/me, /users/{id}, /users/search) - tests/contract/users_test.go
- [x] T007 [P] Create failing contract tests for post endpoints (/posts, /posts/{id}, /posts/{id}/like) - tests/contract/posts_test.go
- [x] T008 [P] Create failing contract tests for comment endpoints (/posts/{id}/comments) - tests/contract/comments_test.go
- [x] T009 [P] Create failing contract tests for connection endpoints (/connections, /connections/{id}) - tests/contract/connections_test.go
- [x] T010 Create failing integration test for complete user journey (registration → login → create post → add connection → like post) - tests/integration/user_journey_test.go
- [x] T011 [P] Create failing k6 performance tests for critical endpoints (registration, login, feed retrieval, post creation) - tests/performance/scripts/

## Phase 3.3: Core Domain Layer ✅ COMPLETED
- [x] T012 [P] Create User entity with validation (internal/domain/entities/user.go)
- [x] T013 [P] Create Post entity with validation (internal/domain/entities/post.go)
- [x] T014 [P] Create Connection entity with status transitions (internal/domain/entities/connection.go)
- [x] T015 [P] Create Like entity (internal/domain/entities/like.go)
- [x] T016 [P] Create Comment entity (internal/domain/entities/comment.go)
- [x] T017 [P] Create repository interfaces for all entities (internal/domain/repositories/)
- [x] T018 [P] Create domain error types and validation rules (pkg/errors/errors.go)

## Phase 3.4: Infrastructure Layer ✅ COMPLETED
- [x] T019 Create PostgreSQL database connection and configuration (internal/infrastructure/database/postgres.go)
- [x] T020 [P] Create database migration files for all tables (internal/infrastructure/database/migrations/)
- [x] T021 [P] Implement User repository with PostgreSQL (internal/infrastructure/repositories/user_repository.go)
- [x] T022 [P] Implement Post repository with PostgreSQL (internal/infrastructure/repositories/post_repository.go)
- [x] T023 [P] Implement Connection repository with PostgreSQL (internal/infrastructure/repositories/connection_repository.go)
- [x] T024 [P] Implement Like repository with PostgreSQL (internal/infrastructure/repositories/like_repository.go)
- [x] T025 [P] Implement Comment repository with PostgreSQL (internal/infrastructure/repositories/comment_repository.go)
- [x] T026 Create configuration loader for config.json (internal/infrastructure/config/config.go)
- [x] T027 Create JWT token service (internal/infrastructure/auth/jwt.go)
- [x] T028 Create password hashing service (internal/infrastructure/auth/password.go)

## Phase 3.5: Application Layer
- [ ] T029 Create authentication service (register/login) (internal/application/services/auth_service.go)
- [ ] T030 [P] Create user service (internal/application/services/user_service.go)
- [ ] T031 [P] Create post service (internal/application/services/post_service.go)
- [ ] T032 [P] Create connection service (internal/application/services/connection_service.go)
- [ ] T033 [P] Create like service (internal/application/services/like_service.go)
- [ ] T034 [P] Create comment service (internal/application/services/comment_service.go)
- [ ] T035 [P] Create use case for user registration (internal/application/usecases/register_user.go)
- [ ] T036 [P] Create use case for user login (internal/application/usecases/login_user.go)
- [ ] T037 [P] Create use case for creating posts (internal/application/usecases/create_post.go)
- [ ] T038 [P] Create use case for getting user feed (internal/application/usecases/get_feed.go)
- [ ] T039 [P] Create use case for adding likes (internal/application/usecases/add_like.go)
- [ ] T040 [P] Create use case for creating comments (internal/application/usecases/create_comment.go)
- [ ] T041 [P] Create use case for managing connections (internal/application/usecases/manage_connections.go)
- [ ] T042 [P] Create use case for deleting posts (internal/application/usecases/delete_post.go)
- [ ] T043 [P] Create use case for deleting comments (internal/application/usecases/delete_comment.go)
- [ ] T044 [P] Create use case for deleting user account (internal/application/usecases/delete_account.go)

## Phase 3.6: Presentation Layer
- [ ] T045 Create HTTP router with middleware setup (internal/presentation/router/router.go)
- [ ] T046 [P] Create JWT authentication middleware (internal/presentation/middleware/auth.go)
- [ ] T047 [P] Create CORS middleware (internal/presentation/middleware/cors.go)
- [ ] T048 [P] Create authentication handlers (internal/presentation/handlers/auth_handler.go)
- [ ] T049 [P] Create user handlers (internal/presentation/handlers/user_handler.go)
- [ ] T050 [P] Create post handlers (internal/presentation/handlers/post_handler.go)
- [ ] T051 [P] Create comment handlers (internal/presentation/handlers/comment_handler.go)
- [ ] T052 [P] Create connection handlers (internal/presentation/handlers/connection_handler.go)
- [ ] T053 Create main application entry point (cmd/server/main.go)

## Phase 3.7: Integration & Polish
- [ ] T054 [P] Create unit tests for all domain entities
- [ ] T055 [P] Create unit tests for all application services
- [ ] T056 [P] Create unit tests for all infrastructure repositories
- [ ] T057 [P] Create unit tests for all presentation handlers
- [ ] T058 [P] Create integration tests for database operations
- [ ] T059 [P] Finalize k6 performance tests with realistic load scenarios
- [ ] T060 Create API documentation endpoint (OpenAPI spec serving)
- [ ] T061 Create health check endpoints
- [ ] T062 Update README.md with setup and usage instructions
- [ ] T063 Create deployment documentation (Docker, environment setup)

## Dependency Graph
```
Phase 3.1 (Setup) → Phase 3.2 (Tests) → Phase 3.3 (Domain) → Phase 3.4 (Infrastructure) → Phase 3.5 (Application) → Phase 3.6 (Presentation) → Phase 3.7 (Integration)
```

## Parallel Execution Examples

### Example 1: Entity Creation (T012-T016)
```bash
# These can run in parallel - different files, no dependencies
Task "Create User entity with validation (internal/domain/entities/user.go)"
Task "Create Post entity with validation (internal/domain/entities/post.go)"
Task "Create Connection entity with status transitions (internal/domain/entities/connection.go)"
Task "Create Like entity (internal/domain/entities/like.go)"
Task "Create Comment entity with validation (internal/domain/entities/comment.go)"
```

### Example 2: Repository Implementation (T021-T025)
```bash
# These can run in parallel - different files, depend only on entities
Task "Implement User repository with PostgreSQL (internal/infrastructure/repositories/user_repository.go)"
Task "Implement Post repository with PostgreSQL (internal/infrastructure/repositories/post_repository.go)"
Task "Implement Connection repository with PostgreSQL (internal/infrastructure/repositories/connection_repository.go)"
Task "Implement Like repository with PostgreSQL (internal/infrastructure/repositories/like_repository.go)"
Task "Implement Comment repository with PostgreSQL (internal/infrastructure/repositories/comment_repository.go)"
```

### Example 3: Handler Creation (T045-T049)
```bash
# These can run in parallel - different files, depend only on services
Task "Create authentication handlers (internal/presentation/handlers/auth_handler.go)"
Task "Create user handlers (internal/presentation/handlers/user_handler.go)"
Task "Create post handlers (internal/presentation/handlers/post_handler.go)"
Task "Create comment handlers (internal/presentation/handlers/comment_handler.go)"
Task "Create connection handlers (internal/presentation/handlers/connection_handler.go)"
```

## Task Execution Notes

### Critical Path Dependencies
1. **Setup must complete first** - All other phases depend on project structure and dependencies
2. **Tests must be written before implementation** - TDD approach required by constitution
3. **Domain layer before infrastructure** - Repositories depend on entity interfaces
4. **Infrastructure before application** - Services depend on repository implementations
5. **Application before presentation** - Handlers depend on service implementations

### Parallel Execution Rules
- **[P] tasks**: Different files, no shared dependencies, can run simultaneously
- **Sequential tasks**: Same file or shared dependencies, must run in order
- **Maximum parallelization**: Up to 5 tasks simultaneously for entity/repository/handler creation

### Validation Criteria
- ✅ All contract tests created and failing correctly (17 tests)
- ✅ All integration tests created and failing correctly (1 test, 10 sub-tests)
- ✅ k6 performance tests created and ready (2 scripts)
- ⏳ All contract tests must pass against implemented API
- ⏳ All integration tests must complete successfully
- ⏳ k6 performance tests must meet MVP targets (<100 concurrent users)
- ⏳ Code coverage must be >90% for business logic
- ⏳ All constitutional requirements must be satisfied

### Estimated Completion
- **Total Tasks**: 63 tasks
- **Estimated Time**: 2-3 weeks with parallel execution
- **Critical Path**: ~1 week (sequential dependencies only)
- **Parallel optimization**: 30-40% time savings with [P] tasks