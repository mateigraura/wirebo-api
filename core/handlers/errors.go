package handlers

import "errors"

var SerializationErr = errors.New("serialization error")

var EntityNotFoundErr = errors.New("could not get entity")

var EntityInsertionErr = errors.New("could not save entity")

var InvalidInputProvidedErr = errors.New("invalid input provided")
