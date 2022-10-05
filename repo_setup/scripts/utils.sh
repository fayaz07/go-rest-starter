#!/bin/bash

function print_status(){
  echo  "$1"
}

function print_error(){
  echo -e "\033[31m$1\033[0m"
}

function print_success(){
  echo -e "\033[32m$1\033[0m"
}

function print_warning(){
  echo -e "\033[33m$1\033[0m"
}

function print_info(){
  echo -e "\033[34m$1\033[0m"
}

function print_bold(){
  echo -e "\033[1m$1\033[0m"
}

function print_header(){
  echo -e "\n\n\033[44m$1\033[0m\n-------------------------------------------------------------------------------------"
}

function print_banner() {
  msg="|$*|"
  msg1="| $* |"
  actual=" $* "
  space=$(echo "$actual" | sed 's/./ /g')
  edge=$(echo "$msg" | sed 's/./-/g')

  echo "+$edge+"
  echo "|$space|"
  echo -e "\033[1m$msg1\033[0m"  
  echo "|$space|"
  echo "+$edge+"

}

function end_line() {
  echo -e ""
}

# print_bold "Hello this is bold"
# print_warning "Hello this is a warning"
# print_info "Hello this is info"
# print_header "hello this is a header"
# print_success "hello this is a Success"
# print_error "hello this is a error"
# print_status "Hello this is status"
