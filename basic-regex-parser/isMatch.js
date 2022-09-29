function isMatch(text, pattern) {
  // your code goes here

  // Input: text and a pattern
  // Output: boolean
  

  // PLAN:
  // 0. initialize match character
  let i = 0, j = 0; 
  let sequenceMatch = "";
  // 1. iterate over the text with i and over pattern with j
  while (i < text.length) {
  // 2. compare text[i] and pattern[j]
    if (text[i] !== pattern[j]) {
    // if there's not a match
      // if pattern[j] is '.'
      if (pattern[j] === '.') {
        // increment both i and j by 1
        i += 1;
        j += 1;
      // if pattern[i] is '*'
      } else if (pattern[j] === '*') {
        // set match character to pattern[j - 1]
        sequenceMatch = pattern[j - 1] === '.' ? text[i] : pattern[j - 1];
        // increment i until text character does not equal match character
        while (text[i] === sequenceMatch) {
          i += 1;
        }
        // increment j
        j += 1;
      // if pattern[j+ 1] is '*'
      } else if (pattern[j + 1] === '*') {
        // increment j + 2
        j += 2;
      // else
      } else { 
        // return false
        return false;
      }
    }
    i += 1;
    j += 1;
  }
  return true;
    // return true
}

console.log(isMatch("acd", "ab*c"))

