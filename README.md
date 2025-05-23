# 🎬 Videostreaming Microservice

This microservice allows you to:

- Accept a video file via a public URL
- Split the video into parts of configurable duration (in seconds)
- Upload the segments to MinIO
- Return the list of accessible URLs for those parts via gRPC

## 🛠 Technologies

- Go (Golang)
- gRPC
- FFmpeg
- MinIO (S3-compatible object storage)
- Docker + Docker Compose

## 📦 Usage

### 1. Build and run

```bash
docker-compose up --build
```

### 2. gRPC Interface

#### Method: `ProcessVideo`

```proto
rpc ProcessVideo (VideoRequest) returns (VideoPartsResponse);
```

**VideoRequest:**
- `video_url` (string): public URL to the video
- `segment_duration` (int): segment duration in seconds
- `output_format` (optional): default is "mp4"
- `filename_prefix` (optional): prefix for split files
- `cleanup` (optional): delete temp file after processing
- `max_duration` (optional): future support

**VideoPartsResponse:**
- `parts` (repeated string): array of part URLs from MinIO

## 📁 Folder Structure

```
videostreaming/
├── cmd/
├── proto/
├── internal/
│   ├── service/
│   ├── downloader/
│   ├── ffmpeg/
│   ├── minio/
│   └── config/
├── output/
├── temp/
├── Dockerfile
├── docker-compose.yml
├── .env
└── README.md
```

## 🧪 Notes

- Only `.mp4` files are supported for now (can be extended)
- FFmpeg does not re-encode (uses `-c copy`)
- Logs are marked via TODO comments (plug your own system)
