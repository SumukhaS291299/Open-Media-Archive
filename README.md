# OpenMediaArchive

Automatically optimize media files for self-hosted storage.

OpenMediaArchive monitors uploaded media, analyzes it using FFprobe, and applies configurable FFmpeg optimizations based on storage tiers such as Hot, Warm, Cold, and Archive.

The goal is simple: reduce storage consumption while preserving the quality level appropriate for the intended storage tier.

## Motivation

Modern devices produce increasingly large media files:

- High bitrate recordings
- 4K and HDR videos
- Screen recordings
- Camera exports

For most personal media collections, storing every file at its original bitrate and resolution is often unnecessary.

OpenMediaArchive helps automate storage optimization using configurable policies rather than manual transcoding workflows.

## Planned Features

### Storage Tiers

Define optimization targets for:

- Hot Storage
- Warm Storage
- Cold Storage
- Archive Storage

Each tier can specify limits such as:

- Maximum bitrate
- Maximum resolution
- Preferred codecs
- Audio settings

### Media Analysis

Use FFprobe to inspect media before processing:

- File size
- Bitrate
- Resolution
- Codec
- Duration
- Audio streams

### Policy-Based Decisions

Examples:

- Skip files already within limits
- Re-encode oversized videos
- Reduce bitrate for archive storage
- Downscale videos exceeding configured resolutions

### Upload Method Agnostic

OpenMediaArchive does not handle uploads directly.

It is designed to work alongside existing solutions such as:

- Syncthing
- rsync
- SMB shares
- NFS shares
- Nextcloud
- SFTP

Any workflow that delivers files into a monitored directory can be supported.

### Resource-Aware Processing

Designed with home servers and NAS environments in mind.

Goals include:

- Controlled CPU usage
- Controlled memory usage
- Safe parallel processing
- Predictable resource consumption

## Non-Goals

OpenMediaArchive is not intended to be:

- A media server
- A file synchronization tool
- A manual transcoding application
- A replacement for FFmpeg

Instead, it focuses on automated, policy-driven media optimization.

## Status

Early development / planning stage.

Feedback and ideas are welcome.
