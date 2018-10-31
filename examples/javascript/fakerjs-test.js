const faker = require('faker')

let sum = (first, second) => { return first+second }

test('adds 2 numbers', () => {
    let first = faker.random.number(100)
    let second = faker.random.number(100)
    let result = first+second
    expect(sum(first, second)).toBe(result);
});
