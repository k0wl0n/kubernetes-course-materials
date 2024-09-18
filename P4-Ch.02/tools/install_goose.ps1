# Download the file using Invoke-WebRequest
$downloadUrl = "https://github.com/pressly/goose/releases/download/v3.22.0/goose_windows_x86_64.exe"
$destinationPath = "$env:USERPROFILE\goose_windows_x86_64.exe"

# Download the goose executable
Invoke-WebRequest -Uri $downloadUrl -OutFile $destinationPath -UseBasicParsing

# Define the folder where you want to move the executable
$targetDirectory = "$env:USERPROFILE\goose-tools"

# Create the directory if it doesn't exist
if (!(Test-Path -Path $targetDirectory)) {
    New-Item -ItemType Directory -Force -Path $targetDirectory
}

# Move the downloaded file to the target directory
Move-Item -Path $destinationPath -Destination $targetDirectory

# Add the directory to the PATH if it's not already there
$existingPath = [System.Environment]::GetEnvironmentVariable("Path", "User")
if ($existingPath -notlike "*$targetDirectory*") {
    [System.Environment]::SetEnvironmentVariable("Path", "$existingPath;$targetDirectory", "User")
    Write-Host "The directory $targetDirectory has been added to your PATH."
} else {
    Write-Host "The directory $targetDirectory is already in your PATH."
}

Write-Host "Goose has been downloaded and added to your PATH."
