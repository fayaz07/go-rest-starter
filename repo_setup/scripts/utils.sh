#!/bin/bash

function print_status(){
  echo  "$1"
}

function print_error(){
  echo -e "\e[31m$1\e[0m"
}

function print_success(){
  echo -e "\e[32m$1\e[0m"
}

function print_warning(){
  echo -e "\e[33m$1\e[0m"
}

function print_info(){
  echo -e "\e[34m$1\e[0m"
}

function print_bold(){
  echo -e "\e[1m$1\e[0m"
}

function print_header(){
  echo -e "\n\n\e[44m$1\e[0m\n-------------------------------------------------------------------------------------"
}
