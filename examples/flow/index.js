// @flow

type someOtherType = {
    a: number | string | boolean,
    b: boolean,
    c: string
}

export type some = {
    a: number,
    b: someOtherType,
}

