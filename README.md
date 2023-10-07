# githubRepoAnalyzer
This application fetches the latest release information for a list of GitHub repositories and predicts the next version number based on the current release.

## Pre-requisites

1. [Go](https://golang.org/doc/install) (To run the application natively)
2. [Docker](https://www.docker.com/products/docker-desktop) (To run the application in a container)
3. A GitHub Personal Access Token. Follow [this guide](https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token) to generate one.

## Steps to Run Natively

1. Clone the repository:
   ```bash
   git clone <your-repository-url>
   cd <repository-directory>
   ```

2. Update the Go code (`githubRepoAnalyzer.go`) with your GitHub Personal Access Token. Replace `YOUR_GITHUB_TOKEN` with your token.

3. Build and run:
   ```bash
   go build -o main
   ./main
   ```

## Steps to Run in Docker

1. Build the Docker image:
   ```bash
   docker build -t github-analyzer .
   ```

2. Run the Docker container:
   ```bash
   docker run github-analyzer
   ```

