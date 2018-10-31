# Fakers?


Small set of resources and exampels of faker libraries for multiple languages

## What are faker libraries?

Faker libraries are resources used to generate random inputs.

## Why to used them?

- Static Data is evil: Static data can cause your test not to be meanningful and leave most of the cases untested.

```js
let sum = (first, second) => { return 3 }

test('adds 1 + 2 to equal 3', () => {
    expect(sum(1, 2)).toBe(3);
});
```

- Test imput complexity: It is posible to generate meaningful tests without fakers but in exchange you will require to pass this data as input to the test. 

```js
let sum = (first, second) => { return first+second }

[
    [1,2,3],
    [1,3,4],
].forEach((param) => {
    test(`adds ${param[0]} + ${param[1]} to equal ${param[2]}`, () => {
        expect(sum(param[0], param[1])).toBe(param[2]);
      });
})
```
- Gives meaning to test coverage: as you can see on the first example the test coverage is 100% and if you read the test is seems meaningful but the code behind is not.

