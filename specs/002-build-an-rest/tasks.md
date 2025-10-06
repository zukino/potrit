# Tasks: Build a REST API for Social Media Application

**Input**: Design documents from `/specs/002-build-an-rest/`
**Prerequisites**: plan.md, research.md, data-model.md, contracts/openapi.yaml, quickstart.md

## Execution Flow (main)
```
1. Load plan.md from feature directory
   → Tech stack: Go, PostgreSQL, standard library + database driver
   → Structure: Clean Architecture with cmd/, internal/, pkg/, tests/
2. Load design documents:
   → data-model.md: 6 entities (User, Post, Connection, Like, Comment, ErrorLog)
   → contracts/openapi.yaml: 17+ API endpoints with authentication
   → research.md: JWT authentication, immediate consistency, k6 testing
   → quickstart.md: Complete user journey walkthrough with migration/seed commands
3. Generate tasks by category:
   → Setup: Go project init, dependencies, linting
   → Tests: contract tests, integration tests, performance tests
   → Core: domain entities, application services, infrastructure
   → Integration: database, middleware, routing
   → Polish: unit tests, performance, docs
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
   → Migration and seed functionality included? ✓
9. Return: SUCCESS (tasks ready for execution)
```

## Format: `[ID] [P?] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- Include exact file paths in descriptions

## Phase 3.1: Setup
- [ ] T001 Create Clean Architecture project structure (cmd/, internal/domain/, internal/application/, internal/infrastructure/, internal/presentation/, pkg/, tests/)
- [ ] T002 Initialize Go module with restricted dependencies: github.com/golang-jwt/jwt/v5, github.com/google/uuid, github.com/lib/pq, golang.org/x/crypto
- [ ] T003 [P] Configure Go linting (golangci-lint) and formatting (gofmt)
- [ ] T004 [P] Set up testing framework: curl for contract testing, k6 for performance testing, test cases for use case testing

## Phase 3.2: Tests First (TDD) ⚠️ MUST COMPLETE BEFORE 3.3
**CRITICAL: These tests MUST be written and MUST FAIL before ANY implementation**
- [ ] T005 [P] Create failing curl-based contract tests for authentication endpoints (/auth/register, /auth/login, /auth/refresh) - tests/contract/test_contract.sh and tests/contract/endpoints/auth_test.sh - **NFR-003 Compliance**: Tests refresh token rotation mechanism
- [ ] T006 [P] Create failing curl-based contract tests for user endpoints (/users/me, /users/{id}, /users/search) - tests/contract/endpoints/users_test.sh
- [ ] T007 [P] Create failing curl-based contract tests for post endpoints (/posts, /posts/{id}, /posts/{id}/like) - tests/contract/endpoints/posts_test.sh
- [ ] T008 [P] Create failing curl-based contract tests for comment endpoints (/posts/{id}/comments) - tests/contract/endpoints/comments_test.sh
- [ ] T009 [P] Create failing curl-based contract tests for connection endpoints (/connections, /connections/{id}) - tests/contract/endpoints/connections_test.sh
- [ ] T009.1 [P] Create failing curl-based contract tests for delete endpoints (DELETE /posts/{id}, DELETE /posts/{id}/comments/{commentId}, DELETE /users/me) - tests/contract/endpoints/delete_test.sh
- [ ] T010 Create failing integration test for complete user journey (registration → login → create post → add connection → like post) - tests/integration/user_journey_test.go
- [ ] T010.1 [P] Create failing integration test for rate limiting behavior (authenticated and unauthenticated limits) - tests/integration/ratelimit_test.go - **NFR-004 Compliance**: Validates 100 req/min authenticated, 10 req/min unauthenticated limits
- [ ] T011 [P] Create failing k6 performance tests for critical endpoints (registration, login, feed retrieval, post creation) - tests/performance/scripts/
- [ ] T012 [P] Create failing integration test for migration commands and seed data functionality - tests/integration/migration_seed_test.go

