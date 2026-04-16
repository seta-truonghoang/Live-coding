const fs = require('fs/promises');

// Example function to read a file (returns a Promise)
async function readFileContent(filename) {
    try {
        const content = await fs.readFile(filename, 'utf-8');
        return content;
    } catch (err) {
        // Throw a new error with custom data, or simply rethrow
        throw new Error(`Could not read file ${filename}: ${err.message}`);
    }
}

async function main() {
    const fileName = 'test.txt';
    
    try {
        // Call the function inside a try-catch block to catch errors
        const data = await readFileContent(fileName);
        console.log(`File read successfully: ${data}`);
    } catch (err) {
        console.error("An error occurred:", err.message);
        // Handle the error or exit the process
    }
}

main();
