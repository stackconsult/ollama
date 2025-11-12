# Ollama Setup & Usage Guide

## What Was Built

This setup includes:
1. **Ollama Server** - Built from source (`~/Desktop/ollama/ollama`)
2. **Web UI** - React/Vite interface at `app/ui/app/`
3. **Model Management** - `code-assistant` model (llama3.2 alias)
4. **Rollback Script** - `scripts/ollama-rollback.sh` for backup/restore

## Running the Application

### 1. Start Ollama Server
```bash
cd ~/Desktop/ollama
nohup ./ollama serve > /tmp/ollama.log 2>&1 &
```

### 2. Start Web UI
```bash
cd ~/Desktop/ollama/app/ui/app
npm run dev
```

### 3. Access UI
Open browser to: http://localhost:5173/

## Model Management

### List Models
```bash
cd ~/Desktop/ollama
./ollama list
```

### Create Model Alias
```bash
./ollama cp <source-model> <new-name>
```

### Remove Model
```bash
./ollama rm <model-name>
```

### Run Model (CLI)
```bash
./ollama run code-assistant
```

## Backup & Rollback

### Create Backup
```bash
~/Desktop/ollama/scripts/ollama-rollback.sh
```

Backups saved to: `~/.ollama-backups/`

### Restore Previous State
1. View backup: `cat ~/.ollama-backups/backup_YYYYMMDD_HHMMSS.txt`
2. Remove unwanted models: `./ollama rm <model-name>`
3. Prune cache: `./ollama prune -f`

## Architecture

- **Backend**: Ollama server on `http://localhost:11434`
- **Frontend**: Vite dev server on `http://localhost:5173`
- **API Proxy**: Vite proxies `/api` requests to backend
- **Models**: Stored in `~/.ollama/models/`

## Current Models

- `code-assistant:latest` (2.0 GB) - llama3.2 alias optimized for coding

## Troubleshooting

### Server Not Responding
```bash
# Check if running
ps aux | grep "ollama serve"

# Restart server
pkill -f "ollama serve"
cd ~/Desktop/ollama && nohup ./ollama serve > /tmp/ollama.log 2>&1 &
```

### UI Not Connecting
```bash
# Test API
curl http://localhost:11434/api/tags

# Restart UI
cd ~/Desktop/ollama/app/ui/app
pkill -f vite
npm run dev
```

### Memory Issues
The model requires ~2GB RAM. If you see memory errors:
- Close other applications
- Use a smaller model variant like `llama3.2:1b`

## Git Repository

Branch: `copilot/install-browser-functionality`
Repo: https://github.com/stackconsult/ollama

### Push Changes
```bash
cd ~/Desktop/ollama
git add .
git commit -m "Your message"
git push origin copilot/install-browser-functionality
```

## Files Added/Modified

- `scripts/ollama-rollback.sh` - Backup/restore script
- `app/ui/app/vite.config.ts` - Added API proxy configuration
- `app/ui/app/package-lock.json` - Dependencies
- `package-lock.json` - Root dependencies

## Next Steps

1. Sign in to Ollama.com to push models remotely
2. Configure git identity: `git config --global user.name "Your Name"`
3. Update browser data: `npx update-browserslist-db@latest`
