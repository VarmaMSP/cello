import React from 'react'
export type Icon =
  | 'arrow-left'
  | 'cheveron-up'
  | 'explore'
  | 'fast-forward'
  | 'fast-rewind'
  | 'feed'
  | 'heart'
  | 'pause'
  | 'play'
  | 'play-outline'
  | 'search'
  | 'user-solid-circle'

export const iconMap: { [key in Icon]: React.SFC<{ className: string }> } = {
  'arrow-left': (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M3.828 9l6.071-6.071-1.414-1.414L0 10l.707.707 7.778 7.778 1.414-1.414L3.828 11H20V9H3.828z" />
    </svg>
  ),

  'cheveron-up': (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M10.707 7.05L10 6.343 4.343 12l1.414 1.414L10 9.172l4.243 4.242L15.657 12z" />
    </svg>
  ),

  explore: (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M10 20a10 10 0 110-20 10 10 0 010 20zM7.88 7.88l-3.54 7.78 7.78-3.54 3.54-7.78-7.78 3.54zM10 11a1 1 0 110-2 1 1 0 010 2z" />
    </svg>
  ),

  'fast-forward': (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M1 5l9 5-9 5V5zm9 0l9 5-9 5V5z" />
    </svg>
  ),

  'fast-rewind': (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M19 5v10l-9-5 9-5zm-9 0v10l-9-5 9-5z" />
    </svg>
  ),

  feed: (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M6 6V2c0-1.1.9-2 2-2h10a2 2 0 012 2v10a2 2 0 01-2 2h-4v4a2 2 0 01-2 2H2a2 2 0 01-2-2V8c0-1.1.9-2 2-2h4zm2 0h4a2 2 0 012 2v4h4V2H8v4zM2 8v10h10V8H2z" />
    </svg>
  ),

  heart: (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M10 3.22l-.61-.6a5.5 5.5 0 00-7.78 7.77L10 18.78l8.39-8.4a5.5 5.5 0 00-7.78-7.77l-.61.61z" />
    </svg>
  ),

  pause: (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M5 4h3v12H5V4zm7 0h3v12h-3V4z" />
    </svg>
  ),

  play: (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M4 4l12 6-12 6z" />
    </svg>
  ),

  'play-outline': (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M2.93 17.07A10 10 0 1117.07 2.93 10 10 0 012.93 17.07zm12.73-1.41A8 8 0 104.34 4.34a8 8 0 0011.32 11.32zM7 6l8 4-8 4V6z" />
    </svg>
  ),

  search: (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M12.9 14.32a8 8 0 111.41-1.41l5.35 5.33-1.42 1.42-5.33-5.34zM8 14A6 6 0 108 2a6 6 0 000 12z" />
    </svg>
  ),

  'user-solid-circle': (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M10 20a10 10 0 110-20 10 10 0 010 20zM7 6v2a3 3 0 106 0V6a3 3 0 10-6 0zm-3.65 8.44a8 8 0 0013.3 0 15.94 15.94 0 00-13.3 0z" />
    </svg>
  ),
}
