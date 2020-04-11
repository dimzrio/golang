#!/bin/bash

protoc model/hellopb.proto --go_out=plugins=grpc:.