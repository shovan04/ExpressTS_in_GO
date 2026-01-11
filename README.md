# ExpressTS in Go 🚀

**Fast. Typed. Opinionated.**

ExpressTS in Go is a modern backend CLI framework written in **Go**, inspired by the simplicity of Express and the discipline of clean architectures.

It helps you bootstrap scalable backend projects with a strong architectural foundation — without unnecessary complexity.

---

## ✨ Why ExpressTS in Go?

Because backend developers want:

- ⚡ Speed and performance with Go
- 🧱 Clear and scalable architecture
- 📁 Clean and predictable project structure
- 😌 A CLI that feels friendly, not overwhelming

This tool focuses on **project scaffolding**, **architecture-first thinking**, and **developer experience**.

---

## 🧠 Supported Architectures

Choose your project architecture during setup:

- **Layered Service Architecture** (recommended)
- **Domain-Driven Design (DDD)**
- **Minimal MVC-style**
- **Clean Architecture** _(under development)_

Each option generates a well-structured and scalable folder layout. The layered architecture scaffolder is implemented and creates a `src/` prefixed set of folders for the new project.

---

## 🖥️ CLI Features

- Interactive step-by-step setup
- Terminal UI powered prompts
- Architecture selection
- Configuration style selection
- Project summary and confirmation
- Optional prefix support (e.g. `src`)
- Non-interactive mode for CI/CD

---

## 🚀 Getting Started

### Clone the repository

```bash
git clone https://github.com/shovan04/ExpressTS-in-GO.git
cd ExpressTS-in-GO
```

### Run the CLI

Run directly from the repository root:

```bash
go run main.go
```

Or build and run the compiled binary:

```bash
go build -o ExpressTS
./ExpressTS
```

---

## 📦 Initialize a Project

```bash
ExpressTS init
```

<!-- --- -->

<!-- ## ⚡ Non-Interactive Mode

```bash
expressts init --yes
``` -->

---

## 📁 Example Generated Structure (Layered)

The scaffolder creates a `src/` directory inside the new project and places architecture folders there:

```text
my-api/
└── src/
	├── bin/
	├── configs/
	├── constants/
	├── controllers/
	├── services/
	├── repositories/
	├── DTOClasses/
	├── exceptions/
	├── middlewares/
	├── routes/
	├── utilities/
	├── interfaces/
	└── mappers/
```

---

## 📜 License

MIT License

---

## ✅ Project Progress

- [x] get project details from user
- [x] check and validate the inserted data

### Layered Folder Structure

- [x] init
- [x] create project folder
- [x] create package.json tsconfig.json and .env
- [x] created all the required data files
- [x] implemented all the data files
- [x] Update the DTO Validator Middleware
- [x] Update the ValidateDto middleware

### DDD Architecture

- [x] add exceptions folder to domain layer
- [x] add exception, dto, and middleware files to scaffolding
- [x] remove try-catch block from user controller
- [x] add validation middleware and error handling to user routes
- [x] add conflict exception
- [x] add error response DTO
- [x] add global error handler middleware
- [x] add generic response DTO
- [x] add DTO validation middleware
- [x] add validation exception response
- [x] correct validateDto import path in userRoutes

### MVC Architecture

- [x] Plan MVC folder structure and templates
- [x] Create templates (Model, Controller, Routes, Middleware, MainApp)
- [x] Implement initialization logic in init.go and folders.go
- [x] Register MVC option in init.go
- [x] Verify with test_init_mvc_architecture
- [x] Add DTOs in utilities/dtos
- [x] Add Generic Response/Error DTOs and Exceptions in utilities
