# Basic Demo

## Configure

This demo leverages dockerized (but locally mounted) ollama, and dockerized (and locally persisted) pgvector to showcase configuration of the rago cli.

## Env Setup

```sh
# set up your rago config
cp ./config ~/.rago/config

# set up ollama and pgvector via docker
docker-compose up

# ensure you have all the required models
ollama pull mistral
ollama pull nomic-embed-text
```

> NOTE: see https://hub.docker.com/r/ollama/ollama for more info on GPU support for dockerized ollama.

You can curl your localhost to ensure all LLMs are present: `curl http://localhost:11434/api/tags`

## Set up the RAG

### Populate the vectorstore with local data

```sh
# Make a local data directory with markdown documents to upload
# Our config file specifies the location under `loaders`
mkdir -p ~/rago/data
echo "# Hello World" > ~/rago/data/my-file.md

# Parse the documents into vectors, add them to the store
rago vectorstore add-documents
```

## Next Step: Plugins

This example showcased a single setup with specific tools, but you'll notice the code is set up in a factory pattern - you can easily add more plugins for all
the pluggable components. As long as they implement the interfaces (and ideally read from config), you can create your own tools to load/store/generate content to your
heart's content! See `/plugins` for additional details.
