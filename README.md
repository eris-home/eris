# Eris - Modular Conversational AI Framework

## Overview

Eris is a highly modular, scalable, and extensible conversational AI framework built in Go. It is designed to provide developers with a robust foundation for building advanced conversational systems that are platform-independent, customizable, and high-performing.

With support for multiple LLM providers, insight generation, behavior management, and extensive storage capabilities, Eris is ideal for creating anything from simple chatbots to complex enterprise-grade conversational agents.

---

## Key Features

### Core Architecture
- **Pluggable Framework**: Dynamic component swapping for flexibility and scalability.
- **Event-Driven Design**: Subscribe to and trigger custom events for extensible functionality.
- **Middleware Support**: Preprocess or modify input dynamically.

### LLM Integration
- **Multi-Provider Abstraction**: Connect to OpenAI, Hugging Face, or custom LLMs seamlessly.
- **Token Management**: Utilities for token counting and input chunking.
- **Mock Providers**: For local testing and development without real API usage.

### State and Context
- **Session Management**: Manage user sessions with isolated and shared data layers.
- **Cache Layer**: In-memory caching with support for TTL and cleanup.
- **Global State**: Maintain a centralized state for cross-manager communication.

### Insight Generation
- **Extractors and Analyzers**: Tools for sentiment analysis, topic detection, and summarization.
- **Reporters**: Generate detailed reports in both human-readable and JSON formats.
- **Formatter Utilities**: Transform raw insights into usable outputs.

### Task and Queue Management
- **Task Queue**: Queue and process background tasks asynchronously.
- **Scheduler**: Run recurring tasks with customizable intervals.
- **Worker Pool**: Execute multiple tasks in parallel efficiently.

### Storage Layer
- **Database Abstraction**: PostgreSQL integration with support for CRUD operations.
- **Indexing System**: Create and manage indices for fast data retrieval.
- **Backup System**: Easily create and manage backups of stored data.

### Observability
- **Logger**: Multi-level logging system (Info, Debug, Error).
- **Analytics**: Track and report system usage, events, and custom metrics.
- **Event System**: Event-driven architecture for extensibility.

---

## Installation

### Prerequisites
- Go >= 1.19
- PostgreSQL (optional, for persistent storage)

### Clone the Repository
```bash
git clone https://github.com/eris-home/eris
cd eris
```

### Configure Environment
Create a `.env` file based on the provided `.env.example`:
```env
DB_URL=postgresql://user:password@localhost:5432/eris
OPENAI_API_KEY=your_openai_api_key
```

### Install Dependencies
```bash
go mod download
```

---

## Quick Start

### Running the Engine
```bash
go run examples/cli/main.go
```

### Example: Adding a Route
Here is a simple example of using the `Engine` to add a route:
```go
package main

import (
    "github.com/eris-home/eris/engine"
)

func main() {
    e := engine.NewEngine()

    // Register a simple route
    e.RegisterRoute("hello", func(input string) {
        fmt.Println("Hello, Eris!")
    })

    e.Run("hello")
}
```

### Example: Event System
```go
e.On("user_logged_in", func(event string, data map[string]interface{}) {
    fmt.Printf("Event: %s triggered with data: %v
", event, data)
})
e.Trigger("user_logged_in", map[string]interface{}{"user_id": 123, "name": "Alice"})
```

---

## Architecture

The project is structured into the following directories:

### **Engine**
- **`engine/engine.go`**: Core execution logic, middleware, lifecycle management.
- **`engine/logger.go`**: Handles structured logging.
- **`engine/router.go`**: Maps user input to specific handlers.
- **`engine/utils.go`**: Provides utility functions like input normalization.
- **`engine/validator.go`**: Validates user input for correctness.

### **LLM**
- **`llm/llm.go`**: Defines interfaces for LLMs and a manager to handle providers.
- **`llm/connector.go`**: Manages HTTP connections to LLM APIs.
- **`llm/helper.go`**: Token counting, chunk splitting utilities.
- **`llm/mock.go`**: Mock provider for testing without real API calls.
- **`llm/provider.go`**: Generic provider abstraction.

### **State**
- **`state/state.go`**: Global state management.
- **`state/context.go`**: Per-session context management.
- **`state/tracker.go`**: Tracks changes and state events.
- **`state/session.go`**: Manages session lifecycle and metadata.
- **`state/cache.go`**: In-memory cache implementation.

### **Insight**
- **`insight/analyzer.go`**: Performs sentiment analysis, topic extraction, etc.
- **`insight/extractor.go`**: Extracts insights from conversations.
- **`insight/formatter.go`**: Formats insights into human-readable or structured formats.
- **`insight/insight.go`**: Manages insights and their storage.
- **`insight/reporter.go`**: Generates reports based on insights.

### **Storage**
- **`storage/storage.go`**: Abstracts common storage operations.
- **`storage/database.go`**: Connects to and interacts with PostgreSQL.
- **`storage/indexer.go`**: Creates and manages indices for fast data lookup.
- **`storage/backup.go`**: Handles backup creation and management.
- **`storage/retriever.go`**: Retrieves data from storage.

### **Managers**
- **`manager/manager.go`**: Base manager system.
- **`manager/analytics.go`**: Tracks system usage and custom events.
- **`manager/observer.go`**: Implements the observer pattern.
- **`manager/scheduler.go`**: Schedules recurring tasks.
- **`manager/worker.go`**: Handles background task execution.

---

## Testing

### Run Unit Tests
```bash
go test ./... -v
```

### Mock Provider
Use the `MockProvider` to test LLM interactions locally:
```go
mock := llm.NewMockProvider("TestMock")
response, _ := mock.GenerateText(llm.LLMRequest{Prompt: "Hello"})
fmt.Println(response.Text)
```

---

## Contribution

We welcome contributions! Please follow these steps:
1. Fork the repository.
2. Create a new branch for your feature.
3. Commit and push your changes.
4. Open a pull request.

---

## License

Eris is licensed under the MIT License. See `LICENSE` for details.

---

## Contact

For questions or support, please contact us at **support@eris-home.com**.
