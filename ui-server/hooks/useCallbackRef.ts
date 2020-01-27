import { useCallback, useState } from 'react'

function useCallbackRef<E extends Element>(): [
  E | null,
  (element: E | null) => void,
] {
  const [element, setElement] = useState<E | null>(null)
  return [element, useCallback(setElement, [])]
}

export default useCallbackRef
