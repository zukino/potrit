
# Implementation Plan: Build a REST API for Social Media Application

**Branch**: `002-build-an-rest` | **Date**: 2025-10-05 | **Spec**: `/specs/002-build-an-rest/spec.md`
**Input**: Feature specification from `/specs/002-build-an-rest/spec.md`

## Execution Flow (/plan command scope)
```
1. Load feature spec from Input path
   → If not found: ERROR "No feature spec at {path}"
2. Fill Technical Context (scan for NEEDS CLARIFICATION)
   → Detect Project Type from file system structure or context (web=frontend+backend, mobile=app+api)
   → Set Structure Decision based on project type
3. Fill the Constitution Check section based on the content of the constitution document.
4. Evaluate Constitution Check section below
   → If violations exist: Document in Complexity Tracking
   → If no justification possible: ERROR "Simplify approach first"
   → Update Progress Tracking: Initial Constitution Check
5. Execute Phase 0 → research.md
   → If NEEDS CLARIFICATION remain: ERROR "Resolve unknowns"
6. Execute Phase 1 → contracts, data-model.md, quickstart.md, agent-specific template file (e.g., `CLAUDE.md` for Claude Code, `.github/copilot-instructions.md` for GitHub Copilot, `GEMINI.md` for Gemini CLI, `QWEN.md` for Qwen Code, or `AGENTS.md` for all other agents).
7. Re-evaluate Constitution Check section
   → If new violations: Refactor design, return to Phase 1
   → Update Progress Tracking: Post-Design Constitution Check
8. Plan Phase 2 → Describe task generation approach (DO NOT create tasks.md)
9. STOP - Ready for /tasks command
```

**IMPORTANT**: The /plan command STOPS at step 7. Phases 2-4 are executed by other commands:
- Phase 2: /tasks command creates tasks.md
- Phase 3-4: Implementation execution (manual or via tools)

## Summary
Build a REST API for social media application with core features including user registration/authentication, post creation/management, user connections, likes, and comments. The system will use Go with PostgreSQL, implement Clean Architecture principles, follow TDD methodology with specialized testing approach: curl for contract testing, k6 for performance testing, and test cases for use case testing. Dependencies are restricted to: github.com/golang-jwt/jwt/v5, github.com/google/uuid, github.com/lib/pq, golang.org/x/crypto.

## Technical Context
**Language/Version**: Go 1.21+
**Primary Dependencies**: github.com/golang-jwt/jwt/v5, github.com/google/uuid, github.com/lib/pq, golang.org/x/crypto
**Storage**: PostgreSQL
**Testing**: curl for contract testing, k6 for performance testing, test cases for use case testing
**Target Platform**: Linux server
**Project Type**: Single project with Clean Architecture
**Performance Goals**: <200ms p95 response time, 100 concurrent users, 50+ rps throughput
**Constraints**: No additional dependencies allowed beyond the four specified above
**Scale/Scope**: MVP for 100 concurrent users, immediate consistency

## Constitution Check
*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

### SOLID Principles Compliance
- ✅ **Single Responsibility**: Clean Architecture layers ensure each component has one responsibility
- ✅ **Open/Closed**: Interface-based design allows extension without modification
- ✅ **Liskov Substitution**: Repository interfaces ensure substitutability
- ✅ **Interface Segregation**: Specific interfaces for each repository type
- ✅ **Dependency Inversion**: High-level modules depend on abstractions

### Clean Architecture Compliance
- ✅ **Domain Layer**: Core entities with no external dependencies
- ✅ **Application Layer**: Use cases orchestrating domain objects
- ✅ **Infrastructure Layer**: Database and external integrations
- ✅ **Presentation Layer**: HTTP handlers and routing
- ✅ **Dependency Rule**: All dependencies point inward

### Test-Driven Development Compliance
- ✅ **TDD Mandatory**: Tests will be written before implementation
- ✅ **100% Use Case Coverage**: All use cases will have test cases
- ✅ **Red-Green-Refactor**: Strict cycle will be enforced
- ✅ **Test Independence**: Test cases will not depend on external systems

### Performance Testing Compliance
- ✅ **k6 Performance Testing**: All critical paths will have k6 scripts
- ✅ **Baseline Metrics**: Performance baselines will be established
- ✅ **Regression Testing**: Performance tests will run in CI
- ✅ **Load Testing**: Tests will cover expected concurrent user load

### Dependency Management Compliance
- ✅ **Interface Contracts**: Dependency injection with explicit interfaces
- ✅ **Version Pinning**: Dependencies will be locked to specific versions
- ✅ **Minimal Dependencies**: Only 4 allowed dependencies specified
- ✅ **Security Scanning**: Dependencies must pass security scans

### Additional Constitutional Requirements
- ✅ **Error Logging**: All errors will be logged to file and database every hour
- ✅ **Integration Tests**: External integrations will have integration tests
- ✅ **Contract Tests**: API contracts will have curl-based contract tests (per constitution)
- ✅ **Testing Methodology**: Using curl for contract testing, k6 for performance testing, test cases for use case testing

## Project Structure

### Documentation (this feature)
```
specs/002-build-an-rest/
├── plan.md              # This file (/plan command output)
├── research.md          # Phase 0 output (/plan command)
├── data-model.md        # Phase 1 output (/plan command)
├── quickstart.md        # Phase 1 output (/plan command)
├── contracts/           # Phase 1 output (/plan command)
└── tasks.md             # Phase 2 output (/tasks command - NOT created by /plan)
```