## Phase 3.3: Core Domain Layer
- [ ] T013 [P] Create User entity with validation (internal/domain/entities/user.go)
- [ ] T014 [P] Create Post entity with validation (internal/domain/entities/post.go)
- [ ] T015 [P] Create Connection entity with status transitions (internal/domain/entities/connection.go)
- [ ] T016 [P] Create Like entity (internal/domain/entities/like.go)
- [ ] T017 [P] Create Comment entity (internal/domain/entities/comment.go)
- [ ] T017.1 [P] Create ErrorLog entity with validation (internal/domain/entities/error_log.go) - **Constitutional Compliance**: Supports NFR-008 error logging requirement
- [ ] T018 [P] Create repository interfaces for all entities (internal/domain/repositories/)
- [ ] T018.1 [P] Create ErrorLogRepository interface with error logging methods (internal/domain/repositories/error_log_repository.go) - **Constitutional Compliance**: Hourly error batch processing requirements
- [ ] T019 [P] Create migration repository interface (internal/domain/repositories/migration_repository.go)
- [ ] T020 [P] Create seed repository interface (internal/domain/repositories/seed_repository.go)
- [ ] T021 [P] Create domain error types and validation rules (pkg/errors/errors.go)

## Phase 3.4: Infrastructure Layer
- [ ] T022 Create PostgreSQL database connection and configuration using config.json settings with connection pooling (max 20 connections, min 5 connections), retry logic (3 attempts with exponential backoff), and health checks (internal/infrastructure/database/postgres.go)
- [ ] T023 [P] Create database migration files for all tables (internal/infrastructure/database/migrations/)
- [ ] T024 [P] Create migration service implementation (internal/infrastructure/database/migrator.go)
- [ ] T025 [P] Create seed data files and templates (internal/infrastructure/database/seed/)
- [ ] T026 [P] Create seed data service implementation (internal/infrastructure/database/seeder.go)
- [ ] T027 [P] Implement User repository with PostgreSQL (internal/infrastructure/repositories/user_repository.go)
- [ ] T028 [P] Implement Post repository with PostgreSQL (internal/infrastructure/repositories/post_repository.go)
- [ ] T029 [P] Implement Connection repository with PostgreSQL (internal/infrastructure/repositories/connection_repository.go)
- [ ] T030 [P] Implement Like repository with PostgreSQL (internal/infrastructure/repositories/like_repository.go)
- [ ] T031 [P] Implement Comment repository with PostgreSQL (internal/infrastructure/repositories/comment_repository.go)
- [ ] T031.1 [P] Implement ErrorLog repository with PostgreSQL (internal/infrastructure/repositories/error_log_repository.go) - **Constitutional Compliance**: Batch error logging every hour
- [ ] T032 Create configuration loader for config.json (internal/infrastructure/config/config.go)
- [ ] T033 Create JWT token service (internal/infrastructure/auth/jwt.go)
- [ ] T033.1 [P] Implement JWT refresh token rotation mechanism (internal/infrastructure/auth/refresh.go) - **NFR-003 Compliance**: 7-day refresh tokens with automatic rotation, httpOnly cookies
- [ ] T034 Create password hashing service (internal/infrastructure/auth/password.go)

## Phase 3.5: CLI Commands Enhancement, Tests First (TDD) ⚠️ MUST COMPLETE BEFORE 3.6
- [ ] T035 [P] Implement migrate command (up, down, status, create) in main application
- [ ] T036 [P] Implement seed command (run, reset, validate) in main application
- [ ] T037 Create CLI command interface documentation
- [ ] T038 Add command-line flag parsing and help system

