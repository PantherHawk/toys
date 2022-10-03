function reverseWords(arr) {
  // your code goes here

  // input: letters of words in a sentence separated by a space
  // output: letters of words in the sentence in reverse order

  // PLAN:
  const stack = [];
  let acc = '';
  // iterate over each character
  for (let c of arr) {
  // push each word onto stack by accumulating words
    if (c.trim() === '') {
      stack.push(acc);
      acc = '';
    } else {
      acc += c;
    }
  }
  stack.push(acc);
  console.log("stack: ", stack)
  arr = [];
  while(stack.length - 1) {
  // pop each word from stack and add to result
    arr.push(...stack.pop(), ' ')
  }
  arr.push(...stack.pop())
  return arr;
}

