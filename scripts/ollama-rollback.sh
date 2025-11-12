#!/bin/bash
# Ollama Rollback & Backup Script
# Created: November 11, 2025
# Purpose: Capture state before tagging operations and provide rollback capability

OLLAMA_BIN="$HOME/Desktop/ollama/ollama"
BACKUP_DIR="$HOME/.ollama-backups"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE="$BACKUP_DIR/backup_$TIMESTAMP.txt"

# Create backup directory
mkdir -p "$BACKUP_DIR"

echo "================================================"
echo "Ollama Backup & Rollback Script"
echo "Timestamp: $TIMESTAMP"
echo "================================================"
echo ""

# Capture current model list
echo "=== Current Models ===" | tee "$BACKUP_FILE"
$OLLAMA_BIN list | tee -a "$BACKUP_FILE"
echo "" | tee -a "$BACKUP_FILE"

# Capture disk usage
echo "=== Disk Usage ===" | tee -a "$BACKUP_FILE"
du -sh ~/.ollama 2>/dev/null | tee -a "$BACKUP_FILE"
echo "" | tee -a "$BACKUP_FILE"

# Capture blob count
echo "=== Storage Info ===" | tee -a "$BACKUP_FILE"
find ~/.ollama/models/blobs -type f 2>/dev/null | wc -l | xargs echo "Total blobs:" | tee -a "$BACKUP_FILE"
echo "" | tee -a "$BACKUP_FILE"

# Export model info to JSON
echo "=== Model Details (JSON) ===" | tee -a "$BACKUP_FILE"
$OLLAMA_BIN list --format json 2>/dev/null | tee -a "$BACKUP_FILE"
echo "" | tee -a "$BACKUP_FILE"

echo "Backup saved to: $BACKUP_FILE"
echo ""
echo "To restore to this state:"
echo "  1. Remove unwanted models: $OLLAMA_BIN rm <model-name>"
echo "  2. Remove unwanted tags: $OLLAMA_BIN rm <tag-name>"
echo "  3. Prune unused data: $OLLAMA_BIN prune -f"
echo ""
echo "Current state captured successfully!"
