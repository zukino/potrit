<!--
Sync Impact Report:
Version change: 1.0.0 → 1.0.0 (initial constitution)
Modified principles: N/A (new constitution)
Added sections: Core Principles (SOLID, Clean Architecture, Testing, Performance), Development Standards, Quality Gates
Removed sections: N/A
Templates requiring updates:
  ✅ plan-template.md (already aligned with SOLID principles)
  ✅ spec-template.md (already aligned with testing requirements)
  ⚠ tasks-template.md (needs k6 performance testing additions)
  ✅ All command templates (generic agent guidance)
Follow-up TODOs: Add k6 performance testing template to tasks-template.md
-->

# Potrit Constitution

## Core Principles

### I. SOLID Principles (Foundation)
Every component MUST follow SOLID principles without exception:
- **Single Responsibility**: Each class/module has one reason to change
- **Open/Closed**: Open for extension, closed for modification
- **Liskov Substitution**: Derived types must be substitutable for base types
- **Interface Segregation**: Clients must not depend on interfaces they don't use
- **Dependency Inversion**: High-level modules must not depend on low-level modules; both depend on abstractions

*Rationale*: SOLID principles ensure maintainable, testable, and flexible code that can evolve without breaking existing functionality.

### II. Clean Architecture (Structural)
All code MUST follow Clean Architecture layering with explicit dependency rules:
- **Domain Layer**: Core business logic, no external dependencies
- **Application Layer**: Use cases orchestrate domain objects, depend only on domain
- **Infrastructure Layer**: External concerns (DB, APIs, frameworks), implement interfaces defined in application
- **Presentation Layer**: UI/CLI controllers, depend only on application layer
- **Dependencies MUST point inward**: Infrastructure → Application → Domain

*Rationale*: Clean architecture isolates business logic from technical concerns, enabling independent testing, evolution, and technology swaps.

### III. Test-Driven Development (Process)
EVERY use case MUST have comprehensive unit tests before implementation:
- **TDD Mandatory**: Write failing test → Implement minimal code → Refactor
- **100% Use Case Coverage**: Every use case in application layer must have unit tests
- **Red-Green-Refactor**: Strict cycle enforced, no implementation before failing tests
- **Test Independence**: Tests must not depend on external systems or test order

*Rationale*: TDD ensures correct behavior, provides living documentation, and guarantees that use cases work as specified.

### IV. Performance Testing with k6 (Quality)
All critical paths MUST have performance validation using k6:
- **Use Case Performance**: Every use case must have k6 test scripts
- **Baseline Metrics**: Establish performance baselines for all endpoints/operations
- **Regression Testing**: Performance tests run in CI to prevent regressions
- **Load Testing**: Test under realistic load conditions based on expected usage

*Rationale*: Performance testing ensures system reliability and user experience under load, preventing production issues.

### V. Dependency Management (Integration)
All dependencies MUST be explicitly managed and versioned:
- **Interface Contracts**: Use dependency injection with explicit interfaces
- **Version Pinning**: All dependencies must be locked to specific versions
- **Minimal Dependencies**: Only include dependencies that are absolutely necessary
- **Security Scanning**: All dependencies must pass security vulnerability scans

*Rationale*: Explicit dependency management prevents version conflicts, security vulnerabilities, and ensures reproducible builds.

## Development Standards

### Code Quality Requirements
- **Static Analysis**: All code must pass linting and formatting tools
- **Code Coverage**: Minimum 90% line coverage for all critical paths
- **Documentation**: All public interfaces must have comprehensive documentation
- **Error Handling**: All error conditions must be explicitly handled and logged

### Testing Strategy
- **Unit Tests**: Every use case must have comprehensive unit tests
- **Integration Tests**: All external integrations must have integration tests
- **Contract Tests**: All API contracts must have contract tests
- **Performance Tests**: All critical paths must have k6 performance tests

### Build and Deployment
- **Automated Builds**: All changes must trigger automated build and test processes
- **Zero-Downtime**: Deployments must not impact system availability
- **Rollback Capability**: All deployments must have immediate rollback capability
- **Environment Parity**: All environments must be identical except for configuration

## Quality Gates

### Pre-commit Requirements
- All code must pass static analysis and formatting checks
- All tests must pass with coverage requirements met
- All security scans must pass without critical vulnerabilities
- All performance tests must meet baseline requirements

### Merge Requirements
- Peer review approval required for all changes
- All automated checks must pass
- Documentation must be updated for all API changes
- Performance impact must be assessed and documented

### Release Requirements
- Full integration test suite must pass
- Performance regression tests must pass
- Security audit must pass
- Rollback plan must be documented and tested

## Governance

This constitution supersedes all other development practices and guidelines. Amendments require:

1. **Proposal**: Written proposal detailing changes and rationale
2. **Review**: Technical review by architecture team
3. **Approval**: Majority approval from development team
4. **Documentation**: Update to this constitution with version increment
5. **Communication**: Team notification and training on changes

All pull requests and code reviews must verify compliance with constitutional principles. Any deviation from these principles must be explicitly justified and approved by the architecture team.

**Version**: 1.0.0 | **Ratified**: 2025-10-05 | **Last Amended**: 2025-10-05