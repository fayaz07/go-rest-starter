#!/bin/bash

echo "------ Setting up your go-project -------"

cp .env.example .env

# generate keys
mkdir keys
cd keys

echo "------ Generating private and public keys-------"
openssl genrsa -out pvt_access.pem 512
openssl rsa -in pvt_access.pem -outform PEM -pubout -out pub_access.pem

openssl genrsa -out pvt_refresh.pem 512
openssl rsa -in pvt_refresh.pem -outform PEM -pubout -out pub_refresh.pem


echo "------ Your project has been setup, please star our repo if you like it. -------"