import React from 'react'

export type TouchOrMouseEvent 
  = React.TouchEvent<HTMLElement>
  | React.MouseEvent<HTMLElement, MouseEvent>

function getClickPosition(e: TouchOrMouseEvent): {clientX: number} {
  if (!!(e as React.TouchEvent<HTMLElement>).touches) {
    return {clientX: (e as React.TouchEvent<HTMLElement>).touches[0].clientX}
  }
  return {clientX: (e as React.MouseEvent<HTMLElement>).clientX}
}

function formatTimeForDisplay(currentTime: number, duration: number): [string, string] {
  let regex = /(\d\d:\d\d:\d\d)/
  if (duration < 60 * 60) {
    regex = /\d\d:(\d\d:\d\d)/
  }
  
  let t = new Date(0, 0, 0, 0, 0, currentTime)
    .toTimeString()
    .match(regex)
    
  let T = new Date(0, 0, 0, 0, 0, duration)
    .toTimeString()
    .match(regex)

  return [t[1], T[1]]
}

export default {
  getClickPosition,
  formatTimeForDisplay,
}