## Phase 3.6: Application Layer
- [ ] T039 Create authentication service (register/login) (internal/application/services/auth_service.go)
- [ ] T040 [P] Create user service (internal/application/services/user_service.go)
- [ ] T041 [P] Create post service (internal/application/services/post_service.go)
- [ ] T042 [P] Create connection service (internal/application/services/connection_service.go)
- [ ] T043 [P] Create like service (internal/application/services/like_service.go)
- [ ] T044 [P] Create comment service (internal/application/services/comment_service.go)
- [ ] T045 [P] Create use case for user registration (internal/application/usecases/register_user.go)
- [ ] T046 [P] Create use case for user login (internal/application/usecases/login_user.go)
- [ ] T047 [P] Create use case for creating posts (internal/application/usecases/create_post.go)
- [ ] T048 [P] Create use case for getting user feed (internal/application/usecases/get_feed.go)
- [ ] T049 [P] Create use case for adding likes (internal/application/usecases/add_like.go)
- [ ] T050 [P] Create use case for creating comments (internal/application/usecases/create_comment.go)
- [ ] T051 [P] Create use case for managing connections (internal/application/usecases/manage_connections.go)
- [ ] T052 [P] Create use case for deleting posts (internal/application/usecases/delete_post.go)
- [ ] T053 [P] Create use case for deleting comments (internal/application/usecases/delete_comment.go)
- [ ] T054 [P] Create use case for deleting user account (internal/application/usecases/delete_account.go)

## Phase 3.6.1: Use Case Testing (TDD) ⚠️ MUST COMPLETE BEFORE 3.7
**CRITICAL: These test cases MUST be written and MUST FAIL before use case implementation**
- [ ] T054.1 [P] Create failing test case for user registration use case (internal/application/testcases/register_user_test.go)
- [ ] T054.2 [P] Create failing test case for user login use case (internal/application/testcases/login_user_test.go)
- [ ] T054.3 [P] Create failing test case for creating posts use case (internal/application/testcases/create_post_test.go)
- [ ] T054.4 [P] Create failing test case for getting user feed use case (internal/application/testcases/get_feed_test.go)
- [ ] T054.5 [P] Create failing test case for adding likes use case (internal/application/testcases/add_like_test.go)
- [ ] T054.6 [P] Create failing test case for creating comments use case (internal/application/testcases/create_comment_test.go)
- [ ] T054.7 [P] Create failing test case for managing connections use case (internal/application/testcases/manage_connections_test.go)
- [ ] T054.8 [P] Create failing test case for deleting posts use case (internal/application/testcases/delete_post_test.go)
- [ ] T054.9 [P] Create failing test case for deleting comments use case (internal/application/testcases/delete_comment_test.go)
- [ ] T054.10 [P] Create failing test case for deleting user account use case (internal/application/testcases/delete_account_test.go)

## Phase 3.7: Presentation Layer
- [ ] T065 Create HTTP router with middleware setup (internal/presentation/router/router.go)
- [ ] T066 [P] Create JWT authentication middleware (internal/presentation/middleware/auth.go)
- [ ] T067 [P] Create CORS middleware (internal/presentation/middleware/cors.go)
- [ ] T067.1 [P] Create rate limiting middleware (internal/presentation/middleware/ratelimit.go) - **NFR-004 Compliance**: 100 req/min authenticated, 10 req/min unauthenticated, burst allowances, rolling 60-second window
- [ ] T068 [P] Create authentication handlers with refresh token support (internal/presentation/handlers/auth_handler.go) - **NFR-003 Compliance**: Includes /auth/refresh endpoint with automatic token rotation
- [ ] T069 [P] Create user handlers (internal/presentation/handlers/user_handler.go)
- [ ] T070 [P] Create post handlers (internal/presentation/handlers/post_handler.go)
- [ ] T071 [P] Create comment handlers (internal/presentation/handlers/comment_handler.go)
- [ ] T072 [P] Create connection handlers (internal/presentation/handlers/connection_handler.go)
- [ ] T072.1 [P] Create delete handlers for posts, comments, and user accounts (internal/presentation/handlers/delete_handler.go) - **FR-013 Compliance**: Handles DELETE /posts/{id}, DELETE /posts/{id}/comments/{commentId}, DELETE /users/me
- [ ] T073 Create main application entry point with CLI command support (cmd/server/main.go)

