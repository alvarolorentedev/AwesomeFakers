# Awesome Fakers

Small set of resources and exampels of faker libraries for multiple languages

## What are faker libraries?

Faker libraries are resources used to generate random inputs.

## Why to used them?

- ### Non Static Data: 

Fakers are very useful for testing as it provides non static data that can help test to be more meanningful and robusts.

```js
let sum = (first, second) => { return 3 }

test('adds 1 + 2 to equal 3', () => {
    expect(sum(1, 2)).toBe(3);
});
```

- ### Simplify complexity of tests:

Fakers allow to simplify the input and complexity of your test by generating random data, In the next 

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



## Faker Libraries

### Javascript

- [Faker.js](https://github.com/marak/Faker.js/)
