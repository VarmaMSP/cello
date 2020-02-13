import React from 'react'

export type TouchOrMouseEvent =
  | React.TouchEvent<HTMLElement>
  | React.MouseEvent<HTMLElement, MouseEvent>

export function getClickPosition(e: TouchOrMouseEvent): { clientX: number } {
  return !!(e as React.TouchEvent<HTMLElement>).touches
    ? { clientX: (e as React.TouchEvent<HTMLElement>).touches[0].clientX }
    : { clientX: (e as React.MouseEvent<HTMLElement>).clientX }
}

export function stopEventPropagation(e: React.SyntheticEvent<HTMLElement>) {
  e.stopPropagation()
  e.nativeEvent.stopImmediatePropagation()
}

export function getImageUrl(urlPath: string): string {
  return process.env.NODE_ENV === 'development'
    ? `http://localhost:8080/thumbnails/${urlPath}.jpg`
    : `https://cdn.phenopod.com/thumbnails/${urlPath}.jpg`
}

export function getAssetUrl(urlPath: string): string {
  return process.env.NODE_ENV === 'development'
    ? `http://localhost:8080/assets/${urlPath}`
    : `https://cdn.phenopod.com/assests/${urlPath}`
}
