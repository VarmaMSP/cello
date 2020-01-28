import { createPopper } from '@popperjs/core'
import * as PopperJS from '@popperjs/core/lib/types'
import { useEffect, useRef, useState } from 'react'
import useCallbackRef from './useCallbackRef'

function usePopper(
  options: PopperJS.Options,
  onPopperClickOutside?: () => void,
) {
  const popperInstance = useRef<PopperJS.Instance>()
  const [reference, referenceRef] = useCallbackRef<HTMLElement>()
  const [popper, popperRef] = useCallbackRef<HTMLElement>()
  const [styles, setStyles] = useState<{
    [key: string]: Partial<CSSStyleDeclaration>
  }>({})

  useEffect(() => {
    const cleanUp = () => {
      setStyles({})
      if (!!popperInstance.current) {
        popperInstance.current.destroy()
        popperInstance.current = undefined
      }
    }

    cleanUp()
    if (!!reference && !!popper) {
      popperInstance.current = createPopper(reference, popper, {
        ...options,
        modifiers: [
          ...(options.modifiers || []),
          {
            name: 'applyStyles',
            fn: ({ state }) => setStyles(state.styles),
          },
        ],
      })

      return cleanUp
    }
  }, [reference, popper, options.placement])

  useEffect(() => {
    const fn: EventListener = (e) => {
      if (!!popper && !popper.contains(e.target as any)) {
        onPopperClickOutside && onPopperClickOutside()
      }
    }

    const cleanUp = () => {
      if (window.PointerEvent) {
        document.removeEventListener('pointerdown', fn)
      } else {
        document.removeEventListener('mousedown', fn)
        document.removeEventListener('touchstart', fn)
      }
    }

    cleanUp()
    if (!!popper) {
      if (window.PointerEvent) {
        document.addEventListener('pointerdown', fn)
      } else {
        document.addEventListener('mousedown', fn)
        document.addEventListener('touchstart', fn)
      }

      return cleanUp
    }
  }, [popper])

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
