#!/bin/bash

source ./repo_setup/scripts/utils.sh
source ./repo_setup/scripts/readp.sh

propertiesFile="./settings.properties"
SLASH="/"

loadProperties $propertiesFile

declare -a PUBLIC_DIRECTORIES=($UPLOADS_PUBLIC_DIR $STATIC_PUBLIC_DIR)
declare -a ENVIRONMENTS=("prod" "staging" "dev" "test")

function getDotEnvsPath() {
  dotEnvsDirPath=$HOME$SLASH$DOT_ENVS_DIR
}

function getKeysPath() {
  keysDirPath=$HOME$SLASH$KEYS_DIR
}

function getPublicServerPath() {
  publicServerDirPath=$HOME$SLASH$PUBLIC_DIR
}

function getLogsPath() {
  logsDirPath=$HOME$SLASH$LOGS_DIR
}

function createDir() {
  print_status "Checking if $1 exists"
  if [ -d $1 ]; then
    print_status "$1 exists"
  else
    print_status "$1 doesn't exist, creating it now"
    mkdir $1
    print_status "Directory created"
  fi
  end_line
}

function createFile() {
  print_status "Checking if $1 exists"
  if [ -f $1 ]; then
    print_status "File $1 exists"
  else
    print_status "File $1 doesn't exist, creating it now"
    touch $1
    print_status "File created"
  fi
  end_line
}

print_info "$(cat $ASCII_FILE)"
end_line
print_status "Setting up project..."
end_line
print_banner "Setting up dotenvs directory"

# Create $HOME/dotenvs directory if doesn't exist
getDotEnvsPath
createDir $dotEnvsDirPath

print_banner "Setting up dotenvs/app-name directory"

# Create $HOME/dotenvs/$APP_NAME directory
appDotEnvDirPath=$dotEnvsDirPath$SLASH$APP_NAME
createDir $appDotEnvDirPath

end_line
print_banner "Setting up .env files for each environment with default properties"

# Creating environment directories inside dotenvs/app-name/... 
for env in  ${ENVIRONMENTS[@]};
do
  createDir $appDotEnvDirPath$SLASH$env
  filePath=$appDotEnvDirPath$SLASH$env$SLASH$DOT_ENV_FILE
  createFile $filePath
  cp ".env.example" $filePath
done

end_line
print_banner "Setting up OpenSSL keys"

# Create $HOME/keys directory if doesn't exist
getKeysPath
createDir $keysDirPath

end_line

# Create $HOME/keys/$APP_NAME directory
appKeysDirPath=$keysDirPath$SLASH$APP_NAME
createDir $appKeysDirPath

# Creating environment directories inside $HOME/keys/app-name/...
for env in  ${ENVIRONMENTS[@]};
do
  createDir $appKeysDirPath$SLASH$env

  # access-token keys
  PvtAccessTokenKeyPath=$appKeysDirPath$SLASH$env$SLASH$PVT_ACCESS_KEY
  PubAccessTokenKeyPath=$appKeysDirPath$SLASH$env$SLASH$PUB_ACCESS_KEY

  openssl genrsa -out $PvtAccessTokenKeyPath $OPEN_SSL_KEY_SIZE_ACCESS_TOKEN
  openssl rsa -in $PvtAccessTokenKeyPath -outform PEM -pubout -out $PubAccessTokenKeyPath

  # refresh-token key
  PvtRefreshTokenKeyPath=$appKeysDirPath$SLASH$env$SLASH$PVT_REFRESH_KEY
  PubRefreshTokenKeyPath=$appKeysDirPath$SLASH$env$SLASH$PUB_REFRESH_KEY

  openssl genrsa -out $PvtRefreshTokenKeyPath $OPEN_SSL_KEY_SIZE_REFRSH_TOKEN
  openssl rsa -in $PvtRefreshTokenKeyPath -outform PEM -pubout -out $PubRefreshTokenKeyPath

done

end_line
print_banner "Setting up public_server directories"

# Creating directories for public/uploads and public/static
getPublicServerPath
createDir $publicServerDirPath

appPublicServerPath=$publicServerDirPath$SLASH$APP_NAME
createDir $appPublicServerPath

for pubDir in  ${PUBLIC_DIRECTORIES[@]};
do
  createDir $appPublicServerPath$SLASH$pubDir
done

end_line
print_banner "Setting up app_logs directory"

getLogsPath
createDir $logsDirPath

appLogsDirPath=$logsDirPath$SLASH$APP_NAME
createDir $appLogsDirPath

for env in  ${ENVIRONMENTS[@]};
do
  createDir $appLogsDirPath$SLASH$env
done

end_line

print_success "Thank you for using Go REST API Template. Happy coding!"