### Source Code (repository root)
```
cmd/
└── server/
    └── main.go          # Application entry point with CLI commands

internal/
├── domain/
│   ├── entities/        # Domain entities (User, Post, Connection, Like, Comment, ErrorLog)
│   └── repositories/    # Repository interfaces
├── application/
│   ├── services/        # Application services
│   ├── usecases/        # Use cases for business operations
│   └── testcases/       # Use case testing (test cases for each use case)
├── infrastructure/
│   ├── config/          # Configuration management
│   ├── database/
│   │   ├── migrations/  # Database migration files
│   │   └── seed/        # Seed data files
│   ├── repositories/    # Repository implementations
│   ├── auth/           # Authentication services
│   └── logging/        # Error and security logging
└── presentation/
    ├── handlers/        # HTTP handlers
    ├── middleware/      # HTTP middleware
    └── router/          # HTTP routing

pkg/
└── errors/              # Shared error types

logs/                     # Log files directory

tests/
├── contract/            # Contract testing using curl scripts
│   ├── test_contract.sh # Main contract test script with curl
│   └── endpoints/       # Individual endpoint contract tests
├── integration/         # Integration tests
├── unit/               # Unit tests
└── performance/         # k6 performance testing
    ├── baseline/        # Baseline performance tests
    ├── load/           # Load testing scripts
    ├── stress/         # Stress testing scripts
    └── regression/     # Performance regression tests
```

**Structure Decision**: Single project with Clean Architecture. The structure follows Go conventions with cmd/ for entry points, internal/ for private application code organized by layers, pkg/ for shared libraries, and specialized testing organization. The testing structure specifically includes curl-based contract testing, k6 performance testing, and use case testing. The logs/ directory supports the constitutional requirement for error logging.

## Phase 0: Outline & Research
1. **Extract unknowns from Technical Context** above:
   - For each NEEDS CLARIFICATION → research task
   - For each dependency → best practices task
   - For each integration → patterns task

2. **Generate and dispatch research agents**:
   ```
   For each unknown in Technical Context:
     Task: "Research {unknown} for {feature context}"
   For each technology choice:
     Task: "Find best practices for {tech} in {domain}"
   ```

3. **Consolidate findings** in `research.md` using format:
   - Decision: [what was chosen]
   - Rationale: [why chosen]
   - Alternatives considered: [what else evaluated]

**Output**: research.md with all NEEDS CLARIFICATION resolved

## Phase 1: Design & Contracts
*Prerequisites: research.md complete*

1. **Extract entities from feature spec** → `data-model.md`:
   - Entity name, fields, relationships
   - Validation rules from requirements
   - State transitions if applicable

2. **Generate API contracts** from functional requirements:
   - For each user action → endpoint
   - Use standard REST/GraphQL patterns
   - Output OpenAPI/GraphQL schema to `/contracts/`

3. **Generate contract tests** from contracts:
   - One test file per endpoint
   - Assert request/response schemas
   - Tests must fail (no implementation yet)

4. **Extract test scenarios** from user stories:
   - Each story → integration test scenario
   - Quickstart test = story validation steps

5. **Update agent file incrementally** (O(1) operation):
   - Run `.specify/scripts/bash/update-agent-context.sh claude`
     **IMPORTANT**: Execute it exactly as specified above. Do not add or remove any arguments.
   - If exists: Add only NEW tech from current plan
   - Preserve manual additions between markers
   - Update recent changes (keep last 3)
   - Keep under 150 lines for token efficiency
   - Output to repository root

**Output**: data-model.md, /contracts/*, failing tests, quickstart.md, agent-specific file

## Phase 2: Task Planning Approach
*This section describes what the /tasks command will do - DO NOT execute during /plan*

**Task Generation Strategy**:
- Load `.specify/templates/tasks-template.md` as base
- Generate tasks from Phase 1 design docs (contracts, data model, quickstart)
- Each contract → contract test task [P]
- Each entity → model creation task [P] 
- Each user story → integration test task
- Implementation tasks to make tests pass

**Ordering Strategy**:
- TDD order: Tests before implementation 
- Dependency order: Models before services before UI
- Mark [P] for parallel execution (independent files)

**Estimated Output**: 25-30 numbered, ordered tasks in tasks.md

**IMPORTANT**: This phase is executed by the /tasks command, NOT by /plan

## Phase 3+: Future Implementation
*These phases are beyond the scope of the /plan command*

**Phase 3**: Task execution (/tasks command creates tasks.md)  
**Phase 4**: Implementation (execute tasks.md following constitutional principles)  
**Phase 5**: Validation (run tests, execute quickstart.md, performance validation)

## Complexity Tracking
*Fill ONLY if Constitution Check has violations that must be justified*

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| [e.g., 4th project] | [current need] | [why 3 projects insufficient] |
| [e.g., Repository pattern] | [specific problem] | [why direct DB access insufficient] |


## Progress Tracking
*This checklist is updated during execution flow*

**Phase Status**:
- [x] Phase 0: Research complete (/plan command) - research.md exists
- [x] Phase 1: Design complete (/plan command) - data-model.md, quickstart.md, contracts/ exist
- [x] Phase 2: Task planning complete (/plan command - describe approach only)
- [ ] Phase 3: Tasks generated (/tasks command)
- [ ] Phase 4: Implementation complete
- [ ] Phase 5: Validation passed

**Gate Status**:
- [x] Initial Constitution Check: PASS - No constitutional violations detected
- [x] Post-Design Constitution Check: PASS - Clean Architecture with updated testing methodology
- [x] All NEEDS CLARIFICATION resolved - Feature spec fully specified
- [x] Complexity deviations documented - No complexity deviations required

**Testing Methodology Updates**:
- ✅ Contract testing: Updated to use curl exclusively (per constitution)
- ✅ Performance testing: Confirmed k6 for all performance testing
- ✅ Use case testing: Specified test cases for use case validation
- ✅ Testing structure: Updated project structure with dedicated testing directories

---
*Based on Constitution v1.0.0 - See `/memory/constitution.md`*
