#!/bin/bash

protoc model/hello.proto --go_out=plugins=grpc:.