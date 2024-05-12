// generate.ts
import { readFileSync, writeFileSync } from 'fs';
import { fileURLToPath } from 'url';
import {dirname, join} from 'path';

interface CustomError {
    message: string;
    code: string;
    httpStatus: number;
}

// Calculate the directory of the current module using new ESM features
const currentDir = dirname(fileURLToPath(import.meta.url));

// Construct the path to the JSON file relative to the current directory
const jsonFilePath = join(currentDir, '../../errors/errors.json');
const errors: CustomError[] = JSON.parse(readFileSync(jsonFilePath, 'utf8'));

let content = `/** GENERATED FILE - DO NOT MODIFY DIRECTLY **/\n\n
export class CustomError extends Error {
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

// Specify the output path for the generated TypeScript file
const outputPath = join(currentDir, '../../ts-errors/src/errors.ts');
writeFileSync(outputPath, content);
