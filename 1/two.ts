const sum = (input: number[]) => {
  return input.reduce((sum, current) => sum + current, 0);
};

export const countWindowedIncreases = (input: number[]): number => {
  return input.reduce((total, current, index) => {
    if (index < 3) {
      return 0;
    }
    const previousSum = sum(input.slice(index - 3, index));
    const currentSum = sum(input.slice(index - 2, index + 1));
    console.log(`${previousSum}:${currentSum}`);
    return currentSum > previousSum ? total + 1 : total;
  }, 0);
};

export const test = async () => {
  const testInput = (await Deno.readTextFile("./test-input.txt"))
    .split("\n")
    .map((i) => parseFloat(i));
  const expectedOutput = 5;
  const realOutput = countWindowedIncreases(testInput);
  if (realOutput !== expectedOutput) {
    throw new Error(`Expected ${expectedOutput}, got ${realOutput}`);
  }
  console.log("Great!");
};

export const process = async () => {
  const input = (await Deno.readTextFile("./one-input.txt"))
    .split("\n")
    .map((i) => parseFloat(i));
  console.log(countWindowedIncreases(input));
};

process();
