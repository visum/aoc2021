// use with deno

export const countDepthIncreases = (input: number[]) => {
  const calc = input.reduce(
    (last, current) => {
      console.log(last, current);
      const [previous, total] = last as [number, number];
      if (previous == null) {
        return [current, 0];
      }
      const newTotal = total + (current > previous ? 1 : 0);
      return [current, newTotal];
    },
    [null, 0]
  );

  return calc[1];
};

export const test = async () => {
  const testInput = (await Deno.readTextFile("./test-input.txt"))
  .split("\n")
  .map((i) => parseFloat(i));

  const testOutput = 7;

  const testAnswer = countDepthIncreases(testInput);
  if (testAnswer === testOutput) {
    console.log("Great!");
  } else {
    throw new Error(`Test failed: got ${testAnswer}, expected ${testOutput}`);
  }
};

export const process = async () => {
  const inputFile = './input.txt';
  const text = await Deno.readTextFile(inputFile);
  console.log(`Result: ${countDepthIncreases(text.split('\n').map(i => parseFloat(i)))}`);
};

process();