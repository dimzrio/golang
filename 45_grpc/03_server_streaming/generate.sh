#!/bin/bash

protoc model/greeting.proto --go_out=plugins=grpc:.