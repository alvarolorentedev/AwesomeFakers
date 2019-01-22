let sumExample1 = (first, second) => { return 3 }
let sumExample2 = (first, second) => { return first+second }

test('adds 1 + 2 to equal 3', () => {
    expect(sumExample1(1, 2)).toBe(3);
  });

  test.each([
    [1,2,3],
    [1,3,4]
  ])(`adds %i + %i to equal %i`,(...param) => {
        expect(sumExample2(param[0], param[1])).toBe(param[2]);
})