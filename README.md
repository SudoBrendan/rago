# RAGo

[![Go Version](https://img.shields.io/github/go-mod/go-version/SudoBrendan/rago)](https://golang.org)
[![Build Status](https://github.com/SudoBrendan/rago/actions/workflows/ci.yml/badge.svg)](https://github.com/SudoBrendan/rago/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/SudoBrendan/rago)](https://goreportcard.com/report/github.com/SudoBrendan/rago)
[![License](https://img.shields.io/github/license/SudoBrendan/rago)](https://github.com/SudoBrendan/rago/blob/main/LICENSE)

> A fast and extensible Retrieval-Augmented Generation (RAG) CLI, written in Go.

## Overview

RAGo is a flexible RAG framework and CLI tool, designed for developers who want full control over:

- **Models:** Which LLMs are used to generate responses
- **Embedders:** How documents are embedded into vector space
- **VectorStores:** Where embeddings are stored and retrieved from
- **Loaders:** How information is fetched, parsed, and chunked for storage

All of these types are **factory-based** and **configurable** via a kubeconfig-like config file at `~/.rago/config`.

Factories allow easy plugin registration and extension without the boilerplate.
Current plugins primarily align with the interfaces from `langchaingo`.

### Goals

RAGo is a pluggable RAG engine, configured by YAML, powered by Go.

- Fast, configurable RAG pipelines
- Easy plugin system for real-world applications with custom logic
- Support for any backend without lock-in
- Powered by Go for performance and portability

## Current Plugins

### Models

- ollama

### Embedders

- ollama

### VectorStores

- pgvector

### Loaders

- markdown

## Configuration

All configuration for all plugins is managed through a single YAML file located at `~/.rago/config`.

A sample config looks like this:

```yaml
models:
- name: mistral
  kind: ollama
  options:
    model: mistral
    serverURL: http://localhost:11434

vectorStores:
- name: pgvector
  kind: pgvector
  options:
    connectionURL: "postgres://testuser:testpass@localhost:5432/testdb"
    preDeleteCollection: true
    embedder:
      name: nomic-embed-text
      kind: ollama
      options:
        model: nomic-embed-text
        serverURL: http://localhost:11434

loaders:
- name: my-wiki-loader
  kind: markdown
  options:
    directory: ~/my-wiki

# Contexts are the glue that map all the pieces together, and allow
# you to easily mix/match models, stores, and loaders in quick, new
# configurations
contexts:
- name: wiki-rag
  model: mistral
  vectorStore: pgvector
  loader: my-wiki-loader

# The current-context is what you are currently set up to use
# on your next rago command.
current-context: wiki-rag
```

## Usage

Basic commands:

```sh
# Add documents to your vectorstore (again, assuming your config choices!)
rago vectorstore add-documents

# More commands comming soon...

## Delete all documents from your vectorstore
#rago vectorstore delete
#
## Start chatting with your LLM, with retrieval augmentation
#rago chat
```

## For Developers - Extensibility

Adding new models, loaders, embedders, or vectorstores is easy:

- Create a new plugin in `/plugins`
- Implement the proper interface based on the factory you want to Register to
- Register your plugin with the Factory via `init()` using the factory system, and add your plugin to the top-level `register.go`

No boilerplate, just what makes your RAG yours.

## Acknowledgements

- `github.com/tmc/langchaingo`: for the foundational interfaces