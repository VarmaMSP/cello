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
  const baseUrl =
    process.env.NODE_ENV === 'development'
      ? 'http://localhost:8080'
      : 'https://phenopod.com'

  switch (size) {
    case 'lg':
    case 'md':
      return `${baseUrl}/img/${id}-500x500.jpg`
    case 'sm':
      return `${baseUrl}/img/${id}-250x250.jpg`
  }
}
