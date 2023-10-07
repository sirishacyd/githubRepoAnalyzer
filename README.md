# GitHub Repository Analyzer

A Go module designed to analyze GitHub repositories, providing insights into their latest releases, upcoming version predictions, and more.

## Features

1. **Interactive User Input**: Allows users to specify a list of GitHub repositories in the format `owner/repo` for analysis.

2. **Repository Cloning**: Efficiently clones each user-specified repository to a local directory for further examination.

3. **Latest Release Analysis**: 
   - Fetches the most recent release for each repository.
   - Displays the list of changes encompassed in the latest release.

4. **Version Prediction**: Utilizing the principles of semantic versioning, the module predicts and displays the upcoming release version for each repository.

5. **YAML Report Generation**:
   - Generates a dedicated YAML file for every repository, capturing all the aforementioned details.
   - Each YAML file is named after its respective repository, wherein any `/` characters are substituted with `_`.

6. **Containerized Application**: The entire application is containerized using Docker, ensuring a consistent runtime environment with all the requisite dependencies.

## Usage

1. **Clone this Repository**:
   ```bash
   git clone [repository-link]
   cd [repository-name]
   ```

2. **Run the Go Module**:
   ```bash
   go run [module-name].go
   ```

3. **Follow On-screen Prompts**: Input the desired GitHub repository names when prompted.

4. **View Results**: Navigate to the generated YAML files for a detailed report on each repository.

5. **Containerization**:
   - Build the Docker image:
     ```bash
     docker build -t github-analyzer .
     ```

   - Run the Docker container:
     ```bash
     docker run github-analyzer
     ```

## Dependencies

Ensure you have the following installed:
- Go (at least version 1.16)
- Docker (if you wish to containerize the application)

