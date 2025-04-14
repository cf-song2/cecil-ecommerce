#!/bin/sh

./wait-for-it.sh db:5432 --timeout=30 --strict -- echo "✅ DB is ready"

echo "Running migrations..."
psql "$DATABASE_URL" -f ./migrations/001_init.up.sql

echo "Seeding data..."
./seed

echo "Starting server..."
./server
