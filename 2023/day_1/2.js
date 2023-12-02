const fs = require("fs");

const digitWords = [
  "one",
  "two",
  "three",
  "four",
  "five",
  "six",
  "seven",
  "eight",
  "nine",
];

const wordDigitToNumberMap = new Map();
digitWords.forEach((value, index) => {
  wordDigitToNumberMap.set(value, (index + 1).toString());
});

function doesWordIncludeDigit(word) {
  for (var i = 0; i < digitWords.length; i++) {
    if (word.includes(digitWords[i])) return [true, digitWords[i]];
  }
  return [false, undefined];
}

fs.open("./1.txt", "r", (_, fd) => {
  fs.readFile(fd, "utf-8", (__, data) => {
    const lines = data.split("\n");
    let sum = 0;
    const isDigit = (char) => /^[0-9]+$/i.test(char);

    lines.forEach((line) => {
      let left = 0;
      let right = line.length - 1;

      let leftStr = "";
      let rightStr = "";

      let firstDigit = "";
      let secondDigit = "";

      while (left < line.length) {
        leftStr = leftStr + line[left];

        if (isDigit(line[left])) {
          firstDigit = line[left];
          break;
        }

        const doesIncludeDigit = doesWordIncludeDigit(leftStr);
        if (doesIncludeDigit[0]) {
          firstDigit = wordDigitToNumberMap.get(doesIncludeDigit[1]);
          break;
        }

        left++;
      }

      while (right >= 0) {
        rightStr = line[right] + rightStr;

        if (isDigit(line[right])) {
          secondDigit = line[right];
          break;
        }

        const doesIncludeDigit = doesWordIncludeDigit(rightStr);
        if (doesIncludeDigit[0]) {
          secondDigit = wordDigitToNumberMap.get(doesIncludeDigit[1]);
          break;
        }

        right--;
      }

      const calibrationVal = Number(firstDigit + secondDigit);

      sum += calibrationVal;
    });

    console.log(sum);
  });
});
