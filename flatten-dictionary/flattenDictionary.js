function flattenDictionary(dict) {
  const flatDictionary = {};
  subroutine("", dict, flatDictionary)
  return flatDictionary;
}

function subroutine(initialKey, input, output) {
  // iterate over each key
  for (const key in input) {
  // for each key, get the value
    const value = input[key];
  // if the value is not an object
    if (value.constructor != Object) {
      // if the initial key is null or empty string
      if (!initialKey || initialKey.length < 1) { 
        // add the key-value to the output dictionary
        output[key] = value;
      } else {
        // otherwise append the key to the initial key separated by a "."
        output[`${initialKey}${key ? "."+key : ""}`] = value;
      }
    } else {
      // else, when the key is an instance of a dictionary
        // if the initial key is null or empty string
        if (!initialKey || initialKey.length < 1) {
        // recurse with key, value, input
          return subroutine(key, value, output);
        }
        // otherwise, recurse with the appended key and value, input
        return subroutine(`${initialKey}${key ? "."+key : ""}`, value, output);
    }
  }
}

let result = flattenDictionary({
            "Key1" : "1",
            "Key2" : {
                "a" : "2",
                "b" : "3",
                "c" : {
                    "d" : "3",
                    "e" : {
                        "" : "1"
                    }
                }
            }
        }
)
console.log(result)
