#!/bin/bash

declare -a environments=("prod" "staging" "dev" "test")
declare -a reqPublicDirs=("uploads" "static")
configDir="config"
keysDir="keys"
publicDir="public"

function isValidDirectory() {
  ASCII_FILE=repo_setup/scripts/ascii.txt
  if test -f "$ASCII_FILE"; then
    echo -e "\e[34m$(cat $ASCII_FILE)\e[0m"
  else
    echo -e "\e[31mYou shouldn't be running it outside a go project's directory\n exit 1\e[0m"
    exit 1
  fi
}

function checkIfProjectIsValid() {

  goModfile=go.mod
  
  if test -f "$ASCII_FILE"; then
    print_success "Found valid go project"
  else
    print_error "Invalid directory/invalid go project found."
    print_status "You must run this script from the project root directory as"
    print_bold "bash repo_setup/scripts/setup_project.sh <your_project_name>"
    print_error "exit 1"
    exit 1
  fi
}

function createPRootAndOtherDirectories() {
  print_info "Creating project root directories"
  originalDir=$(pwd)
  cd $HOME
  
  appDir=$1

  if [ ! -d "proot" ]; then
    print_status "Creating proot directory"
    mkdir "proot"
  fi

  cd proot
  if [ ! -d $appDir ]; then
    print_status "Creating $HOME/$appDir directory"
    mkdir $appDir
  fi

  cd $appDir

  for i in "${environments[@]}"; do

    print_header "Setting up $i environment"
    if [ ! -d "$i" ]; then
      print_status "Creating $HOME/$appDir/$i directory"
      mkdir $i
    fi
    cd $i

    createConfigs "$HOME/$appDir/$i" $originalDir
    createRSAKeys "$HOME/$appDir/$i" $originalDir
    createPublicDirs "$HOME/$appDir/$i" $originalDir

    cd ..
  done
  
  print_success "All requried directories and configs created"
  cd $originalDir
}

function createConfigs(){
  if [ ! -d $configDir ]; then
    print_status "Creating $1/$configDir directory"
    mkdir $configDir
  fi
  cd $configDir
  cp $2/.env.example .env
  cd ..
}

function createRSAKeys(){
  if [ ! -d $keysDir ]; then
    print_status "Creating $1/$keysDir directory"
    mkdir $keysDir
  fi
  cd $keysDir
  openssl genrsa -out pvt_access.pem 512
  openssl rsa -in pvt_access.pem -outform PEM -pubout -out pub_access.pem

  openssl genrsa -out pvt_refresh.pem 512
  openssl rsa -in pvt_refresh.pem -outform PEM -pubout -out pub_refresh.pem
  cd ..
}

function createPublicDirs(){
  mkdir -p $publicDir
  cd $publicDir
  for i in "${reqPublicDirs[@]}"; do
    if [ ! -d "$i" ]; then
      print_status "Creating $1/$i directory"
      mkdir $i
    fi
  done
  cd ..
}

function createConfigsAndEnvs() {
  print_status "Creating docker-compose files"
}

## -- main 

isValidDirectory

. repo_setup/scripts/utils.sh

checkIfProjectIsValid
createPRootAndOtherDirectories $1

# echo "------ Setting up your go-project -------"

# cd $HOME
# echo "Creating environment files and rsa keys in a separate file"

# cp .env.example .env

# # generate keys
# mkdir keys
# cd keys

# echo "------ Generating private and public keys-------"
# openssl genrsa -out pvt_access.pem 512
# openssl rsa -in pvt_access.pem -outform PEM -pubout -out pub_access.pem

# openssl genrsa -out pvt_refresh.pem 512
# openssl rsa -in pvt_refresh.pem -outform PEM -pubout -out pub_refresh.pem


# echo "------ Your project has been setup, please star our repo if you like it. -------"
