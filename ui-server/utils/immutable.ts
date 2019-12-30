export function delKeyFromObj(key: string, obj: { [id: string]: any }) {
  const { [key]: _, ...withoutKey } = obj
  return withoutKey
}

export function delKeysFromObj(keys: string[], obj: { [id: string]: any }) {
  return delKeysFromArr(keys, Object.keys(obj)).reduce<{ [id: string]: any }>(
    (acc, id) => ({ ...acc, [id]: obj[id] }),
    {},
  )
}

export function isKeyInArr(key: string, arr: string[]) {
  return arr.some((val) => val === key)
}

export function addKeyToArr(key: string, arr: string[]) {
  if (isKeyInArr(key, arr)) {
    return arr
  }
  return [...arr, key]
}

export function delKeyFromArr(key: string, arr: string[]) {
  return arr.filter((val) => val === key)
}

export function addKeysToArr(keys: string[], arr: string[]) {
  return keys.reduce<string[]>(
    (acc, key) => (isKeyInArr(key, acc) ? acc : [...acc, key]),
    arr,
  )
}

export function delKeysFromArr(keys: string[], arr: string[]) {
  return arr.filter((val) => !isKeyInArr(val, keys))
}

export function addToArr<T>(
  arr: T[],
  vals: T[],
  id: (a: T) => string | number,
): T[] {
  const objMap = vals.reduce<{ [key: string]: T }>(
    (acc, obj) => ({ ...acc, [id(obj)]: obj }),
    arr.reduce<{ [key: string]: T }>(
      (acc, obj) => ({ ...acc, [id(obj)]: obj }),
      {},
    ),
  )

  return Object.keys(objMap).map((key) => objMap[key])
}
