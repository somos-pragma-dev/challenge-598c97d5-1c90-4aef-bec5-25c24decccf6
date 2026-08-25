#!/bin/sh

# Health check script for Docker
curl -f http://localhost:8080/health || exit 1