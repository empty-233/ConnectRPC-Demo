'use client';

import React, { useState, useEffect } from 'react';
import { calculatorClient } from '../client/connect-client';
import { ConnectError } from '@connectrpc/connect';
import styles from './calculator.module.css';
import { MoonIcon, SunIcon } from './icons';

// 定义操作符类型
type Operator = 'add' | 'subtract' | 'multiply' | 'divide';

// 操作符映射表
const operatorSymbols: Record<Operator, string> = {
  add: '+',
  subtract: '-',
  multiply: '×',
  divide: '÷',
};

export default function Calculator() {
  // 状态管理
  const [number1, setNumber1] = useState<string>('');
  const [number2, setNumber2] = useState<string>('');
  const [operator, setOperator] = useState<Operator>('add');
  const [result, setResult] = useState<number | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [isClient, setIsClient] = useState<boolean>(false);
  const [darkMode, setDarkMode] = useState<boolean>(false);

  // 客户端挂载检查与深色模式初始化 - 修复水合问题
  useEffect(() => {
    setIsClient(true);
    
    // 所有DOM操作移到客户端执行
    const prefersDarkMode = window.matchMedia('(prefers-color-scheme: dark)').matches;
    const savedTheme = localStorage.getItem('calculatorTheme');
    const initialDarkMode = savedTheme ? savedTheme === 'dark' : prefersDarkMode;
    
    setDarkMode(initialDarkMode);
  }, []);
  
  // 切换主题模式
  const toggleDarkMode = () => {
    const newDarkMode = !darkMode;
    setDarkMode(newDarkMode);
    
    // 仅在客户端执行
    if (typeof window !== 'undefined') {
      localStorage.setItem('calculatorTheme', newDarkMode ? 'dark' : 'light');
    }
  };

  // 处理计算逻辑
  const handleCalculate = async () => {
    // 重置状态
    setIsLoading(true);
    setError(null);
    setResult(null);

    // 解析输入值
    const num1 = parseFloat(number1);
    const num2 = parseFloat(number2);

    // 输入验证
    if (isNaN(num1) || isNaN(num2)) {
      setError('请输入有效的数字');
      setIsLoading(false);
      return;
    }

    // 除零检查
    if (operator === 'divide' && num2 === 0) {
      setError('除数不能为零');
      setIsLoading(false);
      return;
    }

    try {
      // 准备请求数据
      const request = { number1: num1, number2: num2 };
      let response;

      // 使用 ConnectRPC 调用后端
      switch (operator) {
        case 'add':
          response = await calculatorClient.add(request);
          break;
        case 'subtract':
          response = await calculatorClient.subtract(request);
          break;
        case 'multiply':
          response = await calculatorClient.multiply(request);
          break;
        case 'divide':
          response = await calculatorClient.divide(request);
          break;
        default:
          throw new Error('无效的操作符');
      }

      // 更新结果
      setResult(response.result);
      
    } catch (err) {
      // 错误处理
      if (err instanceof ConnectError) {
        console.error('Connect Error:', err);
        setError(`计算错误: ${err.message} (Code: ${err.code})`);
      } else {
        console.error('Unexpected Error:', err);
        setError('发生未知错误，请稍后重试');
      }
    } finally {
      setIsLoading(false);
    }
  };

  // 清除计算数据
  const handleClear = () => {
    setNumber1('');
    setNumber2('');
    setResult(null);
    setError(null);
  };

  // 使用简单的骨架屏，确保服务端和客户端初始HTML结构完全一致
  if (!isClient) {
    return (
      <div className={styles.container}>
        <div className={styles.header}>
          <h2 className={styles.title}>ConnectRPC 计算器</h2>
        </div>
      </div>
    );
  }

  return (
    <div className={`${styles.container} ${darkMode ? styles.darkMode : ''}`}>
      <div className={styles.header}>
        <h2 className={styles.title}>ConnectRPC 计算器</h2>
        <button 
          className={styles.themeToggle}
          onClick={toggleDarkMode}
          aria-label={darkMode ? "切换到亮色模式" : "切换到深色模式"}
        >
          {darkMode ? <SunIcon /> : <MoonIcon />}
        </button>
      </div>
      
      <div className={styles.calculatorGrid}>
        <div className={styles.inputGroup}>
          <label htmlFor="num1">数字 1</label>
          <input
            id="num1"
            type="number"
            value={number1}
            onChange={(e) => setNumber1(e.target.value)}
            placeholder="输入第一个数字"
            className={styles.input}
          />
        </div>
        
        <div className={styles.inputGroup}>
          <label htmlFor="operator">操作符</label>
          <select
            id="operator"
            value={operator}
            onChange={(e) => setOperator(e.target.value as Operator)}
            className={styles.select}
          >
            <option value="add">+</option>
            <option value="subtract">-</option>
            <option value="multiply">×</option>
            <option value="divide">÷</option>
          </select>
        </div>
        
        <div className={styles.inputGroup}>
          <label htmlFor="num2">数字 2</label>
          <input
            id="num2"
            type="number"
            value={number2}
            onChange={(e) => setNumber2(e.target.value)}
            placeholder="输入第二个数字"
            className={styles.input}
          />
        </div>
      </div>
      
      <div className={styles.buttonGroup}>
        <button
          onClick={handleCalculate}
          disabled={isLoading}
          className={`${styles.button} ${styles.primaryButton}`}
        >
          {isLoading ? '计算中...' : '计算'}
        </button>
        
        <button
          onClick={handleClear}
          className={styles.button}
          disabled={isLoading}
        >
          清除
        </button>
      </div>
      
      {error && (
        <div className={styles.error}>
          {error}
        </div>
      )}
      
      {result !== null && (
        <div className={styles.result}>
          <h3>计算结果</h3>
          <div className={styles.equation}>
            {number1} {operatorSymbols[operator]} {number2} = <span className={styles.resultValue}>{result}</span>
          </div>
        </div>
      )}
    </div>
  );
}