## Phase 3.8: Integration & Polish
- [ ] T074 [P] Create unit tests for all domain entities
- [ ] T075 [P] Create unit tests for all application services
- [ ] T076 [P] Create unit tests for all infrastructure repositories
- [ ] T077 [P] Create unit tests for all presentation handlers
- [ ] T078 [P] Create unit tests for migration and seed services
- [ ] T079 [P] Create integration tests for database operations
- [ ] T080 [P] Create security logging service for authentication attempts and security events (NFR-008 compliance) - internal/infrastructure/logging/security_logger.go
- [ ] T081 [P] Create k6 baseline performance metrics for all critical endpoints (response time < 200ms p95, throughput > 50 rps) - tests/performance/baseline/ - **Constitutional Compliance**: Implements k6 performance testing requirement from Constitution Section IV
- [ ] T082 [P] Create k6 load testing scenarios for MVP scale (100 concurrent users, 10-minute sustained load, p95 < 500ms) - tests/performance/load/ - **Constitutional Compliance**: Implements k6 load testing requirement from Constitution Section IV
- [ ] T083 [P] Create k6 stress testing scenarios to identify breaking points (ramp to 200 concurrent users, monitoring p95 degradation) - tests/performance/stress/ - **Constitutional Compliance**: Extends k6 performance testing requirement from Constitution Section IV
- [ ] T084 [P] Create k6 performance regression tests with automated pass/fail criteria (CI integration) - tests/performance/regression/ - **Constitutional Compliance**: Implements k6 regression testing requirement from Constitution Section IV
- [ ] T085 [P] Create performance monitoring and reporting dashboard (baseline comparison, trend analysis) - **Constitutional Compliance**: Supports k6 performance testing requirements from Constitution Section IV
- [ ] T086 Create API documentation endpoint (OpenAPI spec serving)
- [ ] T087 Create health check endpoints
- [ ] T088 Update README.md with setup and usage instructions
- [ ] T089 Create deployment documentation (Docker, environment setup)
- [ ] T090 Create migration and seed data documentation

## Dependencies
- Setup (T001-T004) before Tests (T005-T012)
- Tests (T005-T012) before Domain (T013-T021) ⚠️ TDD REQUIREMENT
- Domain (T013-T021) before Infrastructure (T022-T034)
- Infrastructure (T022-T034) before CLI Commands (T035-T038)
- CLI Commands (T035-T038) before Application (T039-T054)
- Application (T039-T054) before Use Case Tests (T054.1-T054.10) ⚠️ TDD REQUIREMENT
- Use Case Tests (T054.1-T054.10) before Presentation (T065-T073)
- Presentation (T065-T073) before Integration & Polish (T074-T090)

## Dependency Graph
```
Phase 3.1 (Setup) → Phase 3.2 (Tests) → Phase 3.3 (Domain) → Phase 3.4 (Infrastructure) → Phase 3.5 (CLI Commands) → Phase 3.6 (Application) → Phase 3.6.1 (Use Case Tests) → Phase 3.7 (Presentation) → Phase 3.8 (Integration)
```

## Parallel Execution Examples

### Example 1: Entity Creation (T013-T017)
```bash
# These can run in parallel - different files, no dependencies
Task "Create User entity with validation (internal/domain/entities/user.go)"
Task "Create Post entity with validation (internal/domain/entities/post.go)"
Task "Create Connection entity with status transitions (internal/domain/entities/connection.go)"
Task "Create Like entity (internal/domain/entities/like.go)"
Task "Create Comment entity with validation (internal/domain/entities/comment.go)"
```

### Example 2: Repository Implementation (T027-T031)
```bash
# These can run in parallel - different files, depend only on entities
Task "Implement User repository with PostgreSQL (internal/infrastructure/repositories/user_repository.go)"
Task "Implement Post repository with PostgreSQL (internal/infrastructure/repositories/post_repository.go)"
Task "Implement Connection repository with PostgreSQL (internal/infrastructure/repositories/connection_repository.go)"
Task "Implement Like repository with PostgreSQL (internal/infrastructure/repositories/like_repository.go)"
Task "Implement Comment repository with PostgreSQL (internal/infrastructure/repositories/comment_repository.go)"
```

### Example 3: Handler Creation (T068-T072.1)
```bash
# These can run in parallel - different files, depend only on services
Task "Create authentication handlers (internal/presentation/handlers/auth_handler.go)"
Task "Create user handlers (internal/presentation/handlers/user_handler.go)"
Task "Create post handlers (internal/presentation/handlers/post_handler.go)"
Task "Create comment handlers (internal/presentation/handlers/comment_handler.go)"
Task "Create connection handlers (internal/presentation/handlers/connection_handler.go)"
```

