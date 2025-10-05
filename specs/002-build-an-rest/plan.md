
# Implementation Plan: Build a REST API for Social Media Application

**Branch**: `002-build-an-rest` | **Date**: 2025-10-05 | **Spec**: [link](spec.md)
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
This plan implements a REST API for a social media application similar to Facebook, designed as an MVP backend service. The system will support user registration, authentication via email/password with JWT tokens, post creation with likes/comments, and user connections. Built with Go and PostgreSQL following Clean Architecture principles, it emphasizes minimal dependencies, immediate data consistency, and comprehensive testing including k6 performance validation.

## Technical Context
**Language/Version**: Go (latest stable)
**Primary Dependencies**: Standard library only + database driver
**Storage**: PostgreSQL (configured via config.json)
**Testing**: Go testing package + k6 for performance testing
**Target Platform**: Linux server
**Project Type**: Single project (REST API backend)
**Performance Goals**: Support < 100 concurrent users (MVP scale)
**Constraints**: Minimal external dependencies, immediate data consistency
**Scale/Scope**: Basic MVP social media API with users, posts, connections

## Constitution Check
*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

### SOLID Principles Compliance
✅ **Single Responsibility**: Go package structure encourages focused modules
✅ **Open/Closed**: Interface-based design enables extension without modification
✅ **Liskov Substitution**: Go interfaces ensure substitutability
✅ **Interface Segregation**: Small, focused interfaces in Go ecosystem
✅ **Dependency Inversion**: Dependency injection pattern supported

### Clean Architecture Compliance
✅ **Layer Separation**: Domain, Application, Infrastructure, Presentation layers
✅ **Dependency Rule**: Dependencies point inward (Infrastructure → Application → Domain)
✅ **Isolation**: Business logic isolated from technical concerns

### TDD Requirements
✅ **Unit Tests**: Go testing package for comprehensive use case coverage
✅ **Red-Green-Refactor**: Strict TDD cycle enforced
✅ **Test Independence**: Tests isolated from external systems

### Performance Testing
✅ **k6 Integration**: Performance testing for all critical paths
✅ **Baseline Metrics**: Performance baselines established
✅ **Load Testing**: Realistic load testing for MVP scale

### Dependency Management
✅ **Interface Contracts**: Explicit interface definitions
✅ **Minimal Dependencies**: Standard library preferred
✅ **Security**: PostgreSQL security with prepared statements

**Status**: PASS - No constitutional violations detected

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
├── server/
│   └── main.go                # Application entry point

internal/
├── domain/                    # Domain layer - core business logic
│   ├── entities/
│   │   ├── user.go
│   │   ├── post.go
│   │   ├── connection.go
│   │   ├── like.go
│   │   └── comment.go
│   └── repositories/
│       ├── user_repository.go
│       ├── post_repository.go
│       └── connection_repository.go
│
├── application/               # Application layer - use cases
│   ├── services/
│   │   ├── auth_service.go
│   │   ├── user_service.go
│   │   ├── post_service.go
│   │   └── connection_service.go
│   └── usecases/
│       ├── register_user.go
│       ├── create_post.go
│       ├── add_like.go
│       └── create_connection.go
│
├── infrastructure/            # Infrastructure layer - external concerns
│   ├── database/
│   │   ├── postgres.go
│   │   └── migrations/
│   ├── config/
│   │   └── config.go
│   └── auth/
│       ├── jwt.go
│       └── password.go
│
└── presentation/              # Presentation layer - HTTP handlers
    ├── handlers/
    │   ├── auth_handler.go
    │   ├── user_handler.go
    │   ├── post_handler.go
    │   └── connection_handler.go
    ├── middleware/
    │   ├── auth.go
    │   └── cors.go
    └── router/
        └── router.go

pkg/                          # Public packages
└── errors/
    └── errors.go

tests/
├── unit/                     # Unit tests for each layer
├── integration/              # Integration tests
├── contract/                 # API contract tests
└── performance/              # k6 performance tests

config.json                   # Application configuration
go.mod
go.sum
README.md
```

**Structure Decision**: Single project with Clean Architecture layering. Domain layer contains core business entities and repository interfaces. Application layer contains use cases and business logic. Infrastructure layer implements external concerns like database and JWT. Presentation layer handles HTTP requests and responses.

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
| JWT library (external dependency) | Secure, industry-standard token-based authentication with proven security track record. Standard library crypto would require implementing complex JWT algorithms (RS256 signing, token validation, expiration handling) increasing security risk and development time. | Manual crypto implementation risks security vulnerabilities, requires extensive testing, and reinvents well-established security patterns. |


## Progress Tracking
*This checklist is updated during execution flow*

**Phase Status**:
- [x] Phase 0: Research complete (/plan command)
- [x] Phase 1: Design complete (/plan command)
- [x] Phase 2: Task planning complete (/plan command - describe approach only)
- [ ] Phase 3: Tasks generated (/tasks command)
- [ ] Phase 4: Implementation complete
- [ ] Phase 5: Validation passed

**Gate Status**:
- [x] Initial Constitution Check: PASS
- [x] Post-Design Constitution Check: PASS
- [x] All NEEDS CLARIFICATION resolved
- [x] Complexity deviations documented

---
*Based on Constitution v1.0.0 - See `/memory/constitution.md`*
