// 从 @connectrpc/connect-web 导入 transport 和 promise client 创建函数
import { createConnectTransport } from "@connectrpc/connect-web";
import { createClient } from "@connectrpc/connect";
// 从生成的代码导入服务定义
import { CalculatorService } from "../gen/proto/calculator_pb";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080", // Go 后端运行的地址
});

export const calculatorClient = createClient(CalculatorService, transport);
