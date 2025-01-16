# Eris - Modular Conversational AI Framework

## Overview
Eris is a scalable, modular AI framework designed to create and manage advanced conversational systems. Built with Go, Eris emphasizes customizability, high performance, and platform flexibility. Its architecture supports:

- Dynamic component swapping through pluggable plugins
- Multi-model AI integration for flexibility and redundancy
- Multi-platform conversation management
- Advanced insight generation and user behavior tracking
- Semantic data storage powered by vector embeddings

---

## Core Features

### Plugin Architecture
- **Core Manager System**: Build custom behaviors with ease.
- **Conversation Manager**: Process conversations and manage their flow.
- **Behavior Manager**: Adapt AI tone, style, and interaction preferences dynamically.
- **Custom Manager Support**: Implement specialized capabilities tailored to your needs.

### State Management
- **Global State Management**: A centralized store for all shared and contextual data.
- **Manager-Specific State**: Isolated state layers for modularity.
- **Dynamic Injection**: Inject external data sources or processing logic on demand.
- **Inter-Manager Communication**: Allow managers to coordinate and share context seamlessly.

### LLM Integration
- **Provider Abstraction**: Support for multiple LLM providers out of the box.
- **Native OpenAI Integration**: Pre-configured for seamless operation.
- **Extensible Interfaces**: Define your own AI providers via the `LLMProvider` interface.
- **Customizable Model Selection**: Choose different models for specific tasks or fallback scenarios.
- **Retry Handling**: Automatic retries for improved reliability.

### Platform Support
- **Platform-Agnostic Core**: Decoupled from platform-specific dependencies.
- **Native Examples**: Examples for CLI and Slack integration.
- **Extensible Platform System**: Define your own platform adapters for integrations like Discord, WhatsApp, or enterprise tools.

### Storage Layer
- **Semantic Data Storage**: Uses PostgreSQL with vector embeddings for advanced search capabilities.
- **ORM-Based Models**: Simplified database interactions.
- **Customizable Storage**: Extend or replace the storage system for your unique needs.

### Toolkit System
- **Pluggable Tools**: Easily add and manage toolkits.
- **Built-in Tool Manager**: Manage and invoke tools during conversations.
- **Function Support**: Call external APIs or services dynamically.
- **Context-Aware Execution**: Tools operate with real-time context for better results.

---

## Quick Start

1. **Clone the repository:**

   ```bash
   git clone https://github.com/eris-home/eris
   ```

2. **Configure environment variables:**
   Copy `.env.example` to `.env` and update settings:

   ```env
   DB_URL=postgresql://user:password@localhost:5432/eris
   OPENAI_API_KEY=your_openai_api_key
   ```

3. **Install dependencies:**

   ```bash
   go mod download
   ```

4. **Run a CLI example:**

   ```bash
   go run examples/cli/main.go
   ```

5. **Run the Slack bot:**

   ```bash
   go run examples/slack/main.go
   ```

---

## Architecture

The project is designed with clean modularity:

- **engine**: The conversation engine core.
- **manager**: Manager framework for modular functionality.
- **managers/**: Built-in manager implementations.
- **state**: Shared state utilities for conversation contexts.
- **llm**: LLM provider interfaces and implementations.
- **storage**: Data storage utilities and vector-based indexing.
- **tools/**: Built-in tools for dynamic functionality.
- **examples/**: Platform-specific examples and demos.

---

## Using Eris as a Module

### Add Eris to your project:

```bash
go get github.com/eris-home/eris
```

### Import the relevant packages:

```go
import (
  "github.com/eris-home/eris/engine"
  "github.com/eris-home/eris/llm"
  "github.com/eris-home/eris/manager"
  "github.com/eris-home/eris/managers/conversation"
  "github.com/eris-home/eris/managers/behavior"
)
```

### Basic Usage Example:

```go
// Initialize the LLM client
llmClient, err := llm.NewClient(llm.Config{
  Provider:    llm.ProviderOpenAI,
  APIKey:      os.Getenv("OPENAI_API_KEY"),
  DefaultModel: "gpt-4",
})

// Create a new engine instance
engine, err := engine.NewEngine(
  engine.WithLLM(llmClient),
  engine.WithStorage(storageClient),
)

// Create a new conversation state
state, err := engine.NewState("actor_id", "session_id", "What’s the weather like?")
if err != nil {
  log.Fatal(err)
}

// Process the user’s input
response, err := engine.Process(state)
if err != nil {
  log.Fatal(err)
}

fmt.Println(response.Text)
```

---

## Available Packages

- **eris/engine**: The conversation engine core.
- **eris/llm**: LLM interfaces and providers.
- **eris/manager**: Base manager system.
- **eris/managers/**: Built-in manager implementations.
- **eris/state**: State utilities for tracking context.
- **eris/storage**: Data storage implementations.
