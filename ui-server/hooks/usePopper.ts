import { createPopper } from '@popperjs/core'
import * as PopperJS from '@popperjs/core/lib/types'
import { useEffect, useRef, useState } from 'react'
import useCallbackRef from './useCallbackRef'

function usePopper(options: PopperJS.Options) {
  const popperInstance = useRef<PopperJS.Instance>()
  const [reference, referenceRef] = useCallbackRef<HTMLElement>()
  const [popper, popperRef] = useCallbackRef<HTMLElement>()
  const [styles, setStyles] = useState<{
    [key: string]: Partial<CSSStyleDeclaration>
  }>({})

  useEffect(() => {
    if (!!popperInstance.current) {
      popperInstance.current.destroy()
      popperInstance.current = undefined
    }

    if (!!reference && !!popper) {
      popperInstance.current = createPopper(reference, popper, {
        ...options,
        modifiers: [
          ...(options.modifiers || []),
          {
            name: 'applyStyles',
            fn: ({ state }) => {
              setStyles(state.styles)
            },
          },
        ],
      })  

      return () => {
        if (popperInstance.current) {
          popperInstance.current.destroy()
          popperInstance.current = undefined
        }
      }
    }
  }, [reference, popper, options.placement])

  return [
    {
      ref: referenceRef,
      styles: {} as React.CSSProperties,
    },
    {
      ref: popperRef,
      styles: (styles['popper'] || {}) as React.CSSProperties,
    },
  ]
}

export default usePopper
