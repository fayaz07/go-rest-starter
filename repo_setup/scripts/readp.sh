#!/bin/bash

function loadProperties() {
  local fileName=$1
  local prefixKey=$2

  if [ ! -f "${fileName}" ]; then
    echo "${fileName} not found!"
    return 1
  fi

  while IFS='=' read -r origKey value; do
    local key=${origKey}
    key=${key//[!a-zA-Z0-9_]/_} 
    if [[ "${origKey}" == "#"*   ]]; then
      local ignoreComments
    elif [ -z "${key}" ]; then
      local emptyLine
    else
      if [[ "${prefixKey}${key}" =~ ^[0-9].* ]]; then
        key=_${key}
      fi
      eval ${prefixKey}${key}=\${value}
    fi
  done < <(grep "" ${fileName})
}

