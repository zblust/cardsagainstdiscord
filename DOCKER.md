# Docker Deployment Guide

This guide explains how to deploy the Cards Against Humanity Discord bot using Docker on your Unraid server.

## Prerequisites

- Docker installed on your Unraid server
- Discord bot token from [Discord Developer Portal](https://discord.com/developers/applications)

## Quick Start with Docker CLI (Recommended for Unraid)

### Step 1: Build the Image

```bash
cd /path/to/cardsagainstdiscord
docker build -t cardsagainstdiscord:latest .
```

### Step 2: Run the Container

**Replace `YOUR_DISCORD_BOT_TOKEN_HERE` with your actual Discord bot token:**

```bash
docker run -d \
  --name cah-discord-bot \
  --restart unless-stopped \
  -e DG_TOKEN="Bot YOUR_DISCORD_BOT_TOKEN_HERE" \
  cardsagainstdiscord:latest
```

**Example with a real token format:**
```bash
docker run -d \
  --name cah-discord-bot \
  --restart unless-stopped \
  -e DG_TOKEN="Bot TOKENHERE" \
  cardsagainstdiscord:latest
```

### Step 3: Verify It's Running

```bash
# Check container status
docker ps | grep cah

# View logs
docker logs -f cah-discord-bot
```

You should see output like:
```
2025/11/19 12:34:56 Running...
```

## Unraid Server Setup (Recommended Method)

### Method 1: Using Unraid Docker UI (Easiest)

1. **SSH into your Unraid server and build the image:**
   ```bash
   cd /mnt/user/appdata/
   git clone https://github.com/zblust/cardsagainstdiscord.git
   cd cardsagainstdiscord
   docker build -t cardsagainstdiscord:latest .
   ```

2. **In Unraid Web Interface:**
   - Go to **Docker** tab
   - Click **"Add Container"**
   - Configure as follows:
     - **Name**: `cah-discord-bot`
     - **Repository**: `cardsagainstdiscord:latest`
     - **Network Type**: `bridge`
     - **Console shell command**: `Bash`
     - Click **"Add another Path, Port, Variable, Label or Device"**
     - Select **Variable**
       - **Key**: `DG_TOKEN`
       - **Value**: `Bot YOUR_DISCORD_BOT_TOKEN_HERE`
     - **Restart Policy**: `unless-stopped`
   - Click **"Apply"**

### Method 2: Using Command Line on Unraid

```bash
# SSH into your Unraid server
cd /mnt/user/appdata/
git clone https://github.com/zblust/cardsagainstdiscord.git
cd cardsagainstdiscord

# Build the image
docker build -t cardsagainstdiscord:latest .

# Run with your token (replace YOUR_TOKEN)
docker run -d \
  --name cah-discord-bot \
  --restart unless-stopped \
  -e DG_TOKEN="Bot YOUR_DISCORD_BOT_TOKEN_HERE" \
  cardsagainstdiscord:latest

# Check if it's running
docker logs -f cah-discord-bot
```

## Managing the Container

### Start the bot
```bash
docker start cah-discord-bot
```

### Stop the bot
```bash
docker stop cah-discord-bot
```

### Restart the bot
```bash
docker restart cah-discord-bot
```

### View logs (live)
```bash
docker logs -f cah-discord-bot
```

### View recent logs
```bash
docker logs --tail 100 cah-discord-bot
```

### Remove the container
```bash
docker stop cah-discord-bot
docker rm cah-discord-bot
```

## Updating the Bot

When you pull new changes from the repository:

```bash
# Navigate to the repo directory
cd /mnt/user/appdata/cardsagainstdiscord

# Pull latest changes
git pull

# Rebuild the image
docker build --no-cache -t cardsagainstdiscord:latest .

# Stop and remove old container
docker stop cah-discord-bot
docker rm cah-discord-bot

# Start new container with same settings (replace YOUR_TOKEN)
docker run -d \
  --name cah-discord-bot \
  --restart unless-stopped \
  -e DG_TOKEN="Bot YOUR_DISCORD_BOT_TOKEN_HERE" \
  cardsagainstdiscord:latest
```

**Tip:** You can create a simple update script to make this easier.

## Troubleshooting

### Check if container is running
```bash
docker ps | grep cah
```

### View recent logs
```bash
docker logs --tail 100 cah-discord-bot
```

### Check environment variables
```bash
docker inspect cah-discord-bot | grep -A 10 Env
```

### Container won't start
1. Check logs: `docker logs cah-discord-bot`
2. Verify token is set correctly
3. Ensure token starts with `Bot `
4. Check Discord bot has proper permissions

## Getting Your Discord Bot Token

1. Go to [Discord Developer Portal](https://discord.com/developers/applications)
2. Create a new application or select existing one
3. Go to "Bot" section
4. Click "Reset Token" to get your token
5. Copy the token (format: `Bot YOUR_TOKEN_HERE`)
6. Enable required intents:
   - Server Members Intent
   - Message Content Intent

## Bot Permissions

When inviting the bot to your server, use these permissions:
- Read Messages/View Channels
- Send Messages
- Embed Links
- Add Reactions
- Read Message History
- Use External Emojis
- Manage Messages (optional, for cleanup)

## Advanced Configuration

### Enable pprof Debugging

Uncomment the ports section in `docker-compose.yml`:
```yaml
ports:
  - "7447:7447"
```

Then access pprof at `http://your-server-ip:7447/debug/pprof/`

### Custom Network

To use a custom Docker network:
```yaml
networks:
  discord-bots:
    external: true

services:
  cah-bot:
    networks:
      - discord-bots
```

## Security Notes

- Never commit your `.env` file with real tokens
- Use `.env` file locally, environment variables in production
- The container runs as non-root user (UID 1000)
- Only expose pprof port (7447) if needed for debugging
- Keep your Discord bot token secure
