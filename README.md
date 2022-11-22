# Web Whisper

Get OpenAI's [Whisper](https://github.com/openai/whisper) right into your browser!

This is a small personal project I am using to learn Golang and Svelte. It is a light web frontend for OpenAI's whisper.

## Contents:

- [Features](#features)
- [Roadmap](#roadmap)
- [Self-hosting](#self-hosting)
    - [With docker](#docker-compose)
    - [From source](#manual)
- [Screenshots](#screenshots)

## ✨ Features:

- [x] Record and transcribe audio right from your browser.
- [x] Upload any media file (video, audio) in any format and transcribe it.
    - [x] Option to cut audio to X seconds before transcription.
    - [x] Option to disable file uploads.
- [x] Download `.srt` subtitle file generated from audio.
- [x] Lightweight and beautiful UI.
- [x] Self-hosted. No 3rd parties.
- [x] **Docker compose** for easy self-hosting
- [x] Select input audio language
- [x] **Privacy respecting**: 
    - All happens locally. No third parties involved.
    - Audio files are deleted immediately after processing.
- [x] Backend written in **Go**
- [x] Frontend written with **Svelte** and **Tailwind CSS**.
- [x] Uses C++ whisper version from [whisper.cpp](https://github.com/ggerganov/whisper.cpp).
    - You don't need a GPU, uses CPU.
    - No need for complex installations.

## 🧭 Roadmap:

- [ ] Translate input audio transcription to english.
- [x] Allow to upload any file (video, audio) in any format and transcribe it.
	- [ ] Limit max file size for server hosting.
    - [x] Allow to cut audio to X seconds before transcription.
    - [x] Option to disable file uploads.
- [ ] Transcription history / save snippets
    - Publish to some **pastebin**-like service.

## 🪺 Self-hosting

> **Warning**
> If you plan to host a production instance, you should use the [docker compose](#docker-compose) method.

### Docker-compose

1. Clone this repo: `git clone https://codeberg.org/pluja/web-whisper`
2. Open and edit the `docker-compose.yml` to fit your needs. You should properly set the variables:

- `DOMAIN_NAME`: This is the domain where you plan to host and access `web-whisper`. In case you want to use it locally, you don't need to change it, leave `https://localhost:3000`. Note that if you change the port, you must also reflect this in this variable.
- `CUT_MEDIA_SECONDS`: This is a variable that lets you limit the seconds of the processed media. If you set it to, for example, `10` all media will be cut after 10 seconds. If you leave it as `0`, no limit is set.
- `ALLOW_FILE_UPLOADS`: This option enables or disables the ability to upload an arbitrary media file to the server for transcription. If set to `false` users will not be able to upload any files. If set to `true` (or any other string), users will be able to upload files.

#### Local usage (without domain):

3. `docker compose up -d`
4. Visit http://localhost:3000

#### Server setup

If you want to set up a reverse proxy, you can just:

3. `docker compose up -d`
5. Point your reverse proxy to the frontend:
    - `whisper:80` or `http://localhost:3000` (or any other port you choose) depending on if you choose to publish ports or not.

---

### Manual

This is a small guide on how to self-host this project. This set up should not be used for production purpose, instead you should use the [docker setup](#docker-compose). If you want to test without docker or just help with development, here is how you can do it:

This project is built of two parts; the `backend` and the `frontend`. Both need to be running at the same time, so you can make use of it.

First step is to clone this repository:

`git clone https://codeberg.org/pluja/web-whisper`

And enter the directory:

`cd web-whisper`

#### Running the backend

You will need `go` (Golang) installed on your computer.

- Now head to the `backend` directory: `cd backend`
- Simply run `go mod download` to get all de dependencies.
- Setup the `whipser.cpp` dependency: `bash setup.sh`
    - You can edit the `setup.sh` to download a better model other than the `small` if you want ([See this link](https://github.com/ggerganov/whisper.cpp#more-audio-samples))
- Run the backend `go run .`
    - The backend starts at `localhost:9090`; you can edit the `main.go` file to change the default port.

#### Running the frontend

You will need `npm` and `yarn` to run the frontend.

- Head to the `frontend` folder: `cd frontend`
- Run `npm i`
- Run `yarn`
- Run the frontend with `yarn dev`
    - The frontend will be available at the address in the console output.

## Screenshots

<p align="center"><sub>*Logo generated with Stable Diffusion*</sub></p>

##### Main page
<img src="https://farside.link/rimgo/GFBHU8V.png" align=center>

##### Recording
<img src="https://farside.link/rimgo/M5pW2BB.png" align=center>

#### Transcription Options
<img src="https://farside.link/rimgo/a4yf4hu.png" align=center>

#### Processing
<img src="https://farside.link/rimgo/SHOTbh8.png" align=center>

#### Result
<img src="https://farside.link/rimgo/8EodxT9.png" align=center>