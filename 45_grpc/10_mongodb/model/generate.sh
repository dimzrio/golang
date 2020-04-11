#!/bin/bash

protoc mongod.proto --go_out=plugins=grpc:.