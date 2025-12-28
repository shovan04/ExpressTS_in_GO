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
- **Clean Architecture** *(under development)*

Each option generates a well-structured and scalable folder layout.

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

```bash
go run ./cmd/expressts
```

Or build once and run:

```bash
go build -o expressts ./cmd/expressts
./expressts
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

```text
my-api/
├── bin/
├── config/
├── constants/
├── controllers/
├── services/
├── repositories/
├── DTO/
├── exceptions/
├── middlewares/
├── routes/
├── utils/
├── interfaces/
└── mappers/
```

---

## 📜 License

MIT License
