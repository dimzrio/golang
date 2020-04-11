#!/bin/bash

protoc model/sumManyNumber.proto --go_out=plugins=grpc:.