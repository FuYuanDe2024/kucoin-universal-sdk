export type LogLevel = 'debug' | 'info' | 'warn' | 'error';

export interface Logger {
    log: (level: LogLevel, message: string, ...meta: any[]) => void;
    info: (message: string, ...meta: any[]) => void;
    warn: (message: string, ...meta: any[]) => void;
    error: (message: string, ...meta: any[]) => void;
    debug: (message: string, ...meta: any[]) => void;
    setGlobalLogger: (newLogger: Logger, newOptions?: LoggerOptions) => void;
}

export interface LoggerOptions {
    level: LogLevel;
    format: (level: LogLevel, message: string, meta?: any) => string;
}

export const logger: Logger;
