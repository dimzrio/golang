#!/bin/bash

protoc model/calc.proto --go_out=plugins=grpc:.