### Example 4: CLI Command Implementation (T035-T038)
```bash
# These can run in parallel - different command modules
Task "Implement migrate command (up, down, status, create) in main application"
Task "Implement seed command (run, reset, validate) in main application"
Task "Create CLI command interface documentation"
Task "Add command-line flag parsing and help system"
```

## Task Execution Notes

### Critical Path Dependencies
1. **Setup must complete first** - All other phases depend on project structure and dependencies
2. **Tests must be written before implementation** - TDD approach required by constitution
3. **Domain layer before infrastructure** - Repositories depend on entity interfaces
4. **Infrastructure before CLI commands** - CLI commands depend on database infrastructure
5. **CLI commands before application** - Application services depend on CLI infrastructure
6. **Application before presentation** - Handlers depend on service implementations
7. **Presentation before polish** - Integration tests require full functionality

### Enhanced Migration & Seed Functionality
- **Migration Tasks**: T023, T024, T035 - Complete migration system with CLI interface
- **Seed Data Tasks**: T025, T026, T036 - Comprehensive seed data generation and management
- **CLI Integration**: T063, T035-T038 - Full command-line interface for database management
- **Test Coverage**: T012, T067 - Dedicated tests for migration and seed functionality

### Parallel Execution Rules
- **[P] tasks**: Different files, no shared dependencies, can run simultaneously
- **Sequential tasks**: Same file or shared dependencies, must run in order
- **Maximum parallelization**: Up to 5 tasks simultaneously for entity/repository/handler creation

### Validation Criteria
- ✅ All curl-based contract tests created and failing correctly (6 main contract scripts including delete and refresh endpoints)
- ✅ All integration tests created and failing correctly (3 tests: user journey, rate limiting, migration/seed, 15+ sub-tests)
- ✅ k6 performance tests created and ready (5 scripts for baseline, load, stress, regression)
- ✅ Use case tests created and failing correctly (10 test cases for all use cases)
- ✅ Migration and seed data tests included
- ✅ Delete functionality tests included (posts, comments, user accounts)
- ✅ Rate limiting tests included (authenticated and unauthenticated limits)
- ✅ JWT refresh token rotation tests included
- ✅ ErrorLog entity and repository tests included (constitutional error logging)
- ✅ Dependency constraints enforced (4 allowed dependencies only)
- ⏳ All curl contract tests must pass against implemented API
- ⏳ All integration tests must complete successfully
- ⏳ Use case tests must pass against implemented business logic
- ⏳ Migration and seed functionality must work correctly
- ⏳ k6 baseline tests must meet targets: response time < 200ms p95 (under 50 concurrent users, 5-minute duration), throughput > 50 rps
- ⏳ k6 load tests must sustain 100 concurrent users for 10 minutes (60-second ramp-up, maximum 1-second response time degradation)
- ⏳ k6 regression tests must pass automated CI criteria
- ⏳ Code coverage must be >90% for business logic
- ⏳ All constitutional requirements must be satisfied

### Estimated Completion
- **Total Tasks**: 97 tasks (increased from 94 due to ErrorLog entity implementation for constitutional compliance)
- **Estimated Time**: 2.5-3 weeks with parallel execution
- **Critical Path**: ~1.9 weeks (sequential dependencies only)
- **Parallel optimization**: 35-45% time savings with [P] tasks

## Notes
- [P] tasks = different files, no dependencies
- Verify tests fail before implementing (TDD requirement)
- Commit after each task
- Avoid: vague tasks, same file conflicts
- All tasks reference exact file paths for clear execution
- curl-based contract testing tasks explicitly support constitutional requirements
- k6 performance testing tasks explicitly support constitutional requirements
- Use case testing ensures comprehensive business logic validation
- Migration and seed data functionality fully integrated with CLI interface
- Enhanced task count reflects comprehensive use case testing implementation
- Dependency constraints enforced throughout all tasks