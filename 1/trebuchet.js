const fs = require('fs');

// Read input file
const input = fs.readFileSync('input.txt', 'utf8');

// Function to get the calibration value from a line
function getCode(line) {
  const numbers = line.split('').filter(char => /\d/.test(char));
  const code = `${numbers[0]}${numbers[numbers.length - 1]}`;
  return parseInt(code, 10);
}

// Main function
function main() {
  const startTime = new Date();
  const lines = input.trim().split('\n');
  const sum = lines.reduce((acc, line) => acc + getCode(line), 0);
  console.log(sum);
  console.log(`Execution time: ${new Date() - startTime} ms`);
}

main();