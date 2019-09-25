import React from 'react'

export type TouchOrMouseEvent =
  | React.TouchEvent<HTMLElement>
  | React.MouseEvent<HTMLElement, MouseEvent>

export function getClickPosition(e: TouchOrMouseEvent): { clientX: number } {
  return !!(e as React.TouchEvent<HTMLElement>).touches
    ? { clientX: (e as React.TouchEvent<HTMLElement>).touches[0].clientX }
    : { clientX: (e as React.MouseEvent<HTMLElement>).clientX }
}

export function getImageUrl(id: string, size: 'sm' | 'md' | 'lg'): string {
  switch (size) {
    case 'lg':
    case 'md':
      return `${process.env.IMAGE_BASE_URL}/${id}-500x500.jpg`
    case 'sm':
      return `${process.env.IMAGE_BASE_URL}/${id}-250x250.jpg`
  }
}
