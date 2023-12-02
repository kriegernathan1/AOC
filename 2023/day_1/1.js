const fs = require("fs");

fs.open("./1.txt", "r", (err, fd) => {
  fs.readFile(fd, "utf-8", (err, data) => {
    const lines = data.split("\n");
    let sum = 0;
    const isDigit = (char) => /^[0-9]+$/i.test(char);

    lines.forEach((line) => {
      let left = 0;
      let right = line.length - 1;

      while (left < line.length) {
        if (isDigit(line[left])) break;
        left++;
      }

      while (right >= 0) {
        if (isDigit(line[right])) break;
        right--;
      }

      let calibrationValue = line[left] + line[right];
      sum = sum += Number(calibrationValue);
    });
    console.log(sum);
  });
});
