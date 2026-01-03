# Project Context

## Purpose
Metrics is a SaaS platform that provides easy to use event tracking and feature flags 
so product managers and developers can easily understand user's behavior and A/B test
their product.

## Tech Stack
- Backend with Go, Fiber, sqlc, PGx5, Golang Migrate;
- Frontend with Bun, Vite, React, Typescript, Mantine UI, React Router;
- Database with Postgres;

## Project Conventions

### Code Style
- Simplicity over fancy patterns, straightforward code without unnecessary abstractions
- Backend with idiomatic go, prefer small interfaces with ER naming like: Reader, Writer, etc.
- Follow the Airbnb JavaScript Style Guide for TypeScript, use slug-case for file names;
- Do not add comments, prefer descriptive variable/function names following language conventions;
- Do not use emojis;

### Architecture Patterns
- Monorepo that builds to a single binary (GO backend project), serving FE static assets embedded
in the binary;
- Local Development runs only database from a docker image, vite proxies api calls to BE;
- Asserts are served from root, all api endpoints are served from /api;
- Makefile with common commands for local development;

### Testing Strategy
- Descriptive tests focused in single or fewer assertions;
- Testify for Backend, tests focused in the behavior, not on functions;
- Bun tests for Frontend - small tests only for complext functions, don't test components;

### Git Workflow
- Main is production ready always;
- Features in short lived branches, named with feature name in slug-case;
- Short commit messages, no descriptions;

## Domain Context
[Add domain-specific knowledge that AI assistants need to understand]
- We provide cheap Analytics and Feature Flags with A/B testing capabilities through our platform.
- We are cheaper than competitors because we care about performance and simplicity, making
our system infrastructure affordable and without unnecessary complexity.

## Important Constraints
- Features focused on good User Experience, focusing users that do not want to be a especialist 
in another tool.
