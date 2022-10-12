function spiralCopy(inputMatrix) {
  // your code goes here

  const result = [];

  let topLayer = 0;
  let bottomLayer = inputMatrix.length - 1;
  let rightLayer = inputMatrix[0].length - 1;
  let leftLayer = 0;

  function validate(r, c) {
    if (r <= bottomLayer && r >= topLayer && c >= leftLayer && c <= rightLayer) {
      return true;
    }
    return false;
  }

  function right(arr, r, c) {
    if (!validate(r, c)) {
      return result;
    }
    if (c === rightLayer) {
      result.push(arr[r][c]);
      topLayer++;
      return down(arr, r + 1, c)
    }
    result.push(arr[r][c])
    return right(arr, r, c + 1);
  }
  
  function down(arr, r, c) {
    if (!validate(r, c)) {
      return result;
    }
    if (r === bottomLayer) {
      result.push(arr[r][c]);
      rightLayer--;
      return left(arr, r, c - 1);
    }
    result.push(arr[r][c]);
    return down(arr, r + 1, c);
  }

  function left(arr, r, c) {
    if (!validate(r, c)) {
      return result;
    }
    if (c === leftLayer) {
      result.push(arr[r][c]);
      bottomLayer--;
      return up(arr, r - 1, c);
    }
    result.push(arr[r][c]);
    return left(arr, r, c - 1);
  }

  function up(arr, r, c) {
    if (!validate(r, c)) {
      return result;
    }
    if (r === topLayer) {
      result.push(arr[r][c]);
      leftLayer++;
      return right(arr, r, c + 1);
    }
    result.push(arr[r][c]);
    return up(arr, r - 1, c);
  }

  return right(inputMatrix, 0, 0);
}

console.log(spiralCopy([ [1,    2,   3,  4,    5],
                         [6,    7,   8,  9,   10],
                         [11,  12,  13,  14,  15],
                         [16,  17,  18,  19,  20] ]))
