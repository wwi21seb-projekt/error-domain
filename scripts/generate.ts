// generate.ts
import { readFileSync, writeFileSync } from 'fs';

interface CustomError {
    message: string;
    code: string;
    httpStatus: number;
}

const errors: CustomError[] = JSON.parse(readFileSync('errors/errors.json', 'utf8'));

let content = `export class CustomError extends Error {
    code: string;
    httpStatus: number;

    constructor(message: string, code: string, httpStatus: number) {
        super(message);
        this.code = code;
        this.httpStatus = httpStatus;
    }
}\n\n`;

errors.forEach(error => {
    content += `export const ${error.code} = new CustomError(\n\t"${error.message}", \n\t"${error.code}", \n\t${error.httpStatus}\n);\n`;
});

writeFileSync('ts-errors/errors.ts', content);
