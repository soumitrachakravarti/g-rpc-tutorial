#!/bin/bash

protoc calculator/calculator-pb/calculator.proto --go_out=plugins=grpc:.