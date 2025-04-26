FROM postgres:16

RUN apt-get update && \
    apt-get install -y --no-install-recommends postgresql-16-pgvector && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN mkdir -p /docker-entrypoint-initdb.d

# NOTE: I don't think we need this, pgvector includes logic
# to add the extension if not present on init of clients.
COPY setup.sql /docker-entrypoint-initdb.d