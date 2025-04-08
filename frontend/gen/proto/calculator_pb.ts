// @generated by protoc-gen-es v2.2.5 with parameter "target=ts"
// @generated from file proto/calculator.proto (package calculator.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file proto/calculator.proto.
 */
export const file_proto_calculator: GenFile = /*@__PURE__*/
  fileDesc("ChZwcm90by9jYWxjdWxhdG9yLnByb3RvEg1jYWxjdWxhdG9yLnYxIjIKDk9wZXJhdGVSZXF1ZXN0Eg8KB251bWJlcjEYASABKAESDwoHbnVtYmVyMhgCIAEoASIhCg9PcGVyYXRlUmVzcG9uc2USDgoGcmVzdWx0GAEgASgBMrgCChFDYWxjdWxhdG9yU2VydmljZRJECgNBZGQSHS5jYWxjdWxhdG9yLnYxLk9wZXJhdGVSZXF1ZXN0Gh4uY2FsY3VsYXRvci52MS5PcGVyYXRlUmVzcG9uc2USSQoIU3VidHJhY3QSHS5jYWxjdWxhdG9yLnYxLk9wZXJhdGVSZXF1ZXN0Gh4uY2FsY3VsYXRvci52MS5PcGVyYXRlUmVzcG9uc2USSQoITXVsdGlwbHkSHS5jYWxjdWxhdG9yLnYxLk9wZXJhdGVSZXF1ZXN0Gh4uY2FsY3VsYXRvci52MS5PcGVyYXRlUmVzcG9uc2USRwoGRGl2aWRlEh0uY2FsY3VsYXRvci52MS5PcGVyYXRlUmVxdWVzdBoeLmNhbGN1bGF0b3IudjEuT3BlcmF0ZVJlc3BvbnNlQjBaLkNvbm5lY3RSUEMtRGVtby9nZW4vY2FsY3VsYXRvci92MTtjYWxjdWxhdG9ydjFiBnByb3RvMw");

/**
 * @generated from message calculator.v1.OperateRequest
 */
export type OperateRequest = Message<"calculator.v1.OperateRequest"> & {
  /**
   * @generated from field: double number1 = 1;
   */
  number1: number;

  /**
   * @generated from field: double number2 = 2;
   */
  number2: number;
};

/**
 * Describes the message calculator.v1.OperateRequest.
 * Use `create(OperateRequestSchema)` to create a new message.
 */
export const OperateRequestSchema: GenMessage<OperateRequest> = /*@__PURE__*/
  messageDesc(file_proto_calculator, 0);

/**
 * @generated from message calculator.v1.OperateResponse
 */
export type OperateResponse = Message<"calculator.v1.OperateResponse"> & {
  /**
   * @generated from field: double result = 1;
   */
  result: number;
};

/**
 * Describes the message calculator.v1.OperateResponse.
 * Use `create(OperateResponseSchema)` to create a new message.
 */
export const OperateResponseSchema: GenMessage<OperateResponse> = /*@__PURE__*/
  messageDesc(file_proto_calculator, 1);

/**
 * @generated from service calculator.v1.CalculatorService
 */
export const CalculatorService: GenService<{
  /**
   * @generated from rpc calculator.v1.CalculatorService.Add
   */
  add: {
    methodKind: "unary";
    input: typeof OperateRequestSchema;
    output: typeof OperateResponseSchema;
  },
  /**
   * @generated from rpc calculator.v1.CalculatorService.Subtract
   */
  subtract: {
    methodKind: "unary";
    input: typeof OperateRequestSchema;
    output: typeof OperateResponseSchema;
  },
  /**
   * @generated from rpc calculator.v1.CalculatorService.Multiply
   */
  multiply: {
    methodKind: "unary";
    input: typeof OperateRequestSchema;
    output: typeof OperateResponseSchema;
  },
  /**
   * @generated from rpc calculator.v1.CalculatorService.Divide
   */
  divide: {
    methodKind: "unary";
    input: typeof OperateRequestSchema;
    output: typeof OperateResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_proto_calculator, 0);

